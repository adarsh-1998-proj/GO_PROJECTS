package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("My string is set to " , myString)
	changeUsingPointer(&myString)
	log.Println("after using pointer and calling the changeUsingPointer func string is set to " , myString)
	
}
func changeUsingPointer(s *string){
	log.Println("my pointer is set to ", s)
	newValue :="Red"
	*s = newValue

}