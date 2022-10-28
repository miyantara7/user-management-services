package user_management

import (
	"time"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/user-management-services/app/adapter/db/user_management"
	"github.com/vins7/user-management-services/app/adapter/entity"
	"github.com/vins7/user-management-services/app/interface/model"
	"github.com/vins7/user-management-services/app/util"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserManagementUsecase struct {
	repo repo.UserManagementRepo
}

func NewUserManagementUsecase(repo repo.UserManagementRepo) UserManagement {
	return &UserManagementUsecase{
		repo: repo,
	}
}

func (u *UserManagementUsecase) Login(in interface{}) (interface{}, error) {
	var req *entity.User

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, err
	}

	res, err := u.repo.Login(req)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password)); err != nil {
		return nil, status.Errorf(codes.Internal, "username atau password salah !")
	}

	if err := u.repo.InsertLoginHistory(&entity.LoginHistory{
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:  time.Now().Format("2006-01-02 15:04:05"),
		Username:    res.Username,
	}); err != nil {
		return nil, err
	}

	out := &model.LoginResponse{
		Username: res.Username,
		UserId:   res.DataUser.UserID,
	}

	return out, nil
}

func (u *UserManagementUsecase) CreateUser(in interface{}) error {
	var req *model.CreateUserReq

	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	data := &entity.User{
		Username: req.Username,
		Password: string(bytes),
		DataUser: entity.DataUser{
			Nama:        req.Nama,
			NoIdentitas: req.NoIdentitas,
			TglLahir:    req.TglLahir,
			Email:       req.Email,
			UserID:      util.GenerateID(),
		},
	}
	if err := u.repo.CreateUser(data); err != nil {
		return err
	}

	return nil
}

func (u *UserManagementUsecase) GetDetailUser(in interface{}) (interface{}, error) {
	var req *model.DetailUserReq

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	res, err := u.repo.DetailUser(req)
	if err != nil {
		return nil, err
	}

	return &model.DataUser{
		Username:    res.Username,
		Nama:        res.DataUser.Nama,
		UserId:      res.DataUser.UserID,
		Email:       res.DataUser.Email,
		NoIdentitas: res.DataUser.NoIdentitas,
		TglLahir:    res.DataUser.TglLahir,
	}, nil
}
