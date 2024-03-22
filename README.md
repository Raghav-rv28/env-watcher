# Environment Watcher

Tired of taking care of your .env files? no more. this script automatically encrypts all your .env files and creates a copy in the same directory which you can keep with the project/repo. any changes in the .env file(s) will automatically be reflected if the service is running in background.
You can easily decrypt those files and it will override the orignal (unencrypted) files or create a new one if not there.

DISCLAIMER: for most scenarios, **YOU DONT NEED THIS**. I have created this for people who want to keep their env files in the same repo and not have to worry about key theft.

## Requirements

- Go programming language (for building the Encryptor and Decryptor).
  If you don't have go, follow these instructions: https://go.dev/doc/install
- Linux or Windows operating system.

## Features

- **File Watcher**: Monitors a directory for changes to `.env` files. Automatically encrypts the `.env` files. it will save the file in the same directory where the original file was located. If a file name contains .env, a copy of that file will be created with a suffix `.enc` Ex: `.env` will be converted to `.env.enc`
- **Encryption**: Encryption using AES cipher in Galois Counter Mode (GCM).
- **Decryption**: Decrypt encrypted files using a encryption key,
- **Scripts** : Easy installation and Deletion scripts for Linux (windows coming soon!).

## Steps (on terminal)

- clone this git repo: `git clone https://github.com/Raghav-rv28/env-watcher`
- cd into your OS folder name (windows and Linux), and start the installation script (make sure you have administer privileges for windows)
- Follow Encryption steps for further setup.

## Usage ( Encryption Service )

#### To start the encryption service follow these steps (only needs to be done once):

1.  **Install**: Run the `install.sh | install.bat` script to configure the environment and install the necessary dependencies.
2.  **Watch Directory**: Specify the directory to watch for `.env` files. (absolute path)
3.  **Encryption Key**: Provide a 16 or 32 character encryption key.
4.  **Start Service**: The File Watcher service is automatically started and will monitor the specified directory for changes.

## Usage ( Decryption Service )

#### To decrypt a particular file use the following command:

```
decrypt <filename> <encryption-key>
```

specify the key you used when starting the encryption service.

The file `~/.file_watcher_env` is located in your `/home/<username>/` (Linux) and `C:/users/<username>/` directory, you can grab the encryption token from there directly.

Grab the encryption key using this:

Linux: `grep -o 'ENCRYPTION_KEY=.*' ~/.file_watcher_env | cut -d '=' -f 2 `

Windows: `for /f "tokens=2 delims==" %i in ('findstr "encryption_key" "%USERPROFILE%\.file_watcher_env"') do @echo %i `

## Directory Structure

- `Encryptor/`: Contains the source code for the Encryptor application.
- `Decryptor/`: Contains the source code for the Decryptor application.
- `Windows/`: Contains the installation and delete scripts for Windows OS.
- `Linux/`: Contains the installation and delete scripts for Linux OS.
  - `install.sh`: Setup script for configuring the environment and installing dependencies.
  - `delete.sh`: if you don't want to have this setup any longer, just use delete.sh to remove all installation files/data. **Once you do this, you will need to use a new key!**
- `README.md`: This file.
- `.env` : sample env file.
- `.env.enc`: sample encrypted file.

## Troubleshooting

### Troubleshooting Steps for the file_watcher Service

#### Checking if file_watcher Service is Running:

1. Open a terminal or Command Prompt window.
2. Run the following command to check the status of the file_watcher service:
   - **Linux (systemd):**
     ```bash
     sudo systemctl status file_watcher
     ```
   - **Windows (Command Prompt):**
     ```batch
     sc query file_watcher
     ```

#### Viewing file_watcher Service Output:

1. If the service is running but encountering issues, you can view its output for debugging purposes.
2. Run the following command to view the service logs:
   - **Linux (systemd):**
     ```bash
     sudo journalctl -u file_watcher
     ```
   - **Windows (Event Viewer):**
     - Open Event Viewer.
     - Navigate to Windows Logs > Application.
     - Look for logs related to the file_watcher service.

#### Checking for Errors:

1. Look for any error messages or warnings in the output/logs.
2. Pay attention to any specific error codes or messages that indicate issues with the service.

#### Restarting the Service:

1. If the service is not running or encountering issues, you can try restarting it.
2. Run the following commands to restart the service:
   - **Linux (systemd):**
     ```bash
     sudo systemctl restart file_watcher
     ```
   - **Windows (Command Prompt):**
     ```batch
     sc stop file_watcher
     sc start file_watcher
     ```

#### Reviewing Configuration:

1. Double-check the configuration settings, including the encryption key and watch directory, to ensure they are correct.
2. Make any necessary adjustments to the configuration if errors are found.

#### Verifying Permissions:

1. Ensure that the service has appropriate permissions to access the directories and files it needs.
2. Check for any permission-related errors in the logs/output.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
