@echo off

rem Remove the file_watcher executable from System32 directory
del "%SystemRoot%\System32\file_watcher.exe" /f /q

rem Remove the decrypt executable from System32 directory
del "%SystemRoot%\System32\decrypt.exe" /f /q

rem Clear environment variables
setx ENCRYPTION_KEY ""
setx WATCH_DIRECTORY ""

echo Uninstallation complete.
echo Environment variables ENCRYPTION_KEY and WATCH_DIRECTORY have been cleared.
echo File watcher and decrypt executables have been removed from System32 directory.
