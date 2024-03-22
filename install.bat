@echo off

rem Prompt the user to enter a 16 or 32 byte character string for encryption key
set /p encryption_key="Enter the encryption key (16 or 32 characters): "

rem Prompt the user to enter the path to the directory to watch for .env files
set /p watch_directory="Enter the path to the directory to watch for .env files: "
set "watch_directory=%watch_directory%"

rem Store the environment variables in a new file
echo encryption_key=%encryption_key% > "%USERPROFILE%\.file_watcher_env"
echo watch_directory=%watch_directory% >> "%USERPROFILE%\.file_watcher_env"

rem Build the encryptor
cd encryptor
go build -o file_watcher.exe main.go

rem Move the file_watcher executable to a directory in the PATH environment variable
move file_watcher.exe "%SystemRoot%\System32"

rem Build the decryptor
cd ../decryptor
go build -o decrypt.exe main.go

rem Move the decryptor to a directory in the PATH environment variable
move decrypt.exe "%SystemRoot%\System32"

rem Create a scheduled task to start file_watcher on user logon
schtasks /create /tn "FileWatcher" /sc onlogon /rl highest /tr "%SystemRoot%\System32\file_watcher.exe %watch_directory%"

echo Installation complete.
echo File watcher is scheduled to start on user logon and will watch the directory: %watch_directory%.
echo The decryption key is set to: %encryption_key%.
echo File_watcher.exe and decrypt.exe are moved to the System32 directory.

