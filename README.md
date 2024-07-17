# libgodot demo

this repo is base on the [libgodot](https://github.com/migeran/libgodot_project): 
## setup environment

```bash
pip install scons
pip install ninja 

git clone git@github.com:JiepengTan/libgodot_llgo_demo.git
cd libgodot_llgo_demo
git submodule update --init --progress --depth 1 godot
```

mac user should install [Vulkan SDK](https://sdk.lunarg.com/sdk/download/latest/mac/vulkan-sdk.dmg) first

## demos

1. [cpp example](/samples/cpp_sample/README.md)
2. [llgo example](/samples/llgo_sample/README.md)
