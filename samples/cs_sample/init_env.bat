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
IF NOT EXIST bin/godot.windows.editor.x86_64.exe (
  call scons platform=windows arch=x86_64 target=editor
)
call "bin/godot.windows.editor.x86_64.exe" --dump-extension-api
copy /Y "extension_api.json"  "..\samples\cs_sample\godot-dotnet\gdextension\" 
copy /Y "core\extension\gdextension_interface.h" "..\samples\cs_sample\godot-dotnet\gdextension/"
cd ../

cd samples/cs_sample
echo "===== run demo ====="
cd godot-dotnet
taskkill /IM dotnet.exe /F
build.cmd --productBuild --pushNupkgsLocal ./nugets --warnAsError false /p:GenerateGodotBindings=true
cd ../
