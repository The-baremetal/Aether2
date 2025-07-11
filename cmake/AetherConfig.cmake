# 🍕 Aether CMake Configuration
# Extraordinary build tool integration for the Aether language

cmake_minimum_required(VERSION 3.16)

# Find Aether compiler
find_program(AETHER_COMPILER
    NAMES aether aether2
    PATHS
        ${CMAKE_CURRENT_SOURCE_DIR}/bin
        ${CMAKE_CURRENT_SOURCE_DIR}/cmd/aether2
        $ENV{HOME}/.local/bin
        /usr/local/bin
        /usr/bin
    DOC "Aether compiler executable"
)

if(NOT AETHER_COMPILER)
    message(FATAL_ERROR "Aether compiler not found! Please install Aether or set AETHER_COMPILER manually.")
endif()

# Get Aether version
execute_process(
    COMMAND ${AETHER_COMPILER} --version
    OUTPUT_VARIABLE AETHER_VERSION_OUTPUT
    ERROR_VARIABLE AETHER_VERSION_ERROR
    OUTPUT_STRIP_TRAILING_WHITESPACE
)

if(AETHER_VERSION_OUTPUT)
    string(REGEX REPLACE "Aether Compiler v([0-9.]+)" "\\1" AETHER_VERSION "${AETHER_VERSION_OUTPUT}")
else()
    set(AETHER_VERSION "unknown")
endif()

message(STATUS "Found Aether compiler: ${AETHER_COMPILER} (version ${AETHER_VERSION})")

# Aether language configuration
set(CMAKE_AETHER_COMPILER ${AETHER_COMPILER})
set(CMAKE_AETHER_COMPILER_ID "Aether")
set(CMAKE_AETHER_COMPILER_VERSION ${AETHER_VERSION})

# Aether file extensions
set(CMAKE_AETHER_SOURCE_FILE_EXTENSIONS ae)

# Aether compiler flags
set(CMAKE_AETHER_FLAGS "")
set(CMAKE_AETHER_FLAGS_DEBUG "-O0 -g")
set(CMAKE_AETHER_FLAGS_RELEASE "-O3")
set(CMAKE_AETHER_FLAGS_RELWITHDEBINFO "-O2 -g")
set(CMAKE_AETHER_FLAGS_MINSIZEREL "-Os")

# Aether library types
set(AETHER_LIBRARY_TYPES "shared;static;both")

# Aether target OS and architecture
if(NOT CMAKE_AETHER_TARGET_OS)
    set(CMAKE_AETHER_TARGET_OS ${CMAKE_SYSTEM_NAME})
endif()

if(NOT CMAKE_AETHER_TARGET_ARCH)
    set(CMAKE_AETHER_TARGET_ARCH ${CMAKE_SYSTEM_PROCESSOR})
endif()

# Aether optimization levels
set(AETHER_OPTIMIZATION_LEVELS "0;1;2;3;s;z")

# Aether linker options
set(CMAKE_AETHER_LINKER "mold")
if(WIN32)
    set(CMAKE_AETHER_LINKER "lld")
endif()

# Aether build types
set(CMAKE_AETHER_BUILD_TYPES "Debug;Release;RelWithDebInfo;MinSizeRel")

# Function to add Aether executable
function(aether_add_executable TARGET)
    set(options)
    set(oneValueArgs OUTPUT)
    set(multiValueArgs SOURCES DEPENDENCIES)
    
    cmake_parse_arguments(ARGS "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})
    
    if(NOT ARGS_OUTPUT)
        set(ARGS_OUTPUT ${TARGET})
    endif()
    
    # Create custom target for Aether compilation
    add_custom_target(${TARGET} ALL
        COMMAND ${CMAKE_AETHER_COMPILER} build
            --output ${ARGS_OUTPUT}
            --target-os ${CMAKE_AETHER_TARGET_OS}
            --target-arch ${CMAKE_AETHER_TARGET_ARCH}
            --linker ${CMAKE_AETHER_LINKER}
            --O ${CMAKE_BUILD_TYPE}
            ${ARGS_SOURCES}
        DEPENDS ${ARGS_DEPENDENCIES}
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building Aether executable: ${TARGET}"
        VERBATIM
    )
    
    # Set output properties
    set_target_properties(${TARGET} PROPERTIES
        OUTPUT_NAME ${ARGS_OUTPUT}
        RUNTIME_OUTPUT_DIRECTORY ${CMAKE_RUNTIME_OUTPUT_DIRECTORY}
    )
endfunction()

# Function to add Aether library
function(aether_add_library TARGET)
    set(options SHARED STATIC BOTH)
    set(oneValueArgs TYPE VERSION DESCRIPTION URL)
    set(multiValueArgs SOURCES DEPENDENCIES REQUIRES CONFLICTS PROVIDES)
    
    cmake_parse_arguments(ARGS "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})
    
    # Determine library type
    if(ARGS_SHARED)
        set(LIB_TYPE "shared")
    elseif(ARGS_STATIC)
        set(LIB_TYPE "static")
    elseif(ARGS_BOTH)
        set(LIB_TYPE "both")
    elseif(ARGS_TYPE)
        set(LIB_TYPE ${ARGS_TYPE})
    else()
        set(LIB_TYPE "shared")
    endif()
    
    # Build command arguments
    set(BUILD_ARGS
        --create-library
        --library-type ${LIB_TYPE}
        --library-name ${TARGET}
        --output lib/${TARGET}
    )
    
    if(ARGS_VERSION)
        list(APPEND BUILD_ARGS --library-version ${ARGS_VERSION})
    endif()
    
    if(ARGS_DESCRIPTION)
        list(APPEND BUILD_ARGS --library-description "${ARGS_DESCRIPTION}")
    endif()
    
    if(ARGS_URL)
        list(APPEND BUILD_ARGS --library-url ${ARGS_URL})
    endif()
    
    if(ARGS_REQUIRES)
        list(APPEND BUILD_ARGS --library-requires "${ARGS_REQUIRES}")
    endif()
    
    if(ARGS_CONFLICTS)
        list(APPEND BUILD_ARGS --library-conflicts "${ARGS_CONFLICTS}")
    endif()
    
    if(ARGS_PROVIDES)
        list(APPEND BUILD_ARGS --library-provides "${ARGS_PROVIDES}")
    endif()
    
    # Add pkg-config generation
    list(APPEND BUILD_ARGS --generate-pkg-config)
    
    # Create custom target
    add_custom_target(${TARGET} ALL
        COMMAND ${CMAKE_AETHER_COMPILER} build ${BUILD_ARGS} ${ARGS_SOURCES}
        DEPENDS ${ARGS_DEPENDENCIES}
        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
        COMMENT "Building Aether library: ${TARGET}"
        VERBATIM
    )
    
    # Set target properties
    set_target_properties(${TARGET} PROPERTIES
        LIBRARY_OUTPUT_DIRECTORY ${CMAKE_LIBRARY_OUTPUT_DIRECTORY}
        ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_ARCHIVE_OUTPUT_DIRECTORY}
    )
    
    # Install rules
    install(TARGETS ${TARGET}
        LIBRARY DESTINATION lib
        ARCHIVE DESTINATION lib
        RUNTIME DESTINATION bin
    )
    
    # Install pkg-config file
    install(FILES ${CMAKE_CURRENT_BINARY_DIR}/${TARGET}.pc
        DESTINATION lib/pkgconfig
    )
endfunction()

# Function to find Aether library
function(aether_find_library TARGET)
    set(options REQUIRED)
    set(oneValueArgs)
    set(multiValueArgs)
    
    cmake_parse_arguments(ARGS "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})
    
    # Try pkg-config first
    find_package(PkgConfig QUIET)
    if(PkgConfig_FOUND)
        pkg_check_modules(PC_${TARGET} QUIET ${TARGET})
    endif()
    
    # Find library files
    find_library(${TARGET}_LIBRARY
        NAMES ${TARGET} lib${TARGET}
        PATHS ${PC_${TARGET}_LIBRARY_DIRS}
        PATH_SUFFIXES lib lib64
    )
    
    # Find include directory
    find_path(${TARGET}_INCLUDE_DIR
        NAMES ${TARGET}.ae
        PATHS ${PC_${TARGET}_INCLUDE_DIRS}
        PATH_SUFFIXES include
    )
    
    # Set variables
    if(${TARGET}_LIBRARY)
        set(${TARGET}_FOUND TRUE)
        set(${TARGET}_LIBRARIES ${${TARGET}_LIBRARY})
        set(${TARGET}_INCLUDE_DIRS ${${TARGET}_INCLUDE_DIR})
        
        if(PC_${TARGET}_FOUND)
            set(${TARGET}_VERSION ${PC_${TARGET}_VERSION})
            set(${TARGET}_CFLAGS ${PC_${TARGET}_CFLAGS_OTHER})
            set(${TARGET}_LDFLAGS ${PC_${TARGET}_LDFLAGS_OTHER})
        endif()
    else()
        set(${TARGET}_FOUND FALSE)
    endif()
    
    # Handle REQUIRED
    include(FindPackageHandleStandardArgs)
    find_package_handle_standard_args(${TARGET}
        REQUIRED_VARS ${TARGET}_LIBRARY ${TARGET}_INCLUDE_DIR
        VERSION_VAR ${TARGET}_VERSION
    )
    
    # Mark as advanced
    mark_as_advanced(${TARGET}_LIBRARY ${TARGET}_INCLUDE_DIR)
endfunction()

# Function to link Aether library
function(aether_target_link_libraries TARGET)
    set(multiValueArgs PRIVATE PUBLIC INTERFACE)
    
    cmake_parse_arguments(ARGS "" "" "${multiValueArgs}" ${ARGN})
    
    # Add libraries to target
    if(ARGS_PUBLIC)
        target_link_libraries(${TARGET} PUBLIC ${ARGS_PUBLIC})
    endif()
    
    if(ARGS_PRIVATE)
        target_link_libraries(${TARGET} PRIVATE ${ARGS_PRIVATE})
    endif()
    
    if(ARGS_INTERFACE)
        target_link_libraries(${TARGET} INTERFACE ${ARGS_INTERFACE})
    endif()
endfunction()

# Function to set Aether compiler flags
function(aether_target_compile_options TARGET)
    set(multiValueArgs PRIVATE PUBLIC INTERFACE)
    
    cmake_parse_arguments(ARGS "" "" "${multiValueArgs}" ${ARGN})
    
    # Add compile options to target
    if(ARGS_PUBLIC)
        target_compile_options(${TARGET} PUBLIC ${ARGS_PUBLIC})
    endif()
    
    if(ARGS_PRIVATE)
        target_compile_options(${TARGET} PRIVATE ${ARGS_PRIVATE})
    endif()
    
    if(ARGS_INTERFACE)
        target_compile_options(${TARGET} INTERFACE ${ARGS_INTERFACE})
    endif()
endfunction()

# Function to enable Aether testing
function(aether_enable_testing)
    enable_testing()
    
    # Find Aether test files
    file(GLOB_RECURSE AETHER_TEST_FILES "*.ae")
    
    foreach(TEST_FILE ${AETHER_TEST_FILES})
        if(TEST_FILE MATCHES ".*_test\\.aeth$")
            get_filename_component(TEST_NAME ${TEST_FILE} NAME_WE)
            
            add_test(NAME ${TEST_NAME}
                COMMAND ${CMAKE_AETHER_COMPILER} build ${TEST_FILE}
                WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
            )
        endif()
    endforeach()
endfunction()

# Function to create Aether package
function(aether_package TARGET)
    set(options)
    set(oneValueArgs VERSION DESCRIPTION URL)
    set(multiValueArgs DEPENDENCIES)
    
    cmake_parse_arguments(ARGS "${options}" "${oneValueArgs}" "${multiValueArgs}" ${ARGN})
    
    # Create package configuration
    set(PACKAGE_CONFIG "${CMAKE_CURRENT_BINARY_DIR}/${TARGET}Config.cmake")
    
    file(WRITE ${PACKAGE_CONFIG}
        "set(${TARGET}_VERSION \"${ARGS_VERSION}\")\n"
        "set(${TARGET}_DESCRIPTION \"${ARGS_DESCRIPTION}\")\n"
        "set(${TARGET}_URL \"${ARGS_URL}\")\n"
        "set(${TARGET}_LIBRARIES ${TARGET})\n"
        "set(${TARGET}_INCLUDE_DIRS \"${CMAKE_CURRENT_SOURCE_DIR}/include\")\n"
    )
    
    # Install package configuration
    install(FILES ${PACKAGE_CONFIG}
        DESTINATION lib/cmake/${TARGET}
    )
    
    # Install export file
    install(EXPORT ${TARGET}Targets
        FILE ${TARGET}Targets.cmake
        NAMESPACE ${TARGET}::
        DESTINATION lib/cmake/${TARGET}
    )
endfunction()

# Export variables for use in other CMake files
set(AETHER_FOUND TRUE)
set(AETHER_VERSION ${AETHER_VERSION})
set(AETHER_COMPILER ${AETHER_COMPILER})

# Include directories for Aether
include_directories(${CMAKE_CURRENT_SOURCE_DIR}/include)

# Set up Aether-specific variables
set(CMAKE_AETHER_COMPILER_WORKS TRUE)
set(CMAKE_AETHER_COMPILER_ID "Aether")
set(CMAKE_AETHER_COMPILER_VERSION ${AETHER_VERSION})

message(STATUS "Aether CMake configuration loaded successfully!") 