package network

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read from client failed,err:", err)
			break
		}
		recvStr := string(buf[:n])
		if len(recvStr) > 0 {
			fmt.Println("get client sent data:", recvStr)
			conn.Write([]byte(recvStr))
		}
	}
}

func CreateTcpServer() {
	listen, err := net.Listen("tcp", "127.0.0.1:7890")

	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	fmt.Printf("server start listener at:http://%v \n", listen.Addr())

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed,err:", err)
			continue
		}

		go process(conn)
	}
}

func CreateTcpClient() {
	connect, err := net.Dial("tcp", "127.0.0.1:7890")

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	defer connect.Close()

	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}
		_, err := connect.Write([]byte(inputInfo))
		if err != nil {
			return
		}
		buf := [512]byte{}
		n, err := connect.Read(buf[:])
		if err != nil {
			fmt.Println("recv failed,err:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}

}
