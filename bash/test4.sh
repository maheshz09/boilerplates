#!/bin/bash

# Define GitLab Repository Info
GIT_REPO_DIR="/etc/httpd/mellon"
GIT_REPO_URL="https://oauth2:$GITLAB_PAT@gitlab.com/your-username/apache-configs.git"

# Define Target Servers
SERVERS=("server01" "server02" "server03")

# Pull latest changes from GitLab
echo "Pulling latest Apache config updates..."
cd "$GIT_REPO_DIR" || exit 1
git fetch origin main
if ! git diff --quiet origin/main; then
    echo "Changes detected, updating..."
    git reset --hard origin/main
    git pull --rebase origin main
    systemctl restart apache2 || systemctl restart httpd
else
    echo "No updates found."
fi

# Sync changes across remote servers
for SERVER in "${SERVERS[@]}"; do
    echo "Syncing Apache config on $SERVER..."
    ssh user@$SERVER <<EOF
        export GITLAB_PAT="$GITLAB_PAT"
        cd $GIT_REPO_DIR || exit 1
        git fetch "$GIT_REPO_URL" main || exit 1
        git reset --hard origin/main || exit 1
        git pull --rebase "$GIT_REPO_URL" main || exit 1
        systemctl restart apache2 || systemctl restart httpd
EOF
done

echo "Apache config sync complete."
