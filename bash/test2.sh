#!/bin/bash

# Define source and destination directories
SRC_DIR="/gpfs_svn/svn/repo/cc1_mng_1002"
DEST_DIR="/gpfs_svn/archival"
LINK_DIR="/var/www/svn/cc1_mng_1002"
LOG_FILE="/var/log/move_and_link.log"

# Ensure link directory exists
mkdir -p "$LINK_DIR"

# Check if test.txt exists
if [ ! -f test.txt ]; then
    echo "Error: test.txt not found!" | tee -a "$LOG_FILE"
    exit 1
fi

# Process each directory
while IFS= read -r dir; do
    # Skip empty lines
    [[ -z "$dir" ]] && continue

    if [ -d "$SRC_DIR/$dir" ]; then
        if [ -d "$DEST_DIR/$dir" ]; then
            echo "Warning: $DEST_DIR/$dir already exists. Skipping move." | tee -a "$LOG_FILE"
            continue
        fi

        # Move the directory
        mv "$SRC_DIR/$dir" "$DEST_DIR/"
        if [ $? -ne 0 ]; then
            echo "Error: Failed to move $SRC_DIR/$dir to $DEST_DIR/" | tee -a "$LOG_FILE"
            exit 1
        fi

        # Create symbolic link
        ln -s "$DEST_DIR/$dir" "$LINK_DIR/$dir"
        if [ $? -ne 0 ]; then
            echo "Error: Failed to create symlink $LINK_DIR/$dir -> $DEST_DIR/$dir" | tee -a "$LOG_FILE"
            exit 1
        fi

        echo "Successfully moved $dir to $DEST_DIR and created symlink in $LINK_DIR" | tee -a "$LOG_FILE"
    else
        echo "Warning: Directory $SRC_DIR/$dir does not exist, skipping." | tee -a "$LOG_FILE"
    fi
done < test.txt
