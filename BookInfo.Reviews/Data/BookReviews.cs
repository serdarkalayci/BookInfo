namespace BookInfo.Reviews.Data
{
    using System.Collections.Generic;
    
    public static class BookReviews 
    {
        public static IEnumerable<Models.BookReview> Reviews = new List<Models.BookReview> {
            new Models.BookReview() {
                BookId = 1,
                Reviewer = "Serdar Kalaycı",
                ReviewDate =  System.DateTime.Parse("2019-08-18T07:22:16.0000000+02:00"),
                ReviewText = "Çok güzel bir kitap"
            },
            new Models.BookReview() {
                BookId = 1,
                Reviewer = "Mahmut Tuncer",
                ReviewDate =  System.DateTime.Parse("2019-09-10T07:11:13.0000000+02:00"),
                ReviewText = "Ben hiç beğenmedim"
            }
        };

        
    }
}