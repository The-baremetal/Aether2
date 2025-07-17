#!/bin/bash
set -e

# Remove symlink
if [ -L /usr/local/bin/aether ]; then
    rm -f /usr/local/bin/aether
fi

# Remove profile script
if [ -f /etc/profile.d/aether.sh ]; then
    rm -f /etc/profile.d/aether.sh
fi
