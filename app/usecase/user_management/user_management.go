package user_management

type UserManagement interface {
	Login(interface{}) (interface{}, error)
}
