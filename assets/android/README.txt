Aether Programming Language for Android
=======================================

This is a cross-compiled binary for Android systems.

Installation on Android with Termux:
1. Install Termux from F-Droid or Google Play
2. Extract the archive: tar -xvf aether-android_arm64.tar
3. Copy the binary: cp usr/local/aether/bin/aether $PREFIX/bin/
4. Copy packages: cp -r usr/local/aether/packages $PREFIX/share/aether
5. Run: aether --help

Note: This binary requires Android API level 21 or higher.

For more information, visit: https://github.com/The-baremetal/Aether2
