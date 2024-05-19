#!/bin/bash

echo "Installing Watchlist CLI" 



HOME_DIR="$USERPROFILE"
build_exe="Watchlist-CLI.exe"
watch_script="watch"
config_yaml=".watchcli.yaml"

if [ ! -e "$build_exe" ]; then
    echo "No build found, run 'go build'"
    exit 1
fi

echo "..installing in $HOME_DIR\bin and $HOME_DIR\templates"

exe_path="$HOME_DIR\templates\\$build_exe"
script_path="$HOME_DIR\bin\\$watch_script"
config_path="$HOME_DIR\\$config_yaml"

if [ -e "$exe_path" ]; then
    echo "Removing old exe: $exe_path"
    rm "$exe_path"
fi

if [ -e "$script_path" ]; then
    echo "Removing old script: $script_path"
    rm "$script_path"
fi

if [ -e "$config_path" ]; then
    echo "Old configs found: $config_path"
    echo "Delete them and run installation again to install new configs"
else
    rm "$config_path"
    cp ".watchcli.yaml" "$config_path"
    echo "Added configs in $config_path"
fi


cp "$build_exe" "$exe_path"
echo "Added exe in $exe_path"
cp "$watch_script" "$script_path"
echo "Added script in $script_path"

