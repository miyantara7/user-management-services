package user_management

type UserManagement interface {
	Login(interface{}) (interface{}, error)
	CreateUser(interface{}) error
	GetDetailUser(interface{}) (interface{}, error)
}
