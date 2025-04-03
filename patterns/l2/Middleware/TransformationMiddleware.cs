using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace l2.Middleware
{
    // Middleware для трансформации данных (например, приведение строки к верхнему регистру)
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
