namespace BookInfo.Stock.Dto
{
    using System.Text.Json.Serialization;
    public class Stock
    {
        [JsonPropertyName("currentStock")]
        public int CurrentStock { get; set; }
    }
}