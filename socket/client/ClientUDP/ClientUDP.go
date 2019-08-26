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
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	// Localhost at port 1200
	service := ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	conn, err := net.DialUDP("udp", nil, addr)
	shared.CheckError(err)

	//defer conn.Close()

	number := os.Args[1]
	request := []byte(number)

	for i := 0; i < shared.SAMPLE_SIZE; i++ {
		t1 := time.Now()

		_, err = conn.Write(request)
		shared.CheckError(err)

		response := make([]byte, 1024)

		_, err = conn.Read(response)
		shared.CheckError(err)

		t2 := time.Now()
		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("Sample: %d - %f\n", i, x)
		fmt.Println(s)
	}

	os.Exit(0)
}
