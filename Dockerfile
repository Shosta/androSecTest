# Build the AndroSecTest App on the golang latest image.
FROM golang:latest as go-builder

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



# Download the Required Hacking Tools
FROM ubuntu:20.04 as ubuntu-downloader

ENV HACKTOOLS_DIR=/home/Developpement/HackingTools

WORKDIR $HACKTOOLS_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    wget \
    unzip

# Install SignApk
RUN mkdir -p SignApkUtils && \
    wget --no-check-certificate --quiet -O ./SignApkUtils/sign.jar https://github.com/techexpertize/SignApk/blob/master/signapk.jar

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



# Pull Ubuntu LTS image.
FROM ubuntu:20.04

# Labels and Credits
LABEL \
    name="AndroSecTest on Ubuntu" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static."

ENV SRC_DIR=/root/go/src/github.com/Shosta/androSecTest
ENV HACKTOOLS_DIR=/home/Developpement/HackingTools
WORKDIR $SRC_DIR

# Expose default ADB port
EXPOSE 5555

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y --no-install-recommends \
    openjdk-11-jdk \
    usbutils \
    android-tools-adb \
    bash-completion

# Copy SignApk to Container
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/SignApkUtils/ $HACKTOOLS_DIR/SignApkUtils/

# Copy jadx and apktool
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/DecompilingAndroidAppUtils $HACKTOOLS_DIR/DecompilingAndroidAppUtils

# Copy Humpty-dumpty
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/humpty-dumpty-android-master $HACKTOOLS_DIR/humpty-dumpty-android-master

# Copy the built executable from the go-builder container and add it to this container.
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/androSecTest $SRC_DIR

# Copy the User Settings
RUN mkdir  $SRC_DIR/.res
COPY --from=go-builder /go/src/github.com/Shosta/androSecTest/res/ $SRC_DIR/.res/

# Start the server by default
# CMD ["adb", "-a", "-P", "5037", "server", "nodaemon"]