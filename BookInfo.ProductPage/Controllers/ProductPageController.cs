using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text.Json;

namespace BookInfo.ProductPage.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ProductPageController : ControllerBase
    {

        private readonly ILogger<ProductPageController> _logger;

        public ProductPageController(ILogger<ProductPageController> logger)
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
        public async Task<IActionResult> GetSingle(int bookId)
        {
            Dto.BookReviewResult result = await GetReview(bookId);
            return Ok(result);
            //return Ok(Data.BookReviews.Reviews.Where(c => c.BookId == bookId));
        }

        private async Task<Dto.BookReviewResult> GetReview(int bookId) 
        {
            HttpClient client = new HttpClient();
            client.DefaultRequestHeaders.Accept.Clear();
            client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
            client.DefaultRequestHeaders.Add("User-Agent", ".NET Foundation Product Page Service");
            string serviceURL = System.Environment.GetEnvironmentVariable("REVIEW_URL") ?? "http://localhost:5111";
            serviceURL += "/bookreview/" + bookId;
            var streamTask = client.GetStreamAsync(serviceURL);
            var result = await JsonSerializer.DeserializeAsync<Dto.BookReviewResult>(await streamTask);
            return result;       
        }
    }
}
