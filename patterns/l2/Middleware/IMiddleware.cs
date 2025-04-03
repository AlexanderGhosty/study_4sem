using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace l2.Middleware
{
    public interface IMiddleware
    {
        IMiddleware SetNext(IMiddleware next);
        void Handle(Models.DataContext context);
    }
}
