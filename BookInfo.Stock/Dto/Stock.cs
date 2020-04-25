namespace BookInfo.Stock.Dto
{
    using System.Text.Json.Serialization;
    public class Stock
    {
        [JsonPropertyName("currentstock")]
        public int CurrentStock { get; set; }
    }
}