# go-pool-demo

// boot rabbitmq and redis services

docker-compose up -d 

// configure your pool by editing .env file

vim .env

// run app

go run ./

// test app performance (you can get wrk here https://github.com/wg/wrk)

 wrk -s json.lua -t8 -c200 -d30s http://127.0.0.1:8888
