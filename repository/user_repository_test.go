package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
	"unit-test/database"
)

func TestUserRepository_GetUserById(t *testing.T) {
	convey.Convey("TestUserRepository_GetUserById", t, func() {
		userRepository := NewUserRepository()

		//创建sqlmock
		sqldb, mock, err := sqlmock.New()
		if nil != err {
			log.Fatalf("Init sqlmock failed, err %v", err)
		}
		//基于sqlmock创建db实例
		db, err := gorm.Open(mysql.New(mysql.Config{
			SkipInitializeWithVersion: true,
			Conn:                      sqldb,
		}), &gorm.Config{})
		if nil != err {
			log.Fatalf("Init DB with sqlmock failed, err %v", err)
		}

		//对全局变量database.Db打桩，替换为基于sqlmock创建的db实例
		patches := gomonkey.ApplyGlobalVar(&database.Db, db)
		defer patches.Reset()

		convey.Convey("一般性用例", func() {
			userId := 1
			rows := mock.NewRows([]string{"id", "name"}).
				AddRow(1, "张三")

			//注意实际使用的sql语句必须和预期的sql完全吻合，才会匹配，甚至大小写也需要完全一致。
			//若语句中包含*或?，记得做转义
			mock.ExpectQuery("^SELECT \\* FROM `users` WHERE id = \\? ORDER BY `users`.`id` LIMIT 1").
				WithArgs(userId).
				WillReturnRows(rows)

			user, err := userRepository.GetUserById(userId)
			convey.So(err, convey.ShouldEqual, nil)
			convey.So(user, convey.ShouldNotEqual, nil)
		})

		convey.Convey("异常用例", func() {
			mock.ExpectQuery("SELECT").WillReturnError(errors.New("失败"))

			userId := 1
			user, err := userRepository.GetUserById(userId)
			convey.So(err, convey.ShouldNotEqual, nil)
			convey.So(user, convey.ShouldEqual, nil)
		})

	})
}
