using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace l3.Interfaces
{
    // Интерфейс доступа к данным
    public interface IDataRepository
    {
        string GetData(string key);
    }
}
