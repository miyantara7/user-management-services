package main

import (
	grpc "github.com/vins7/user-management-services/app/infrastructure/grpc"
)

func main() {
	grpc.RunServer()
}
