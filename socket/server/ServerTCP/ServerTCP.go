package main

import (
	"encoding/json"
	"net"
	"strconv"

	"github.com/arma29/mid-socket-perf/application"
	"github.com/arma29/mid-socket-perf/shared"
)

func main() {

	// Listening in all interfaces , port number 7171
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(shared.TCP_PORT))
	shared.CheckError(err)

	// fmt.Println("Fibonacci,From,Time")

	// Infinite loop to listen to connections
	for {

		conn, err := listener.Accept()
		shared.CheckError(err)

		defer conn.Close()

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	var msgFromClient string

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	// Same loop from Client Side
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		// Recieve + Deserializes
		err := jsonDecoder.Decode(&msgFromClient)
		shared.CheckError(err)

		// t1 := time.Now()
		// Prepare the response
		number, _ := strconv.Atoi(msgFromClient)
		msgToClient := application.Fibbonacci(number)
		// t2 := time.Now()

		// Serializes + Send
		err = jsonEncoder.Encode(strconv.Itoa(msgToClient))
		shared.CheckError(err)

		// x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		// s := fmt.Sprintf("%d,%s,%f", number, conn.RemoteAddr(), x)
		// fmt.Println(s)
	}
}
