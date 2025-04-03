namespace l2.Models
{
    // Класс, инкапсулирующий обрабатываемые данные
    public class DataContext
    {
        public string InputData { get; set; }
        public bool IsValid { get; set; } = true;
    }
}
