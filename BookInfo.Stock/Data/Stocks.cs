using BookInfo.Stock.RedisDatabase;
namespace BookInfo.Stock.Data
{
    using System.Collections.Generic;
    
    public class BookStocks
    {
        public int GetStock(IRedisDatabaseProvider redisDatabaseProvider, int bookId)
        {
            // Get new score from redis and add to original score
            var db = redisDatabaseProvider.GetDatabase();
            var value = db.StringGet(bookId.ToString());
            int currentStock = 0;
            if (value != StackExchange.Redis.RedisValue.Null)
                currentStock = (int)value;
            return currentStock;
        } 
    }
}