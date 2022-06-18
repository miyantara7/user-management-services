package user_mangement

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	ucUser "github.com/vins7/user-management-services/app/usecase/user_management"
)

type UserManagementService struct {
	uc ucUser.UserManagement
}

func NewUserManagementService(uc ucUser.UserManagement) *UserManagementService {
	return &UserManagementService{
		uc: uc,
	}
}

func (u *UserManagementService) Login(ctx context.Context, req *proto.LoginRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}
