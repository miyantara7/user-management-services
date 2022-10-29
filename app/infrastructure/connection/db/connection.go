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
	d := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "mysql")
	dbTemp, err := gorm.Open(mysql.Open(d), &gorm.Config{})
	if err != nil {
		return dbTemp, err
	}
	CreateDB(dbTemp, cfg.Dbname)

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	MigrateSchema(db)

	return db, err
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}

func CreateDB(db *gorm.DB, database string) {
	db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %v", database))
}
