package user_management

import (
	"time"

	"github.com/mitchellh/mapstructure"
	repo "github.com/vins7/user-management-services/app/adapter/db/user_management"
	"github.com/vins7/user-management-services/app/adapter/entity"
	"github.com/vins7/user-management-services/app/interface/model"
	"github.com/vins7/user-management-services/app/usecase/credential"
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

	t, err := credential.GenerateToken(res.Username)
	if err != nil {
		return nil, err
	}

	if err := u.repo.InsertLoginHistory(&entity.LoginHistory{
		CreatedDate: time.Now().Format("2006-01-02 15:04:05"),
		UpdateDate:  time.Now().Format("2006-01-02 15:04:05"),
		Token:       t,
		Username:    res.Username,
	}); err != nil {
		return nil, err
	}

	out := &model.LoginResponse{
		Username: res.Username,
		Token:    t,
	}

	return out, nil
}
