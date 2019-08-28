package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/arma29/mid-socket-perf/shared"
)

func main() {

	// Get Argument from command Line
	if len(os.Args) != 3 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	ipContainer := os.Args[1]

	service := ipContainer + ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	// Conect to server ipContainer : 7171
	conn, err := net.DialUDP("udp", nil, addr)
	shared.CheckError(err)

	number := os.Args[2]
	fmt.Println("Fibonacci, Sample, Time")
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()

		// Serializes the request: string -> byte
		request := []byte(number)

		// Sends the request
		_, err = conn.Write(request)
		shared.CheckError(err)

		// Byte structure to pass as argument in Read
		response := make([]byte, 1024)

		// Response now has the payload of recieved UDP datagram
		n, _, err := conn.ReadFromUDP(response)
		shared.CheckError(err)

		// Deserializes the response
		_ = string(response[:n])

		t2 := time.Now()

		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("%s, %d, %f", number, i, x)
		fmt.Println(s)
	}

	os.Exit(0)
}
