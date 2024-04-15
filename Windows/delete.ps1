# Function to remove a file or directory if it exists
function RemoveIfExists($path) {
    if (Test-Path $path) {
        Remove-Item -Path $path -Recurse -Force
    }
}

# Stop and remove the file_watcher service
Stop-Service -Name 'file_watcher' -Force
Unregister-ScheduledTask -TaskName 'file_watcher' -Confirm:$false

# Remove the file_watcher executable
RemoveIfExists 'C:\Program Files\file_watcher.exe'

# Remove the cryptor executable
RemoveIfExists 'C:\Program Files\cryptor.exe'

# Remove the .file_watcher_env file
RemoveIfExists "$env:USERPROFILE\.file_watcher_env"

# Go to the home directory
Set-Location $env:USERPROFILE

# Remove the env-watcher directory
RemoveIfExists 'env-watcher'

Write-Host "Uninstallation complete."
Write-Host "All files, services, and environment variables created by the installation script have been removed."
