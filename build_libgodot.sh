#!/bin/bash

set -eux

BASE_DIR="$( cd "$(dirname "$0")" ; pwd -P )"

GODOT_DIR="$BASE_DIR/godot"
BUILD_DIR=$BASE_DIR/build

host_system="$(uname -s)"
host_arch="$(uname -m)"
host_target="template_release"
target=""
target_arch=""
lib_suffix="so"

case "$host_system" in
    Linux)
        host_platform="linuxbsd"
        cpus="$(nproc)"
        target_platform="linuxbsd"
    ;;
    Darwin)
        host_platform="macos"
        cpus="$(sysctl -n hw.logicalcpu)"
        target_platform="macos"
        lib_suffix="dylib"
    ;;
    *)
        echo "System $host_system is unsupported"
        exit 1
    ;;
esac

target="$host_target"
if [ "$target_arch" = "" ]
then
    target_arch="$host_arch"
fi

host_godot_suffix="$host_platform.$host_target"
host_godot_suffix="$host_godot_suffix.$host_arch"

target_godot_suffix="$target_platform.$target.$target_arch"
target_godot="$GODOT_DIR/bin/libgodot.$target_godot_suffix.$lib_suffix"

mkdir -p $BUILD_DIR
cd $GODOT_DIR
scons platform=$target_platform arch=$target_arch library_type=shared_library optimize=size target=template_release
cp -v $target_godot $BUILD_DIR/libgodot.$lib_suffix