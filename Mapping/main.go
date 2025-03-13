package main

import "fmt"

// Service interface defines a method for replying to messages.
type Service interface {
	ReplyMessage()
}

// MyService is an implementation of the Service interface.
type MyService struct{}

// ReplyMessage prints a specific reply message for MyService.
func (s MyService) ReplyMessage() {
	fmt.Println("Message Replied.")
}

// GeneralService is another implementation of the Service interface.
type GeneralService struct{}

// ReplyMessage prints a general reply message for GeneralService.
func (s GeneralService) ReplyMessage() {
	fmt.Println("General Message Replied.")
}

func main() {
	fmt.Println("Message Inbox") // Print header for the inbox

	// Create a map to associate service IDs with their corresponding Service implementations.
	userMessages := map[string]Service{
		"SERVICE-ID-1": MyService{},      // Service ID 1 maps to MyService
		"SERVICE-ID-2": GeneralService{}, // Service ID 2 maps to GeneralService
	}

	// Iterate over the map and call ReplyMessage for each service.
	for id, service := range userMessages {
		fmt.Println(id)        // Print the service ID
		service.ReplyMessage() // Call ReplyMessage method for the service
	}
}
