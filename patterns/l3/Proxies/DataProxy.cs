using l3.Interfaces;
using System.Runtime.Caching;

namespace l3.Proxies
{
    // Прокси-класс, который кэширует результаты и перехватывает вызовы к DataRepository
    public class DataProxy : IDataRepository
    {
        private readonly IDataRepository _realRepository;
        private readonly MemoryCache _cache;
        private readonly CacheItemPolicy _policy;

        public DataProxy(IDataRepository realRepository)
        {
            _realRepository = realRepository;
            _cache = MemoryCache.Default;

            // Данные будут истекать через 10 секунд
            _policy = new CacheItemPolicy
            {
                SlidingExpiration = TimeSpan.FromSeconds(10)
            };
        }

        public string GetData(string key)
        {
            // Сначала проверяем наличие данных в кэше
            var cachedValue = _cache.Get(key) as string;
            if (cachedValue != null)
            {
                Console.WriteLine(">>> Данные взяты из кэша");
                return cachedValue;
            }

            // Если нет – получаем данные у реального репозитория и сохраняем в кэше
            var realData = _realRepository.GetData(key);
            _cache.Set(key, realData, _policy);
            return realData;
        }
    }
}
