# go-pool-demo

Runs a web server, publishes sent json messages to rabbitmq and persists them to redis with unique keys. Check out https://github.com/serqol/go-pool/blob/master/pool.go to write your own connectors. Here we only use default connections,  but you can manage multiple connection configurations within your project. You can use rabbitmq admin panel and redis stats to verify that configuration is applied.

// boot rabbitmq and redis services

docker-compose up -d 

// configure your pool by editing .env file

vim .env

// run app

go run ./

// test app performance (you can get wrk here https://github.com/wg/wrk)

 wrk -s json.lua -t2 -c2 -d30s http://127.0.0.1:8888
