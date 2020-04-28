using System.Text.Json.Serialization;

namespace BookInfo.Reviews.Dto
{
    public class ReviewResult 
    {
        [JsonPropertyName("bookId")]
        public int BookId { get; set; }
        [JsonPropertyName("currentRating")]
        public decimal Rating { get; set; }
        [JsonPropertyName("voteCount")]
        public int VoteCount { get; set; }
        [JsonPropertyName("reviews")]
        public Review[] Reviews { get; set; }
    }

    public struct Review
    {
        [JsonPropertyName("reviewer")]
        public string Reviewer { get; set; }
        [JsonPropertyName("reviewDate")]
        public System.DateTime ReviewDate { get; set; }
        [JsonPropertyName("reviewText")]
        public string ReviewText { get; set; }

        public Review(Models.Review original) 
        {
            this.Reviewer = original.Reviewer;
            this.ReviewDate = original.ReviewDate;
            this.ReviewText = original.ReviewText;
        }

        public static explicit operator Review(Models.Review original) => new Review(original);
    }
    
}