package main

import (
	"fmt"
	"net"

	"./worker"
)

//main. MCE - Managed Calculator Extensibility
func main() {
	listener, _ := net.Listen("tcp", ":80")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handlerConnection(conn)
	}

	// var temp *worker.Worker = worker.GetDefaultWorker()
}

func handlerConnection(conn net.Conn) {
	defer conn.Close()
	var commingMessage string
	for {
		input := make([]byte, (1024 * 4))
		count, err := conn.Read(input)
		if count == 0 || err != nil {
			fmt.Println(count)
			break
		}
		commingMessage += string(input)
	}
	
	var worker := worker.GetDefaultWorker(1)
	
}
