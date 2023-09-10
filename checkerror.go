package main

import "log"

// CheckError function takes in a string and the error code!
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
