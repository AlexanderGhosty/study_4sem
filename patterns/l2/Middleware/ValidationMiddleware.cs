namespace l2.Middleware
{
    // Middleware для валидации входных данных
    public class ValidationMiddleware : Middleware
    {
        public override void Handle(Models.DataContext context)
        {
            Console.WriteLine("ValidationMiddleware: Проверяем входные данные...");

            if (string.IsNullOrWhiteSpace(context.InputData))
            {
                Console.WriteLine("ValidationMiddleware: Данные невалидны (пустая строка). Прерываем цепочку.");
                context.IsValid = false;
                return; // Прерываем цепочку, не вызывая следующий обработчик
            }

            base.Handle(context);
        }
    }
}
