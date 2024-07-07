set BASE_DIR=%~dp0
cd %BASE_DIR%\godot
cd ..
md build
cd ./godot

REM Build libgodot
call scons platform=windows arch=x86_64 library_type=shared_library optimize=size target=template_release
copy /y bin\godot.windows.template_release.x86_64.dll ..\build\libgodot.dll

cd ..