build: setup tidy
	@echo "--- Building binary file ---"
	@go build -o ./main server/grpc/main.go

grpc:
	@echo "--- running gRPC server in dev mode ---"
	@go run server/grpc/main.go

tidy:
	@go mod tidy

setup:
	@echo " --- Setup and generate configuration --- "
	@cp config/example/mysql.yml.example config/db/mysql.yml
	@cp config/example/server.yml.example config/server/server.yml

build-docker: build
	@docker build --tag bussiness-logic-services .

protoc-docker: build
	@docker container create --name bl-services -p 9901:9901 bussiness-logic-services