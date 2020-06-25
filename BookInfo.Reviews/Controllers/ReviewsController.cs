using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Extensions.Logging;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text.Json;
using BookInfo.Reviews.Data;
using Microsoft.EntityFrameworkCore;
using BatMap;

namespace BookInfo.Reviews.Controllers
{
    [ApiController]
    [Route("[controller]")]
    public class ReviewsController : ControllerBase
    {
        private readonly ILogger<ReviewsController> _logger;
        private readonly ReviewContext _reviewContext;

        public ReviewsController(ILogger<ReviewsController> logger, ReviewContext reviewContext)
        {
            _logger = logger;
            _reviewContext = reviewContext;
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
            Dto.ReviewResult result = await GetRating(bookId);
            var reviews = _reviewContext.Reviews.Where(x => x.BookId == bookId);
            var reviewsDto = reviews.Map<Models.Review, Dto.Review>();
            result.Reviews = reviewsDto.ToArray();
            return Ok(result);
        }

        private async Task<Dto.ReviewResult> GetRating(int bookId) 
        {
            HttpClient client = new HttpClient();
            client.DefaultRequestHeaders.Accept.Clear();
            client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
            client.DefaultRequestHeaders.Add("User-Agent", ".NET Foundation Review Service");
            string serviceURL = System.Environment.GetEnvironmentVariable("RATING_URL") ?? "http://localhost:5112";
            serviceURL += "/ratings/" + bookId;
            var streamTask = client.GetStreamAsync(serviceURL);
            var result = await JsonSerializer.DeserializeAsync<Dto.ReviewResult>(await streamTask);
            return result;        
        }
    }
}
