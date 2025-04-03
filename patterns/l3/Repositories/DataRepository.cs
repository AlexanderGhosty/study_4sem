using l3.Interfaces;

namespace l3.Repositories
{
    // Реальный класс-репозиторий, который "дорого" получает данные
    public class DataRepository : IDataRepository
    {
        public string GetData(string key)
        {
            // Имитируем дорогостоящую операцию (например, запрос к БД)
            Console.WriteLine(">>> Выполняем обращение к реальному репозиторию...");
            Thread.Sleep(2000);

            // Возвращаем данные с меткой времени для наглядности
            return $"Data for {key} (retrieved at {DateTime.Now:HH:mm:ss})";
        }
    }
}
