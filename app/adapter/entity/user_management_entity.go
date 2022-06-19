package entity

type User struct {
	BaseModel
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}

type LoginHistory struct {
	BaseModel
	CreatedDate string
	UpdateDate  string
	Token       string
	Username    string
}
