package user_management

import "gorm.io/gorm"

type UserManagementDB struct {
	db *gorm.DB
}

func NewUserManagementDB(db *gorm.DB) *UserManagementDB {
	return &UserManagementDB{
		db: db,
	}
}

func (u *UserManagementDB) Login() {

}
