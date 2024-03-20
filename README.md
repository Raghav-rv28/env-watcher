# Environment Watcher

This repository contains a set of scripts for monitoring environment files in a specified directory and performing actions based on changes.

## Requirements

-   Go programming language (for building the Encryptor and Decryptor). 
If you don't have go, follow these instructions: https://go.dev/doc/install
-   Linux (Windows coming soon) operating system.

## Features

-   **File Watcher**: Monitors a directory for changes to `.env` files. Automatically encrypts  the `.env` files. it will save the file in the same directory where the original file was located. If a file name contains .env, a copy of that file will be created with a suffix `.enc` Ex: `.env` will be converted to `.env.enc`
-   **Encryption**: Encryption using AES cipher in Galois Counter Mode (GCM). 
-   **Decryption**: Decrypt encrypted files using a encryption key,
-  **Scripts** : Easy installation and Deletion scripts for Linux (windows coming soon!).

## Usage ( Encryption Service )
#### To start the encryption service follow these steps:
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
## Directory Structure

-   `Encryptor/`: Contains the source code for the Encryptor application.
-   `Decryptor/`: Contains the source code for the Decryptor application.
-   `install.sh`: Setup script for configuring the environment and installing dependencies.
-  `delete.sh`: if you don't want to have this setup any longer, just use delete.sh to remove all installation files/data. **Once you do this, you will need to use a new key!** 
-   `README.md`: This file.
-   `.env.enc`: sample encrypted file.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
