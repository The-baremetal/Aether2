#!/bin/bash
set -e

# Remove symlink
if [ -L /usr/local/bin/aether ]; then
    rm -f /usr/local/bin/aether
fi

# Remove profile script that sets PATH and AETHERROOT
if [ -f /etc/profile.d/aether.sh ]; then
    rm -f /etc/profile.d/aether.sh
fi

# Clean PATH and AETHERROOT from common shell profiles
SHELLS=("$HOME/.bash_profile" "$HOME/.bashrc" "$HOME/.zshrc")
for shell in "${SHELLS[@]}"; do
    if [ -f "$shell" ]; then
        sed -i '/\/usr\/local\/aether\/bin/d' "$shell"
        sed -i '/AETHERROOT/d' "$shell"
    fi
done

echo "Aether has been fully uninstalled!"
echo "PATH and AETHERROOT cleaned from user profiles and system environment."
