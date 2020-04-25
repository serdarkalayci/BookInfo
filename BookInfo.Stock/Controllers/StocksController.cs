using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using BookInfo.Stock.RedisDatabase;

namespace BookInfo.Stock.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class StocksController : ControllerBase
    {
        private readonly ILogger<StocksController> _logger;
        private readonly IRedisDatabaseProvider _redisDatabaseProvider;

        public StocksController(ILogger<StocksController> logger, IRedisDatabaseProvider redisDatabaseProvider)
        {
            _logger = logger;
            _redisDatabaseProvider = redisDatabaseProvider;
        }

        [HttpGet]
        [Route("")]
        public IActionResult GetAll()
        {
            return BadRequest();
        }

        [HttpGet]
        [Route("{bookId:int}")]
        public IActionResult GetSingle(int bookId)
        {
            Data.BookStocks stockData = new Data.BookStocks();
            int currentStock = stockData.GetStock(_redisDatabaseProvider, bookId);

            var result = new Dto.Stock() {
                CurrentStock = currentStock
            };
            return Ok(result);
        }
    }
}
