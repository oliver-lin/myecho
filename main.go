package main

import (
	"echo/boot"
)

func main() {

	e := boot.GetRouter()

	//启动http server, 并监听8080端口，冒号（:）前面为空的意思就是绑定网卡所有Ip地址，本机支持的所有ip地址都可以访问。
	e.Start(":8080")
}
