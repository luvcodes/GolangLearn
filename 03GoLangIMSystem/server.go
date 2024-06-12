package main

import (
	"fmt"
	"net"
)

// Server 定义Server类
type Server struct {
	Ip   string
	Port int
}

// NewServer 创建一个server接口
// 这个方法就是构造方法，但是在Go中构造方法不需要与Server结构体进行关联 (this *Server) 不用这样做
// 返回值是一个*Server指针，这样可以避免内存开销，防止反复复制结构体
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}

	return server
}

// Start 启动服务器的接口
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	// close listen socket
	defer listener.Close()

	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)

	}

}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接简历成功")
}
