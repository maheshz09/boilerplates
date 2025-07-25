#!/bin/bash

echo "🔍 Scanning for submodules..."

# Step 1: Check if .gitmodules exists
if [ ! -f .gitmodules ]; then
    echo "✅ No submodules found (.gitmodules missing)"
    exit 0
fi

# Step 2: Extract submodule paths
submodules=$(git config -f .gitmodules --get-regexp path | awk '{ print $2 }')

# Step 3: Loop through each submodule and remove it
for path in $submodules; do
    echo "🔥 Removing submodule: $path"

    # Remove from Git index
    git rm --cached "$path"

    # Remove the actual folder
    rm -rf "$path"

    # Remove any lingering submodule config
    git config -f .git/config --remove-section submodule."$path" 2>/dev/null
done

# Step 4: Delete .gitmodules
rm -f .gitmodules

# Step 5: Final commit and push
echo "📦 Committing changes..."
git add .
git commit -m "🔥 Removed all submodules completely"
git push

echo "✅ All submodules removed and repo cleaned successfully!"

