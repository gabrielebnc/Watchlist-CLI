#!/bin/bash

echo "Installing Watchlist CLI" 



HOME_DIR="$USERPROFILE"
build_exe="Watchlist-CLI.exe"
watch_script="watch"

if [ ! -e "$build_exe" ]; then
    echo "No build found, run 'go build'"
    exit 1
fi

echo "..installing in $HOME_DIR\bin and $HOME_DIR\templates"

exe_path="$HOME_DIR\templates\\$build_exe"
script_path="$HOME_DIR\bin\\$watch_script"

if [ -e "$exe_path" ]; then
    echo "Removing old exe: $exe_path"
    rm "$exe_path"
fi

if [ -e "$script_path" ]; then
    echo "Removing old script: $script_path"
    rm "$script_path"
fi


cp "$build_exe" "$exe_path"
echo "Added exe in $exe_path"
cp "$watch_script" "$script_path"
echo "Added script in $script_path"

