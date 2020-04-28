using System.Collections.Generic;
using Microsoft.EntityFrameworkCore;

namespace BookInfo.Reviews.Data
{
    
    public class ReviewContext : DbContext
    {
        public ReviewContext(DbContextOptions<ReviewContext> options)
            : base(options)
        {
        }
        public DbSet<Models.Review> Reviews { get; set; }
        // protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        //     => optionsBuilder.UseNpgsql(Environment.GetEnvironmentVariable("ReviewConnStr"));
    }
}