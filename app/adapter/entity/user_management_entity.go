package entity

type User struct {
	BaseModel
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
