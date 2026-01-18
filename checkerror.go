package tardigrade

// Version 0.3.0 - Sun Jan 18 09:38:18 PM GMT 2026

import "log"

// CheckError function takes in a string and the error code!
func CheckError(msg string, err error) {
	if err != nil {
		log.Println(msg)
		panic(err)
	}
}
