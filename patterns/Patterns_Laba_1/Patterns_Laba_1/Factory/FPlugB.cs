using Patterns.PluginsFactory;

namespace Application.PluginsFactory;

public class FPlugB : Factory
{
    public FPlugB(string path) : base(path) { }

    public override IPlugin Create()
    {
        IPlugin plugin = LoadPlugin(path);
        Console.WriteLine("Do some business logic with plugin B");
        return plugin;
    }
}