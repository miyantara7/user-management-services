package db

import (
	"fmt"
	"log"

	"github.com/vins7/user-management-services/app/adapter/entity"
	"github.com/vins7/user-management-services/config"
	"github.com/vins7/user-management-services/config/db"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var tables = []interface{}{
	&entity.DataUser{},
	&entity.User{},
	&entity.LoginHistory{},
}

var (
	UserDB *gorm.DB
)

func init() {
	var err error
	cfg := config.GetConfig()

	UserDB, err = Conn(cfg.Database.UserManagement)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func Conn(cfg db.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	MigrateSchema(db)
	return db, err
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
