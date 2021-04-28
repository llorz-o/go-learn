package network

import (
	"fmt"
	"net"
)

func CreateUdpServer() {
	udp, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 7890,
	})
	if err != nil {
		fmt.Println("udp create error:", err)
		return
	}

	defer udp.Close()

	for {
		var data [1024]byte
		fromUDP, u, err := udp.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp data error:", err)
			continue
		}
		fmt.Printf("data:%v addr:%v count:%v \n", string(data[:fromUDP]), u, fromUDP)
		_, err = udp.WriteToUDP(data[:fromUDP], u)
		if err != nil {
			fmt.Println("write to udp error:", err)
			continue
		}
	}
}

func CreateUdpClient() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 7890,
	})

	if err != nil {
		fmt.Println("connect to udp failed,err:", err)
		return
	}

	defer socket.Close()

	sendData := []byte("hello world!")

	_, err = socket.Write(sendData)

	if err != nil {
		fmt.Println("write udp error:", err)
		return
	}

	data := make([]byte, 4096)

	n, remoteAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println("read udp error:", err)
		return
	}

	fmt.Printf("recv:%v addr:%v count%v \n", string(data[:n]), remoteAddr, n)

}
