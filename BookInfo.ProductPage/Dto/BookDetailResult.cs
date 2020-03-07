namespace BookInfo.ProductPage.Dto
{
    using System.Text.Json.Serialization;
    using System;

    public class BookDetailResult 
    {
        [JsonPropertyName("bookid")]
        public int BookId { get; set; }
        [JsonPropertyName("name")]
        public string Name { get; set; }
        [JsonPropertyName("isbn")]
        public string ISBN { get; set; }
        [JsonPropertyName("author")]
        public string Author { get; set; }
        [JsonPropertyName("publishdate")]
        public DateTime PublishDate { get; set; }
        [JsonPropertyName("price")]
        public decimal Price { get; set; }
        [JsonPropertyName("currentstock")]
        public decimal CurrentStock { get; set; }

    }
}