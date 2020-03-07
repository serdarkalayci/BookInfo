namespace BookInfo.Stock.Data
{
    using System.Collections.Generic;
    
    public static class BookStocks
    {
        public static List<Models.BookStock> Stocks = new List<Models.BookStock>() {
            new Models.BookStock() {
                BookID = 1,
                StockCount = 5
            },
            new Models.BookStock() {
                BookID = 2,
                StockCount = 4
            }
        };
    }
}