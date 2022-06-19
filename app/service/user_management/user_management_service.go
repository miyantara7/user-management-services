package user_mangement

import (
	"context"

	"github.com/mitchellh/mapstructure"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	"github.com/vins7/user-management-services/app/interface/model"
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

func (u *UserManagementService) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	res, err := u.uc.Login(req)
	if err != nil {
		return nil, err
	}

	var data *model.LoginResponse
	if err := mapstructure.Decode(res, &data); err != nil {
		return nil, err
	}

	return &proto.LoginResponse{
		Username: data.Username,
		Token:    data.Token,
	}, nil
}
