using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Elastic.CommonSchema.Serilog;
using Microsoft.AspNetCore.Http;
using Serilog;

namespace BookInfo.Stock
{
    public class Program
    {
        public static void Main(string[] args)
        {
            CreateHostBuilder(args).Build().Run();
        }

        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .UseSerilog((ctx, config) =>
                {
                    config.ReadFrom.Configuration(ctx.Configuration);
                    var formatterConfig = new EcsTextFormatterConfiguration();
                    formatterConfig.MapHttpContext(ctx.Configuration.Get<HttpContextAccessor>());
                    var formatter = new EcsTextFormatter(formatterConfig);
                    config.WriteTo.Console(formatter);
                })
                .ConfigureWebHostDefaults(webBuilder =>
                {
                    webBuilder.UseStartup<Startup>().UseUrls(Environment.GetEnvironmentVariable("BASE_URL") ?? "http://*:5114");
                });
    }
}
