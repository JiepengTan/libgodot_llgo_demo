
# build godot
../../build_libgodot.bat

# build demo
mkdir build
cmake -S . -B build -G Ninja
cd build
ninja

# run demo
sample.exe