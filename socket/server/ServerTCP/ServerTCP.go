package main

import (
	"fmt"
	"net"
	"strconv"
	"time"

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

	fmt.Println("Fibonacci, From, Time")

	// Infinite loop to listen to connections
	for {

		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Byte structure to pass as argument in Read
	request := make([]byte, 1024)

	// for i := 0; i < shared.SAMPLE_SIZE; i++ {

	// Read Client Request
	n, err := conn.Read(request)
	shared.CheckError(err)

	strRequest := string(request[:n])
	// Deserializate the request
	number, err := strconv.Atoi(strRequest)
	shared.CheckError(err)

	t1 := time.Now()

	response := application.Fibbonacci(number)
	// Serializate and Sends the response
	conn.Write([]byte(strconv.Itoa(response)))

	t2 := time.Now()

	x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
	s := fmt.Sprintf("%d, %s, %f", number, conn.RemoteAddr(), x)
	fmt.Println(s)
	// }
	// for {
	// }

}
