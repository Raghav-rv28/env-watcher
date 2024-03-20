#!/bin/bash

# Prompt the user to enter a 16 or 32 byte character string for encryption key
read -p "Enter the encryption key (16 or 32 characters): " ENCRYPTION_KEY

# Prompt the user to enter the path to the directory to watch for .env files
read -p "Enter the path to the directory to watch for .env files: " WATCH_DIRECTORY
export WATCH_DIRECTORY="$WATCH_DIRECTORY"

# Store the environment variables in a new file with restricted permissions
echo "ENCRYPTION_KEY=$ENCRYPTION_KEY" >~/.file_watcher_env
echo "WATCH_DIRECTORY=$WATCH_DIRECTORY" >>~/.file_watcher_env
chmod 600 ~/.file_watcher_env

# Build the Encryptor
cd Encryptor
go build -o file_watcher main.go

# Move the file_watcher executable to /usr/local/bin
sudo mv file_watcher /usr/local/bin/

# Make file_watcher executable
sudo chmod +x /usr/local/bin/file_watcher

# Automatically start file_watcher on PC startup
# Create a systemd service file
echo "[Unit]
Description=File Watcher Service
After=network.target

[Service]
Type=simple
EnvironmentFile=/home/$(whoami)/.file_watcher_env
ExecStart=/usr/local/bin/file_watcher \$WATCH_DIRECTORY
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target" | sudo tee /etc/systemd/system/file_watcher.service

# Reload systemd daemon and start the service
sudo systemctl daemon-reload
sudo systemctl enable file_watcher
sudo systemctl start file_watcher

# Build the Decryptor
cd ../Decryptor
go build -o decrypt main.go

# Copy the decrypt executable to /usr/local/bin
sudo cp decrypt /usr/local/bin/

# Make decrypt executable
sudo chmod +x /usr/local/bin/decrypt

echo "Installation complete."
echo "File Watcher is set to start on PC startup and will watch the directory: $WATCH_DIRECTORY."
echo "The decryption key is set to: $ENCRYPTION_KEY."
echo "file_watcher executable and decrypt executable are copied to /usr/local/bin"
