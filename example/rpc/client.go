package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, _ := rpc.Dial("tcp", ":1234")

	var reply string
	err := client.Call("Say.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
