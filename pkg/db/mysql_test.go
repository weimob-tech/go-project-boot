package db

import (
	"context"
	"github.com/spf13/viper"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNewMysql(t *testing.T) {
	viper.SetDefault("mysql.address", "127.0.0.1")
	viper.SetDefault("mysql.username", "root")
	viper.SetDefault("mysql.password", "root")
	viper.SetDefault("mysql.database", "test")
	viper.SetDefault("mysql.true", true)

	type model struct {
		Id   int64
		Name string
	}

	Convey("Create mysql client should success", t, func() {
		db, err := NewMysql(context.Background())
		SoMsg("create mysql client fail", err, ShouldBeNil)

		_ = db.Exec(`create table if not exists test_create_mysql (
    id int(11) not null auto_increment,
    name varchar(255) not null,
    primary key (id))`).Error
		_ = db.Exec(`truncate table test_create_mysql`)

		row := model{Name: "test"}
		err = db.Table("test_create_mysql").Create(&row).Error
		SoMsg("insert fail", err, ShouldBeNil)
		_, _ = Print("row insert: ", row)

		Convey("Mysql query should success", func() {
			var result []model
			db.Table("test_create_mysql").Select("name = 'test'").Find(&result)
			SoMsg("result not empty", result, ShouldNotBeEmpty)
		})
	})
}
