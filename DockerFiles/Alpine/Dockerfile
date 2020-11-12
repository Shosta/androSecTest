# Build the AndroSecTest App on the golang latest image.
FROM golang:1.15-alpine as go-builder
LABEL go-builder=true

# Environmentn variables
ENV SRC_DIR=/go/src/github.com/Shosta/androSecTest
ENV GIT_SSL_NO_VERIFY=1

# Set the Current Working Directory inside the container
WORKDIR $SRC_DIR

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Dowload the Go Dependancies
RUN apk add git && go get ./...

# Build the Go app for a Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o androSecTest .



# Download the Required Hacking Tools
FROM alpine:latest as alpine-downloader

ENV HACKTOOLS_DIR=/home/Developpement/HackingTools

WORKDIR $HACKTOOLS_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apk add wget unzip

# Install SignApk
RUN mkdir -p SignApkUtils && \
    wget --no-check-certificate --quiet -O ./SignApkUtils/signapk.jar https://github.com/techexpertize/SignApk/blob/master/signapk.jar

# Install jadx
RUN wget --no-check-certificate --quiet https://github.com/skylot/jadx/releases/download/v1.1.0/jadx-1.1.0.zip && \
    mkdir -p ./DecompilingAndroidAppUtils/jadx && \
    unzip jadx-1.1.0.zip -d ./DecompilingAndroidAppUtils/jadx && rm jadx-1.1.0.zip

# Download apktool-2 & Rename downloaded jar to apktool.jar
RUN mkdir -p ./DecompilingAndroidAppUtils/apktool && \
    wget --no-check-certificate --quiet -O ./DecompilingAndroidAppUtils/apktool/apktool.jar https://bitbucket.org/iBotPeaches/apktool/downloads/apktool_2.4.1.jar

# Install Humpty-dumpty
RUN mkdir -p ./humpty-dumpty-android-master && \
    wget --no-check-certificate --quiet -O ./humpty-dumpty-android-master/humpty.sh https://github.com/Pixplicity/humpty-dumpty-android/blob/master/humpty.sh



FROM alpine:latest

# Labels and Credits
LABEL \
    name="AndroSecTest on Alpine" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static."

ENV SRC_DIR=/root/go/src/github.com/Shosta/androSecTest
ENV HACKTOOLS_DIR=/home/Developpement/HackingTools
WORKDIR $SRC_DIR

RUN apk add openjdk8 \
    usbutils \
    unzip
RUN apk --update-cache --repository http://dl-3.alpinelinux.org/alpine/edge/testing/ add android-tools

# Copy jadx and apktool
COPY --from=alpine-downloader $HACKTOOLS_DIR/DecompilingAndroidAppUtils $HACKTOOLS_DIR/DecompilingAndroidAppUtils
RUN chmod +x $HACKTOOLS_DIR/DecompilingAndroidAppUtils/apktool/apktool.jar

# Copy Humpty-dumpty
COPY --from=alpine-downloader $HACKTOOLS_DIR/humpty-dumpty-android-master $HACKTOOLS_DIR/humpty-dumpty-android-master

# Copy the built executable from the go-builder container and add it to this container.
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/androSecTest $SRC_DIR

# Copy the User Settings
RUN mkdir  $SRC_DIR/.res
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/res/ $SRC_DIR/.res/

# Copy SignApk to Proper Location in Container
RUN mv $SRC_DIR/.res/SignApk $HACKTOOLS_DIR/SignApkUtils/ && \
    chmod +x $HACKTOOLS_DIR/SignApkUtils/signapk.jar