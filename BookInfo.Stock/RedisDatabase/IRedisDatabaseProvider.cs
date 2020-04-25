using StackExchange.Redis;

namespace BookInfo.Stock.RedisDatabase
{
    public interface IRedisDatabaseProvider
    {
        IDatabase GetDatabase();
    }
}