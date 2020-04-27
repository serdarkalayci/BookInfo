using System.Text.Json.Serialization;

namespace BookInfo.ProductPage.Dto
{
    public class ProductPageResponse
    {
        [JsonPropertyName("bookId")]
        public int BookId { get; set; }
        public BookDetailResult bookDetailResult { get; set; }
        public BookReviewResult bookReviewResult { get; set; }
    }
}