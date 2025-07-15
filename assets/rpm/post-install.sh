#!/bin/bash
set -e

# Add /usr/local/bin to PATH if not already present
if [ -f /etc/profile.d/aether.sh ]; then
    rm -f /etc/profile.d/aether.sh
fi

cat > /etc/profile.d/aether.sh <<EOF
# Add Aether to PATH
export PATH="/usr/local/aether/bin:\$PATH"
EOF

# Create symlink if it doesn't exist
if [ ! -L /usr/local/bin/aether ]; then
    ln -sf /usr/local/aether/bin/aether /usr/local/bin/aether
fi

echo "Aether has been installed successfully!"
echo "Run 'aether --help' to get started."
