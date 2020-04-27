using System.Text.Json.Serialization;

namespace BookInfo.ProductPage.Dto
{
    public class BookReviewResult 
    {
        [JsonPropertyName("currentRating")]
        public decimal Rating { get; set; }
        [JsonPropertyName("voteCount")]
        public int VoteCount { get; set; }
        [JsonPropertyName("reviews")]
        public BookReview[] Reviews { get; set; }
    }

    public struct BookReview
    {
        [JsonPropertyName("reviewer")]
        public string Reviewer { get; set; }
        [JsonPropertyName("reviewDate")]
        public System.DateTime ReviewDate { get; set; }
        [JsonPropertyName("reviewText")]
        public string ReviewText { get; set; }
    }
    
}