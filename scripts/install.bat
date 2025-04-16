@echo off
setlocal EnableDelayedExpansion

:: Config
set DOWNLOAD_URL=https://github.com/PaulSteve005/clin/releases/download/stable/clin-windows-x64.exe
set INSTALL_DIR=C:\Program Files\clin
set EXE_NAME=clin.exe
set FULL_PATH=%INSTALL_DIR%\%EXE_NAME%

:: Check if PowerShell is available
where powershell >nul 2>&1
if errorlevel 1 (
    echo PowerShell is required but not found. Exiting.
    exit /b 1
)

:: Create temp download folder
set TMP_FILE=%TEMP%\%EXE_NAME%

:: Download the file
echo Downloading %DOWNLOAD_URL% to %TMP_FILE% ...
powershell -NoProfile -Command "Invoke-WebRequest -Uri '%DOWNLOAD_URL%' -OutFile '%TMP_FILE%'"
if errorlevel 1 (
    echo Failed to download file.
    exit /b 1
)

:: Create install directory
if not exist "%INSTALL_DIR%" (
    echo Creating directory %INSTALL_DIR% ...
    mkdir "%INSTALL_DIR%"
    if errorlevel 1 (
        echo Failed to create install directory.
        exit /b 1
    )
)

:: Move the file
echo Moving %TMP_FILE% to %FULL_PATH% ...
move /Y "%TMP_FILE%" "%FULL_PATH%" >nul
if errorlevel 1 (
    echo Failed to move file to install directory.
    exit /b 1
)

:: Add to PATH (user scope)
echo Adding %INSTALL_DIR% to user PATH if not already present...
for /f "tokens=*" %%A in ('powershell -NoProfile -Command "[Environment]::GetEnvironmentVariable('Path', 'User')"') do set "CUR_PATH=%%A"

echo %CUR_PATH% | find /I "%INSTALL_DIR%" >nul
if errorlevel 1 (
    powershell -NoProfile -Command "[Environment]::SetEnvironmentVariable('Path', '%CUR_PATH%;%INSTALL_DIR%', 'User')"
    echo User PATH updated.
) else (
    echo PATH already contains %INSTALL_DIR%.
)

:: Add to Windows Defender exclusion list (optional, silently fails if no access)
echo Attempting to add Windows Defender exclusion for %INSTALL_DIR% ...
powershell -NoProfile -Command "Try { Add-MpPreference -ExclusionPath '%INSTALL_DIR%' } Catch { Write-Host 'Could not add Defender exclusion (insufficient permissions or Defender disabled).' }"

:: Done
echo.
echo [✓] Installation complete!
echo clin installed at: %FULL_PATH%
echo You may need to restart your terminal or PowerShell session for PATH changes to apply.

endlocal
pause
