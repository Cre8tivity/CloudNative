package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongodbEndpoint = /*"mongodb://0.0.0.0:27017"*/ "mongodb://10.0.2.15:31450" // Find this from the Mongo container
)

type item struct {
	ID        primitive.ObjectID `bson:"_id"`
	Item      string             `bson:"item"`
	Price     float32            `bson:"price"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.Mutex
	db   *mongo.Collection
	cntx context.Context
	canc context.CancelFunc
}

func main() {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongodbEndpoint),
	)
	checkError(err)
	// select collection from database
	col := client.Database("blog").Collection("posts")
	db := database{db: col}

	db.cntx, db.canc = context.WithTimeout(context.Background(), 1500*time.Second)
	err = client.Connect(db.cntx)
	// disconnect from db when main returns
	defer db.canc()
	defer client.Disconnect(db.cntx)
	checkError(err)

	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	mux.HandleFunc("/create", db.create)
	mux.HandleFunc("/delete", db.delete)
	mux.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe(":8000", mux))
}

func checkError(err error) {
	if errors.Is(err, context.DeadlineExceeded) {
		log.Println("ContextDeadlineExceeded: true")
	}
	if os.IsTimeout(err) {
		log.Println("IsTimeoutError: true")
	}
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			return
		default:
			log.Fatal(err)
		}
	}
}

func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]

	if !check || len(item_list) > 1 {
		fmt.Println("[LOG]: more than one item queried; undefined")
		fmt.Fprintf(w, "[SERVER]: more than one item queried; undefined")
		return
	} else if len(item_list) == 0 {
		fmt.Println("[LOG]: item not found")
		fmt.Fprintf(w, "[SERVER]: item not found")
		return
	}

	filter := bson.M{"item": bson.M{"$elemMatch": bson.M{"$eq": item_list[0]}}}
	var i item
	var err error
	db.Lock()
	if err = db.db.FindOne(db.cntx, filter).Decode(&i); err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "[SERVER]: item doesn't exist in database\n")
		fmt.Printf("[LOG]: item doesn't exist in database\n")
	} else {
		_, err := db.db.DeleteOne(db.cntx, bson.M{"item": item_list[0]}) // delete function
		db.Unlock()
		checkError(err)
		fmt.Fprintf(w, "[SERVER]: %s was deleted from db\n", item_list[0])
		fmt.Printf("[LOG]: %s was deleted from db\n", item_list[0])
	}
}

func (db *database) list(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("here0")
	curs, err := db.db.Find(db.cntx, bson.M{})
	checkError(err)
	var query []bson.M
	if err = curs.All(db.cntx, &query); err != nil {
		fmt.Println("empty database")
		panic(err)
	}
	for _, entry := range query {
		fmt.Fprintf(w, "%s: %.2f\n", entry["item"].(string), entry["price"].(float64))
	}
}

func (db *database) update(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]
	price_list, checkP := query["price"]

	if !check || len(item_list) > 1 {
		fmt.Println("[LOG]: more than one item queried; undefined")
		fmt.Fprintf(w, "[SERVER]: more than one item queried; undefined")
		return
	} else if len(item_list) == 0 {
		fmt.Println("[LOG]: item not found")
		fmt.Fprintf(w, "[SERVER]: item not found")
		return
	}
	if !checkP || len(price_list) > 1 {
		fmt.Println("[LOG]: more than one price queried; undefined")
		fmt.Fprintf(w, "[SERVER]: more than one price queried; undefined")
		return
	} else if !checkP || len(price_list) == 0 {
		fmt.Println("[LOG]: no price assiciated w item")
		fmt.Fprintf(w, "[SERVER]: no price assiciated w item")
		return
	}

	//bson.M{"$set": bson.M{"price": price_list[0]}}
	var i item
	var err error
	price, _ := strconv.ParseFloat(price_list[0], 32)

	filter := bson.M{"item": bson.M{"$elemMatch": bson.M{"$eq": item_list[0]}}}
	// update := bson.M{"$set": bson.M{"price": price_list[0]}}

	db.Lock()                                                                         // lock db to access db
	if err = db.db.FindOne(db.cntx, filter).Decode(&i); err != mongo.ErrNoDocuments { // if not in db
		w.WriteHeader(http.StatusNotFound) // 404
		checkError(err)
		fmt.Printf("[LOG]: item doesn't exist in database\n")
		fmt.Fprintf(w, "[SERVER]: item doesn't exist in database\n")
	} else { // if already in db
		checkError(err)
		// _, err := db.db.UpdateOne(db.cntx, filter, update)
		t := i.CreatedAt
		_, err := db.db.DeleteOne(db.cntx, bson.M{"item": item_list[0]})
		_, err2 := db.db.InsertOne(db.cntx, &item{
			ID:        primitive.NewObjectID(),
			Item:      item_list[0],
			Price:     float32(price),
			CreatedAt: t,
			UpdatedAt: time.Now(),
		})
		db.Unlock()
		checkError(err)
		checkError(err2)
		fmt.Printf("[LOG]: %s updated in database with value %.2f\n", item_list[0], dollars(price))
		fmt.Fprintf(w, "[SERVER]: %s updated in database with value %.2f\n", item_list[0], dollars(price))
	}
}

func (db *database) price(w http.ResponseWriter, req *http.Request) {
	itemq := req.URL.Query().Get("item")
	var i bson.M
	var err error
	filter := bson.M{"item": bson.M{"$elemMatch": bson.M{"$eq": itemq}}}
	if err = db.db.FindOne(db.cntx, filter).Decode(&i); err != mongo.ErrNoDocuments {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "[SERVER]: no such item: %q\n", itemq)
		fmt.Printf("[SERVER]: no such item: %q\n", itemq)
	} else {
		fmt.Fprintf(w, "[SERVER]: Price is %s\n", dollars(i["price"].(float32)))
		fmt.Printf("[LOG]: Price Request for %s", itemq)
	}
}

func (db *database) create(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	item_list, check := query["item"]
	price_list, checkP := query["price"]

	if !check || len(item_list) > 1 {
		fmt.Println("[LOG]: more than one item queried; undefined")
		fmt.Fprintf(w, "[SERVER]: more than one item queried; undefined")
		return
	} else if len(item_list) == 0 {
		fmt.Println("[LOG]: item not found")
		fmt.Fprintf(w, "[SERVER]: item not found")
		return
	}
	if !checkP || len(price_list) > 1 {
		fmt.Println("[LOG]: more than one price queried; undefined")
		fmt.Fprintf(w, "[SERVER]: more than one price queried; undefined")
		return
	} else if !checkP || len(price_list) == 0 {
		fmt.Println("[LOG]: no price assiciated w item")
		fmt.Fprintf(w, "[SERVER]: no price assiciated w item")
		return
	}

	filter := bson.M{"item": bson.M{"$elemMatch": bson.M{"$eq": item_list[0]}}}
	var i item
	var err error
	price, _ := strconv.ParseFloat(price_list[0], 32)

	db.Lock()                                                                         // lock db to access db
	if err = db.db.FindOne(db.cntx, filter).Decode(&i); err != mongo.ErrNoDocuments { // if not in db
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "[SERVER]: item already exists in database\n")
		fmt.Printf("[LOG]: item already exists in database\n")
		checkError(err)
	} else { // if already in db
		res, err := db.db.InsertOne(db.cntx, &item{
			ID:        primitive.NewObjectID(),
			Item:      item_list[0],
			Price:     float32(price),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		checkError(err)
		db.Unlock()
		fmt.Printf("[LOG]: inserted object id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
		fmt.Fprintf(w, "[SERVER]: %s added to database with value %s\n", item_list[0], dollars(price))
		fmt.Printf("[LOG]: %s added to database with value %s\n", item_list[0], dollars(price))
	}
}
