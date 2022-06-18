package db

import (
	"fmt"
	"log"

	"github.com/vins7/user-management-services/app/adapter/entity"
	"github.com/vins7/user-management-services/config"
	"github.com/vins7/user-management-services/config/db"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var tables = []interface{}{
	&entity.User{},
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
	pg := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta", cfg.Host, cfg.Username, cfg.Password, cfg.Dbname, cfg.Port)
	db, err := gorm.Open(postgres.Open(pg), &gorm.Config{})
	MigrateSchema(db)
	return db, err
}

func MigrateSchema(db *gorm.DB) {
	db.AutoMigrate(tables...)
}
