namespace BookInfo.Reviews.Models
{
    public class Review 
    {
        public int Id { get; set; }
        public int BookId { get; set; }
        public string Reviewer { get; set; }
        public System.DateTime ReviewDate { get; set; }
        public string ReviewText { get; set; }
    }

    
}