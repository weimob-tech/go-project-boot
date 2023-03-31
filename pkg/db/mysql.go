package db

import (
	"context"
	"fmt"
	"github.com/weimob-tech/go-project-base/pkg/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMysql(ctx context.Context) (db *gorm.DB, err error) {
	// todo: logger, timeout
	db, err = gorm.Open(mysql.Open(mysqlDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// enable debug
	if config.GetBool("debug") || config.GetBool("mysql.debug") {
		db = db.Debug()
	}
	// enable connection test
	if config.GetString("mysql.test") != "false" {
		err := db.WithContext(ctx).Exec(`select 1 = 1`).Error
		if err != nil {
			panic(fmt.Errorf("redis connection test failed: %w", err))
		}
	}
	return
}

func mysqlDsn() string {
	template := "%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local"
	return fmt.Sprintf(template,
		config.GetString("mysql.username"),
		config.GetString("mysql.password"),
		config.GetString("mysql.address"),
		config.GetString("mysql.database"),
	)
}
