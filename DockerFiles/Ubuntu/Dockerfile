# Build the AndroSecTest App on the golang latest image.
FROM golang:latest as go-builder

# Environmentn variables
ENV SRC_DIR=/go/src/github.com/shosta/androSecTest
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

ENV GIT_SSL_NO_VERIFY=1

WORKDIR /home

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    wget \
    unzip


# Install SignApk
RUN mkdir -p /home/Developpement/HackingTools/SignApkUtils && \
    wget --no-check-certificate --quiet -O /home/Developpement/HackingTools/SignApkUtils/sign.jar https://github.com/techexpertize/SignApk/blob/master/signapk.jar

# Install jadx
RUN wget --no-check-certificate --quiet https://github.com/skylot/jadx/releases/download/v1.1.0/jadx-1.1.0.zip && \
    mkdir -p /home/Developpement/HackingTools/DecompilingAndroidAppUtils/bin/jadx
RUN unzip jadx-1.1.0.zip -d /home/Developpement/HackingTools/DecompilingAndroidAppUtils/bin/jadx && rm jadx-1.1.0.zip

# Install Humpty-dumpty
RUN mkdir -p /home/Developpement/HackingTools/humpty-dumpty-android-master && \
    wget --no-check-certificate --quiet -O /home/Developpement/HackingTools/humpty-dumpty-android-master/humpty.sh https://github.com/Pixplicity/humpty-dumpty-android/blob/master/humpty.sh



# Pull Ubuntu LTS image.
FROM ubuntu:20.04

# Labels and Credits
LABEL \
    name="AndroSecTest on Ubuntu" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static."

WORKDIR /go/src/github.com/shosta/androSecTest/

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    apktool \
    android-tools-adb \
    bash-completion

# Copy SignApk to Container
COPY --from=ubuntu-downloader /home/Developpement/HackingTools/SignApkUtils/ /home/Developpement/HackingTools/SignApkUtils/

# Copy jadx
COPY --from=ubuntu-downloader /home/Developpement/HackingTools/DecompilingAndroidAppUtils /home/Developpement/HackingTools/DecompilingAndroidAppUtils

# Copy Humpty-dumpty
COPY --from=ubuntu-downloader /home/Developpement/HackingTools/humpty-dumpty-android-master /home/Developpement/HackingTools/humpty-dumpty-android-master

# Copy the built executable from the go-builder container and add it to this container.
COPY --from=go-builder /go/src/github.com/shosta/androSecTest/androSecTest /go/src/github.com/shosta/androSecTest

# You can pass args to the docker image if you launch the app like this : `docker run --rm androsectest-ubuntu:latest --appname yourApp`
# ENTRYPOINT [ "/go/src/github.com/shosta/androSecTest/androSecTest" ]
# CMD ["--help"]
CMD ["/bin/bash"]