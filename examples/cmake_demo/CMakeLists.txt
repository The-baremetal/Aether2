# 🍕 Aether CMake Demo
# Extraordinary build tool integration example

cmake_minimum_required(VERSION 3.16)
project(AetherDemo VERSION 1.0.0 LANGUAGES CXX)

# Include Aether CMake configuration
include(cmake/AetherConfig.cmake)

# Set C++ standard
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Enable testing
aether_enable_testing()

# Create math library
aether_add_library(mathlib
    SHARED
    SOURCES
        src/mathlib.ae
    VERSION "1.0.0"
    DESCRIPTION "Aether Math Library - Fast mathematical operations"
    URL "https://github.com/aether-lang/mathlib"
    REQUIRES "c"
    PROVIDES "math"
)

# Create crypto library
aether_add_library(cryptolib
    BOTH
    SOURCES
        src/cryptolib.ae
    VERSION "1.0.0"
    DESCRIPTION "Aether Crypto Library - Cryptographic operations"
    URL "https://github.com/aether-lang/cryptolib"
    REQUIRES "c;openssl"
    PROVIDES "crypto"
)

# Create main executable
aether_add_executable(demo
    SOURCES
        src/main.ae
    DEPENDENCIES
        mathlib
        cryptolib
    OUTPUT "aether-demo"
)

# Link libraries to executable
aether_target_link_libraries(demo
    PRIVATE
        mathlib
        cryptolib
)

# Set compiler options
aether_target_compile_options(demo
    PRIVATE
        --O3
        --verbose
)

# Create package
aether_package(AetherDemo
    VERSION "1.0.0"
    DESCRIPTION "Aether Demo Application"
    URL "https://github.com/aether-lang/demo"
    DEPENDENCIES
        mathlib
        cryptolib
)

# Install rules
install(TARGETS demo
    RUNTIME DESTINATION bin
)

install(DIRECTORY include/
    DESTINATION include
    FILES_MATCHING PATTERN "*.h"
)

# Create package configuration
include(CMakePackageConfigHelpers)
write_basic_package_version_file(
    "${CMAKE_CURRENT_BINARY_DIR}/AetherDemoConfigVersion.cmake"
    VERSION ${PROJECT_VERSION}
    COMPATIBILITY AnyNewerVersion
)

install(FILES
    "${CMAKE_CURRENT_BINARY_DIR}/AetherDemoConfigVersion.cmake"
    DESTINATION lib/cmake/AetherDemo
)

# Configure and install config file
configure_package_config_file(
    "${CMAKE_CURRENT_SOURCE_DIR}/cmake/AetherDemoConfig.cmake.in"
    "${CMAKE_CURRENT_BINARY_DIR}/AetherDemoConfig.cmake"
    INSTALL_DESTINATION lib/cmake/AetherDemo
)

install(FILES
    "${CMAKE_CURRENT_BINARY_DIR}/AetherDemoConfig.cmake"
    DESTINATION lib/cmake/AetherDemo
)

# Export targets
install(EXPORT AetherDemoTargets
    FILE AetherDemoTargets.cmake
    NAMESPACE AetherDemo::
    DESTINATION lib/cmake/AetherDemo
)

# Create namespace targets
export(EXPORT AetherDemoTargets
    FILE "${CMAKE_CURRENT_BINARY_DIR}/AetherDemoTargets.cmake"
    NAMESPACE AetherDemo::
)

# Add custom targets for development
add_custom_target(format
    COMMAND ${CMAKE_AETHER_COMPILER} format src/
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Formatting Aether source files"
)

add_custom_target(lint
    COMMAND ${CMAKE_AETHER_COMPILER} lint src/
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Linting Aether source files"
)

add_custom_target(test
    COMMAND ${CMAKE_AETHER_COMPILER} test
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Running Aether tests"
    DEPENDS demo
)

# Add documentation target
add_custom_target(docs
    COMMAND ${CMAKE_AETHER_COMPILER} docs
    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    COMMENT "Generating Aether documentation"
)

# Add clean target
add_custom_target(clean-all
    COMMAND ${CMAKE_COMMAND} -E remove_directory ${CMAKE_BINARY_DIR}
    COMMENT "Cleaning all build artifacts"
)

# Set up development environment
if(CMAKE_BUILD_TYPE STREQUAL "Debug")
    aether_target_compile_options(demo
        PRIVATE
            --debug-info
            --debug-symbols
    )
endif()

# Cross-compilation support
if(DEFINED CMAKE_AETHER_TARGET_OS)
    set(CMAKE_AETHER_TARGET_OS ${CMAKE_AETHER_TARGET_OS})
endif()

if(DEFINED CMAKE_AETHER_TARGET_ARCH)
    set(CMAKE_AETHER_TARGET_ARCH ${CMAKE_AETHER_TARGET_ARCH})
endif()

# Print configuration summary
message(STATUS "Aether Demo Configuration:")
message(STATUS "  Version: ${PROJECT_VERSION}")
message(STATUS "  Build Type: ${CMAKE_BUILD_TYPE}")
message(STATUS "  Target OS: ${CMAKE_AETHER_TARGET_OS}")
message(STATUS "  Target Arch: ${CMAKE_AETHER_TARGET_ARCH}")
message(STATUS "  Compiler: ${CMAKE_AETHER_COMPILER}")
message(STATUS "  Linker: ${CMAKE_AETHER_LINKER}")
message(STATUS "  Libraries: mathlib, cryptolib")
message(STATUS "  Executable: demo") 