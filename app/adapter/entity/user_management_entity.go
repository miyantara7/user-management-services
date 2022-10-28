package entity

type User struct {
	BaseModel
	Username   string   `gorm:"username"`
	Password   string   `gorm:"password"`
	DataUserId string   `gorm:"column:data_user_id;size:100"`
	DataUser   DataUser `gorm:"foreignKey:DataUserId"`
}

func (User) TableName() string {
	return "t_user"
}

type DataUser struct {
	BaseModel
	UserID      string `gorm:"user_id"`
	Nama        string `gorm:"nama"`
	NoIdentitas string `gorm:"no_identitas"`
	TglLahir    string `gorm:"tgl_lahir"`
	Email       string `gorm:"email"`
}

func (DataUser) TableName() string {
	return "t_data_user"
}

type LoginHistory struct {
	BaseModel
	CreatedDate string
	UpdateDate  string
	Username    string
}

func (LoginHistory) TableName() string {
	return "t_login_history"
}
