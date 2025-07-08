#!/bin/bash
if [ "$#" -ne 2 ]; then
    echo "Usage: $0 <runner-number> <github-token>"
    exit 1
fi
RUNNER_NUMBER=$1
TOKEN=$2
RUNNER_DIR="/home/$USER/actions-runner-${RUNNER_NUMBER}"
SERVICE_FILE="/tmp/github-actions-${RUNNER_NUMBER}.service"

# Remove existing directory if exists
rm -rf "$RUNNER_DIR"

# Create fresh runner directory
mkdir -p "$RUNNER_DIR"
cd "$RUNNER_DIR" || exit

# Download and extract runner
curl -o actions-runner-linux-x64-2.321.0.tar.gz -L https://github.com/actions/runner/releases/download/v2.321.0/actions-runner-linux-x64-2.321.0.tar.gz
tar xzf ./actions-runner-linux-x64-2.321.0.tar.gz

# Configure runner
./config.sh --url https://github.com/REPLACE_GH_USERNAME/REPLACE_GH_REPO \
    --token "$TOKEN" \
    --name "runner-${RUNNER_NUMBER}" \
    --unattended \
    --replace

# Create service file in temp location
cat > "$SERVICE_FILE" << EOL
[Unit]
Description=GitHub Actions Runner ${RUNNER_NUMBER}
After=network.target

[Service]
Type=simple
User=${USER}
Group=${USER}
WorkingDirectory=${RUNNER_DIR}
ExecStart=${RUNNER_DIR}/run.sh
Restart=always
RestartSec=10
KillSignal=SIGTERM

[Install]
WantedBy=multi-user.target
EOL

# Move service file and configure systemd (these need sudo)
sudo mv "$SERVICE_FILE" "/etc/systemd/system/"
sudo systemctl daemon-reload
sudo systemctl enable "github-actions-${RUNNER_NUMBER}"
sudo systemctl start "github-actions-${RUNNER_NUMBER}"

echo "Runner ${RUNNER_NUMBER} has been created and started"