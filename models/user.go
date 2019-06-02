package models

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    First_Name string `json:"first_name"`
    Last_Name string `json:"last_name"`
    Email string `json:"email"`
}
