#!/bin/bash

cd ..
# prompt the user to enter a 16 or 32 byte character string for encryption key
read -p "enter the encryption key (16 or 32 characters): " encryption_key

# prompt the user to enter the path to the directory to watch for .env files
read -p "enter the path to the directory to watch for .env files: " watch_directory
export watch_directory="$watch_directory"

# store the environment variables in a new file with restricted permissions
echo "encryption_key=$encryption_key" >~/.file_watcher_env
echo "watch_directory=$watch_directory" >>~/.file_watcher_env
chmod 600 ~/.file_watcher_env

# build the encryptor
cd Auto-Encryptor
go build -o file_watcher main.go

# move the file_watcher executable to /usr/local/bin
sudo mv file_watcher /usr/local/bin/

# make file_watcher executable
sudo chmod +x /usr/local/bin/file_watcher

# automatically start file_watcher on pc startup
# create a systemd service file
cat <<EOF | sudo tee /etc/systemd/system/file_watcher.service
[Unit]
Description=File Watcher Service
After=network.target

[Service]
Type=simple
EnvironmentFile=/home/$(whoami)/.file_watcher_env
ExecStart=/usr/local/bin/file_watcher \$watch_directory
Restart=always
RestartSec=3

[Install]
WantedBy=multi-user.target
EOF
# reload systemd daemon and start the service
sudo systemctl daemon-reload
sudo systemctl enable file_watcher
sudo systemctl start file_watcher

# build the decryptor
cd ../Cryptor/
go build -o cryptor main.go

# copy the decrypt executable to /usr/local/bin
sudo cp cryptor /usr/local/bin/

# make decrypt executable
sudo chmod +x /usr/local/bin/cryptor

echo "installation complete."
echo "file watcher is set to start on pc startup and will watch the directory: $watch_directory."
echo "the decryption key is set to: $encryption_key."
echo "file_watcher executable and cryptor executable are copied to /usr/local/bin"
