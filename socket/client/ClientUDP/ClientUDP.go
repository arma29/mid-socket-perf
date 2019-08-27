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
	fmt.Println(ipContainer)

	// container ipContainer at port 1200
	service := ipContainer + ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	shared.CheckError(err)

	defer conn.Close()

	number := os.Args[2]
	// Serializes the request: string -> byte
	request := []byte(number)

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		t1 := time.Now()

		// Sends the request
		_, err = conn.Write(request)
		shared.CheckError(err)

		// Byte structure to pass as argument in Read
		response := make([]byte, 1024)

		// Response now has the payload of recieved UDP datagram
		_, _, err = conn.ReadFromUDP(response)
		shared.CheckError(err)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("Fibonacci: %s - %s - Sample: %d - %f", number, string(response), i, x)
		fmt.Println(s)
	}

	os.Exit(0)
}
