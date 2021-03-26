cmake_minimum_required(VERSION 2.8.12)

# configure project
# -- set module path to cmake folder to make go compiler settings visible
#    for the project during declarationn
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/cmake")
project(ipmctl-exporter Go)

get_filename_component(ROOT ./ ABSOLUTE)

# -- define output directory for all build artifacts
set(OUTPUTDIR ${ROOT}/output)
set(INCLUDEDIR ${ROOT}/collector/nvm/include/)
file(MAKE_DIRECTORY ${OUTPUTDIR})
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${ROOT}/output/${BUILD_TYPE})
set(CMAKE_BINARY_DIR ${OUTPUTDIR})
set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR})

# -- include custom settings
if(WIN32) 
    include(cmake/golang_windows.cmake)
else()
    include(cmake/golang_linux.cmake)
endif()
include(cmake/flags.cmake)

# -- get console parameters
if(RELEASE)
    set(BUILD_TYPE release)
    set(CMAKE_BUILD_TYPE Release)
else()
    set(BUILD_TYPE debug)
    set(CMAKE_BUILD_TYPE Debug)
    add_definitions(-DDEBUG_BUILD)
endif()

if(NOT IPMCTL_EXPORTER_VERSION_MAJOR)
    set(IPMCTL_EXPORTER_VERSION_MAJOR 0)
endif()
if(NOT IPMCTL_EXPORTER_VERSION_MINOR)
    set(IPMCTL_EXPORTER_VERSION_MINOR 0)
endif()
if(NOT IPMCTL_EXPORTER_VERSION_PATCH)
    set(IPMCTL_EXPORTER_VERSION_PATCH 0)
endif()

# build project
# -- define build version number
message("CMAKE_VERSION: " ${CMAKE_VERSION})
message("CMAKE_MODULE_PATH: " ${CMAKE_MODULE_PATH})
message("BUILD_TYPE: " ${CMAKE_BUILD_TYPE})
message("COMPILER: " ${CMAKE_Go_COMPILER})
message("OUTPUTDIR: " ${OUTPUTDIR})

if(BUILDNUM)
    set(IPMCTL_EXPORTER_VERSION_STRING ${BUILDNUM})
else()
    find_package(Git)
    if(GIT_FOUND)
        execute_process(COMMAND ${GIT_EXECUTABLE} describe --always --tags
            OUTPUT_VARIABLE IPMCTL_EXPORTER_VERSION_STRING
            RESULT_VARIABLE GIT_RETURN_CODE)
        # -- strip leading string in git tag
        string(REGEX MATCH "([0-9]+\\.[0-9]+\\.[0-9]+)" IPMCTL_EXPORTER_VERSION_STRING "${IPMCTL_EXPORTER_VERSION_STRING}")
        # -- strip ending newline
        string(REGEX REPLACE "\n" "" IPMCTL_EXPORTER_VERSION_STRING "${IPMCTL_EXPORTER_VERSION_STRING}")
    endif()

    if(NOT GIT_FOUND OR ${GIT_RETURN_CODE})
        set(IPMCTL_EXPORTER_VERSION_STRING ${IPMCTL_EXPORTER_VERSION_MAJOR}.${IPMCTL_EXPORTER_VERSION_MINOR}.${IPMCTL_EXPORTER_VERSION_PATCH})
    endif()
endif()

# -- define libraries that have to be linked during the build
if(UNIX)
    set(LNX_BUILD 1)
    set(OS_TYPE linux)
    set(FILE_PREFIX lnx)
    set(CMAKE_PREFIX_PATH /usr/lib/x86_64-linux-gnu/)

    find_package(PkgConfig REQUIRED)
    set_property(GLOBAL PROPERTY FIND_LIBRARY_USE_LIB64_PATHS TRUE)

    include(GNUInstallDirs)
    set(LIBDIR ${CMAKE_INSTALL_PREFIX}/${CMAKE_INSTALL_LIBDIR})

    message("Finding libraries need by linker...")
    message("NOTICE: to point libraries directory by hand set
        -DCMAKE_INSTALL_LIBDIR parameter for cmake. Otherwise cmake will
        try to find these libs accross the system files and link it to the
        default directory pointed by MAKE_INSTALL_LIBDIR env var.")
    message("libs:")
    find_library(IPMCTL_LIBRARY REQUIRED
        NAMES libipmctl.so libipmctl.so.4
        HINTS /usr/lib64 ${LIBDIR}
    )
    find_library(NDCTL_LIBRARY REQUIRED
        NAMES libndctl.so libndctl.so.6
        HINTS /usr/lib64 ${LIBDIR}
    )
    find_library(DAXCTL_LIBRARY REQUIRED
        NAMES libdaxctl.so libdaxctl.so.1
        HINTS /usr/lib64 ${LIBDIR}
    )
    find_library(KMOD_LIBRARY REQUIRED
        NAMES libkmod.so libkmod.so.2
        HINTS /usr/lib64 ${LIBDIR}
    )
    message("  -- NDCTL_LIBRARY: " ${NDCTL_LIBRARY})
    message("  -- DAXCTL_LIBRARY: " ${DAXCTL_LIBRARY})
    message("  -- KMOD_LIBRARY: " ${KMOD_LIBRARY})
elseif(WIN32)
    set(WIN_BUILD 1)
    set(OS_TYPE win)
    set(FILE_PREFIX win)

    find_package(PkgConfig REQUIRED)

    FILE(GLOB BUILD_DEPS
        $ENV{PMM_HOME}/export_api.h
        $ENV{PMM_HOME}/nvm_management.h
        $ENV{PMM_HOME}/nvm_types.h
        $ENV{PMM_HOME}/NvmSharedDefs.h
        $ENV{PMM_HOME}/libipmctl.dll)

    FILE(COPY ${BUILD_DEPS}
        DESTINATION ${OUTPUTDIR})

    message("PMM_HOME: $ENV{PMM_HOME}")
    message("LIBDIR: ${OUTPUTDIR}")
    set(CMAKE_INSTALL_PREFIX "${OUTPUTDIR}")
    set(CMAKE_INSTALL_LIBDIR "${OUTPUTDIR}")
    set(CMAKE_INSTALL_BINDIR "${OUTPUTDIR}")
    set(CMAKE_INSTALL_INCLUDEDIR "${OUTPUTDIR}")
    include(GNUInstallDirs)
    set(CMAKE_FIND_LIBRARY_PREFIXES "")
    set(CMAKE_FIND_LIBRARY_SUFFIXES ".lib" ".dll")

    find_library(IPMCTL_LIBRARY REQUIRED
        NAMES ipmctl libipmctl
        HINTS $ENV{PMM_HOME} ${OUTPUTDIR}
    )
endif()

message("  -- IPMCTL_LIBRARY: " ${IPMCTL_LIBRARY})
message("IPMCTL_EXPORTER_VERSION_STRING: " ${IPMCTL_EXPORTER_VERSION_STRING})

include_directories(${INCLUDEDIR})
add_go_executable(ipmctl-exporter)

# -- run build command
if(LNX_BUILD)
    message("INCLUDEDIR: " ${INCLUDEDIR})
    # -- link libraries
    message("LIBDIR: " ${LIBDIR})
    if(NOT EXISTS ${LIBDIR}/libipmctl.so)
        message("LIB " ${LIBDIR} "/libipmctl.so - NOT FOUND")
        message("LINK " ${LIBDIR}/libipmctl.so " to " ${IPMCTL_LIBRARY})
        execute_process(COMMAND sudo ln -s ${IPMCTL_LIBRARY} ${LIBDIR}/libipmctl.so)
    endif()
    if(NOT EXISTS ${LIBDIR}/libndctl.so)
        message("LIB " ${LIBDIR} "/libndctl.so - NOT FOUND")
        message("LINK " ${LIBDIR}/libndctl.so " to " ${NDCTL_LIBRARY})
        execute_process(COMMAND sudo ln -s ${NDCTL_LIBRARY} ${LIBDIR}/libndctl.so)
    endif()
    if(NOT EXISTS ${LIBDIR}/libdaxctl.so)
        message("LIB " ${LIBDIR} "/libdaxctl.so - NOT FOUND")
        message("LINK " ${LIBDIR}/libdaxctl.so " to " ${DAXCTL_LIBRARY})
        execute_process(COMMAND sudo ln -s ${DAXCTL_LIBRARY} ${LIBDIR}/libdaxctl.so)
    endif()
    if(NOT EXISTS ${LIBDIR}/libkmod.so)
        message("LIB " ${LIBDIR} "/libkmod.so - NOT FOUND")
        message("LINK " ${LIBDIR}/libkmod.so " to " ${KMOD_LIBRARY})
        execute_process(COMMAND sudo ln -s ${KMOD_LIBRARY} ${LIBDIR}/libkmod.so)
    endif()
    # -- configure package config
    configure_file(${ROOT}/install/linux/libipmctl.pc.in ${OUTPUTDIR}/libipmctl.pc @ONLY)
    set(ENV{PKG_CONFIG_PATH} "${OUTPUTDIR}")
    message("PKG_CONFIG_PATH: " $ENV{PKG_CONFIG_PATH})
endif()

if(WIN_BUILD)
    # -- configure package config
    configure_file(${ROOT}/install/win/libipmctl.pc.in ${OUTPUTDIR}/libipmctl.pc @ONLY)
    set(ENV{PKG_CONFIG_PATH} "${OUTPUTDIR}")
	message("PKG_CONFIG_PATH: " $ENV{PKG_CONFIG_PATH})
endif()

# Windows installer build
if(WIN_BUILD AND RELEASE)

    # IPMCTL release link, used to download IPMCTL Windows installer
    set(IPMCTL_RELEASE_URL https://github.com/intel/ipmctl/releases/download)
    # VERSION NUMBER - REMEMBER TO UPDATE WHEN NEW RELEASE COMES OUT
    set(IPMCTL_VERSION 02.00.00.3852)
    set(IPMCTL_NAME "ipmctl_windows_install-${IPMCTL_VERSION}.exe")
    
    # Downloads IPMCTL Windows installer that is embedded into our exporter installer
    if (NOT EXISTS ${IPMCTL_NAME})
        message(DOWNLOADING_IPMCTL_WINDOWS_SETUP)
        FILE(DOWNLOAD ${IPMCTL_RELEASE_URL}/v${IPMCTL_VERSION}/${IPMCTL_NAME} ${OUTPUTDIR}/${IPMCTL_NAME}
             SHOW_PROGRESS)
    endif()
    install(PROGRAMS ${OUTPUTDIR}/${IPMCTL_NAME}
            DESTINATION "PMem Software"
            COMPONENT ipmctl
    )
    
    set(CPACK_PACKAGE_VERSION ${IPMCTL_EXPORTER_VERSION_STRING})
    set(CPACK_RESOURCE_FILE_LICENSE ${ROOT}/LICENSE.txt)
    set(CPACK_NSIS_DISPLAY_NAME "Intel(R) Optane(TM) PMC Exporter")
    
    # Checks for already installed versions
    set(CPACK_NSIS_ENABLE_UNINSTALL_BEFORE_INSTALL ON)
    
    # Extra command that runs the embedded Ipmctl Windows installer
    set(CPACK_NSIS_EXTRA_INSTALL_COMMANDS 
        " ExecWait '$INSTDIR\\\\PMem Software\\\\${IPMCTL_NAME}' "
    )
    
    # Default installation path
    set(CPACK_PACKAGE_INSTALL_DIRECTORY "Ipmctl-exporter")
    set(CPACK_NSIS_INSTALL_ROOT $PROGRAMFILES\\\\Intel)
    
    include(CPack)
    
    cpack_add_component(ipmctl_exporter
                        DISPLAY_NAME "Ipmctl-Exporter"
                        DESCRIPTION  "Utility for exposing health and performance metrics 
                                      from Intel(R) Optane(TM) Persistent Memory modules (PMem)
                                      for Prometheus."
                        REQUIRED 1
    ) 
                                                    
    cpack_add_component(ipmctl
                        DISPLAY_NAME "Ipmctl"
                        DESCRIPTION "Utility for configuring and managing Intel(R) Optane(TM) Persistent Memory modules (PMem)."
    )

endif()