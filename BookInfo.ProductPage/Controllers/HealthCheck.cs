using System.Net;
using System.Net.Http;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.Extensions.Diagnostics.HealthChecks;

namespace BookInfo.ProductPage.Controllers
{
    public class ReviewsHealthCheck : IHealthCheck
    {
        public async Task<HealthCheckResult> CheckHealthAsync(HealthCheckContext context, CancellationToken cancellationToken = default(CancellationToken))
        {
            string reviewApiUri = (System.Environment.GetEnvironmentVariable("REVIEW_URL") ?? "http://localhost:5111") + "/readiness";
            HttpClient client = new HttpClient();
            HttpResponseMessage response = await client.GetAsync(reviewApiUri);
            if (response.IsSuccessStatusCode)
            {
                return HealthCheckResult.Healthy();
            }
            else
            {
                return HealthCheckResult.Unhealthy();
            }
        }
    }

    public class DetailsHealthCheck : IHealthCheck
    {
        public async Task<HealthCheckResult> CheckHealthAsync(HealthCheckContext context, CancellationToken cancellationToken = default(CancellationToken))
        {
            string detailsApiUri = (System.Environment.GetEnvironmentVariable("DETAIL_URL") ?? "http://localhost:5113") + "/health/ready";
            HttpClient client = new HttpClient();
            HttpResponseMessage response = await client.GetAsync(detailsApiUri);
            if (response.IsSuccessStatusCode)
            {
                return HealthCheckResult.Healthy();
            }
            else
            {
                return HealthCheckResult.Unhealthy();
            }
        }
    }
}