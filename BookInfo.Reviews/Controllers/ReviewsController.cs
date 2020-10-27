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
            _logger.LogInformation($"Fetching reviews for BookId:{bookId}");
            Dto.ReviewResult result = await GetRating(bookId);

            if (result == null)
            {
                _logger.LogError($"Can not retrieve ratings for BookId:{bookId}");
                return NoContent();
            }

            try
            {
                var reviews = _reviewContext.Reviews.Where(x => x.BookId == bookId);
                var reviewsDto = reviews.Map<Models.Review, Dto.Review>();
                result.Reviews = reviewsDto.ToArray();
                _logger.LogInformation($"Success fetching reviews for BookId:{bookId}");
                return Ok(result);
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Can not retrieve reviews for BookId:{bookId}");
                return NoContent();
            }
        }

        private async Task<Dto.ReviewResult> GetRating(int bookId)
        {
            _logger.LogInformation($"Fetching ratings for BookId:{bookId}");

            try
            {
                HttpClient client = new HttpClient();
                client.DefaultRequestHeaders.Accept.Clear();
                client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
                client.DefaultRequestHeaders.Add("User-Agent", ".NET Foundation Review Service");
                string serviceURL = System.Environment.GetEnvironmentVariable("RATING_URL") ?? "http://localhost:5112";
                serviceURL += "/ratings/" + bookId;
                var streamTask = client.GetStreamAsync(serviceURL);
                var result = await JsonSerializer.DeserializeAsync<Dto.ReviewResult>(await streamTask);
                _logger.LogInformation($"Success fetching ratings for BookId:{bookId}");
                return result;
            }
            catch (Exception ex)
            {
                _logger.LogError(ex, $"Ratings service does not return values for BookId:{bookId}");
                return null;
            }
        }
    }
}
