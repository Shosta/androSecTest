# Download the Required Hacking Tools
FROM ubuntu:20.04 as ubuntu-downloader
LABEL ubuntu-downloader=true

ENV HACKTOOLS_DIR=/home/Developpement/HackingTools

WORKDIR $HACKTOOLS_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    wget \
    unzip && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get purge   --auto-remove && \
    apt-get clean

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



# Build the AndroSecTest App on the golang latest image.
FROM golang:latest as go-builder
LABEL go-builder=true

# Environmentn variables
ENV SRC_DIR=/go/src/github.com/Shosta/androSecTest
ENV GIT_SSL_NO_VERIFY=1

# Set the Current Working Directory inside the container
WORKDIR $SRC_DIR

# Copy the source from the current directory to the Working Directory inside the container
COPY . $SRC_DIR

# Dowload the Go Dependancies
RUN go get $SRC_DIR/...

# Build the Go app for a Linux target
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o androSecTest .



# Pull Ubuntu LTS image.
FROM ubuntu:20.04

# Labels and Credits
LABEL \
    name="AndroSecTest on Ubuntu" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static security assessment."

ENV SRC_DIR=/root/go/src/github.com/Shosta/androSecTest
ENV HACKTOOLS_DIR=/home/Developpement/HackingTools
WORKDIR $SRC_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apt-get update -y && apt-get install -y --no-install-recommends \
    openjdk-8-jdk \
    usbutils \
    unzip \
    android-tools-adb \
    bash-completion && \
    rm -rf /var/lib/apt/lists/* && \
    apt-get purge --auto-remove && \
    apt-get clean

# Copy jadx and apktool and signapk
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/ $HACKTOOLS_DIR/
RUN chmod +x $HACKTOOLS_DIR/DecompilingAndroidAppUtils/apktool/apktool.jar && chmod +x $HACKTOOLS_DIR/SignApkUtils/signapk.jar

# Copy Humpty-dumpty
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/humpty-dumpty-android-master $HACKTOOLS_DIR/humpty-dumpty-android-master

# Copy the built executable from the go-builder container and add it to this container.
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/androSecTest $SRC_DIR

# Copy the User Settings
RUN mkdir  $SRC_DIR/.res
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/res/ $SRC_DIR/.res/

# Copy SignApk to Proper Location in Container
RUN mv $SRC_DIR/.res/SignApk $HACKTOOLS_DIR/SignApkUtils/ && \
    chmod +x $HACKTOOLS_DIR/SignApkUtils/signapk.jar
