Aether Programming Language for OpenBSD
=======================================

Installation:
1. Extract the archive: tar -xvf aether-openbsd_amd64.tar
2. Copy files to system: doas cp -r usr/local/* /usr/local/
3. Run: aether --help

The binary will be installed to: /usr/local/aether/bin/aether
A symlink will be created at: /usr/local/bin/aether

Uninstallation:
Run: doas rm -rf /usr/local/aether
Run: doas rm -f /usr/local/bin/aether

For more information, visit: https://github.com/The-baremetal/Aether2
