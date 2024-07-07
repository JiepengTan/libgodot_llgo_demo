set BASE_DIR=%~dp0
cd %BASE_DIR%\godot

REM Build editor and export extension API
REM call scons platform=windows arch=x86_64 msvc=yes optimize=size
REM start /wait bin\godot.windows.editor.x86_64.exe --dump-extension-api
REM copy /y extension_api.json ..\build\extension_api.json

REM Build libgodot
call scons platform=windows arch=x86_64 library_type=shared_library optimize=size target=template_release
copy /y bin\godot.windows.template_release.x86_64.dll ..\build\libgodot.dll

REM Copy extension API and interface to godot-cpp
copy /y ..\build\extension_api.json ..\godot-cpp\gdextension
copy /y core\extension\gdextension_interface.h ..\godot-cpp\gdextension

cd ..