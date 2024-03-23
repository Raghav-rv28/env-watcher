@echo off

rem Go to the root of the project.
cd ..
rem Prompt the user to enter a 16 or 32 byte character string for encryption key
set /p encryption_key="Enter the encryption key (16 or 32 characters): "

rem Prompt the user to enter the path to the directory to watch for .env files
set /p watch_directory="Enter the path to the directory to watch for .env files: "
set "watch_directory=%watch_directory%"


rem Set environment variables
setx ENCRYPTION_KEY "%encryption_key%"
setx WATCH_DIRECTORY "%watch_directory%"

rem Store the environment variables in a new file
echo encryption_key=%encryption_key% > "%USERPROFILE%\.file_watcher_env"
echo watch_directory=%watch_directory% >> "%USERPROFILE%\.file_watcher_env"

rem Build the encryptor
cd encryptor
go build -ldflags -H=windowsgui -o file_watcher.exe main.go

rem Move the file_watcher executable to System32 directory
move file_watcher.exe "%SystemRoot%\System32"

rem Build the decryptor
cd ../decryptor
go build -ldflags -H=windowsgui -o decrypt.exe main.go

rem Move the decryptor executable to System32 directory
move decrypt.exe "%SystemRoot%\System32"

echo Installation complete.
echo Environment variables ENCRYPTION_KEY and WATCH_DIRECTORY have been set.
echo File watcher is set to start on PC startup and will watch the directory: %watch_directory%.
echo The decryption key is set to: %encryption_key%.
echo File_watcher.exe and decrypt.exe are moved to System32 directory.
