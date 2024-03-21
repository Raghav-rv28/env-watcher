# Environment Watcher
Tired of taking care of your .env files? no more. this script automatically encrypts all your .env files and creates a copy in the same directory which you can keep with the project/repo. any changes in the .env file(s) will automatically be reflected if the service is running in background.
You can easily decrypt those files and it will override the orignal (unencrypted) files or create a new one if not there.

DISCLAIMER: for most scenarios, **YOU DONT NEED THIS**. I have created this for people who want to keep their env files in the same repo and not have to worry about key theft.
 
## Requirements

-   Go programming language (for building the Encryptor and Decryptor). 
If you don't have go, follow these instructions: https://go.dev/doc/install
-   Linux (Windows coming soon) operating system.

## Features

-   **File Watcher**: Monitors a directory for changes to `.env` files. Automatically encrypts  the `.env` files. it will save the file in the same directory where the original file was located. If a file name contains .env, a copy of that file will be created with a suffix `.enc` Ex: `.env` will be converted to `.env.enc`
-   **Encryption**: Encryption using AES cipher in Galois Counter Mode (GCM). 
-   **Decryption**: Decrypt encrypted files using a encryption key,
-  **Scripts** : Easy installation and Deletion scripts for Linux (windows coming soon!).

## Steps
- clone this git repo: `git clone https://github.com/Raghav-rv28/env-watcher`
- open the repo in terminal (make sure the terminal is at the repo level where the install.sh script is present).
- Follow Encryption steps for further setup.

## Usage ( Encryption Service )
#### To start the encryption service follow these steps (only needs to be done once): 
1.  **Install**: Run the `install.sh` script to configure the environment and install the necessary dependencies.
2.  **Watch Directory**: Specify the directory to watch for `.env` files.
3.  **Encryption Key**: Provide a 16 or 32 character encryption key.
4.  **Start Service**: The File Watcher service is automatically started and will monitor the specified directory for changes.

## Usage ( Decryption Service )
#### To decrypt a particular file use the following command:
```
decrypt <filename> <encryption-key>
```
specify the key you used when starting the encryption service.

The file `~/.file_watcher_env` is located in your `/home/<username>/` directory, you can grab the encryption token from there directly.

Grab the encryption key using this:

```grep -o 'ENCRYPTION_KEY=.*' ~/.file_watcher_env | cut -d '=' -f 2 ``` 
## Directory Structure

-   `Encryptor/`: Contains the source code for the Encryptor application.
-   `Decryptor/`: Contains the source code for the Decryptor application.
-   `install.sh`: Setup script for configuring the environment and installing dependencies.
-  `delete.sh`: if you don't want to have this setup any longer, just use delete.sh to remove all installation files/data. **Once you do this, you will need to use a new key!** 
-   `README.md`: This file.
-   `.env.enc`: sample encrypted file.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
