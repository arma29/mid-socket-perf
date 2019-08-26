package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"github.com/arma29/mid-socket-perf/shared"
)

func main() {

	// Get Argument from command Line
	if len(os.Args) != 2 {
		fmt.Printf("Missing arguments: %s number\n", os.Args[0])
		os.Exit(1)
	}

	// Localhost at port 7171
	service := ":" + strconv.Itoa(shared.TCP_PORT)

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	shared.CheckError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	shared.CheckError(err)

	number := os.Args[1]
	request := []byte(number)

	_, err = conn.Write(request)
	shared.CheckError(err)

	result, err := ioutil.ReadAll(conn)
	shared.CheckError(err)

	fmt.Println(string(result))
	os.Exit(0)

}