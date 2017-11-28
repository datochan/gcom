package utils

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAddDays(t *testing.T) {
	Convey("测试日期的加减法", t, func() {
		// 默认加载股票日历数据
		target1 := AddDays("20170618", 37)
		So(target1, ShouldEqual, "20170725")


		target2 := AddDays("20170618", -17)
		So(target2, ShouldEqual, "20170601")
	})

}

func TestAddDaysExceptWeekend(t *testing.T) {
	Convey("测试跳过周末的日期加减法", t, func() {
		target1 := AddDaysExceptWeekend("20171126", 7)
		So(target1, ShouldEqual, "20171205")


		target2 := AddDaysExceptWeekend("20171126", -7)
		So(target2, ShouldEqual, "20171116")
	})
}
