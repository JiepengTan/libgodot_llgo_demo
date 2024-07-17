# build godot
./../../build_libgodot.sh

# build demo
mkdir -p build && cd build
cmake .. && make

# run demo
./sample