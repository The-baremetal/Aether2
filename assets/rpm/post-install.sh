#!/bin/bash
set -e

# Add /usr/local/bin to PATH if not already present
if [ -f /etc/profile.d/aether.sh ]; then
    rm -f /etc/profile.d/aether.sh
fi

cat > /etc/profile.d/aether.sh <<EOF
# Add Aether to PATH
export PATH="/usr/local/aether/bin:\$PATH"

# Set AETHERROOT environment variable
export AETHERROOT="/usr/local/aether"
EOF

# Create symlink if it doesn't exist
if [ ! -L /usr/local/bin/aether ]; then
    ln -sf /usr/local/aether/bin/aether /usr/local/bin/aether
fi

# Source it for current shell (if bash, zsh, etc.)
if [ -f /etc/profile.d/aether.sh ]; then
    source /etc/profile.d/aether.sh
fi

echo "Aether has been installed successfully!"
echo "Run 'aether --help' to get started."
echo "AETHERROOT is now set to /usr/local/aether"
