# 🍕 Aether CMake Support - Extraordinary Build Tool Integration

Aether provides **extraordinary CMake support** that makes it the ultimate language for modern C++ projects! This demo shows how Aether integrates seamlessly with CMake for professional-grade builds.

## 🚀 Features

### **Complete CMake Integration**
- **Native CMake functions** for Aether projects
- **Automatic dependency management**
- **Cross-platform support**
- **Professional build system**

### **Library Management**
- **Shared libraries** (.so/.dll/.dylib)
- **Static libraries** (.a/.lib)
- **Both library types** simultaneously
- **pkg-config integration**

### **Development Tools**
- **Formatting** with `make format`
- **Linting** with `make lint`
- **Testing** with `make test`
- **Documentation** with `make docs`

## 📁 Project Structure

```
cmake_demo/
├── CMakeLists.txt          # Main CMake configuration
├── cmake/
│   └── AetherConfig.cmake  # Aether CMake integration
├── src/
│   ├── main.ae            # Main executable
│   ├── mathlib.ae         # Math library
│   └── cryptolib.ae       # Crypto library
└── include/               # Header files (if needed)
```

## 🛠️ Building with CMake

### **Basic Build**
```bash
mkdir build
cd build
cmake ..
make
```

### **Debug Build**
```bash
mkdir build-debug
cd build-debug
cmake -DCMAKE_BUILD_TYPE=Debug ..
make
```

### **Release Build**
```bash
mkdir build-release
cd build-release
cmake -DCMAKE_BUILD_TYPE=Release ..
make
```

### **Cross-Compilation**
```bash
mkdir build-arm
cd build-arm
cmake -DCMAKE_AETHER_TARGET_OS=linux -DCMAKE_AETHER_TARGET_ARCH=arm64 ..
make
```

## 📚 CMake Functions

### **aether_add_executable()**
```cmake
aether_add_executable(demo
    SOURCES
        src/main.ae
    DEPENDENCIES
        mathlib
        cryptolib
    OUTPUT "aether-demo"
)
```

### **aether_add_library()**
```cmake
aether_add_library(mathlib
    SHARED                    # or STATIC, BOTH
    SOURCES
        src/mathlib.ae
    VERSION "1.0.0"
    DESCRIPTION "Math Library"
    URL "https://github.com/aether-lang/mathlib"
    REQUIRES "c"
    PROVIDES "math"
)
```

### **aether_target_link_libraries()**
```cmake
aether_target_link_libraries(demo
    PRIVATE
        mathlib
        cryptolib
)
```

### **aether_target_compile_options()**
```cmake
aether_target_compile_options(demo
    PRIVATE
        --O3
        --verbose
)
```

### **aether_enable_testing()**
```cmake
aether_enable_testing()
# Automatically finds *_test.ae files
```

### **aether_package()**
```cmake
aether_package(AetherDemo
    VERSION "1.0.0"
    DESCRIPTION "Aether Demo"
    URL "https://github.com/aether-lang/demo"
    DEPENDENCIES
        mathlib
        cryptolib
)
```

## 🎯 Development Workflow

### **Format Code**
```bash
make format
```

### **Lint Code**
```bash
make lint
```

### **Run Tests**
```bash
make test
```

### **Generate Documentation**
```bash
make docs
```

### **Clean Everything**
```bash
make clean-all
```

## 📦 Package Management

### **Install Package**
```bash
make install
```

### **Create Package**
```bash
cpack
```

### **Find Aether Library**
```cmake
aether_find_library(mathlib REQUIRED)
target_link_libraries(myapp ${mathlib_LIBRARIES})
target_include_directories(myapp ${mathlib_INCLUDE_DIRS})
```

## 🔧 Advanced Features

### **Custom Targets**
```cmake
add_custom_target(analyze
    COMMAND ${CMAKE_AETHER_COMPILER} analyze src/
    COMMENT "Analyzing Aether code"
)
```

### **Conditional Compilation**
```cmake
if(CMAKE_BUILD_TYPE STREQUAL "Debug")
    aether_target_compile_options(demo
        PRIVATE
            --debug-info
            --debug-symbols
    )
endif()
```

### **Cross-Platform Support**
```cmake
if(WIN32)
    set(CMAKE_AETHER_LINKER "lld")
else()
    set(CMAKE_AETHER_LINKER "mold")
endif()
```

## 🏆 Benefits

### **Professional Integration**
- **CMake ecosystem** compatibility
- **IDE support** (CLion, VSCode, etc.)
- **CI/CD integration**
- **Package managers** support

### **Performance**
- **Fast linking** with mold
- **LLVM optimization**
- **Incremental builds**
- **Parallel compilation**

### **Simplicity**
- **Declarative syntax**
- **Automatic dependency resolution**
- **Cross-platform builds**
- **Standard CMake patterns**

## 🚀 The Future

Aether's CMake support enables:
- **Enterprise adoption**
- **Large-scale projects**
- **Professional toolchains**
- **Modern C++ workflows**

**Aether = Go's simplicity + Rust's performance + CMake's ecosystem!** 🍕✨

## 📖 Examples

### **Simple Executable**
```cmake
aether_add_executable(hello
    SOURCES src/main.ae
)
```

### **Library with Dependencies**
```cmake
aether_add_library(mylib
    SHARED
    SOURCES src/lib.ae
    REQUIRES "c;openssl"
)
```

### **Complex Project**
```cmake
# Multiple libraries
aether_add_library(core SHARED SOURCES src/core.ae)
aether_add_library(utils SHARED SOURCES src/utils.ae)

# Main executable
aether_add_executable(app
    SOURCES src/main.ae
    DEPENDENCIES core utils
)

# Link libraries
aether_target_link_libraries(app PRIVATE core utils)
```

**Aether makes CMake extraordinary!** 🎯🚀 