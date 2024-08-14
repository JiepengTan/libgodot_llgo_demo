@echo off
cd ../../ 
build_libgodot.bat

cd ./godot
call scons platform=windows arch=x86_64 target=editor
bin/godot.windows.editor.x86_64.exe --dump-extension-api
cd ../

REM REM clone godot-dotnet
cd samples/csharp
git clone git@github.com:JiepengTan/godot-dotnet.git

REM export gdextension json
pwd
copy /Y "..\..\godot\extension_api.json"  ".\godot-dotnet\gdextension\" 
copy /Y "..\..\godot\core\extension\gdextension_interface.h" ".\godot-dotnet\gdextension/"

REM Build godot-dotnet
cd godot-dotnet
./build.cmd --productBuild --pushNupkgsLocal ./nugets --warnAsError false /p:GenerateGodotBindings=true

REM Build samples
cd ../Summator
dotnet publish Extension -r win-x64 -o Game/lib
