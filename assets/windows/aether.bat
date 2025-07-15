@echo off
setlocal enabledelayedexpansion

REM Get the directory where this batch file is located
set "AETHERROOT=%~dp0"
set "AETHERROOT=%AETHERROOT:~0,-1%"

REM Set environment variables
set "AETHER_BIN=%AETHERROOT%\bin"
set "AETHER_PACKAGES=%AETHERROOT%\packages"

REM Add to PATH for this session
set "PATH=%AETHER_BIN%;%PATH%"

REM Run aether with all arguments
"%AETHER_BIN%\aether.exe" %*
