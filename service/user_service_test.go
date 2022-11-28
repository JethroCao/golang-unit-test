package service

import (
	"errors"
	. "github.com/agiledragon/gomonkey"
	"go.uber.org/zap"
	"reflect"
	"testing"
	log2 "unit-test/log"
	"unit-test/model"
	"unit-test/repository"

	//使用官方推荐的方式导入 GoConvey 的辅助包以减少冗余的代码：. "github.com/smartystreets/goconvey/convey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUserService_GetUserById(t *testing.T) {
	Convey("TestUserService_GetUserById", t, func() {
		userService := NewUserService()

		Convey("一般性用例", func() {
			//注意不要丢掉桩方法的第1个形参
			patches := ApplyMethod(reflect.TypeOf(&repository.UserRepository{}), "GetUserById", func(_ *repository.UserRepository, userId int) (*model.User, error) {
				user := &model.User{
					Id:   1,
					Name: "李四",
				}

				log2.Logger.Info("GetUserById: 执行成功, ApplyMethod repository.UserRepository GetUserById", zap.Any("user", user))
				return user, nil
			})
			defer patches.Reset()

			user, err := userService.GetUserById(1)
			So(err, ShouldEqual, nil)
			So(user, ShouldNotEqual, nil)
		})

		Convey("异常用例", func() {
			//注意不要丢掉桩方法的第1个形参
			patches := ApplyMethod(reflect.TypeOf(&repository.UserRepository{}), "GetUserById", func(_ *repository.UserRepository, userId int) (*model.User, error) {
				log2.Logger.Info("GetUserById: 执行失败, ApplyMethod repository.UserRepository GetUserById", zap.Any("userId", userId))
				return nil, errors.New("失败")
			})
			defer patches.Reset()

			_, err := userService.GetUserById(1)
			So(err, ShouldNotEqual, nil)
		})
	})
}
