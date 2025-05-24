using Application.PluginsFactory;

           Console.WriteLine("context begin\n");

            var fA = new FPlugA("D:\\Programming\\study_4sem\\patterns\\Patterns_Laba_1\\Patterns_Laba_1\\bin\\Debug\\net8.0\\PluginA.dll");
            var fB = new FPlugB("D:\\Programming\\study_4sem\\patterns\\Patterns_Laba_1\\Patterns_Laba_1\\bin\\Debug\\net8.0\\PluginB.dll");

            var pluginA = fA.Create();
            var pluginB = fB.Create();

            Console.WriteLine();

            pluginA.Initialize();
            pluginA.Execute();
            pluginA.Terminate();

            Console.WriteLine();

            pluginB.Initialize();
            pluginB.Execute();
            pluginB.Terminate();

            Console.WriteLine("\ncontext end");