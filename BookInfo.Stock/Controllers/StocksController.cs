using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text.Json;
using BookInfo.Stock.Data;
using BookInfo.Stock.Dto;

namespace BookInfo.Stock.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class StocksController : ControllerBase
    {
        private readonly ILogger<StocksController> _logger;

        public StocksController(ILogger<StocksController> logger)
        {
            _logger = logger;
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
            var stock = Data.BookStocks.Stocks.Where(c => c.BookID == bookId).FirstOrDefault();
            if (stock == null) 
            {
                return NotFound();
            }
            else
            {
                var result = new Dto.Stock() {
                    CurrentStock = stock.StockCount
                };
                return Ok(result);
            }
            
            //return Ok(Data.BookReviews.Reviews.Where(c => c.BookId == bookId));
        }
    }
}
