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
    }
])