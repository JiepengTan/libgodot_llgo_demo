@echo off

echo "===== build godot-dotnet ====="
cd godot-dotnet
IF NOT EXIST artifacts (
./build.cmd --productBuild --pushNupkgsLocal ./nugets --warnAsError false /p:GenerateGodotBindings=true
)

echo "===== build samples ====="
rd /s /q artifacts\bin\Godot.BindingsGenerator\
cd ../Summator
dotnet publish Extension -r win-x64 -o Game/lib

echo "===== init projects ====="
cd Game
md .godot
copy /Y "extension_list.cfg"  ".godot"
cd ../

echo "===== run demo ====="
cd ../
call "../../godot/bin/godot.windows.editor.x86_64.exe" --path Summator/Game