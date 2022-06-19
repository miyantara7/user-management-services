package user_management

import "github.com/vins7/user-management-services/app/adapter/entity"

type UserManagementRepo interface {
	Login(interface{}) (*entity.User, error)
	CreateUser(interface{}) error
	InsertLoginHistory(interface{}) error
}
