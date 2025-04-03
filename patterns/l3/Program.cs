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

            while (true)
            {
                // Запрашиваем ключ у пользователя
                Console.Write("Введите ключ (или оставьте пустым для выхода): ");
                string key = Console.ReadLine();
                if (string.IsNullOrWhiteSpace(key))
                    break;

                // Первый запрос
                Console.WriteLine($"Запрашиваем {key} ...");
                Console.WriteLine(proxy.GetData(key));
                Console.WriteLine();
            }
        }
    }
}

