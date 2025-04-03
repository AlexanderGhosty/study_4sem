namespace l2.Middleware
{
    public interface IMiddleware
    {
        IMiddleware SetNext(IMiddleware next);
        void Handle(Models.DataContext context);
    }
}
