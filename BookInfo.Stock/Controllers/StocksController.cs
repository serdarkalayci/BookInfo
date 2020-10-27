using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using BookInfo.Stock.RedisDatabase;
using System;

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
            _logger.LogInformation($"Fetching stocks for BookId:{bookId}");
            Data.BookStocks stockData = new Data.BookStocks();
            try
            {
                int currentStock = stockData.GetStock(_redisDatabaseProvider, bookId);
                var result = new Dto.Stock()
                {
                    CurrentStock = currentStock
                };
                _logger.LogInformation($"Success fetching stocks for BookId:{bookId}");
                return Ok(result);
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Can not retrieve stocks for BookId:{bookId}");
                return NoContent();
            }
        }
    }
}
