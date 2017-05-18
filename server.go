package main

import (
"fmt"
"log"
"os"
"net"
"bufio"
)

func main(){

	userCount := 0 

	users := make(map[net.Conn] int) // Map of active connections

	newUser := make(chan net.Conn) // New connection
	deadUser := make(chan net.Conn) // Users that have left

	messages := make(chan string) // channel that recieves messages from all users
	announcements := make(chan string) // channel that recieves user announcements

	server, err := net.Listen("tcp", ":6000")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	go func(){ // Launch thread that will accept connections forever
		for{
			conn, err := server.Accept()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			newUser <- conn
			announcements <- fmt.Sprintf("Accepted new user, #%d\n", userCount)
		}
	}()

	for{

		select{

		case conn := <-newUser: // If new connection
			log.Printf("Accepted new user, #%d", userCount)
			fmt.Print('\a')
			// announcements <- fmt.Sprintf("Accepted new user, #%d", userCount)

			users[conn] = userCount // Add connection
			userCount++ // Increment the usercount

			go func(conn net.Conn, userId int) { // launch a thread that handles messages (1 thread per user)
				reader := bufio.NewReader(conn)
				for {
					newMessage, err := reader.ReadString('\n')
					if err != nil{
						break
					}
					messages <- fmt.Sprintf("User #%d : %s\a\n", userId, newMessage) // Send to messages channel and ring every user

				}

				deadUser <- conn // If error occurs, connection has been terminated 
			}(conn, users[conn])

		case message := <- messages: // If message recieved from any user

			for conn, _ := range users { // Send to all users
				go func(conn net.Conn, message string){
					_, err := conn.Write([]byte(message)) // Write to all user connections

					if err != nil{
						deadUser <- conn
					}
				}(conn, message)
			log.Printf("New message: %s", message)
			log.Printf("Sent to %d users", len(users))
			}

		case conn := <- deadUser: // Handle dead users
			announcements <- fmt.Sprintf("Client %d disconnected\n", users[conn]) // Announce that user has left
			log.Printf("Client %d disconnected", users[conn])
			delete(users, conn)

		case message := <- announcements: // Send announcement to all users

			for conn, _ := range users {
				go func(conn net.Conn, message string){
					_, err := conn.Write([]byte(message))

					if err != nil{
						deadUser <- conn
					}
				}(conn, message)
			}
		}
	}
}
