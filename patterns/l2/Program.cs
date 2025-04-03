using System;
using l2.Middleware;
using l2.Models;

namespace l2
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.Write("Введите строку для обработки: ");
            string userInput = Console.ReadLine();

            // Создаём наш DataContext
            var context = new DataContext
            {
                InputData = userInput
            };

            // Создаём обработчики
            var validation = new ValidationMiddleware();
            var transform = new TransformationMiddleware();
            var logging = new LoggingMiddleware();

            // Строим цепочку: validation -> transform -> logging
            validation
                .SetNext(transform)
                .SetNext(logging);

            // Запускаем обработку с первого звена цепочки
            validation.Handle(context);

            Console.WriteLine("Обработка завершена.");
            Console.WriteLine($"Результат: \"{context.InputData}\", IsValid = {context.IsValid}");
            Console.ReadLine();
        }
    }
}
