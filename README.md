# Command-Line Messaging Application

This is a simple command-line messaging application written in Go. It allows users to send messages between each other, broadcast messages to all users, and view their message logs.
Installation

    Make sure you have Go installed on your system. If not, you can download and install it from the official Go website: https://golang.org/dl/
    Clone this repository to your local machine:

bash

 git clone https://github.com/yourusername/command-line-messaging.git

    Navigate to the project directory:

bash

cd command-line-messaging

    Run the application:

bash

go run message.go

# Usage

Upon running the application, you will be presented with a menu where you can choose from the following options:

    Send Message between two users: Allows you to send a message from one user to another.

    Broadcast Message to all users: Sends a message to all users in the system.

    View Message Log of a User: Displays the message log of a specific user.

    Exit: Exits the application.

# Example Usage
Sending a Message

To send a message between two users, choose option 1 from the menu. You will be prompted to enter the sender's ID, receiver's ID, and the message content.

plaintext

Enter your choice: 1
Enter Sender's ID: 1
Enter Receiver's ID: 2
Enter Message Content: Hello, how are you?
Message sent successfully!

# Broadcasting a Message

To broadcast a message to all users, choose option 2 from the menu. You will be prompted to enter the sender's ID and the message content.

plaintext

Enter your choice: 2
Enter Sender's ID: 1
Enter Message Content: Hello everyone!
Message broadcasted successfully!

# Viewing Message Logs

To view the message log of a specific user, choose option 3 from the menu. You will be prompted to enter the user's ID.

plaintext

Enter your choice: 3
Enter User's ID: 1
Message log for User 1:
From: 1, To: 2 - Hello, how are you?
From: 1, To: 0 - Hello everyone!

# Contribution

Contributions are welcome! If you have any suggestions or find any issues, feel free to open an issue or submit a pull request.

![image](https://github.com/AKarun0047/messaging-application/assets/68144786/cc7e2083-afd3-449e-aab7-e8846981f55d)
