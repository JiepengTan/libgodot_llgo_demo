# libgodot llgo demo

this repo is base on the [libgodot](https://github.com/migeran/libgodot_project): 
## setup environment

```bash
pip install scons
pip install ninja 

git clone git@github.com:JiepengTan/libgodot_llgo_demo.git
cd libgodot_llgo_demo
git submodule update --init --progress --depth 1 godot
```

## windows
```bash

# build godot
build_libgodot.bat

# build demo
cd samples/cpp_sample
mkdir build
cmake -S . -B build -G Ninja
cd build
ninja

# run demo
sample.exe
```

## linux or mac
```bash
# build godot
./build_libgodot.sh

# build demo
cd samples/cpp_sample
mkdir -p build && cd build
cmake .. && make

# run demo
./sample
```

