// 测速程序使用go语言编写
// 测速程序仿照360测速原理，以外网测速点下载文件测量接入速度
package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	url1 = "http://sw.bos.baidu.com/sw-search-sp/software/509cca8b8ae/BaiduYunGuanjia_5.4.5.4.exe"
	//百度测速点
	url2 = "http://down.360safe.com/yunpan/360wangpan_setup_6.5.6.1282.exe"
	//360测速点
	url3 = "http://mail.lzu.edu.cn/cab/cmplugin_setup.exe"
	//校内测速点
)

func main() {
	fmt.Println("正在测试连接淘宝网.....")
	time.Sleep(time.Second * 1)
	_, err1 := net.Dial("ip4:icmp", "www.taobao.com")
	//使用icmp协议向对端服务器发送请求，如果错误为空表示请求成功返回
	if err1 != nil {
		fmt.Println("无法连接淘宝网，请先确认网络连接正常！")
	} else {
		fmt.Println("连接淘宝网正常！")
	}
	time.Sleep(time.Second * 1)

	fmt.Println("正在测试连接优酷网.....")
	time.Sleep(time.Second * 1)
	_, err2 := net.Dial("ip4:icmp", "www.youku.com")
	if err2 != nil {
		fmt.Println("无法连接优酷网，请先确认网络连接正常！")
	} else {
		fmt.Println("连接优酷网正常！")
	}
	time.Sleep(time.Second * 1)

	fmt.Println("正在测试连接兰大主页.....")
	time.Sleep(time.Second * 1)
	_, err3 := net.Dial("ip4:icmp", "www.lzu.edu.cn")
	if err3 != nil {
		fmt.Println("无法连接兰大主页，请直接致电：8914088报修")
	} else {
		fmt.Println("连接兰大主页正常！")
	}
	time.Sleep(time.Second * 1)

	fmt.Println("开始测试校内连接速度.....")
	t3 := time.Now()
	//打下开始时间戳
	res3, err := http.Get(url3)
	//从测速点下载文件，以测试下载速度，如果无法连接则panic退出程序
	if err != nil {
		panic(err)
	}
	f3, err := os.Create("cmplugin_setup.exe") //cmplugin_setup.exe:2.1M
	if err != nil {
		panic(err)
	}
	io.Copy(f3, res3.Body)
	const m3 float32 = 2.1
	s3 := (m3 / float32(time.Since(t3))) * 10e8
	//用文件尺寸除以下载用时，得到平均下载速度，并换算成MB/S
	fmt.Printf("您的校内连接速度为：%5.2fMB/秒。\n", s3)
	time.Sleep(time.Second * 1)

	fmt.Println("开始测试外网连接速度，可能会持续几分钟.....")
	t1 := time.Now()
	res, err := http.Get(url1)
	if err != nil {
		panic(err)
	}
	f, err := os.Create("baiduyun.exe") //baiduyun:15.1M
	if err != nil {
		panic(err)
	}
	io.Copy(f, res.Body)
	const m1 float32 = 15.1
	s1 := (m1 / float32(time.Since(t1))) * 10e8

	t2 := time.Now()
	res2, err := http.Get(url2)
	if err != nil {
		panic(err)
	}
	f2, err := os.Create("360yunpan.exe") //360yunpan:18.4M
	if err != nil {
		panic(err)
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
	//选取测速点中下载速度快的，作为外网连接速度
	fmt.Printf("您的外网连接速度为：%5.2fMB/秒。\n", sz)

	switch {
	case sz <= 0.5:
		fmt.Println("您的网速这么慢，还是别上网，好好看书吧~~")
	case 0.5 < sz && sz <= 2:
		fmt.Println("您的网速可以流畅的打开网页~~")
	case 2 < sz && sz <= 5:
		fmt.Println("您的网速可以流畅的观看视频~~")
	case sz > 5:
		fmt.Println("您是从网络中心拉的专线吧？不然怎么这么快~~")
	}
	//对网速做比较
}
