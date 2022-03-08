# go-pool-demo

# boot rabbitmq and redis services
docker-compose up -d 

# run app
go run ./

# configure your pool by editing .env file
vim .env

# test app performance (you can get wrk here https://github.com/wg/wrk)
 wrk -s json.lua -t2 -c2 -d30s http://127.0.0.1:8888
