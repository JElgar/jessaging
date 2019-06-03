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

func GetUser(u User) (User, error) {
    var dbUser User
    collection := db.Collection("users")
    fmt.Print("Connected")
    filter := bson.M{"username": u.Username}
    fmt.Print("Collecting")
    // TODO this doesnt work as dbUser is empty
    err := collection.FindOne(context.TODO(), filter).Decode(&dbUser)
    if err != nil {
        fmt.Printf("There was an error trying to get the user")
        fmt.Print(err)
        return dbUser, err
    }else {
        fmt.Print("Printing the user that was found in the DB")
        fmt.Print(dbUser)
        fmt.Print("Printed")
        return dbUser, nil
    }
}

func AddUser(u User) error {
    fmt.Println("Adding user function is a go")
    collection := db.Collection("users")
    // If the user doesnt already exist them do some stuff
    if _, e := GetUser(u); e == nil {
        fmt.Println("The error is not nil so i think the user already exists")
        // TODO return a custom error to say user already exists
        return e
    }else {
        fmt.Println("The user is being added as we speak")
	    insertedUser , err := collection.InsertOne(context.TODO(), u)
	    if err != nil {
            return err
	    }else{
            fmt.Println(insertedUser)
            return nil
        }
    }
}

func TestPrint() {
	fmt.Println("Printing...")
}
