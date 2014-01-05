package main

import (
	"../../pusher"
	"fmt"
	"time"
)

func main() {
	client := pusher.NewClient("4115", "23ed642e81512118260e", "cd72de5494540704dcf1")

	done := make(chan bool)

	go func() {
		channel, err := client.Channel("common", nil)
		if err != nil {
			fmt.Printf("Error %s\n", err)
		} else {
			fmt.Println(channel)
		}

		authInfo, errAuth := channel.Authenticate("1234.1234", nil)
		if errAuth != nil {
			fmt.Printf("Error %s\n", errAuth)
		} else {
			fmt.Println(authInfo)
		}
		done <- true
	}()

	select {
	case <-done:
		fmt.Println("Done :-)")
	case <-time.After(1 * time.Minute):
		fmt.Println("Timeout :-(")
	}
}
