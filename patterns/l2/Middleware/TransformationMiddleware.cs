namespace l2.Middleware
{
    // Middleware для трансформации данных 
    public class TransformationMiddleware : Middleware
    {
        public override void Handle(Models.DataContext context)
        {
            if (context.IsValid)
            {
                Console.WriteLine("TransformationMiddleware: Преобразуем данные в верхний регистр...");
                context.InputData = context.InputData.ToUpper();
            }

            base.Handle(context);
        }
    }
}
