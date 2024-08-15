
using System;
using System.Collections.Generic;
using System.Linq;

public class Program
{
    public static void Main(string[] args)
    {
        Console.WriteLine("Hello Libgodot-csharp ! ");
        string program = "";
        string projectPath = "../../Summator/Game";
        if (args.Length > 0)
        {
            projectPath = args[0];
            Console.WriteLine("Project path: " + projectPath);
        }
        List<string> arguments = new List<string> { program, "--path", projectPath, "--rendering-method", "gl_compatibility", "--rendering-driver", "opengl3" };

        IntPtr instance = LibGodot.libgodot_create_godot_instance(arguments.Count, arguments.ToArray(), IntPtr.Zero);
        if (instance == IntPtr.Zero)
        {
            Console.Error.WriteLine("Error creating Godot instance");
            Environment.Exit(1);
        }
        bool success = LibGodot.libgodot_start_godot_instance(instance);
        while (LibGodot.libgodot_iteration_godot_instance(instance))
        {
        }
        LibGodot.libgodot_destroy_godot_instance(instance);
        Environment.Exit(0);
    }
}
