package ipc

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

// 定义接口
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

// 匿名组合，可以使用不同的server来构建对象
type IpcServer struct {
	Server  // 可以使用实现了Server接口的类型对象赋值
}

// 构造函数
func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}
}

// 客户端调用，返回channel用于通信
func (server *IpcServer) Connect() chan string {

	session := make(chan string, 0)

	// 创建goroutine用于读取
	go func(c chan string) {

		for {
			request := <- c

			if request == "CLOSE" {
				break
			}

			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("Invalid request format:", request)
			}

			rsp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(rsp)

			c <- string(b)
		}

		fmt.Println("Session closed")
	}(session)

	fmt.Println("A new session has been created succfully")
	return session
}