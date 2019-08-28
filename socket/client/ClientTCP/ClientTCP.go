package main

import (
	"fmt"
	"io/ioutil"
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

	// container ipContainer at port 7171
	service := ipContainer + ":" + strconv.Itoa(shared.TCP_PORT)

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	shared.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
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

		// Recieve the request
		result, err := ioutil.ReadAll(conn)
		shared.CheckError(err)

		// Deserializes the response
		_ = string(result)

		t2 := time.Now()

		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("%s, %d, %f", number, i, x)
		fmt.Println(s)
	}
	os.Exit(0)

}
