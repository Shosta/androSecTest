
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Ask me anything" src="https://img.shields.io/static/v1.svg?label=Ask%20me&message=anything&color=green"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Maintained" src="https://img.shields.io/static/v1.svg?label=Maintained?&message=Yes&color=Blue"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/Shosta/androSecTest.svg?style=social"></a>
<a href="https://github.com/Shosta/androSecTest/network"><img alt="GitHub forks" src="https://img.shields.io/github/forks/Shosta/androSecTest.svg?style=social"></a>
<a href="https://github.com/Shosta/androSecTest/blob/master/LICENSE.md"><img alt="GitHub license" src="https://img.shields.io/github/license/Shosta/androSecTest.svg?color=green&style=flat-square"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Pentest" src="https://img.shields.io/static/v1.svg?label=Pentest&message=Your%20App&color=green&logo=Android"></a>

 
 # Android-Static-Security-Audit

Here is a quick Cheat Sheet to test the security of an Android app that AndroSecTest is doing.

You can have a quick look at how the application is pentesting an Android app on Youtube : https://youtu.be/zzyTFjnwolo

## Easiest Way to Try It 

### Use the docker Container

1. Build the Docker Container that has all the dependencies and tools already installed.
    > `docker build .`

2. Connect your Android Device

    2.1. Be sure that the "adb server" is **not** running on the host machine as an android phone can only be connected to one adb server at a given time.
    
    2.2. USB connection is not working from host device to Container on MacOS, so it is only working on a Linux host for the time being.

3. Run the Docker Container
    > `docker run -it --privileged -v /dev/bus/usb:/dev/bus/usb "The Container ID"`

    3.1 `-it` is here so that we can have an iteractive session.

    3.2. `--privileged` is required to use a USB device.

    3.3. `-v /dev/bus/usb:/dev/bus/usb` defines a shared volume between the host machine and the Container in order to share the USB device (*the android phone*) information

⚠️ The results from the SAST is not persisted outside of the Docker Container at the moment.
I am planning to add a shared volume to persist it in the near future.

## The first part of the Security testing is to :
1. Get the application from the Store,
1. Pull it from the device,
1. Unpackaged it,
1. Look for some unsecure behavior,
1. Make it debuggable,
1. Repackage it and reinstall it on the device.

### 1. Get the application from your device, using the `adb` command
#### 1.1. List the applications' package names on your device :
> `adb shell pm list packages | grep “hint from the app you are looking for”`

#### 1.2. Get the path of the desired application on the device : 
> `adb shell pm path app.package.name.apk`

#### 1.3. Pull it from your device to your computer :
> `adb pull app.path`


#### 1.4. Change the file name from ".apk" to ".zip".
Unzip the file.
You now have access to the application's file system.

### 2. Look for interesting strings or files in the application 
#### 2.1. Locate interesting files or strings
Run the following commands at the root of the application file system.
* `find . -name "*key"`
* `find . -name "*cer*"`
* `find . -name "*pass*"'''`

If you find some files whose name contains 'key' try these commands :
* `hexdump ./path/to/.appkey  -vC`
* `more ./path/to/.appkey `


#### 2.2. Check the application signature.

Verify the signature : 
> `apksigner verify --verbose Application.apk`

or
> `jarsigner -verify -certs -verbose app.apk`

and

Move to the META.INF folder and check the signature with openssl : 
> `openssl pkcs7 -inform DER -in CERT.RSA -noout -print_certs -text`

Extract CERT.RSA from the package and display the certificate with keytool. 
> `keytool -printcert -file CERT.RSA `

You can then check the type of encryption used (hint, [SHA-1 is no more secure](https://shattered.io)).


### 2. Make the application debuggable and ready for penetration testing

Now that you have the apk file from the application you want, you must disassemble the app to make it debuggable.

#### 1. To disassemble the application, you can use the tool 'apktool'.

>`apktool d -o localAppFolder/ app.package.name.apk`

#### 2. Make the application debuggable and allow backup

In the `"<application”`, in the manifest file, add a `android:debuggable="true”` value to make the app debuggable.

In the `"<application”`, in the manifest file, add a `android:allowBackup="true”` value to allow backup from the app.

#### 3. Intercept and decrypt network requests

Edit the app Manifest to be able to intercept and decrypt encrypted requests from the app later on:
In the `"<application”` node, in the manifest file, add a `android:networkSecurityConfig="@xml/network_security_config"` value to be sure that the user added certificate are going to be trusted on a debug configuration.

Add a “network_security_config.xml” file in the “xml” folder with the following content or append the content to the existing file:
```xml
<!-- The "network_security_config.xml" -->

<?xml version="1.0" encoding="utf-8"?>
    <network-security-config>
        <debug-overrides>
            <trust-anchors>
                <!-- Trust user added CAs while debuggable only -->
                <certificates src="user" />
            </trust-anchors>
        </debug-overrides>
    ...
```

#### 4. Add the certificate to the device.
Download it from Burp, Charles, etc… and add it to your device following your preferred method (add push to the sdcard is the method I use).
You can use Bettercap to monitor the UDP traffic.


#### 5. Repackage and sign the app:
1. Repackage the app:
```
apk tool b -o app.package.name.apk localAppFolder/
```

2. Generate a signing key :
```
keytool -genkey -v -keystore resign.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000
```
3. then sign the app with it : 
```
jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore resign.keystore app.package.name.apk alias_name
```
or
```
apksigner sign -ks resign.keystore app.package.name.apk
```

#### 6. Install the app on the device : 

Run the following command to install the repackage app to the device: 
```
adb install app.package.name.apk
```

## The next steps of the security testing areto use some static test tool

I want to use some Man in the Middle attack while the user is using the application. It will jsute intercept all the requests/responses for later analysis.
I plan to use Bettercap or mitmproxy to do it.

We are going to use [MobSF](https://github.com/MobSF/Mobile-Security-Framework-MobSF) (MobSF stands for Mobile Security Framework) to test some part of the security of the app. 

As described in the Github page of the Project :
> Mobile Security Framework (MobSF) is an automated, all-in-one mobile application (Android/iOS/Windows) pen-testing framework capable of performing static, dynamic and malware analysis. It can be used for effective and fast security analysis of Android, iOS and Windows mobile applications and support both binaries (APK, IPA & APPX ) and zipped source code. MobSF can do dynamic application testing at runtime for Android apps and has Web API fuzzing capabilities powered by CapFuzz, a Web API specific security scanner. MobSF is designed to make your CI/CD or DevSecOps pipeline integration seamless.

I personnaly use the Docker container to use MobSF for Android security audit.
So you could just launch that command `docker run -it -p 8000:8000 -v <your_local_dir>:/root/.MobSF opensecurity/mobile-security-framework-mobsf:latest`

MobSF is going to automate a lot of the process of static security analysis and deliver a report that will make it easier to start the dynamic security audit.

