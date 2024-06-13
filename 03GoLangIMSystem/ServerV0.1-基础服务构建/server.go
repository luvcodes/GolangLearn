package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

// 启动服务器的接口
func (this *Server) Start() {

	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	// close listen socket
	// 确保在函数退出时关闭监听, 因为defer是在最后一步执行的
	defer listener.Close()

	// 进入一个无限循环，不断接收客户端连接
	for {
		// accept
		conn, err := listener.Accept()
		// 出现了error, 打印错误信息, 继续下一次循环
		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}

		// do handler
		// 每接收到一个连接，启动一个新的goroutine来处理连接
		go this.Handler(conn)
	}
}

// ...当前链接的业务, 给出业务链接是否成功的打印消息
func (this *Server) Handler(conn net.Conn) {
	fmt.Println("链接建立成功")
}
