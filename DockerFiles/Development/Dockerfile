# The Container that allows to develop the Go app, androSecTest, within a container, without any install.

FROM ubuntu:20.04 as ubuntu-downloader

ENV HACKTOOLS_DIR=/home/Developpement/HackingTools

WORKDIR $HACKTOOLS_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y  --no-install-recommends \
    wget \
    unzip

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




# Pull Ubuntu LTS image.
FROM ubuntu:20.04

# Labels and Credits
LABEL \
    name="AndroSecTest as Dev Env" \
    author="Rémi Lavedrine <remi@github.com>" \
    maintainer="Rémi Lavedrine <remi@github.com>" \
    description="Android Security Test (AndroSecTest) is an automated, all-in-one mobile application (Android) security assessment framework capable of performing static."

ENV SRC_DIR=/root/go/src/github.com/Shosta/androSecTest
ENV HACKTOOLS_DIR=/home/Developpement/HackingTools
WORKDIR $SRC_DIR

# Install adb tools, unzip, wget, signapk and apktool
RUN apt update -y && apt install -y --no-install-recommends \
    openjdk-8-jdk \
    usbutils \
    unzip \
    wget \
    android-tools-adb \
    bash-completion \
    golang-go \
    gcc

# Copy jadx and apktool
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/DecompilingAndroidAppUtils $HACKTOOLS_DIR/DecompilingAndroidAppUtils
RUN chmod +x $HACKTOOLS_DIR/DecompilingAndroidAppUtils/apktool/apktool.jar

# Copy Humpty-dumpty
COPY --from=ubuntu-downloader $HACKTOOLS_DIR/humpty-dumpty-android-master $HACKTOOLS_DIR/humpty-dumpty-android-master

RUN apt install -y --no-install-recommends git zsh && chsh -s /usr/bin/zsh

RUN sh -c "$(wget -O- https://github.com/deluan/zsh-in-docker/releases/download/v1.1.1/zsh-in-docker.sh)" -- \
    -t https://github.com/denysdovhan/spaceship-prompt \
    -a 'SPACESHIP_PROMPT_ADD_NEWLINE="false"' \
    -a 'SPACESHIP_PROMPT_SEPARATE_LINE="false"' \
    -p git \
    -p sudo \
    -p https://github.com/zsh-users/zsh-autosuggestions \
    -p https://github.com/zsh-users/zsh-completions

RUN mkdir  $SRC_DIR/.res
COPY ./res/ $SRC_DIR/.res/