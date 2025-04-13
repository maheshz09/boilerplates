#!/bin/bash

# Set variables
GIT_REPO_DIR="/etc/httpd/mellon"
GIT_REPO_URL="git@gitlab.com:your-username/apache-configs.git"
SERVERS=("server01" "server02" "server03")

# Function to update Git repository
update_git_repo() {
    echo "Checking for updates in $GIT_REPO_DIR..."
    cd "$GIT_REPO_DIR" || { echo "Failed to access $GIT_REPO_DIR"; exit 1; }

    git fetch origin main
    if ! git diff --quiet origin/main; then
        echo "Changes detected, updating repository..."
        git reset --hard origin/main
        git pull --rebase origin main
        restart_apache
    else
        echo "No updates found, skipping restart."
    fi
}

# Function to restart Apache
restart_apache() {
    echo "Restarting Apache service..."
    if systemctl is-active --quiet apache2; then
        systemctl restart apache2
    elif systemctl is-active --quiet httpd; then
        systemctl restart httpd
    else
        echo "Apache service not found. Check your system configuration."
    fi
}

# Sync updates on local server
update_git_repo

# Sync updates across remote servers
for SERVER in "${SERVERS[@]}"; do
    echo "Syncing Apache config on $SERVER..."
    ssh user@$SERVER <<EOF
        cd "$GIT_REPO_DIR" || exit 1
        git fetch origin main
        if ! git diff --quiet origin/main; then
            echo "Updating repository on $SERVER..."
            git reset --hard origin/main
            git pull --rebase origin main
            systemctl restart apache2 || systemctl restart httpd
        else
            echo "No updates found on $SERVER."
        fi
EOF
done

echo "Apache config sync completed successfully!"
