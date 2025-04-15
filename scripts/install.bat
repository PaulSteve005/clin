@echo off
powershell -NoProfile -ExecutionPolicy Bypass -Command "Start-Process -Verb RunAs -FilePath \"%SystemRoot%\System32\WindowsPowerShell\v1.0\powershell.exe\" -ArgumentList '-NoProfile -ExecutionPolicy Bypass -File \"%~dp0install_clin.ps1\"'"
```powershell
# --- Configuration ---
$DownloadUrl = "[https://github.com/PaulSteve005/clin/releases/download/stable/clin-windows-x64](https://github.com/PaulSteve005/clin/releases/download/stable/clin-windows-x64)"
$InstallDir = "C:\Program Files\clin"
$ExeName = "clin.exe"
$FullPath = "$InstallDir\$ExeName"

# --- Helper Functions ---
function Test-Admin {
    $IsAdmin = ([Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
    if (-not $IsAdmin) {
        Write-Error "This script requires administrator privileges. Please run it as administrator."
        exit 1
    }
}

function Download-File {
    param(
        [string]$Url,
        [string]$OutFile
    )
    Write-Host "Downloading from '$Url' to '$OutFile'..."
    try {
        Invoke-WebRequest -Uri $Url -OutFile $OutFile -ErrorAction Stop
        Write-Host "Download successful."
    } catch {
        Write-Error "Failed to download file: $($_.Exception.Message)"
        exit 1
    }
}

function Install-File {
    param(
        [string]$InstallPath,
        [string]$FileName
    )
    Write-Host "Creating directory '$InstallPath'..."
    if (-not (Test-Path -Path $InstallPath -PathType 'Container')) {
        try {
            New-Item -ItemType Directory -Path $InstallPath -ErrorAction Stop
            Write-Host "Directory created."
        } catch {
            Write-Error "Failed to create directory: $($_.Exception.Message)"
            exit 1
        }
    }

    Write-Host "Moving file to '$InstallPath'..."
     try{
        Move-Item -Path $FileName -Destination $InstallPath -ErrorAction Stop
     }
     catch{
        Write-Error "Failed to move file: $($_.Exception.Message)"
        exit 1
     }
    Write-Host "File installation successful."
}

function Add-To-Path {
    param(
        [string]$PathToAdd
    )
    Write-Host "Adding '$PathToAdd' to PATH..."
    $CurrentPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ($CurrentPath -notcontains $PathToAdd) {
        try {
            [Environment]::SetEnvironmentVariable("Path", "$CurrentPath;$PathToAdd", "User") -ErrorAction Stop
            Write-Host "Path updated.  You may need to restart your terminal/PowerShell session."
        } catch {
            Write-Error "Failed to update PATH: $($_.Exception.Message)"
            exit 1
        }
    } else {
        Write-Host "Path '$PathToAdd' is already in PATH."
    }
}

function Add-Defender-Exception {
    param(
        [string]$FilePath
    )
    # Check if the Defender module is available (requires admin)
    if (Get-Module -Name Defender -ListAvailable) {
        Write-Host "Adding Windows Defender exclusion for '$FilePath'..."
        try {
            Add-MpPreference -ExclusionPath $FilePath -ErrorAction Stop
            Write-Host "Defender exclusion added."
        } catch {
            Write-Warning "Failed to add Defender exclusion: $($_.Exception.Message).  This is not critical, but Defender might block clin."
            # Don't exit, continue with installation
        }
    }
    else{
       Write-Warning "The Defender module is not available.  Cannot add Defender exclusion.  This is not critical, but Defender might block clin."
    }
}

# --- Main Script ---

# Check for administrator privileges
Test-Admin

# Download the file
Download-File -Url $DownloadUrl -OutFile $ExeName

# Install the file
Install-File -InstallPath $InstallDir -FileName $ExeName

# Add to PATH
Add-To-Path -PathToAdd $InstallDir

# Add Defender exception
Add-Defender-Exception -FilePath $FullPath

Write-Host "Installation complete. clin is located at '$FullPath'."
Write-Host "You may need to restart your terminal/PowerShell session for the PATH changes to take effect."


