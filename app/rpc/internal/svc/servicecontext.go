package svc

import (
	"goZeroDemo/app/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to db error")
	}
	return &ServiceContext{
		Config: c,
		Db:     db,
	}

}
