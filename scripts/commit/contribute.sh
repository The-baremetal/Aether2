#!/bin/sh
set -e
echo "Aether contribution script"
echo "DO NOT PROCEED UNLESS YOU KNOW WHAT YOU ARE DOING! First it's recommended to read the script to prevent surprises. You may regret what you did later on. It is recommended to make a complete copy of this folder before proceeding."
echo "Create a backup of your current work? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  echo "Creating backup..."
  cp -r . ../aether_backup_$(date +%Y%m%d_%H%M%S)
  echo "Backup created successfully."
else
  echo "No backup created. Proceed with caution!"
fi
echo "Add all untracked files to track list? This is recommended as a measure to prevent you from destroying your work later on. If you stash when you say no to this prompt, it may cause confusion and you may end up deleting your stash. It is an overall good practice to add your untracked files. (y/n)"
read ans
if [ "$ans" = "y" ]; then
  git add .
fi
echo "Stash uncommitted changes before proceeding? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  git stash
fi
echo "Run make publish? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  make publish
fi
echo "Run tests? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  make test
fi
echo "Move .tar files to build/bin? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  mv *.tar build/bin/ 2>/dev/null || true
fi
echo "Stage all changes and commit? (y/n)"
read ans
if [ "$ans" = "y" ]; then
  git add .
  echo "Enter commit message:"
  read msg
  git commit -m "$msg"
  if command -v gh >/dev/null 2>&1; then
    echo "Open a pull request? (y/n)"
    read ans
    if [ "$ans" = "y" ]; then
      gh pr create --fill --title "$msg"
      echo "PR created successfully."
    else
      echo "Your code has been committed."
    fi
  else
    echo "Your code has been committed. Install GitHub CLI (gh) to open PRs automatically."
  fi
else
  echo "No commit made."
fi 