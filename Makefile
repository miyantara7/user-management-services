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
	@docker build --tag user-management-services .

protoc-docker: build
	@docker container create --name user-services -p 9902:9902/tcp user-management-services