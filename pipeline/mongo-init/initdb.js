db.ratings.drop();
db.ratings.insertMany([
    {
        "bookId" : 1,
        "currentRating" : 8.0,
        "voteCount": 120
    },
    {
        "bookId" : 2,
        "currentRating" : 6.1,
        "voteCount": 85
    }]
);
db.details.insertMany([
    {
		"bookId":       1,
        "name":        "Lord of the Rings: The Fellowship of the Ring",
		"isbn":        "123AS123",
		"author":      "J.R.R. Tolkien",
		"publishDate": new Date(1954, 7, 29, 0, 0, 0),
		"price":       55.29,
    },
    {
		"bookId":      2,
        "name":        "Lord of the Rings: : The Two Towers",
		"isbn":        "123AS124",
		"author":      "J.R.R. Tolkien",
		"publishDate": new Date(1954, 11, 11, 0, 0, 0),
		"price":       55.29,
	}
]);