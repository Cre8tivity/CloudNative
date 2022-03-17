package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{db: map[string]dollars{"shoes": 50, "socks": 5}}
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.Mutex
	db map[string]dollars
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db.db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db.db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]
	price_list, checkP := query["price"]

	if !check || len(item_list) > 1 {
		fmt.Println("more than one item queried; undefined")
	} else if len(item_list) == 0 {
		fmt.Println("item not found")
	}
	if !checkP || len(item_list) > 1 {
		fmt.Println("more than one price queried; undefined")
	} else if !checkP || len(price_list) == 0 {
		fmt.Println("no price assiciated w item")
	}
	db.Lock()
	if _, ok := db.db[item_list[0]]; !ok {
		price, _ := strconv.ParseFloat(price_list[0], 32)
		db.db[item_list[0]] = dollars(price)
		db.Unlock()
		fmt.Fprintf(w, ": %s added to database with value %s\n", item_list[0], dollars(price))
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item already exists in database\n")
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]
	if !check || len(item_list) > 1 {
		fmt.Println("more than one item queried; undefined")
	} else if len(item_list) == 0 {
		fmt.Println("item not found")
	}
	db.Lock()
	if _, ok := db.db[item_list[0]]; ok {
		delete(db.db, item_list[0])
		db.Unlock()
		fmt.Fprintf(w, ": %s was deleted from db\n", item_list[0])
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item doesn't exist in database\n")
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]
	price_list, checkP := query["price"]
	if !check || len(item_list) > 1 {
		fmt.Println("more than one item queried; undefined")
	} else if len(item_list) == 0 {
		fmt.Println("item not found")
	}
	if !checkP || len(item_list) > 1 {
		fmt.Println("more than one price queried; undefined")
	} else if !checkP || len(price_list) == 0 {
		fmt.Println("no price assiciated w item")
	}
	db.Lock()
	if _, ok := db.db[item_list[0]]; ok {
		price, _ := strconv.ParseFloat(price_list[0], 32)
		db.db[item_list[0]] = dollars(price)
		db.Unlock()
		fmt.Fprintf(w, ": %s updated in database with value %s\n", item_list[0], dollars(price))
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "item doesn't exist in database\n")
	}
}
