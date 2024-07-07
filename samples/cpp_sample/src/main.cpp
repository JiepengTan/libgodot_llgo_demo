#include <stdlib.h>
#include <stdio.h>
#if defined(__APPLE__) || defined(__unix__)
#include <dlfcn.h>
#elif defined(_WIN32)
#include <windows.h>
#endif
#include <string>
#include <vector>
#include <iostream>
extern "C"
{
    typedef enum
    {
        GDEXTENSION_INITIALIZATION_CORE,
        GDEXTENSION_INITIALIZATION_SERVERS,
        GDEXTENSION_INITIALIZATION_SCENE,
        GDEXTENSION_INITIALIZATION_EDITOR,
        GDEXTENSION_MAX_INITIALIZATION_LEVEL,
    } GDExtensionInitializationLevel;

    typedef struct
    {
        /* Minimum initialization level required.
         * If Core or Servers, the extension needs editor or game restart to take effect */
        GDExtensionInitializationLevel minimum_initialization_level;
        /* Up to the user to supply when initializing */
        void *userdata;
        /* This function will be called multiple times for each initialization level. */
        void (*initialize)(void *userdata, GDExtensionInitializationLevel p_level);
        void (*deinitialize)(void *userdata, GDExtensionInitializationLevel p_level);
    } GDExtensionInitialization;

    typedef unsigned char   uint8_t;
    typedef uint8_t GDExtensionBool;
    typedef void *GDExtensionObjectPtr;
    typedef void (*GDExtensionInterfaceFunctionPtr)();
    typedef void *GDExtensionClassLibraryPtr;
    typedef GDExtensionInterfaceFunctionPtr (*GDExtensionInterfaceGetProcAddress)(const char *p_function_name);
    typedef GDExtensionBool (*GDExtensionInitializationFunction)(GDExtensionInterfaceGetProcAddress p_get_proc_address,
                                                                 GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);
}
#define LIBGODOT_LIBRARY_NAME "libgodot.dll"
class LibGodot
{
public:
    LibGodot(std::string p_path = LIBGODOT_LIBRARY_NAME)
    {
        LPCSTR libgodot_library_name = reinterpret_cast<LPCSTR>(LIBGODOT_LIBRARY_NAME);
        handle = LoadLibrary(libgodot_library_name);
        if (handle == NULL)
        {
            fprintf(stderr, "Error opening alibgodot: %lu\n", GetLastError());
            return;
        }
        func_libgodot_create_godot_instance = (GDExtensionObjectPtr(*)(int, char *[], GDExtensionInitializationFunction, void *))GetProcAddress(handle, "libgodot_create_godot_instance");
        if (func_libgodot_create_godot_instance == NULL)
        {
            fprintf(stderr, "Error acquiring function: %lu\n", GetLastError());
            FreeLibrary(handle);
            return;
        }
        func_libgodot_destroy_godot_instance = (void (*)(GDExtensionObjectPtr))GetProcAddress(handle, "libgodot_destroy_godot_instance");
        func_libgodot_godot_instance_start = (bool (*)(GDExtensionObjectPtr))GetProcAddress(handle, "libgodot_godot_instance_start");
        func_libgodot_godot_instance_iteration = (bool (*)(GDExtensionObjectPtr))GetProcAddress(handle, "libgodot_godot_instance_iteration");
    }

    ~LibGodot()
    {
        if (is_open())
        {
            FreeLibrary(handle);
        }
    }

    bool is_open()
    {
        return handle != NULL && func_libgodot_create_godot_instance != NULL;
    }

    GDExtensionObjectPtr create_godot_instance(int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func = nullptr)
    {
        if (!is_open())
        {
            return nullptr;
        }
        GDExtensionObjectPtr instance = func_libgodot_create_godot_instance(p_argc, p_argv, p_init_func, handle);
        if (instance == nullptr)
        {
            return nullptr;
        }
        return instance;
    }

    void destroy_godot_instance(GDExtensionObjectPtr instance)
    {
        func_libgodot_destroy_godot_instance(instance);
    }
    bool start_godot_instance(GDExtensionObjectPtr instance)
    {
        return func_libgodot_godot_instance_start(instance);
    }
    bool iteration_godot_instance(GDExtensionObjectPtr instance)
    {
        return func_libgodot_godot_instance_iteration(instance);
    }

private:
    HINSTANCE handle = NULL;
    GDExtensionObjectPtr (*func_libgodot_create_godot_instance)(int, char *[], GDExtensionInitializationFunction, void *) = NULL;
    void (*func_libgodot_destroy_godot_instance)(GDExtensionObjectPtr) = NULL;
    bool (*func_libgodot_godot_instance_start)(GDExtensionObjectPtr) = NULL;
    bool (*func_libgodot_godot_instance_iteration)(GDExtensionObjectPtr) = NULL;
};

int main(int argc, char **argv)
{

    LibGodot libgodot;

    std::string program;
    if (argc > 0)
    {
        program = std::string(argv[0]);
    }
    std::vector<std::string> args = {program, "--path", "../../project/", "--rendering-method", "gl_compatibility", "--rendering-driver", "opengl3"};

    std::vector<char *> argvs;
    for (const auto &arg : args)
    {
        argvs.push_back((char *)arg.data());
    }
    argvs.push_back(nullptr);

    GDExtensionObjectPtr instance = libgodot.create_godot_instance(argvs.size(), argvs.data());
    if (instance == nullptr)
    {
        fprintf(stderr, "Error creating Godot instance\n");
        return EXIT_FAILURE;
    }
    bool succ = libgodot.start_godot_instance(instance);
    while (!libgodot.iteration_godot_instance(instance))
    {
    }
    libgodot.destroy_godot_instance(instance);

    return EXIT_SUCCESS;
}