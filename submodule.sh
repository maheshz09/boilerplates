#!/bin/bash

# Make sure you're in the root of your git repo
echo "Cleaning all submodules..."

# Step 1: Loop through submodules listed in .gitmodules
git config -f .gitmodules --get-regexp path | while read key path; do
  echo "Processing $path ..."

  # Remove the submodule from Git index
  git rm --cached "$path"

  # Clone the actual repo into the folder
  repo_url=$(git config -f .gitmodules --get submodule."$path".url)
  rm -rf "$path"
  git clone "$repo_url" "$path"

  # Remove the .git folder from the sub-cloned repo
  rm -rf "$path/.git"
done

# Step 2: Clean up .gitmodules
rm -f .gitmodules

# Step 3: Add everything back and commit
git add .
git commit -m "Converted all submodules to regular folders"
git push

echo "✅ All submodules converted to regular folders."

