#!/bin/bash

current_dir="$( cd "$(dirname "$0")" ; pwd -P )"
if [[ "$1" == "-b" ]]; then
    cd ../../
    ./build_libgodot.sh
    cd $current_dir
fi
build_dir=${current_dir}/build
mkdir -p build
cp ../../build/libgodot.so build
cp config/godot.pc build
sed -i "1s@^prefix=.*@prefix=${build_dir}/@" build/godot.pc


export PKG_CONFIG_PATH="$build_dir:$PKG_CONFIG_PATH"
export LD_LIBRARY_PATH="$build_dir:$LD_LIBRARY_PATH"

nm -D build/libgodot.so > build/api_dump.txt

cd logic
llgo run .
cd ..