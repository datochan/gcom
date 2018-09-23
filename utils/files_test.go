package utils

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCurrentPath(t *testing.T) {
	Convey("测试获取文件输出目录", t, func() {
		// 默认加载股票日历数据
		_, err := CurrentPath()
		So(err, ShouldEqual, nil)
	})

}
