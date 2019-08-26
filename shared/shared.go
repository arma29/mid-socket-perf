package shared

import (
	"fmt"
	"os"
)

const UDP_PORT = 1200
const TCP_PORT = 7171
const SAMPLE_SIZE = 3000
const PORT_LIMIT = 6000

func CheckError(err error) {
	if err != nil {
		fmt.Println("Fatal error: ", err.Error())
		os.Exit(1)
	}
}
