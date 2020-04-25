using StackExchange.Redis;
using System.Text;
using System;

namespace BookInfo.Stock.RedisDatabase 
{
    public class RedisDatabaseProvider : IRedisDatabaseProvider
    {
    
        private ConnectionMultiplexer _redisMultiplexer;
    
        public IDatabase GetDatabase()
        {
            if (_redisMultiplexer == null)
            {
                string redisAddress = Environment.GetEnvironmentVariable("RedisAddress");
                string redisPassword = Environment.GetEnvironmentVariable("RedisPassword");
                string databaseName = Environment.GetEnvironmentVariable("DatabaseName");
                StringBuilder sb = new StringBuilder();
                sb.AppendFormat("{0},defaultDatabase={1},password={2}", redisAddress, databaseName, redisPassword);
                _redisMultiplexer = ConnectionMultiplexer.Connect(sb.ToString());
            }
            return _redisMultiplexer.GetDatabase();
        }
    }
}