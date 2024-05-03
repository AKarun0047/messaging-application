package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// User represents a user in the system
type User struct {
	UserID   int
	UserName string
}

// Message represents a message sent between users
type Message struct {
	SenderID   int
	ReceiverID int
	Content    string
}

// messageLogs stores the message logs for each user
var messageLogs map[int][]Message

// Initialize the random number generator
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Function to check if a message is empty
func isEmptyMessage(content string) bool {
	return content == ""
}

// Function to fetch a random fact from an API
func getRandomFact() string {
	// In a real application, you would make an API call to fetch a random fact
	// For simplicity, we'll return a hardcoded fact here
	return "Did you know that cats can rotate their ears 180 degrees?"
}

// Function to display the menu and prompt for user input
func displayMenu() int {
	fmt.Println("Welcome to the Command-Line Messaging App!")
	fmt.Println("1. Send Message between two users")
	fmt.Println("2. Broadcast Message to all users")
	fmt.Println("3. View Message Log of a User")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("Enter your choice: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Invalid input. Please enter a number.")
		return -1
	}
	return choice
}

// Function to prompt user for input with a custom message
func promptInput(message string) string {
	fmt.Print(message + ": ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

// Function to get an integer input from the user
func promptIntInput(message string) int {
	input := promptInput(message)
	num, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return -1
	}
	return num
}

// Function to send a message between two users
func sendMessage() {
	senderID := promptIntInput("Enter Sender's ID")
	receiverID := promptIntInput("Enter Receiver's ID")
	content := promptInput("Enter Message Content")

	// Create a new message
	message := Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
	}

	// Log the message
	messageLogs[senderID] = append(messageLogs[senderID], message)
	messageLogs[receiverID] = append(messageLogs[receiverID], message)

	fmt.Println("Message sent successfully!")
}

// Function to broadcast a message to all users
func broadcastMessage() {
	senderID := promptIntInput("Enter Sender's ID")
	content := promptInput("Enter Message Content")

	// Create a new message
	message := Message{
		SenderID: senderID,
		Content:  content,
	}

	// Log the message for each user
	for userID := range messageLogs {
		messageLogs[userID] = append(messageLogs[userID], message)
	}

	fmt.Println("Message broadcasted successfully!")
}

// Function to view message log of a user
func viewMessageLog() {
	userID := promptIntInput("Enter User's ID")

	// Check if the user exists in the messageLogs map
	messages, ok := messageLogs[userID]
	if !ok {
		fmt.Println("User ID not found.")
		return
	}

	// Display the message log for the user
	fmt.Printf("Message log for User %d:\n", userID)
	for _, msg := range messages {
		fmt.Printf("From: %d, To: %d - %s\n", msg.SenderID, msg.ReceiverID, msg.Content)
	}
}

// Main function to drive the program
func main() {
	// Initialize messageLogs map
	messageLogs = make(map[int][]Message)

	for {
		choice := displayMenu()

		switch choice {
		case 1:
			sendMessage()
		case 2:
			broadcastMessage()
		case 3:
			viewMessageLog()
		case 4:
			fmt.Println("Exiting...")
			// Display all message logs before exiting
			for userID, messages := range messageLogs {
				fmt.Printf("Message log for User %d:\n", userID)
				for _, msg := range messages {
					fmt.Printf("From: %d, To: %d - %s\n", msg.SenderID, msg.ReceiverID, msg.Content)
				}
			}
			return
		default:
			fmt.Println("Invalid choice. Please enter a valid option.")
		}
	}
}
