All database instances needed by this project can be started by typing following command in terminal (assuming you have Docker running)

`docker-compose -f db-compose.yaml -p BookInfoDb up -d`

and all applications including Jaeger, Prometheus and Grafana by typing

`docker-compose -f docker-compose.yaml -p BookInfo up -d`

The Redis databases by default are arranged as follows:

0: Rating

1: Stock

Redis databases can be initialized by connecting to the container's console by 

`docker exec -it bookinfo-redis bash`

and connecting Rating database through redis-cli by entering

`redis-cli -u redis://127.0.0.1:6379/0`

and inserting Rating records
```
HMSET bookId:1 currentRating 9.79 voteCount 19785
HMSET bookId:2 currentRating 9.82 voteCount 21562
```
disconnect from Rating database by typing `exit` and connect Stock database by typing

`redis-cli -u redis://127.0.0.1:6379/1`

and inserting Stock records
```
SET "1" "127"
SET "2" "85"
```
