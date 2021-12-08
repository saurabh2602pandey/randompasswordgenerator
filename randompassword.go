package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 25)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]

	}
	newpassword := string(b)
	fmt.Println("password", newpassword)
	fmt.Println(b)

}
