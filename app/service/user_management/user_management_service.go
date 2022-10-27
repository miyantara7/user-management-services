package user_mangement

import (
	"context"

	"github.com/mitchellh/mapstructure"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	"github.com/vins7/user-management-services/app/interface/model"
	ucUser "github.com/vins7/user-management-services/app/usecase/user_management"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
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
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &proto.LoginResponse{
		Username: data.Username,
		UserID:   data.UserId,
	}, nil
}

func (u *UserManagementService) Register(ctx context.Context, req *proto.RegisterRequest) (*emptypb.Empty, error) {

	if err := u.uc.CreateUser(req); err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (u *UserManagementService) UserInformation(ctx context.Context, req *proto.UserInformationReq) (*proto.UserInformationRes, error) {
	res, err := u.uc.GetDetailUser(req)
	if err != nil {
		return nil, err
	}

	var out *proto.UserInformationRes
	if err := mapstructure.Decode(res, &out); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return out, nil
}
