#!/bin/bash

# Configuration
SNAPSHOT_DIR="/root/etcd-snapshots"
ETCD_DATA_DIR="/var/lib/etcd"
RESTORE_DATA_DIR="/var/lib/etcd-from-backup"
CERT_DIR="/etc/kubernetes/pki/etcd"

# Ensure snapshot directory exists
mkdir -p "$SNAPSHOT_DIR"

# Function to take snapshot
take_snapshot() {
    TIMESTAMP=$(date +%F-%H-%M-%S)
    SNAPSHOT_FILE="$SNAPSHOT_DIR/snapshot-$TIMESTAMP.db"

    echo "📦 Taking snapshot at: $SNAPSHOT_FILE"
    ETCDCTL_API=3 etcdctl snapshot save "$SNAPSHOT_FILE" \
        --endpoints=https://127.0.0.1:2379 \
        --cacert="$CERT_DIR/ca.crt" \
        --cert="$CERT_DIR/server.crt" \
        --key="$CERT_DIR/server.key"

    if [[ $? -eq 0 ]]; then
        echo "✅ Snapshot saved successfully."
    else
        echo "❌ Failed to take snapshot."
    fi
}

# Function to restore snapshot
restore_snapshot() {
    echo "📂 Available snapshots:"
    select SNAPSHOT in "$SNAPSHOT_DIR"/*.db; do
        if [[ -f "$SNAPSHOT" ]]; then
            echo "🔁 Restoring from: $SNAPSHOT"
            break
        else
            echo "❌ Invalid selection."
        fi
    done

    echo "🛑 Stopping kubelet and etcd..."
    systemctl stop kubelet
    docker stop $(docker ps -q --filter name=etcd) 2>/dev/null

    echo "🧹 Cleaning old restore dir and restoring snapshot..."
    rm -rf "$RESTORE_DATA_DIR"
    ETCDCTL_API=3 etcdctl snapshot restore "$SNAPSHOT" \
        --data-dir="$RESTORE_DATA_DIR"

    echo "✏️ Updating etcd manifest to point to restored data dir..."
    ETCD_MANIFEST="/etc/kubernetes/manifests/etcd.yaml"
    sed -i "s|--data-dir=.*|--data-dir=$RESTORE_DATA_DIR|g" "$ETCD_MANIFEST"

    echo "🚀 Restarting kubelet (etcd will auto-restart)..."
    systemctl start kubelet

    echo "✅ Cluster restored from snapshot."
}

# Main menu
echo "============================="
echo "   ETCD Snapshot Manager"
echo "============================="
echo "1️⃣  Take Snapshot"
echo "2️⃣  Restore Snapshot"
echo "3️⃣  Exit"
echo

read -p "Select an option [1-3]: " CHOICE

case $CHOICE in
    1)
        take_snapshot
        ;;
    2)
        restore_snapshot
        ;;
    3)
        echo "👋 Bye!"
        exit 0
        ;;
    *)
        echo "❌ Invalid choice."
        ;;
esac
