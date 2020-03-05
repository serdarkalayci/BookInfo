namespace BookInfo.ProductPage.Dto
{
    using System.Text.Json.Serialization;

    public class BookReviewResult 
    {
        [JsonPropertyName("bookid")]
        public int BookId { get; set; }
        [JsonPropertyName("rating")]
        public decimal Rating { get; set; }
        [JsonPropertyName("votecount")]
        public int VoteCount { get; set; }
        [JsonPropertyName("reviews")]
        public BookReview[] Reviews { get; set; }
    }

    public struct BookReview
    {
        [JsonPropertyName("reviewer")]
        public string Reviewer { get; set; }
        [JsonPropertyName("reviewdate")]
        public System.DateTime ReviewDate { get; set; }
        [JsonPropertyName("reviewtext")]
        public string ReviewText { get; set; }
    }
    
}