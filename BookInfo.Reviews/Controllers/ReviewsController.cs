using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text.Json;

namespace BookInfo.Reviews.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ReviewsController : ControllerBase
    {
        private readonly ILogger<ReviewsController> _logger;

        public ReviewsController(ILogger<ReviewsController> logger)
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
            Dto.BookReviewResult result = await GetRating(bookId);
            List<Dto.BookReview> reviews = new List<Dto.BookReview>();
            foreach (var review in Data.BookReviews.Reviews.Where(c => c.BookId == bookId))
            {
                reviews.Add((Dto.BookReview)review);
            }
            result.Reviews = reviews.ToArray();
            return Ok(result);
            //return Ok(Data.BookReviews.Reviews.Where(c => c.BookId == bookId));
        }

        private async Task<Dto.BookReviewResult> GetRating(int bookId) 
        {
            HttpClient client = new HttpClient();
            client.DefaultRequestHeaders.Accept.Clear();
            client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
            client.DefaultRequestHeaders.Add("User-Agent", ".NET Foundation Review Service");
            string serviceURL = System.Environment.GetEnvironmentVariable("RATING_URL") ?? "http://localhost:5112";
            serviceURL += "/ratings/" + bookId;
            var streamTask = client.GetStreamAsync(serviceURL);
            var result = await JsonSerializer.DeserializeAsync<Dto.BookReviewResult>(await streamTask);
            return result;        
        }
    }
}
