package db

import (
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    "context"
    "fmt"
     . "github.com/jelgar/jessage-back/models"
   "log"
)

var db *mongo.Database

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://Jelgar:TGf7eLiFyHrjSg5@cluster0-2two8.mongodb.net")

	fmt.Println("Connecting...")
    client, err := mongo.Connect(context.TODO(), clientOptions)

	fmt.Println("Pinging...")
	err = client.Ping(context.TODO(), nil)
	
	if err != nil {
		log.Fatal(err)	
	}
	
    db = client.Database("Jessaging")

	fmt.Println("Connected")
}

func SendMessage(m MessageStruct) {
    collection := db.Collection("messages")
    fmt.Println("Start Collection connection")
    fmt.Println("Done Collection connection")
	fmt.Println(collection)
    fmt.Println("Collection connection :D")
	
	insertResult, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("Added to database")
	fmt.Println(insertResult)
}

func GetMessages() []*MessageStruct{

    var messages []*MessageStruct

    collection := db.Collection("messages")

    findOptions := options.Find()
    
    cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
    if err != nil {
        log.Fatal(err)
    }

    for cur.Next(context.TODO()) {
        var m MessageStruct
        err := cur.Decode(&m)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("The message is: " + m.Message)

        messages = append(messages, &m)
    }

    cur.Close(context.TODO())

    return messages
}

func TestPrint() {
	fmt.Println("Printing...")
}
