package user_management

import (
	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/vins7/user-management-services/app/adapter/entity"
	"github.com/vins7/user-management-services/app/interface/model"
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

	if err := u.db.Debug().Where("username = ?", req.Username).Joins("DataUser").First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "username atau password anda salah !")
		}
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &data, nil
}

func (u *UserManagementDB) CreateUser(in interface{}) error {

	var req *entity.User

	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	data := &entity.User{}
	if err := u.db.Debug().Where("username = ?", req.Username).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if err := u.db.Debug().Create(req).Error; err != nil {
				return status.Errorf(codes.Internal, err.Error())
			}
			return nil
		}
		return status.Errorf(codes.Internal, err.Error())
	}

	if data.Id != "" {
		return status.Errorf(codes.FailedPrecondition, "username already exist !")
	}

	return nil
}

func (u *UserManagementDB) DetailUser(req *model.DetailUserReq) (*entity.User, error) {

	data := &entity.User{}
	if err := u.db.Debug().
		Joins("DataUser").
		Where("username = ? and DataUser.user_id = ?", req.Username, req.UserId).
		First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user tidak ditemukan !")
		}
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return data, nil
}

func (u *UserManagementDB) InsertLoginHistory(in interface{}) error {
	var (
		req *entity.LoginHistory
	)

	if err := mapstructure.Decode(in, &req); err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}

	if err := u.db.Debug().Create(req).Error; err != nil {
		return status.Errorf(codes.NotFound, err.Error())
	}

	return nil
}
