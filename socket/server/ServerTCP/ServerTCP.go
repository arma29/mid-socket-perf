package main

import (
	"fmt"
	"net"
	"strconv"
	"github.com/arma29/mid-socket-perf/application"
	"github.com/arma29/mid-socket-perf/shared"
)

func main() {

	// Localhost at port 7171
	service := ":" + strconv.Itoa(shared.TCP_PORT)

	// Get TCP Address
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	shared.CheckError(err)

	// Prepare to Listen
	listener, err := net.ListenTCP("tcp", tcpAddr)
	shared.CheckError(err)
	fmt.Println("Server listening at", service)

	// Infinite loop to listen to connections
	for {

		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	request := make([]byte, 1024)

	n, err := conn.Read(request)
	shared.CheckError(err)

	strRequest := string(request[:n])

	number, err := strconv.Atoi(strRequest)
	shared.CheckError(err)

	response := application.Fibbonacci(number)

	conn.Write([]byte(strconv.Itoa(response)))

}