package chapters

import (
	"bytes"
	"context"
	"fmt"
	pb "golang/src/go_learn/protocolbuf/go"
	"google.golang.org/grpc"
	"io"
	"net"
	"net/http"
	"net/rpc"
	"os"
)

func checkSum(msg []byte) uint16 {

	sum := 0
	// 先假设为偶数
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var answer uint16 = uint16(^sum)
	return answer
}

func checkError(err error) {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) ([]byte, error) {

	defer conn.Close()

	result := bytes.NewBuffer(nil)
	var buf [1024]byte

	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}
	}

	return result.Bytes(), nil
}

func Icmp_test() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: ", os.Args[0], "host")
		os.Exit(1)
	}

	host := os.Args[1]

	conn, err := net.Dial("ip4:icmp", host)
	checkError(err)

	var msg [512]byte
	msg[0] = 8  // echo
	msg[1] = 0  // code 0
	msg[2] = 0  // checksum
	msg[3] = 0  // checksum
	msg[4] = 0  // identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  // sequence[0]
	msg[7] = 37 // sequence[1]
	len := 8

	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("Got response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[7] == 37 {
		fmt.Println("Sequence matches")
	}

	os.Exit(0)
}

func Tcp_test() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: ", os.Args[0], "host")
		os.Exit(1)
	}

	host := os.Args[1]

	conn, err := net.Dial("tcp", host)
	checkError(err)

	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	result, err := readFully(conn)
	checkError(err)

	fmt.Println(string(result))
	os.Exit(0)
}

func Http_test() {

	resp, err := http.Get("http://www.jd.com")
	checkError(err)

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}

//rpc
type Args struct {
	A, B int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func Rpc_Server_test() {

	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	checkError(err)

	go http.Serve(l, nil)
}

func Rpc_Client_test() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	checkError(err)

	var args Args
	args.A = 1
	args.B = 2
	var reply int
	err = client.Call("Arith.Multiply", &args, reply)
	checkError(err)

	fmt.Println("Arith:%d * %d = %d", args.A, args.B, reply)
}

type HelloServiceImpl struct {
}

func (h *HelloServiceImpl) Hello(ctx context.Context, args *pb.String) (*pb.String, error) {
	reply := &pb.String{Value: "hello " + args.GetValue()}
	fmt.Println("req:", args.GetValue(), "  reply:", reply)
	return reply, nil
}

func GRpc_Servce_test() {
	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))

	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen err")
	} else {
		grpcServer.Serve(lis)
	}
}

func GRpc_Client_test() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn err", err)
		return
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &pb.String{Value: "hello world"})
	if err != nil {
		fmt.Println("Hello err", err)
		return
	}
	fmt.Println("Hello reply:", reply.GetValue())
}
