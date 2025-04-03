using l3.Interfaces;
using l3.Proxies;
using l3.Repositories;

namespace l3
{
    public class Program
    {
        public static void Main()
        {
            // Создаем реальный репозиторий
            IDataRepository repo = new DataRepository();

            // Создаем прокси, передавая туда реальный репозиторий
            IDataRepository proxy = new DataProxy(repo);

            // Первый запрос
            Console.WriteLine("Запрашиваем key1 первый раз...");
            Console.WriteLine(proxy.GetData("key1"));
            Console.WriteLine();

            // Второй запрос (ожидается, что данные будут из кэша)
            Console.WriteLine("Запрашиваем key1 второй раз (должно взяться из кэша)...");
            Console.WriteLine(proxy.GetData("key1"));
            Console.WriteLine();

            // Ожидание для истечения срока действия кэша (10 секунд + небольшой запас)
            Console.WriteLine("Ждем 12 секунд, чтобы кэш протух...");
            Thread.Sleep(12000);

            // Третий запрос – кэш уже истек, будет вызов реального репозитория
            Console.WriteLine("Снова запрашиваем key1 (должен опять обратиться к репозиторию)...");
            Console.WriteLine(proxy.GetData("key1"));
            Console.WriteLine("Нажмите любую клавишу для выхода...");
            Console.ReadKey();
        }
    }
}

