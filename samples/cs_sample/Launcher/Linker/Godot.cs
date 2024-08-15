using System;
using System.Runtime.InteropServices;


public enum GDExtensionInitializationLevel
{
    GDEXTENSION_INITIALIZATION_CORE,
    GDEXTENSION_INITIALIZATION_SERVERS,
    GDEXTENSION_INITIALIZATION_SCENE,
    GDEXTENSION_INITIALIZATION_EDITOR,
    GDEXTENSION_MAX_INITIALIZATION_LEVEL
}

[StructLayout(LayoutKind.Sequential)]
public struct GDExtensionInitialization
{
    public GDExtensionInitializationLevel minimum_initialization_level;
    public IntPtr userdata;
    public Action<IntPtr, GDExtensionInitializationLevel> initialize;
    public Action<IntPtr, GDExtensionInitializationLevel> deinitialize;
}

public class LibGodot
{
    const string LIBGODOT_LIBRARY_NAME = "libgodot.dll";

    [DllImport(LIBGODOT_LIBRARY_NAME)]
    public static extern IntPtr libgodot_create_godot_instance(int argc, string[] argv, IntPtr handle);

    [DllImport(LIBGODOT_LIBRARY_NAME)]
    public static extern void libgodot_destroy_godot_instance(IntPtr instance);

    [DllImport(LIBGODOT_LIBRARY_NAME)]
    [return: MarshalAs(UnmanagedType.I1)]
    public static extern bool libgodot_start_godot_instance(IntPtr instance);

    [DllImport(LIBGODOT_LIBRARY_NAME)]
    [return: MarshalAs(UnmanagedType.I1)]
    public static extern bool libgodot_iteration_godot_instance(IntPtr instance);
}
