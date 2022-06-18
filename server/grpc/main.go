package main

import (
	"fmt"

	grpc "github.com/vins7/user-management-services/app/infrastructure/grpc"
)

func main() {
	fmt.Println("Hello World")
	grpc.RunServer()
}
