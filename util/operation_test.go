package util

import (
	"log"
	"testing"
	//使用官方推荐的方式导入 GoConvey 的辅助包以减少冗余的代码：. "github.com/smartystreets/goconvey/convey"
	. "github.com/agiledragon/gomonkey"
	. "github.com/smartystreets/goconvey/convey"
)

//每个单元测试的名称需要以 Test 开头，例如：TestAdd，并需要接受一个类型为 *testing.T 的参数。
func TestAdd(t *testing.T) {
	//每个测试用例需要使用 Convey 函数包裹起来。它接受的第一个参数为 string 类型的描述；
	// 第二个参数一般为 *testing.T，即本例中的变量 t；
	// 第三个参数为不接收任何参数也不返回任何值的函数（习惯以闭包的形式书写）。
	Convey("将两数相加", t, func() {
		patches := ApplyFunc(Add, func(a, b int) int {
			log.Println("执行ApplyFunc")
			return a + b
		})
		defer patches.Reset()

		//断言So 参数的理解，总共有三个参数：actual: 输入\assert：断言\expected：期望值
		So(Add(1, 2), ShouldEqual, 3)
	})
}

func TestSubtract(t *testing.T) {
	Convey("将两数相减", t, func() {
		So(Subtract(1, 2), ShouldEqual, -1)
	})
}

func TestMultiply(t *testing.T) {
	Convey("将两数相乘", t, func() {
		So(Multiply(3, 2), ShouldEqual, 6)
	})
}

func TestDivision(t *testing.T) {
	Convey("将两数相除", t, func() {
		//Convey 语句同样可以无限嵌套，以体现各个测试用例之间的关系
		Convey("除以非 0 数", func() {
			num, err := Division(10, 2)
			So(err, ShouldBeNil)
			So(num, ShouldEqual, 5)
		})

		Convey("除以 0", func() {
			_, err := Division(10, 0)
			So(err, ShouldNotBeNil)
		})
	})
}
