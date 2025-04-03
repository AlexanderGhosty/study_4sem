using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace l2.Middleware
{
    // Базовый класс Middleware, реализующий логику "передать следующему"
    public abstract class Middleware : IMiddleware
    {
        private IMiddleware _next;

        // Метод для связывания обработчиков в цепочку
        public IMiddleware SetNext(IMiddleware next)
        {
            _next = next;
            return next;  // Для возможности установки цепочки "по цепочке"
        }

        // Основной метод обработки
        public virtual void Handle(Models.DataContext context)
        {
            // Если текущий обработчик не прервал цепочку, вызываем следующий
            _next?.Handle(context);
        }
    }
}
