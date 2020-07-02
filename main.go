package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Item struct
type Item struct {
	Title string
	Body  string
}

var database []Item

// API object
type API int

// GetByName get the item for the database and return Item
func (a *API) GetByName(title string, reply *Item) error {
	var getItem Item

	for _, val := range database {
		if val.Title == title {
			getItem = val
		}
	}
	*reply = getItem
	return nil
}

// AddItem return void
func (a *API) AddItem(item Item, reply *Item) error {
	database = append(database, item)
	*reply = item
	return nil
}

// EditItem return Item
func (*API) EditItem(edit Item, reply *Item) error {
	var changed Item
	for idx, val := range database {
		if val.Title == edit.Title {
			database[idx] = Item{edit.Title, edit.Body}
			changed = edit
		}
	}
	*reply = changed
	return nil
}

// DeleteItem return item
func (*API) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range database {
		if val.Title == item.Title && val.Body == item.Body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del

	return nil
}

// GetDB return a database wraped
func (a *API) GetDB(title string, reply *[]Item) error {
	*reply = database
	return nil
}

// main return void
func main() {
	var api = new(API)
	err := rpc.Register(api)

	if err != nil {
		log.Fatal("error registering API to RPC", err)
	}

	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", ":4040")

	if err != nil {
		log.Fatal("Listener error ", err)
	}

	log.Printf("serving rpc on por %d", 4040)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("error serving: ", err)
	}

	// fmt.Println("Initial database: ", database)
	// a := Item{"first", "This is the first item on the db"}
	// b := Item{"second", "This is the second item on the db"}
	// c := Item{"third", "This is the third item on the db"}

	// api.AddItem(a, reply)
	// AddItem(b)
	// AddItem(c)
	// fmt.Println("Second database is: ", database)

	// DeleteItem(b)
	// fmt.Println("Third database is: ", database)

	// EditItem("third", Item{"third new", "The new Item"})
	// fmt.Println("fourth database is: ", database)

	// x := GetByName("first")
	// y := GetByName("third")

	// fmt.Println(x, y)

}
