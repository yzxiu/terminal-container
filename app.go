package main

import (
	"fmt"
	"time"

	"github.com/containerd/console"
)

func main() {
	con := console.Current()
	defer con.Reset()
	fmt.Println(con)
	for {
		con.Write([]byte(time.Now().String() + "\n"))
		time.Sleep(time.Second)
	}
}
