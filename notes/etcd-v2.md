#!/bin/bash

# Configuration
SNAPSHOT_DIR="/root/etcd-snapshots"
ETCD_DATA_DIR="/var/lib/etcd"
RESTORE_DATA_DIR="/var/lib/etcd-from-backup"
CERT_DIR="/etc/kubernetes/pki/etcd"
ETCD_MANIFEST="/etc/kubernetes/manifests/etcd.yaml"

# Ensure snapshot directory exists
mkdir -p "$SNAPSHOT_DIR"

# Detect container runtime (docker or containerd)
detect_runtime() {
    if command -v docker &>/dev/null; then
        RUNTIME="docker"
    elif command -v crictl &>/dev/null; then
        RUNTIME="containerd"
    else
        echo "❌ Neither Docker nor crictl found. Cannot stop etcd container."
        exit 1
    fi
}

# Check if etcdctl is available
check_etcdctl() {
    if ! command -v etcdctl &>/dev/null; then
        echo "❌ etcdctl not found in PATH."
        echo "👉 Please install it: https://github.com/etcd-io/etcd/releases"
        exit 1
    fi
}

# Stop etcd container
stop_etcd_container() {
    echo "🛑 Stopping kubelet and etcd..."
    systemctl stop kubelet

    if [ "$RUNTIME" == "docker" ]; then
        docker ps --filter name=etcd -q | xargs -r docker stop
    else
        crictl ps | grep etcd | awk '{print $1}' | xargs -r crictl stop
    fi
}

# Take etcd snapshot
take_snapshot() {
    check_etcdctl
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
        echo "❌ Snapshot failed."
    fi
}

# Restore etcd snapshot
restore_snapshot() {
    check_etcdctl

    echo "📂 Available snapshots:"
    select SNAPSHOT in "$SNAPSHOT_DIR"/*.db; do
        if [[ -f "$SNAPSHOT" ]]; then
            echo "🔁 Restoring from: $SNAPSHOT"
            break
        else
            echo "❌ Invalid selection."
        fi
    done

    stop_etcd_container

    echo "🧹 Cleaning and restoring snapshot..."
    rm -rf "$RESTORE_DATA_DIR"
    ETCDCTL_API=3 etcdctl snapshot restore "$SNAPSHOT" \
        --data-dir="$RESTORE_DATA_DIR"

    echo "✏️ Updating etcd manifest to use restored data..."
    sed -i "s|--data-dir=.*|--data-dir=$RESTORE_DATA_DIR|g" "$ETCD_MANIFEST"

    echo "🚀 Restarting kubelet..."
    systemctl start kubelet

    echo "✅ Restore process complete. Wait a few seconds and run:"
    echo "   kubectl get pods -A"
}

# Main Menu
echo "============================="
echo "   🚀 ETCD Snapshot Manager"
echo "============================="
echo "1️⃣  Take Snapshot"
echo "2️⃣  Restore Snapshot"
echo "3️⃣  Exit"
echo

read -p "Select an option [1-3]: " CHOICE

detect_runtime

case $CHOICE in
    1)
        take_snapshot
        ;;
    2)
        restore_snapshot
        ;;
    3)
        echo "👋 Exiting."
        exit 0
        ;;
    *)
        echo "❌ Invalid choice."
        ;;
esac
