package main

import (
	"encoding/json"
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

	// Conect to server ipContainer : 7171
	conn, err := net.Dial("tcp", ipContainer+":"+
		strconv.Itoa(shared.TCP_PORT))
	shared.CheckError(err)

	var msgFromServer string

	jsonDecoder := json.NewDecoder(conn)
	jsonEncoder := json.NewEncoder(conn)

	fmt.Println("Fibonacci,Sample,Time")
	for i := 0; i < shared.SAMPLE_SIZE; i++ {

		t1 := time.Now()
		// Prepares the request
		msgToServer := os.Args[2]

		// Serializes + Send
		err = jsonEncoder.Encode(msgToServer)
		shared.CheckError(err)

		// Recieve + Deserializes
		err = jsonDecoder.Decode(&msgFromServer)
		shared.CheckError(err)
		t2 := time.Now()

		x := float64(t2.Sub(t1).Nanoseconds()) / 1000000
		s := fmt.Sprintf("%s,%d,%f", msgToServer, i, x)
		fmt.Println(s)
	}
	os.Exit(0)

}
