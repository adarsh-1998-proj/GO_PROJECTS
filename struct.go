package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate time.Time
	PhoneNumber string
	Age int
}

func main(){
	user := User{
		FirstName:   "Adarsh",
		LastName:    "Das",
		PhoneNumber: "8637203980",
		Age:         25,
	} 
	log.Println(user.FirstName , user.LastName , user.BirthDate , user.Age , user.PhoneNumber)
}