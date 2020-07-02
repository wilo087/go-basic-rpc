package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Item struc
type Item struct {
	Title string
	Tody  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"Apple", "This is the apple description"}

	client.Call("API.AddItem", a, &reply)
	fmt.Println("Response: ", reply)

	client.Call("API.GetDB", "", &db)
	fmt.Println("Database: ", db)
}
