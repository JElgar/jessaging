package main

import (
	"github.com/gin-gonic/gin"
    "fmt"
    "github.com/jelgar/jessage-back/db"
    . "github.com/jelgar/jessage-back/models"
    "encoding/json"
    "log"
)

func CORSMiddleware() gin.HandlerFunc {
     return func(c *gin.Context) {
         print("Using middleware")
         c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
         c.Writer.Header().Set("Access-Control-Max-Age", "86400")
         c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
         c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
         c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
         c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

         if c.Request.Method == "OPTIONS" {
             c.AbortWithStatus(201)
         } else {
             c.Next()
         }
     }
 }

//var messages []message

func ping(c *gin.Context){
    c.Header("Content-Type", "application/json")
    c.JSON(200, gin.H{
        "word": "Hello",
	})
}

func send(c *gin.Context){
    var m MessageStruct
	c.BindJSON(&m)
    fmt.Printf("Message to send: %v\n", m)
    db.SendMessage(m)

    var messages []*MessageStruct
    messages = db.GetMessages()
    fmt.Println("The first id is: " + messages[0].ID)
    fmt.Println("The first message is: " + messages[0].Message)
    messagesJson, err := json.Marshal(messages[0])
    if err != nil {
        log.Fatal(err)
    }

    fmt.Print(messagesJson)
    //c.JSON(200, gin.H{
	//	"response": m.Message,
	//})
    c.JSON(200, messagesJson)
}

func main() {
        db.Connect()
		r := gin.Default()
        r.Use(CORSMiddleware())
        r.GET("/ping", ping)
        r.POST("/send", send)
        r.Run(":8080")
}
