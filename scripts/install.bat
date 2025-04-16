@echo off
:: Elevate to admin if needed
>nul 2>&1 "%SYSTEMROOT%\system32\cacls.exe" "%SYSTEMROOT%\system32\config\system"
if '%errorlevel%' NEQ '0' (
    echo Requesting administrator access...
    powershell -Command "Start-Process '%~f0' -Verb RunAs"
    exit /b
)

:: Call PowerShell with a here-string script
powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"& {
$DownloadUrl = 'https://github.com/PaulSteve005/clin/releases/download/stable/clin-windows-x64'
$InstallDir = 'C:\Program Files\clin'
$ExeName = 'clin.exe'
$FullPath = Join-Path $InstallDir $ExeName

function Download-File {
    param([string]$Url, [string]$OutFile)
    if (-Not (Test-Path $OutFile)) {
        Write-Host 'Downloading from' $Url 'to' $OutFile '...'
        try {
            Invoke-WebRequest -Uri $Url -OutFile $OutFile -ErrorAction Stop
            Write-Host 'Download successful.'
        } catch {
            Write-Error 'Failed to download file: ' + $_.Exception.Message
            exit 1
        }
    } else {
        Write-Host 'File' $OutFile 'already exists. Skipping download.'
    }
}

function Install-File {
    param([string]$InstallPath, [string]$FileName)
    Write-Host 'Creating directory' $InstallPath '...'
    if (-not (Test-Path -Path $InstallPath -PathType Container)) {
        try {
            New-Item -ItemType Directory -Path $InstallPath -ErrorAction Stop | Out-Null
            Write-Host 'Directory created.'
        } catch {
            Write-Error 'Failed to create directory: ' + $_.Exception.Message
            exit 1
        }
    }
    Write-Host 'Moving file to' $InstallPath '...'
    try {
        Move-Item -Path $FileName -Destination $InstallPath -Force -ErrorAction Stop
        Write-Host 'File installation successful.'
    } catch {
        Write-Error 'Failed to move file: ' + $_.Exception.Message
        exit 1
    }
}

function Add-To-Path {
    param([string]$PathToAdd)
    Write-Host 'Adding' $PathToAdd 'to PATH...'
    $CurrentPath = [Environment]::GetEnvironmentVariable('Path', 'User')
    $PathList = $CurrentPath.Split(';')
    if (-not ($PathList -contains $PathToAdd)) {
        try {
            [Environment]::SetEnvironmentVariable('Path', "$CurrentPath;$PathToAdd", 'User')
            Write-Host 'Path updated. Restart your terminal for changes to take effect.'
        } catch {
            Write-Error 'Failed to update PATH: ' + $_.Exception.Message
            exit 1
        }
    } else {
        Write-Host 'Path' $PathToAdd 'is already in PATH.'
    }
}

function Add-Defender-Exception {
    param([string]$FilePath)
    if (Get-Command -Name Add-MpPreference -ErrorAction SilentlyContinue) {
        Write-Host 'Adding Windows Defender exclusion for' $FilePath '...'
        try {
            Add-MpPreference -ExclusionPath $FilePath -Force -ErrorAction Stop
            Write-Host 'Defender exclusion added.'
        } catch {
            Write-Warning 'Failed to add Defender exclusion: ' + $_.Exception.Message
        }
    } else {
        Write-Warning 'Defender cmdlet not available. Skipping exclusion.'
    }
}

Download-File -Url $DownloadUrl -OutFile $ExeName
Install-File -InstallPath $InstallDir -FileName $ExeName
Add-To-Path -PathToAdd $InstallDir
Add-Defender-Exception -FilePath $FullPath

Write-Host ''
Write-Host '[✓] Installation complete. clin is located at:' $FullPath -ForegroundColor Green
Write-Host 'You may need to restart your terminal or PowerShell session.'
}"
