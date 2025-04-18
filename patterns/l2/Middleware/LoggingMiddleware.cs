﻿namespace l2.Middleware
{
    // Middleware для логирования
    public class LoggingMiddleware : Middleware
    {
        public override void Handle(Models.DataContext context)
        {
            Console.WriteLine($"LoggingMiddleware: Текущее состояние данных: \"{context.InputData}\", IsValid = {context.IsValid}");
            base.Handle(context);
        }
    }
}
