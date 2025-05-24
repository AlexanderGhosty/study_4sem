using Application.PluginsFactory;

Console.WriteLine("context begin\n");

var fA = new FactoryPlugA("D:\\Programming\\study_4sem\\patterns\\L1\\PluginA\\bin\\Debug\\net8.0\\PluginA.dll");
var fB = new FactoryPlugB("D:\\Programming\\study_4sem\\patterns\\L1\\PluginB\\obj\\Debug\\net8.0\\PluginB.dll");

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
