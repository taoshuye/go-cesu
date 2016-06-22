package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/samples/flags"
)

var (
	url1 = "http://sw.bos.baidu.com/sw-search-sp/software/509cca8b8ae/BaiduYunGuanjia_5.4.5.4.exe"
	//百度测速点
	url2 = "http://down.360safe.com/yunpan/360wangpan_setup_6.5.6.1282.exe"
	//360测速点
)

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver) //添加样式

	lianjie1 := theme.CreateLinearLayout()
	button := theme.CreateButton()
	lianjie1.SetHorizontalAlignment(gxui.AlignCenter) //确定位置
	lianjie1.AddChild(button)                         //添加按钮
	button.SetText("ping to taobao.com")
	label2 := theme.CreateLabel() //添加文字说明
	label2.SetText("Start measuring")
	toggle := func() { //按钮触发的函数
		_, err1 := net.Dial("tcp", "www.taobao.com:80")
		if err1 != nil {
			label2.SetText("connection fail")
		} else {
			label2.SetText("The network connection is normal to taobao.com")
		}
	}
	button.OnClick(func(gxui.MouseEvent) { toggle() })
	lianjie1.AddChild(label2)

	lianjie2 := theme.CreateLinearLayout()
	button2 := theme.CreateButton()
	lianjie2.SetHorizontalAlignment(gxui.AlignCenter)
	lianjie2.AddChild(button2)
	button2.SetText("ping to youku.com")
	label3 := theme.CreateLabel()
	label3.SetText("Start measuring")
	toggle2 := func() {
		_, err2 := net.Dial("tcp", "www.youku.com:80")
		if err2 != nil {
			label3.SetText("connection fail")
		} else {
			label3.SetText("The network connection is normal to youku.com")
		}
	}
	button2.OnClick(func(gxui.MouseEvent) { toggle2() })
	lianjie2.AddChild(label3)

	lianjie3 := theme.CreateLinearLayout()
	button3 := theme.CreateButton()
	lianjie3.SetHorizontalAlignment(gxui.AlignCenter)
	lianjie3.AddChild(button3)
	button3.SetText("ping to lzu.edu.cn")
	label4 := theme.CreateLabel()
	label4.SetText("Start measuring")
	toggle3 := func() {
		_, err3 := net.Dial("tcp", "www.lzu.edu.cn:80")
		if err3 != nil {
			label4.SetText("connection fail,Please call 8914088 for maintenance")
		} else {
			label4.SetText("The network connection is normal to lzu.edu.cn")
		}
	}
	button3.OnClick(func(gxui.MouseEvent) { toggle3() })
	lianjie3.AddChild(label4)

	cesu := theme.CreateLinearLayout()
	button4 := theme.CreateButton()
	cesu.SetHorizontalAlignment(gxui.AlignCenter)
	button4.SetText("Test speed")
	label5 := theme.CreateLabel()
	label5.SetText("If start please wait for a few minutes...")
	toggle4 := func() {
		_, err4 := net.Dial("tcp", "www.baidu.com:80")
		if err4 != nil {
			label5.SetText("connection fail")
		} else {
			t1 := time.Now()
			res, err5 := http.Get(url1)
			if err5 != nil {
				label5.SetText("connection fail")
			}
			f, err6 := os.Create("baiduyun.exe") //baiduyun:15.1M
			if err6 != nil {
				label5.SetText("Failed to download the test file")
				//panic(err6)
			}
			io.Copy(f, res.Body)
			const m1 float32 = 15.1
			s1 := (m1 / float32(time.Since(t1))) * 10e8

			t2 := time.Now()
			res2, err7 := http.Get(url2)
			if err7 != nil {
				label5.SetText("connection fail")
				//panic(err)
			}
			f2, err8 := os.Create("360yunpan.exe") //360yunpan:18.4M
			if err8 != nil {
				label5.SetText("Failed to download the test file")
				//panic(err)
			}
			io.Copy(f2, res2.Body)
			const m2 float32 = 18.4
			s2 := (m2 / float32(time.Since(t2))) * 10e8

			var sz float32
			if s1 > s2 {
				sz = s1
			} else {
				sz = s2
			}
			jieguo := fmt.Sprintf("Your Internet connection is:%5.2fMB/S", sz)
			label5.SetText(jieguo)
		}
	}
	button4.OnClick(func(gxui.MouseEvent) { toggle4() })
	cesu.AddChild(button4)
	cesu.AddChild(label5)

	//绘制边框
	label1 := theme.CreateLabel()
	label1.SetColor(gxui.Black)
	label1.SetText("Speed measuring software v0.1")
	hengkuang := theme.CreateLinearLayout()
	hengkuang.SetBackgroundBrush(gxui.CreateBrush(gxui.White))
	hengkuang.SetHorizontalAlignment(gxui.AlignCenter)
	hengkuang.AddChild(label1)
	label6 := theme.CreateLabel()
	label6.SetColor(gxui.Black)
	label6.SetText("L contact us: taozw@lzu.edu.cn")
	hengkuang2 := theme.CreateLinearLayout()
	hengkuang2.SetBackgroundBrush(gxui.CreateBrush(gxui.White))
	hengkuang2.SetHorizontalAlignment(gxui.AlignRight)
	hengkuang2.AddChild(label6)
	shukuang := theme.CreateLinearLayout()
	shukuang.SetBackgroundBrush(gxui.CreateBrush(gxui.White))
	shukuang2 := theme.CreateLinearLayout()
	shukuang2.SetBackgroundBrush(gxui.CreateBrush(gxui.White))

	//以格子的方式确定窗口板式
	table := theme.CreateTableLayout()
	table.SetGrid(7, 6) // columns, rows

	// row, column, horizontal span, vertical span
	//确定各个子项位置
	table.SetChildAt(0, 0, 7, 1, hengkuang)
	table.SetChildAt(0, 5, 7, 1, hengkuang2)
	table.SetChildAt(0, 1, 1, 4, shukuang)
	table.SetChildAt(6, 1, 1, 4, shukuang2)
	table.SetChildAt(1, 1, 5, 1, lianjie1)
	table.SetChildAt(1, 2, 5, 1, lianjie2)
	table.SetChildAt(1, 3, 5, 1, lianjie3)
	table.SetChildAt(1, 4, 5, 1, cesu)

	window := theme.CreateWindow(420, 380, "web speed")
	window.AddChild(table)
	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
