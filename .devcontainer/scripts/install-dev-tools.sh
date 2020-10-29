#!/bin/bash

# Get Go Dependancies
go get ./...

# Copy signapk.jar to the proper location and make it executable
cp -R ./res/SignApk /home/Developpement/HackingTools/SignApkUtils/
chmod +x /home/Developpement/HackingTools/SignApkUtils/signapk.jar

 # Copy the required Settings.
cp res/settings/usersettings.json .res/settings
cp -R res/watermark .res/watermark 

# Add GOPATH to the PATH in the ~/.zshrc file as I come from bash.
sed -i '1i\\'\"export PATH=$HOME/bin:/usr/local/bin:/root/go/bin:$PATH\" ~/.zshrc
