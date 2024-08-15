@echo off
cd ../../ 
echo "===== build libgodot ====="
IF NOT EXIST build/libgodot.dll (
    build_libgodot.bat
)
cd samples/cs_sample
echo "===== clone godot-dotnet ====="
pwd
IF NOT EXIST godot-dotnet (
  git clone git@github.com:JiepengTan/godot-dotnet.git 
)
cd ../../


echo "===== export extension api ====="
cd ./godot
REM call scons platform=windows arch=x86_64 target=editor
REM bin/godot.windows.editor.x86_64.exe --dump-extension-api
REM copy /Y "..\..\godot\extension_api.json"  ".\godot-dotnet\gdextension\" 
REM copy /Y "..\..\godot\core\extension\gdextension_interface.h" ".\godot-dotnet\gdextension/"
cd ../

cd samples/cs_sample
echo "===== run demo ====="
run.bat 
