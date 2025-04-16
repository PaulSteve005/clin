@echo off
setlocal EnableDelayedExpansion

:: Config
set INSTALL_DIR=C:\Program Files\clin
set EXE_NAME=clin.exe
set FULL_PATH=%INSTALL_DIR%\%EXE_NAME%

echo Uninstalling clin...

:: Remove the clin.exe file
if exist "%FULL_PATH%" (
    echo Deleting %FULL_PATH% ...
    del "%FULL_PATH%" >nul
    if errorlevel 1 (
        echo Failed to delete %FULL_PATH%.
    ) else (
        echo Deleted %FULL_PATH%.
    )
) else (
    echo File not found: %FULL_PATH%
)

:: Remove the install directory if empty
if exist "%INSTALL_DIR%" (
    rd "%INSTALL_DIR%" >nul 2>&1
    if errorlevel 1 (
        echo Could not remove %INSTALL_DIR% (might not be empty).
    ) else (
        echo Removed folder %INSTALL_DIR%.
    )
)

:: Remove clin path from User PATH
echo Checking for PATH entry...
for /f "tokens=*" %%A in ('powershell -NoProfile -Command "[Environment]::GetEnvironmentVariable('Path', 'User')"') do set "CUR_PATH=%%A"

echo %CUR_PATH% | find /I "%INSTALL_DIR%" >nul
if errorlevel 1 (
    echo No PATH entry found for %INSTALL_DIR%.
) else (
    echo Removing %INSTALL_DIR% from PATH...
    powershell -NoProfile -Command "[Environment]::SetEnvironmentVariable('Path', ($env:Path -split ';' | Where-Object { $_ -ne '%INSTALL_DIR%' }) -join ';', 'User')"
    echo PATH entry removed.
)

:: Remove Defender exclusion
echo Attempting to remove Windows Defender exclusion for %INSTALL_DIR% ...
powershell -NoProfile -Command "Try { Remove-MpPreference -ExclusionPath '%INSTALL_DIR%' } Catch { Write-Host 'Could not remove exclusion (may require admin or Defender may be disabled).' }"

echo.
echo [✓] clin has been uninstalled.
echo You may need to restart your terminal or PowerShell for PATH changes to apply.

endlocal
pause
@echo off
setlocal EnableDelayedExpansion

:: Config
set INSTALL_DIR=C:\Program Files\clin
set EXE_NAME=clin.exe
set FULL_PATH=%INSTALL_DIR%\%EXE_NAME%

echo Uninstalling clin...

:: Remove the clin.exe file
if exist "%FULL_PATH%" (
    echo Deleting %FULL_PATH% ...
    del "%FULL_PATH%" >nul
    if errorlevel 1 (
        echo Failed to delete %FULL_PATH%.
    ) else (
        echo Deleted %FULL_PATH%.
    )
) else (
    echo File not found: %FULL_PATH%
)

:: Remove the install directory if empty
if exist "%INSTALL_DIR%" (
    rd "%INSTALL_DIR%" >nul 2>&1
    if errorlevel 1 (
        echo Could not remove %INSTALL_DIR% (might not be empty).
    ) else (
        echo Removed folder %INSTALL_DIR%.
    )
)

:: Remove clin path from User PATH
echo Checking for PATH entry...
for /f "tokens=*" %%A in ('powershell -NoProfile -Command "[Environment]::GetEnvironmentVariable('Path', 'User')"') do set "CUR_PATH=%%A"

echo %CUR_PATH% | find /I "%INSTALL_DIR%" >nul
if errorlevel 1 (
    echo No PATH entry found for %INSTALL_DIR%.
) else (
    echo Removing %INSTALL_DIR% from PATH...
    powershell -NoProfile -Command "[Environment]::SetEnvironmentVariable('Path', ($env:Path -split ';' | Where-Object { $_ -ne '%INSTALL_DIR%' }) -join ';', 'User')"
    echo PATH entry removed.
)

:: Remove Defender exclusion
echo Attempting to remove Windows Defender exclusion for %INSTALL_DIR% ...
powershell -NoProfile -Command "Try { Remove-MpPreference -ExclusionPath '%INSTALL_DIR%' } Catch { Write-Host 'Could not remove exclusion (may require admin or Defender may be disabled).' }"

echo.
echo [✓] clin has been uninstalled.
echo You may need to restart your terminal or PowerShell for PATH changes to apply.

endlocal
pause

