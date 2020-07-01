package main

import "fmt"

// Item struct
type Item struct {
	title string
	body  string
}

var database []Item

// GetByName get the item for the database and return Item
func GetByName(title string) Item {
	var getItem Item

	for _, val := range database {
		if val.title == title {
			getItem = val
		}
	}
	return getItem
}

// CreateItem return Item
func CreateItem(item Item) Item {
	database = append(database, item)
	return item
}

// AddItem return void
func AddItem(item Item) Item {
	database = append(database, item)
	return item
}

// EditItem return Item
func EditItem(title string, edit Item) Item {
	var changed Item
	for idx, val := range database {
		if val.title == edit.title {
			database[idx] = edit
			changed = edit
		}
	}

	return changed
}

// DeleteItem return item
func DeleteItem(item Item) Item {
	var del Item
	for idx, val := range database {
		if val.title == item.title && val.body == item.body {
			database = append(database[:idx], database[idx+1:]...)
			del = item
			break
		}
	}
	return del
}

// main return void
func main() {
	fmt.Println("Initial database: ", database)
	a := Item{"first", "This is the first item on the db"}
	b := Item{"second", "This is the second item on the db"}
	c := Item{"third", "This is the third item on the db"}

	AddItem(a)
	AddItem(b)
	AddItem(c)
	fmt.Println("Second database is: ", database)

	DeleteItem(b)
	fmt.Println("Third database is: ", database)

	EditItem("third", Item{"third new", "The new Item"})
	fmt.Println("fourth database is: ", database)

	x := GetByName("first")
	y := GetByName("third")

	fmt.Println(x, y)

}
