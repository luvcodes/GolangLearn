package main

import (
	"fmt"
	"net"
	"sync"
)

// Server 定义Server类
type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	// 消息广播的channel
	Message chan string
}

// NewServer 创建一个server接口
// 这个方法就是构造方法，但是在Go中构造方法不需要与Server结构体进行关联 (this *Server) 不用这样做
// 返回值是一个*Server指针，这样可以避免内存开销，防止反复复制结构体
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
		// 初始化在线用户列表
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
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

	// 启动监听Message的goroutine
	go this.ListenMessage()

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

// 属性: 由哪个用户发起，消息内容是什么
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ": " + msg

	// 将消息广播到Message中, 也就是总体的负责广播的channel
	this.Message <- sendMsg
}

// 监听Message广播消息channel的goroutine，一旦有消息就发送给全部在线用户
func (this *Server) ListenMessage() {
	for {
		msg := <-this.Message

		// 将msg发送给全部在线用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.Channel <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	// fmt.Println("连接简历成功")

	user := NewUser(conn)

	// 用户上线，将用户加入到OnlineMap中
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	// 广播当前用户上线消息
	this.BroadCast(user, "已上线")

	// 当前handler阻塞
	select {}
}
