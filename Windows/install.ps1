# Function to check if a software is installed
function CheckInstalled($software) {
    if (Get-Command $software -ErrorAction SilentlyContinue) {
        return $true
    } else {
        return $false
    }
}

function RemoveIfExistsWithElevation($path) {
    if (Test-Path $path) {
        # Use Start-Process with -Verb RunAs to run the command with elevated permissions
        Start-Process powershell -ArgumentList "-Command & { Remove-Item -Path '$path' -Recurse -Force }" -Verb RunAs
    }
}
# Check if Chocolatey is installed, if not install
if (-not (CheckInstalled 'choco')) {
    Write-Host "Chocolatey is not installed. Installing..."
    Set-ExecutionPolicy Bypass -Scope Process -Force; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))
    Write-Host "Please restart PowerShell and run the script again to continue the installation."
    exit
}

# Check if OpenSSL is installed, if not install
if (-not (CheckInstalled 'openssl')) {
    Write-Host "OpenSSL is not installed. Installing..."
    choco install openssl.light -y
    Write-Host "Please restart PowerShell and run the script again to continue the installation."
    exit
}

# Go to the user's home directory
Set-Location $env:USERPROFILE\Desktop

# Clone the code temporarily
git clone https://github.com/Raghav-rv28/env-watcher

Set-Location $env:USERPROFILE\Desktop\env-watcher

# Prompt the user to enter an encryption key
$encryption_key = Read-Host "Enter the encryption key (16 or 32 characters, press Enter for random key)"

# Prompt the user to enter the path to the directory to watch for .env files
$watch_directory = Read-Host "Enter the path to the directory to watch for .env files"

# Prompt the user to enter folders to ignore
$ignore_dir = Read-Host "Enter folders you want to ignore separated by ; (only absolute paths)"

# If the user pressed Enter without typing a key, generate a random key
if (-not $encryption_key) {
    $encryption_key = (openssl rand -base64 32) -replace "[^a-zA-Z0-9]", ""
    Write-Host "Randomly generated encryption key: $encryption_key"
    } else {
    Write-Host "Entered encryption key: $encryption_key"
}

# Store the environment variables in a new file
@"
encryption_key=$encryption_key
watch_directory=$watch_directory
ignore_dir=$ignore_dir
"@ | Set-Content -Path "$env:USERPROFILE\.file_watcher_env"

# Build the encryptor
Set-Location "$env:USERPROFILE\Desktop\env-watcher\Auto-Encryptor"
go build -o file_watcher main.go

# Move the file_watcher executable to C:\Program Files
Move-Item -Path .\file_watcher.exe -Destination 'C:\Program Files\file_watcher.exe'

# Automatically start file_watcher on PC startup
# Create a scheduled task
$action = New-ScheduledTaskAction -Execute 'C:\Program Files\file_watcher.exe' -Argument $watch_directory
$trigger = New-ScheduledTaskTrigger -AtStartup
Register-ScheduledTask -Action $action -Trigger $trigger -TaskName 'file_watcher' -Description 'File Watcher Service'

# Build the decryptor
Set-Location ..\Cryptor\
go build -o cryptor main.go

# Copy the decryptor executable to C:\Program Files
Move-Item -Path .\cryptor.exe -Destination 'C:\Program Files\cryptor.exe'

Write-Host "Installation complete."
Write-Host "The Auto-Encryptor will automatically ignore node_modules and directories starting with ."
Write-Host "File watcher is set to start on PC startup and will watch the directory: $watch_directory."
Write-Host "The decryption key is set to: $encryption_key."
Write-Host "File_watcher executable and cryptor executable are copied to C:\Program Files"

# Return to the home directory and delete the temporary code
Set-Location $env:USERPROFILE\Desktop
RemoveIfExistsWithElevation "$env:USERPROFILE\Desktop\env-watcher"
