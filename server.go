package main

import (
"fmt"
"log"
"os"
"net"
"bufio"
"strings"
)

func main(){

	userCount := 1
	maxUsers := 2 // By default

	users := make(map[net.Conn] string) // Map of active connections
	newUser := make(chan net.Conn) // Handle new connection
	addedUser := make(chan net.Conn) // Add new connection
	deadUser := make(chan net.Conn) // Users that have left
	messages := make(chan string) // channel that recieves messages from all users

	server, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func(){ // Launch routine that will accept connections forever
		for{
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if userCount > maxUsers {
				os.Exit(1)
			}
			newUser <- conn // Send to handle new user
		}
	}()

	for{

		select{
		case conn := <-newUser:
			go func(conn net.Conn){ // Ask user for name and information
				reader := bufio.NewReader(conn)
				conn.Write([]byte("Enter name: "))
				userName, _ := reader.ReadString('\n')
				userName = strings.Trim(userName, "\r\n")
				log.Printf("Accepted new user : %s", userName)
				messages <- fmt.Sprintf("Accepted user : [%s]\n\n", userName)

				users[conn] = userName // Add connection
				userCount++ // Increment the usercount

				addedUser <- conn // Add user to pool
			}(conn)

		case conn := <-addedUser: // Launch a new go routine for the newly added user

			go func(conn net.Conn, userName string) {
				reader := bufio.NewReader(conn)
				for {
					newMessage, err := reader.ReadString('\n')
					newMessage = strings.Trim(newMessage, "\r\n")
					if err != nil{
						break
					}
					// Send to messages channel therefore ring every user
					messages <- fmt.Sprintf(">%s: %s \a\n", userName, newMessage)
				}

				deadUser <- conn // If error occurs, connection has been terminated
				messages <- fmt.Sprintf("%s disconnected\n", userName)
			}(conn, users[conn])

		case message := <- messages: // If message recieved from any user

			for conn, _ := range users { // Send to all users
				go func(conn net.Conn, message string){ // Write to all user connections
					_, err := conn.Write([]byte(message))

					if err != nil{
						deadUser <- conn
					}
				}(conn, message)
			log.Printf("New message: %s", message)
			log.Printf("Sent to %d users", len(users))
			}

		case conn := <- deadUser: // Handle dead users
			log.Printf("Client disconnected")
			delete(users, conn)
			userCount--
		}
	}
}
