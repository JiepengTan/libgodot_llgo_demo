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

extern "C" {
typedef enum {
	GDEXTENSION_INITIALIZATION_CORE,
	GDEXTENSION_INITIALIZATION_SERVERS,
	GDEXTENSION_INITIALIZATION_SCENE,
	GDEXTENSION_INITIALIZATION_EDITOR,
	GDEXTENSION_MAX_INITIALIZATION_LEVEL,
} GDExtensionInitializationLevel;

typedef struct {
	/* Minimum initialization level required.
	 * If Core or Servers, the extension needs editor or game restart to take effect */
	GDExtensionInitializationLevel minimum_initialization_level;
	/* Up to the user to supply when initializing */
	void *userdata;
	/* This function will be called multiple times for each initialization level. */
	void (*initialize)(void *userdata, GDExtensionInitializationLevel p_level);
	void (*deinitialize)(void *userdata, GDExtensionInitializationLevel p_level);
} GDExtensionInitialization;

typedef unsigned char uint8_t;
typedef uint8_t GDExtensionBool;
typedef void *GDExtensionObjectPtr;
typedef void (*GDExtensionInterfaceFunctionPtr)();
typedef void *GDExtensionClassLibraryPtr;
typedef GDExtensionInterfaceFunctionPtr (*GDExtensionInterfaceGetProcAddress)(const char *p_function_name);
typedef GDExtensionBool (*GDExtensionInitializationFunction)(GDExtensionInterfaceGetProcAddress p_get_proc_address,
		GDExtensionClassLibraryPtr p_library, GDExtensionInitialization *r_initialization);
}

#ifdef __APPLE__
#define LIBGODOT_LIBRARY_NAME "libgodot.dylib"
#elif defined(__unix__)
#define LIBGODOT_LIBRARY_NAME "libgodot.so"
#elif defined(_WIN32)
#define LIBGODOT_LIBRARY_NAME "libgodot.dll"
#endif


class LibGodot {
private:
	HINSTANCE handle = NULL;
    void print_lib_error(const char* errorMessage) {
#if defined(__APPLE__) || defined(__unix__)
        fprintf(stderr, "%s %s\n", errorMessage,  dlerror());
#elif defined(_WIN32)
        fprintf(stderr, "%s %lu\n", errorMessage,  GetLastError());
#endif
    }
    void freelib() {
        if (handle != NULL) {
#if defined(__APPLE__) || defined(__unix__)
            dlclose(handle);
#elif defined(_WIN32)
            FreeLibrary(handle);
#endif
            handle = NULL;
        }
    }
    void loadlib(){
#if defined(__APPLE__) || defined(__unix__)
        handle = dlopen(LIBGODOT_LIBRARY_NAME, RTLD_LAZY);
#elif defined(_WIN32)
		LPCSTR libname = reinterpret_cast<LPCSTR>(LIBGODOT_LIBRARY_NAME);
		handle = LoadLibrary(libname);
#endif
		if (handle == NULL) {
			print_lib_error( "Error opening alibgodot:");
			return;
		}
    }
    void* getlibfunc(const char* funcname){
#if defined(__APPLE__) || defined(__unix__)
        return (void*)dlsym(handle, funcname);
#elif defined(_WIN32)
        return (void*)GetProcAddress(handle, funcname);
#endif
    }

public:
	LibGodot(std::string p_path = LIBGODOT_LIBRARY_NAME) {
        loadlib();
		load_symbols();
	}

	~LibGodot() {
		if (is_open()) {
			FreeLibrary(handle);
		}
	}

	bool is_open() {
		return handle != NULL && func_libgodot_create_godot_instance != NULL;
	}

	GDExtensionObjectPtr create_instance(int p_argc, char *p_argv[], GDExtensionInitializationFunction p_init_func) {
		return func_libgodot_create_godot_instance(p_argc, p_argv, p_init_func, handle);
	}

	void destroy(GDExtensionObjectPtr instance) {
		func_libgodot_destroy_godot_instance(instance);
	}

	bool start(GDExtensionObjectPtr instance) {
		return func_libgodot_start_godot_instance(instance);
	}

	bool iteration(GDExtensionObjectPtr instance) {
		return func_libgodot_iteration_godot_instance(instance);
	}

#define GET_PROC_ADDRESS_AND_CHECK(funcName)                                    \
    do {                                                                        \
        *(void**)(&func_##funcName) = getlibfunc(#funcName);                    \
        if (func_##funcName == NULL){                                           \
            print_lib_error("Error acquiring function: ");                      \
            freelib();                                                          \
            return;                                                             \
        }                                                                       \
    } while (0)

private:
	void load_symbols() {
		GET_PROC_ADDRESS_AND_CHECK(libgodot_create_godot_instance);
		GET_PROC_ADDRESS_AND_CHECK(libgodot_destroy_godot_instance);
		GET_PROC_ADDRESS_AND_CHECK(libgodot_start_godot_instance);
		GET_PROC_ADDRESS_AND_CHECK(libgodot_iteration_godot_instance);
	}

private:
	GDExtensionObjectPtr (*func_libgodot_create_godot_instance)(int, char *[], GDExtensionInitializationFunction, void *) = NULL;
	void (*func_libgodot_destroy_godot_instance)(GDExtensionObjectPtr) = NULL;
	bool (*func_libgodot_start_godot_instance)(GDExtensionObjectPtr) = NULL;
	bool (*func_libgodot_iteration_godot_instance)(GDExtensionObjectPtr) = NULL;
};

int main(int argc, char **argv) {
	LibGodot libgodot;

	std::string program;
	std::string project_path = "../../project/";
	if (argc > 0) {
		program = std::string(argv[0]);
	}
	if (argc > 1) {
		project_path = std::string(argv[1]);
	}
	std::vector<std::string> args = { program, "--path", project_path, "--rendering-method", "gl_compatibility", "--rendering-driver", "opengl3" };

	std::vector<char *> argvs;
	for (const auto &arg : args) {
		argvs.push_back((char *)arg.data());
	}
	argvs.push_back(nullptr);
	GDExtensionObjectPtr instance = libgodot.create_instance(argvs.size(), argvs.data(), nullptr);
	if (instance == nullptr) {
		fprintf(stderr, "Error creating Godot instance\n");
		return EXIT_FAILURE;
	}
	bool succ = libgodot.start(instance);
	while (!libgodot.iteration(instance)) {
	}
	libgodot.destroy(instance);
	return EXIT_SUCCESS;
}
