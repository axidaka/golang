package ipc

import (
	"time"
	"fmt"
	"testing"
)
type EchoServer struct {
}

func (server *EchoServer)Handle(mehod, params string) *Response {
	return &Response{"200", "hello from echo server"}
}

func (server *EchoServer)Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {

	server := NewIpcServer(&EchoServer{})
	client1 := NewIpcClient(server)
	
	resp1, ok := client1.Call("From Client1", "klklj")
	if ok != nil {
		t.Error("IpcClient.Call failed. resp1:", resp1)
	}

	fmt.Println(resp1)

	client1.Close()

	time.Sleep(5e9)

}
