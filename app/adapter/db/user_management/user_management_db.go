package user_management

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/vins7/user-management-services/app/adapter/entity"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserManagementDB struct {
	db *gorm.DB
}

func NewUserManagementDB(db *gorm.DB) *UserManagementDB {
	return &UserManagementDB{
		db: db,
	}
}

func (u *UserManagementDB) Login(in interface{}) (*entity.User, error) {
	var (
		req  *entity.User
		data entity.User
	)

	if err := mapstructure.Decode(in, &req); err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	if err := u.db.Debug().Where("username = ? and password = ?", req.Username, req.Password).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "username atau password tidak ditemukan !")
		}
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &data, nil
}

func (u *UserManagementDB) CreateUser(in interface{}) error {
	var (
		req  *entity.User
		data entity.User
	)

	if err := mapstructure.Decode(in, &req); err != nil {
		return err
	}

	if err := u.db.Debug().Where("username = ? and password = ?").Find(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return status.Errorf(codes.NotFound, "username atau password tidak ditemukan !")
		}
		return err
	}

	return nil
}

func (u *UserManagementDB) InsertLoginHistory(in interface{}) error {
	var (
		req *entity.LoginHistory
	)

	if err := mapstructure.Decode(in, &req); err != nil {
		return err
	}

	if err := u.db.Debug().Create(req).Error; err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	return nil
}
