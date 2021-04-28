package network

import (
	"bufio"
	"os"
	"strings"
)

func Network() {

	var tasks = map[string]func(){
		"tcp server":  CreateTcpServer,
		"tcp client":  CreateTcpClient,
		"udp server":  CreateUdpServer,
		"udp client":  CreateUdpClient,
		"http server": CreateHttpServer,
		"http client": CreateHttpClient,
	}

	println("tasks:")
	for k, _ := range tasks {
		println("    ", k)
	}
	println("run app by input key")

	for {

		inputReader := bufio.NewReader(os.Stdin)
		input, _ := inputReader.ReadString('\n')
		inputInfo := strings.Trim(input, "\r\n")

		fn := tasks[inputInfo]

		if fn != nil {
			fn()
			break
		} else {
			println("the key not found in tasks")
		}
	}

}
