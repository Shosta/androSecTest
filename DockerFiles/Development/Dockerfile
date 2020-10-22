# The Container that allows to develop the Go app, androSecTest, within a container, without any install.
# Pull Ubuntu LTS image.
FROM ubuntu:20.04

# Labels and Credits
LABEL \
    name="AndroSecTest as Dev Env" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static."

WORKDIR /go/src/github.com/shosta/androSecTest/

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    unzip \
    wget \
    signapk \
    apktool \
    android-tools-adb \
    golang-go

# Install jadx
RUN wget --quiet https://github.com/skylot/jadx/releases/download/v1.1.0/jadx-1.1.0.zip
RUN unzip jadx-1.1.0.zip -d /usr/local/bin && rm jadx-1.1.0.zip
