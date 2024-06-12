package main

import "net"

// User 用户结构体，包含用户名、地址、消息管道、链接
type User struct {
	Name       string
	Addr       string
	Channel    chan string // 跟用户绑定的channel
	connection net.Conn    // 客户端连接，用于发送消息
}

// 创建一个user的API
func NewUser(connection net.Conn) *User {

	// 客户端链接地址作为Name和Addr
	userAddr := connection.RemoteAddr().String()

	user := &User{
		Name:       userAddr,
		Addr:       userAddr,
		Channel:    make(chan string),
		connection: connection,
	}

	// 启动监听当前user channel消息的goroutine
	go user.ListenMessage()

	return user
}

// 当前user channel的方法，一旦有消息就发送给客户端
func (user *User) ListenMessage() {
	for {
		// 这个操作是阻塞的，直到通道中有消息为止。
		msg := <-user.Channel
		// 将读取到的消息发送给客户端，并在消息后面加上换行符。
		user.connection.Write([]byte(msg + "\n"))
	}

}
