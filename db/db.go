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
	fmt.Println("Pung")
	
    db = client.Database("Jessaging")

	fmt.Println("Connected")
}

func SendMessage(m MessageStruct) {
    collection := db.Collection("messages")
	
	insertResult, err := collection.InsertOne(context.TODO(), m)
	if err != nil {
	    log.Fatal(err)
	}
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
        
        messages = append(messages, &m)
    }

    cur.Close(context.TODO())

    return messages
}

func GetUser(u User) error {
    var dbUser User
    collection := db.Collection("users")
    fmt.Print("Connected")
    filter := bson.M{"username": u.Username}
    fmt.Print("Collecting")
    err := collection.FindOne(context.TODO(), filter).Decode(&dbUser)
    fmt.Print("Err")
    if err != nil {
        fmt.Print(err)
        return err
    }else {
        fmt.Print("Printing")
        fmt.Print(dbUser)
        fmt.Print("Printed")
        return nil
    }
}

func TestPrint() {
	fmt.Println("Printing...")
}
