package models

import (
    "github.com/dgrijalva/jwt-go"
)

type User struct {
    Username string `json:"username" bson:"username"`
    Password string `json:"password" bson:"password"`
    First_Name string `json:"first_name" bson:"first_name"`
    Last_Name string `json:"last_name" bson:"last_name"`
    Email string `json:"email" bson:"email"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
