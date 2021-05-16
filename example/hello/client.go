package main

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

func main() {
	for {
		conn, err := net.Dial("tcp", ":1234")
		if err != nil {
			fmt.Println(err)
			time.Sleep(time.Second)
			continue
		}
		for {
			var value []byte
			fmt.Scanln(&value)

			if reflect.DeepEqual(value, []byte("exit")) {
				conn.Close()
				break
			}

			if _, err := conn.Write(value); err != nil {
				fmt.Println(err)
				break
			}
		}
	}
}
