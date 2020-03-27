FROM mcr.microsoft.com/dotnet/core/sdk:3.1 AS build-env
WORKDIR /app

COPY . ./
RUN dotnet restore 
RUN dotnet publish -c Release -o out

FROM mcr.microsoft.com/dotnet/core/aspnet:3.1
WORKDIR /app
ENV BASE_URL http://+:5114
ENV TZ=Europe/Istanbul
COPY --from=build-env /app/out .

ENTRYPOINT ["dotnet", "BookInfo.Stock.dll"]