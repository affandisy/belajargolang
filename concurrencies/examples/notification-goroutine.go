package main

import (
	"fmt"
	"sync"
	"time"
)

type Notification struct {
	UserID  int
	Message string
}

func SendEmailSync(userID int, message string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 1)
	fmt.Printf("Email notification sent to user %d: %s \n", userID, message)
}

func main() {
	notifications := []Notification{
		{UserID: 101, Message: "Your order has been confirmed."},
		{UserID: 202, Message: "Your account has been created."},
		{UserID: 303, Message: "Your payment was successful."},
	}

	var wg sync.WaitGroup
	wg.Add(len(notifications))

	for _, notification := range notifications {
		go SendEmailSync(notification.UserID, notification.Message, &wg)
	}

	wg.Wait()

	fmt.Println("Main application continues...")

	time.Sleep(time.Second * 2)

	fmt.Println("Main application finished")
}
