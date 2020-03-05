using Microsoft.AspNetCore.Mvc;
using System.Net.Http;
using System.Net.Http.Headers;
using System.Text.Encodings.Web;
using System.Threading.Tasks;

namespace HelloServerless.WebApp.Controllers
{
    public class WeatherController : Controller
    {
        private static readonly HttpClient client = new HttpClient();
        
        // 
        // GET: /Weather/
        public async Task<IActionResult> Index(string name, int numTimes = 1)
        {
            string weather = await ProcessWeather();
            ViewData["Message"] = "Hello" + name;
            ViewData["NumTimes"] = numTimes;
            ViewData["Weather"] = weather;
            return View();
        }

        // 
        // GET: /Weather/Welcome/ 
        public IActionResult Welcome(string name, int ID = 1)
        {
            ViewData["Message"] = "Hello" + name;
            ViewData["ID"] = ID;
            return View();
        }

        private async Task<string> ProcessWeather() 
        {
            client.DefaultRequestHeaders.Accept.Clear();
            client.DefaultRequestHeaders.Accept.Add(new MediaTypeWithQualityHeaderValue("application/json"));
            client.DefaultRequestHeaders.Add("User-Agent", ".NET Foundation Weather Client");
            string serviceURL = System.Environment.GetEnvironmentVariable("WEATHER_URL") ?? "http://localhost:5001";
            var stringTask = client.GetStringAsync(serviceURL + "/weatherforecast");

            var msg = await stringTask;
            return msg;
            //Console.Write(msg);            
        }
    }
}