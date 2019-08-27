package main

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"time"

	"github.com/arma29/mid-socket-perf/application"
	"github.com/arma29/mid-socket-perf/shared"
)

func main() {

	service := ":" + strconv.Itoa(shared.UDP_PORT)

	addr, err := net.ResolveUDPAddr("udp", service)
	shared.CheckError(err)

	// Listening in all interfaces , port number 1200
	conn, err := net.ListenUDP("udp", addr)
	shared.CheckError(err)

	fmt.Println("Fibonacci, From, Time")

	// Signalling channel
	done := make(chan struct{})

	// Parallel -> starts multiple go routines
	// Each one do the ReadFromUDP loop
	for i := 0; i < runtime.NumCPU(); i++ {
		// go routine -> thread
		go handleConnection(conn, done)
	}
	// Wait for the loop finishes -> until error
	<-done
}

func handleConnection(conn *net.UDPConn, done chan struct{}) {

	// Byte structure to pass as argument in ReadFromUDP
	request := make([]byte, 1024)
	n, addr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		// Reads a payload of the recieved UDP datagram
		// Copy the payload into 'request'
		// n -> number of bytes copied into 'request'
		// addr -> source address
		n, addr, err = conn.ReadFromUDP(request)

		// Deserialization: byte -> string -> int
		number, _ := strconv.Atoi(string(request[:n]))

		t1 := time.Now()

		response := application.Fibbonacci(number)
		// Sends the serialized response: int -> string -> byte to addr
		conn.WriteToUDP([]byte(strconv.Itoa(response)), addr)

		t2 := time.Now()

		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("%d, %s, %f", number, addr, x)
		fmt.Println(s)
	}
	// Error
	fmt.Println("Listener failed - ", err)
	done <- struct{}{}

}
