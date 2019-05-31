package main

import (
	"github.com/gin-gonic/gin"
    "fmt"
    "github.com/jelgar/jessage-back/db"
    . "github.com/jelgar/jessage-back/models"
//    "encoding/json"
//    "log"
)

func CORSMiddleware() gin.HandlerFunc {
     return func(c *gin.Context) {
         //print("Using middleware")
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

func ping(c *gin.Context){
    c.Header("Content-Type", "application/json")
    c.JSON(200, gin.H{
        "word": "Hello",
	})
}

func createAccount(c *gin.Context) {
    var u User
    c.BindJSON(&u)
}

func login(c *gin.Context) {
    //Check if credentials are correct
    var u User
    c.BindJSON(&u)
    e := db.GetUser(u)
    if e != nil {
        // Return phat error so react can display user not found message
        fmt.Println("User not found")
    }else {
        fmt.Println("User found")
        c.JSON(200, u)
    }
}

func send(c *gin.Context){
    var m MessageStruct
	c.BindJSON(&m)
    fmt.Printf("Message to send: %v\n", m)
    db.SendMessage(m)

    var messages []*MessageStruct
    messages = db.GetMessages()
    
    c.Header("Content-Type", "application/json")
    c.JSON(200, messages)
}

func main() {
        db.Connect()
		r := gin.Default()
        r.Use(CORSMiddleware())
        r.GET("/ping", ping)
        r.POST("/send", send)
        r.POST("/login", login)
        r.POST("/createAccount", createAccount)
        r.Run(":8080")
}
