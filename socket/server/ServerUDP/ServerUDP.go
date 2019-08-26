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

	conn, err := net.ListenUDP("udp", addr)
	shared.CheckError(err)
	fmt.Println("UDP Server listening at", addr)

	quit := make(chan struct{})
	for i := 0; i < runtime.NumCPU(); i++ {
		go handleConnection(conn, quit)
	}
	<-quit // hang until an error

}

func handleConnection(conn *net.UDPConn, quit chan struct{}) {

	request := make([]byte, 1024)

	n, addr, err := 0, new(net.UDPAddr), error(nil)
	for err == nil {
		n, addr, err = conn.ReadFromUDP(request)
		if err != nil {
			return
		}

		number, _ := strconv.Atoi(string(request[:n]))

		t1 := time.Now()

		response := application.Fibbonacci(number)
		conn.WriteToUDP([]byte(strconv.Itoa(response)), addr)

		t2 := time.Now()

		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("Fibonacci: %d - Process Time: %f", number, x)
		fmt.Println(s)
		fmt.Println("from", addr, "-", request[:n])
	}
	fmt.Println("Listener failed - ", err)
	quit <- struct{}{}

}
