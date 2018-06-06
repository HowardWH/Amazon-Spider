/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package log

import (
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestNode(t *testing.T) {
	convey.Convey("测试node", t, func() {

		l := &LoggerConf{Name: ""}
		root := newNode("", nil, l)
		root.addChild("a/b/c", &LoggerConf{Name: "a/b/c"})
		// convey.Convey("测试 添加子节点", func() {

		c, ok := root.children["a"]
		convey.So(ok, convey.ShouldEqual, true)

		c, ok = c.children["b"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current, convey.ShouldEqual, nil)

		c, ok = c.children["c"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b/c")
		convey.So(len(c.children), convey.ShouldEqual, 0)

		// convey.Convey("在已有节点前面插入logger", func() {

		root.addChild("a/b", &LoggerConf{Name: "a/b"})
		c, ok = root.children["a"]
		convey.So(ok, convey.ShouldEqual, true)

		c, ok = c.children["b"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current, convey.ShouldNotEqual, nil)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b")

		c, ok = c.children["c"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b/c")
		convey.So(len(c.children), convey.ShouldEqual, 0)
		// })

		// convey.Convey("在已有节点后面插入logger", func() {
		root.addChild("a/b/c/d/e", &LoggerConf{Name: "a/b/c/d/e"})
		c, ok = c.children["d"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current, convey.ShouldEqual, nil)

		c, ok = c.children["e"]
		convey.So(ok, convey.ShouldEqual, true)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b/c/d/e")
		// })
		// })

		// convey.Convey("测试获取子节点", func() {
		child := root.child("a/b/c")
		convey.So(child, convey.ShouldNotEqual, nil)
		convey.So(child.current.Name, convey.ShouldEqual, "a/b/c")

		child = root.child("a/b/e")
		convey.So(child, convey.ShouldEqual, nil)
		// })

		// convey.Convey("测试获取父节点", func() {
		// {

		child = root.child("a/b/c/d/e")
		c = child.higher()
		convey.So(c, convey.ShouldNotEqual, nil)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b/c")

		c = c.higher()
		convey.So(c, convey.ShouldNotEqual, nil)
		convey.So(c.current.Name, convey.ShouldEqual, "a/b")
		// }
		// })

		convey.Convey("测试获取配置", func() {
			l := &LoggerConf{Name: "", Levels: map[int]bool{ERROR: true}}
			root := newNode("", nil, l)
			root.addChild("a/b", &LoggerConf{Name: "a/b", Appenders: []Appender{NewConsoleAppender("test")}})
			root.addChild("a/b/c", &LoggerConf{Name: "a/b/c"})
			root.addChild("a/b/c/d/e", &LoggerConf{Name: "a/b/c/d/e", Levels: map[int]bool{DEBUG: true}})

			cur := root.generate("a/b/c/d/e/f/g")
			cfg := cur.inheritConf()
			convey.So(cfg.Levels[DEBUG], convey.ShouldBeTrue)
			convey.So(len(cfg.Appenders), convey.ShouldEqual, 1)
		})
	})

}
