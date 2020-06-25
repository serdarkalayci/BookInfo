using System;
using Microsoft.AspNetCore.Builder;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Configuration;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.OpenApi.Models;
using OpenTracing;
using OpenTracing.Util;
using Prometheus;
using BookInfo.Stock.RedisDatabase;


namespace BookInfo.Stock
{
    public class Startup
    {
        public Startup(IConfiguration configuration)
        {
            Configuration = configuration;
        }

        public IConfiguration Configuration { get; }

        // This method gets called by the runtime. Use this method to add services to the container.
        public void ConfigureServices(IServiceCollection services)
        {
            services.AddControllers();
            // Redis database service
            if (Environment.GetEnvironmentVariable("RedisAddress") == null) 
                    Environment.SetEnvironmentVariable("RedisAddress", "127.0.0.1:6379");
            if (Environment.GetEnvironmentVariable("RedisPassword") == null) 
                    Environment.SetEnvironmentVariable("RedisPassword", "");
            if (Environment.GetEnvironmentVariable("DatabaseName") == null) 
                    Environment.SetEnvironmentVariable("DatabaseName", "1");
            services.AddSingleton<IRedisDatabaseProvider, RedisDatabaseProvider>();

            // Swagger
            services.AddSwaggerGen(c =>
            {
                c.SwaggerDoc("v1", new OpenApiInfo { Title = "Book Stocks API", Version = "v1" });
            });

            // OpenTracing
            services.AddOpenTracing();

            // Adds the Jaeger Tracer.
            services.AddSingleton<ITracer>(serviceProvider =>
            {
                string serviceName = serviceProvider.GetRequiredService<Microsoft.AspNetCore.Hosting.IWebHostEnvironment>().ApplicationName;

                if (Environment.GetEnvironmentVariable("JAEGER_SERVICE_NAME") == null) 
                    Environment.SetEnvironmentVariable("JAEGER_SERVICE_NAME", serviceName);
                if (Environment.GetEnvironmentVariable("JAEGER_AGENT_HOST") == null) 
                    Environment.SetEnvironmentVariable("JAEGER_AGENT_HOST", "localhost");                
                if (Environment.GetEnvironmentVariable("JAEGER_AGENT_PORT") == null) 
                    Environment.SetEnvironmentVariable("JAEGER_AGENT_PORT", "6831");                
                if (Environment.GetEnvironmentVariable("JAEGER_SAMPLER_TYPE") == null) 
                    Environment.SetEnvironmentVariable("JAEGER_SAMPLER_TYPE", "const");

                var loggerFactory = new LoggerFactory();

                var config = Jaeger.Configuration.FromEnv(loggerFactory);
                var tracer = config.GetTracer();

                if (!GlobalTracer.IsRegistered())
                {
                    // Allows code that can't use DI to also access the tracer.
                    GlobalTracer.Register(tracer);
                }

                return tracer;
            }); 
        }

        // This method gets called by the runtime. Use this method to configure the HTTP request pipeline.
        public void Configure(IApplicationBuilder app, IWebHostEnvironment env)
        {
            if (env.IsDevelopment())
            {
                app.UseDeveloperExceptionPage();
            }

            
            // Enable middleware to serve generated Swagger as a JSON endpoint.
            app.UseSwagger();

            // Enable middleware to serve swagger-ui (HTML, JS, CSS, etc.),
            // specifying the Swagger JSON endpoint.
            app.UseSwaggerUI(c =>
            {
                c.SwaggerEndpoint("/swagger/v1/swagger.json", "Book Stocks V1");
            });

            app.UseRouting();

            app.UseAuthorization();

            app.UseEndpoints(endpoints =>
            {
                endpoints.MapControllers();
                endpoints.MapMetrics();
            });
        }
    }
}
