package main

import (
	"fmt"
	"log"
)

func init() {
	fmt.Print("hello init\n")
	log.SetPrefix("log prefixxx\n")
	log.Println("init started\n")
}

func init() {
	fmt.Println("___________________")
}

func main() {
	fmt.Println("hello main\n")
	log.Println("loggger in golang")
}
