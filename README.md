# Environment Watcher

Tired of taking care of your .env files? no more. this script automatically encrypts all your .env files and creates a copy in the same directory which you can keep with the project/repo. any changes in the .env file(s) will automatically be reflected if the service is running in background.
You can easily decrypt those files and it will override the orignal (unencrypted) files or create a new one if not there.

DISCLAIMER: for most scenarios, **YOU DONT NEED THIS**. I have created this for people who want to keep their env files in the same repo and not have to worry about key theft.

## Requirements

- Go programming language (for building the Encryptor and Decryptor).
  If you don't have go, follow these instructions: https://go.dev/doc/install
- Linux or Windows operating system.

## Features

- **Auto Encryptor**: Monitors a directory for changes to `.env` files. Automatically encrypts the `.env` files. it will save the file in the same directory where the original file was located. If a file name contains .env, a copy of that file will be created with a suffix `.enc` Ex: `.env` will be converted to `.env.enc`
- **Cryptor**:  Manual Encryption/Decryption script using AES cipher in Galois Counter Mode (GCM). **Share with your friends** using manual encryption mode and provide them with the key used during the manual process (the file wont be generated without a new key, it is advised you use a different key then the default one). 
- **Scripts** : Easy installation and Deletion scripts for Linux and Windows.

## Installation Steps (on terminal)

- clone this git repo: `git clone https://github.com/Raghav-rv28/env-watcher`
- cd into your OS folder name (windows and Linux), and start the installation script (make sure you have administer privileges for windows)
- **Install**: Run the `install.sh | install.bat` script to configure the environment and install the necessary dependencies.
-  **Watch Directory**: Specify the directory to watch for `.env` files. (absolute path)
-  **Encryption Key**: Provide a 16 or 32 character encryption key.
-  **Start Service**: The File Watcher service is automatically started and will monitor the specified directory for changes. (on Windows you might need to restart)

## Usage ( Cryptor Service )

Arugments:
 1. process: "encrypt" / "decrypt", tells the script to either encrypt the file or decrypt.
 2. filename: name of the file
 3. encryption-key: encryption key to be used to either encrypt a file or decrypt an encrypted file. 
#### To Encrypt a file manually

```sh
cryptor encrypt <filename> <encryption-key>
```

#### To Decrypt a file manually

```sh
cryptor decrypt <filename> <encryption-key>
```

If the file name contains .share, it was created using a custom encryption key and only that key will unencrypt the file. 
If a file is encrypted and doesn't have .share (only .enc) specify the default key you used when starting the encryption service.

The default key can be found using commands below:
The file `~/.file_watcher_env` is located in your `/home/<username>/` (Linux) and `C:/users/<username>/` directory, you can grab the encryption token from there directly.

Grab the default encryption key using this:

Linux: `grep -o 'ENCRYPTION_KEY=.*' ~/.file_watcher_env | cut -d '=' -f 2 `

Windows: `for /f "tokens=2 delims==" %i in ('findstr "encryption_key" "%USERPROFILE%\.file_watcher_env"') do @echo %i `

## Directory Structure

- `Auto Encryptor/`: Contains the source code for the Auto Encryptor application.
- `Cryptor/`: Contains the source code for the Manual Cryptor application.
- `Windows/`: Contains the installation and delete scripts for Windows OS.
- `Linux/`: Contains the installation and delete scripts for Linux OS.
  - `install.sh`: Setup script for configuring the environment and installing dependencies.
  - `delete.sh`: if you don't want to have this setup any longer, just use delete.sh to remove all installation files/data. **Once you do this, you will need to use a new key!**
- `README.md`: This file.
- `.env` : sample env file.
- `.env.enc`: sample encrypted file.
- `.env.enc.share` : sample encrypted file using manual encryption.

## Troubleshooting

### Troubleshooting Steps for the file_watcher Service

#### Checking if file_watcher Service is Running:

1. Open a terminal or Command Prompt window.
2. Run the following command to check the status of the file_watcher service:
   - **Linux (systemd):**
     ```sh
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
     ```sh
     sudo journalctl -u file_watcher
     ```
     By default it will show old logs, press `shift + g` to reach the bottom.
   - **Windows (Event Viewer):**
     - Open Event Viewer.
     - Navigate to Windows Logs > Application.
     - Look for logs related to the file_watcher service.

#### Checking for Errors

1. Look for any error messages or warnings in the output/logs.
2. Pay attention to any specific error codes or messages that indicate issues with the service.

#### Restarting the Service

1. If the service is not running or encountering issues, you can try restarting it.
2. Run the following commands to restart the service:
   - **Linux (systemd):**

```sh
     sudo systemctl restart file_watcher
```

- **Windows (Command Prompt):**

```batch
     sc stop file_watcher
     sc start file_watcher
```

#### Reviewing Configuration

1. Double-check the configuration settings, including the encryption key and watch directory, to ensure they are correct.
2. Make any necessary adjustments to the configuration if errors are found.

#### Verifying Permissions

1. Ensure that the service has appropriate permissions to access the directories and files it needs.
2. Check for any permission-related errors in the logs/output.

#### Reset Env Watcher

1. use the delete scripts for your respective OS and start with a fresh installation.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## NOTES FOR AUTHOR

TODO:

1. create installation scripts only for cryptor service. (Linux and Windows)
2. add more configure options for the files being watched.
