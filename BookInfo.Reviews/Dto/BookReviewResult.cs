namespace BookInfo.Reviews.Dto
{
    using System.Text.Json.Serialization;

    public class BookReviewResult 
    {
        [JsonPropertyName("bookId")]
        public int BookId { get; set; }
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
        [JsonPropertyName("reviewdate")]
        public System.DateTime ReviewDate { get; set; }
        [JsonPropertyName("reviewtext")]
        public string ReviewText { get; set; }

        public BookReview(Models.BookReview original) 
        {
            this.Reviewer = original.Reviewer;
            this.ReviewDate = original.ReviewDate;
            this.ReviewText = original.ReviewText;
        }

        public static explicit operator BookReview(Models.BookReview original) => new BookReview(original);
    }
    
}