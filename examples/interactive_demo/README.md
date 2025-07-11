# Interactive Aether Demo 🍕

This demo shows the new interactive features of Aether!

## Interactive Init

Run `aether init` and answer the questions:

```bash
🍕 Initializing Aether project 'my-awesome-app'...

🍕 Where do you want your code to be? (src): lib
🍕 Where do you want build outputs? (bin): dist
🍕 Who are you? (optional): Pizza Lover
🍕 What does your project do? (optional): A super customizable app
🍕 Create a main.aeth file? (y/n): y

🍕 Project 'my-awesome-app' initialized successfully!
🍕 Source directory: lib
🍕 Output directory: dist
🍕 Run 'aether build' to compile your project!
```

## Creating Packages

Run `aether package` to create new packages:

```bash
🍕 What's your package name? math-utils
🍕 What type of package? (library/app): library
🍕 What does this package do? Mathematical utilities for Aether
🍕 Who created this package? (optional): Math Wizard
🍕 Package version? (0.1.0): 1.0.0

🍕 Package 'math-utils' created successfully!
🍕 Location: packages/math-utils
🍕 Added to project dependencies
🍕 Run 'aether build' to test your package!
```

## Generated Structure

After running both commands, you'll have:

```
my-awesome-app/
├── aether.toml          # Custom configuration
├── lib/                 # Custom source directory
│   └── main.aeth       # Your main file
├── dist/               # Custom output directory
├── packages/           # Local packages
│   └── math-utils/     # Your new package
│       ├── aether.toml
│       ├── src/
│       ├── examples/
│       └── README.md
└── .gitignore
```

## Custom Configuration

The `aether.toml` file will be automatically configured:

```toml
[project]
name = "my-awesome-app"
version = "0.1.0"
author = "Pizza Lover"
description = "A super customizable app"

[build]
source_directories = ["lib"]
output_directory = "dist"
target = "native"
optimization = "2"
linker = "mold"
create_library = false
library_type = "shared"

[dependencies]
math-utils = "packages/math-utils"
```

**SUPER CUSTOMIZABLE AND USER-FRIENDLY!** 🚀✨ 