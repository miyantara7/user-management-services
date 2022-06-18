package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	rkboot "github.com/rookie-ninja/rk-boot"
	rkgrpc "github.com/rookie-ninja/rk-grpc/boot"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	dbUser "github.com/vins7/user-management-services/app/adapter/db/user_management"
	"github.com/vins7/user-management-services/app/infrastructure/connection/db"
	svcUser "github.com/vins7/user-management-services/app/service/user_management"
	ucUser "github.com/vins7/user-management-services/app/usecase/user_management"
	cfg "github.com/vins7/user-management-services/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func RunServer() {

	boot := rkboot.NewBoot()

	config := cfg.GetConfig()
	grpcServer := grpc.NewServer()

	reflection.Register(grpcServer)

	// ***************************************
	// ******* Register GRPC & Gateway *******
	// ***************************************

	// Get grpc entry with name
	grpcEntry := boot.GetEntry("greeter").(*rkgrpc.GrpcEntry)
	// Register grpc registration function
	grpcEntry.AddRegFuncGrpc(Apply)
	// Register grpc-gateway registration function
	// grpcEntry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())

	svcHost := config.Server.Grpc.Host
	svcPort := config.Server.Grpc.Port

	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", svcHost, svcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to start notpool Service gRPC server: %v", err)
		}
	}()

	fmt.Printf("gRPC server is running at %s:%d\n", svcHost, svcPort)

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("process killed with signal: %v\n", signal.String())

}

func Apply(server *grpc.Server) {
	proto.RegisterTransactionNotpoolServiceServer(server, svcUser.NewUserManagementService(ucUser.NewUserManagementUsecase(dbUser.NewUserManagementDB(db.UserDB))))
}
