#!/bin/bash

# Remove environment variables file
sudo rm -f /home/$(whoami)/.file_watcher_env

# Stop and disable the file_watcher service
sudo systemctl stop file_watcher
sudo systemctl disable file_watcher

# Remove the file_watcher executable from /usr/local/bin
sudo rm -f /usr/local/bin/file_watcher

# Remove the decrypt executable from /usr/local/bin
sudo rm -f /usr/local/bin/decrypt

# Remove the file_watcher service file
sudo rm -f /etc/systemd/system/file_watcher.service

# Reload systemd daemon
sudo systemctl daemon-reload

echo "Deletion complete."
