<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Ask me anything" src="https://img.shields.io/static/v1.svg?label=Ask%20me&message=anything&color=green"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Maintained" src="https://img.shields.io/static/v1.svg?label=Maintained?&message=Yes&color=Blue"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/Shosta/androSecTest.svg?style=social"></a>
<a href="https://github.com/Shosta/androSecTest/network"><img alt="GitHub forks" src="https://img.shields.io/github/forks/Shosta/androSecTest.svg?style=social"></a>
<a href="https://github.com/Shosta/androSecTest/blob/master/LICENSE.md"><img alt="GitHub license" src="https://img.shields.io/github/license/Shosta/androSecTest.svg?color=green&style=flat-square"></a>
<a href="https://github.com/Shosta/androSecTest/stargazers"><img alt="Pentest" src="https://img.shields.io/static/v1.svg?label=Pentest&message=Your%20App&color=green&logo=Android"></a>

 
 # Android-Static-Security-Audit

AndroSecTest is a [Go](https://golang.org/) application that helps you during your Android application pentesting.

It removes all the hassle of extracting an app, disassembling it and looking for low hanging fruits.

You can have a quick look at how the application is pentesting an Android app on Youtube : 

[![Watch it here](https://img.youtube.com/vi/zzyTFjnwolo/hqdefault.jpg)](https://youtu.be/zzyTFjnwolo)


## How to Use It ?  🤔

Run it in a Docker Container (*nothing to install or configure. I did everything for you* 😉) 

 `docker run -it --privileged -v /dev/bus/usb:/dev/bus/usb -v ./android/security:/home/androSecTest-Results <container_id>`

1. Clone the AndroSecTest repo
    > `git clone git@github.com:Shosta/androSecTest.git`

1. Build the Docker Container that has all the dependencies and tools already installed.
    > `docker build .`

1. Connect your Android Device

    3.1. Be sure that the "adb server" is **NOT** running on the host machine as an android phone can only be connected to one adb server at a given time.
    
    3.2. USB connection is not working from host device to Container on MacOS, so it is only working on a Linux host for the time being.

1. On Linux 🐧 - Run the Docker Container
    > `docker run -it --privileged -v /dev/bus/usb:/dev/bus/usb "The Container ID"`

    4.1 `-it` is here so that we can run an iteractive session.

    4.2. `--privileged` is required to use a USB device.

    4.3. `-v /dev/bus/usb:/dev/bus/usb` defines a shared volume between the host machine and the Container in order to share the USB device (*the android phone*) information

1. On Mac 🍏 - Run the Docker Container
    > `docker run -it --privileged "The Container ID"`

    5.1 `-it` is here so that we can run an iteractive session.

    5.2. `--privileged` is required to use a USB device.

    5.3. `-v <the folder to persist the Pentest Results>:/home/androSecTest-Results` defines a shared volume between the host machine and the Container in order to share the Pentest results.

Well done, you are good to go.
Get into your Docker Container, fire the androSecTest, connect your device and start assessing the security of your Android application.

---

## What is this app is doing ? 🤔
(*in case you want to do a step by step, to understand the method it is using*)

I made this app, in order to automate a part of the process of pentesting and android app. Every time, I had to exfiltrate the app from the device and disassemble the code and then repackage the app with some instrumentation.

androSecTest is doing that for you automatically (*and much more*).

### Here are the steps, it is doing :

1. Get the application you want to pentest from the Store,
1. Pull it from the device,
1. Unpackaged it,
1. Look for some unsecure behavior,
1. Make it debuggable,
1. Repackage it and reinstall it on the device.

🎉 You have a debuggable application on your device, with Backup available.
🎉 You have the source code of the application on your file system, ready for analysis. 

---

## How to Do that Manually in your Terminal ? 🤔

It is a good idea to do that once, in order to understand what an Android app Security Assessment looks like.

### 1. Pull the application from your device, using the `adb` command
#### 1.1. List the applications' package names on your device :
> `adb shell pm list packages | grep “hint from the app you are looking for”`

#### 1.2. Get the path of the desired application on the device : 
> `adb shell pm path app.package.name.apk`

#### 1.3. Pull it from your device to your computer :
> `adb pull app.path`

#### 1.4. Change the file name from ".apk" to ".zip".
> `mv <the_app_your_pentesting.apk> <the_app_your_pentesting.zip>`

#### 1.5. Unzip the file.
> `unzip <the_app_your_pentesting.zip>`

🎉 You have access to the application's filesystem.

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
> `apksigner verify --verbose <the_app_your_pentesting.apk>`

or
> `jarsigner -verify -certs -verbose <the_app_your_pentesting.apk>`

and

Move to the META.INF folder and check the signature with openssl : 
> `openssl pkcs7 -inform DER -in CERT.RSA -noout -print_certs -text`

Extract CERT.RSA from the package and display the certificate with keytool. 
> `keytool -printcert -file CERT.RSA `

You can then check the type of encryption used (hint, [SHA-1 is no more secure](https://shattered.io)).


### 3. Make the application debuggable and ready for penetration testing

Now that you have the apk file from the application you want, you must disassemble the app to make it debuggable.

#### 3.1. To disassemble the application, you can use the tool 'apktool'.

> `apktool d -o localAppFolder/ app.package.name.apk`

#### 3.2. Make the application debuggable and allow backup

In the `"<application”`, in the manifest file, add a `android:debuggable="true”` value to make the app debuggable. 🐛

In the `"<application”`, in the manifest file, add a `android:allowBackup="true”` value to allow backup from the app. 💾

#### 4. Intercept and decrypt network requests

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

#### 5. Add the certificate to the device.
Download it from [Burp](https://portswigger.net/burp), [Charles](https://www.charlesproxy.com/), etc… and add it to your device following your preferred method (*add `push` to the sdcard is the method I use*).
You can use Bettercap to monitor the UDP traffic.


#### 6. Repackage and sign the app:
1. Repackage the app:
> `apk tool b -o app.package.name.apk localAppFolder/`

2. Generate a signing key :
> `keytool -genkey -v -keystore resign.keystore -alias alias_name -keyalg RSA -keysize 2048 -validity 10000`

3. then sign the app with it : 
> `jarsigner -verbose -sigalg SHA1withRSA -digestalg SHA1 -keystore resign.keystore app.package.name.apk alias_name`
or
> `apksigner sign -ks resign.keystore app.package.name.apk`

#### 6. Install the app on the device : 

Run the following command to install the repackage app to the device: 
> `adb install app.package.name.apk`

## Next Steps and Future Features 🚀

I want to use some Man in the Middle attack while the user is using the application. It will jsute intercept all the requests/responses for later analysis.
I plan to use Bettercap or mitmproxy to do it.

We are going to use [MobSF](https://github.com/MobSF/Mobile-Security-Framework-MobSF) (MobSF stands for Mobile Security Framework) to test some part of the security of the app. 

As described in the Github page of the Project :
> Mobile Security Framework (MobSF) is an automated, all-in-one mobile application (Android/iOS/Windows) pen-testing framework capable of performing static, dynamic and malware analysis. It can be used for effective and fast security analysis of Android, iOS and Windows mobile applications and support both binaries (APK, IPA & APPX ) and zipped source code. MobSF can do dynamic application testing at runtime for Android apps and has Web API fuzzing capabilities powered by CapFuzz, a Web API specific security scanner. MobSF is designed to make your CI/CD or DevSecOps pipeline integration seamless.

I personnaly use the Docker container to use MobSF for Android security audit.
So you could just launch that command `docker run -it -p 8000:8000 -v <your_local_dir>:/root/.MobSF opensecurity/mobile-security-framework-mobsf:latest`

MobSF is going to automate a lot of the process of static security analysis and deliver a report that will make it easier to start the dynamic security audit.

