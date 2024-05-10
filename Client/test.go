package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Connect to the server
	conn, err := net.Dial("tcp", "192.168.205.57:8081")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// Get user input for the command
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter command: ")
	command, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading command:", err)
		return
	}

	// Send the command to the server
	_, err = conn.Write([]byte(command))
	if err != nil {
		fmt.Println("Error sending command:", err)
		return
	}
	// Read and print the response from the server
	if strings.HasPrefix(strings.ToUpper(command), "SELECT") {
		response := make([]byte, 1024)
		n, err := conn.Read(response)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}
		fmt.Println("Response from server:", string(response[:n]))
	} else {
		fmt.Println("Command Excuted Succesfully")
		return
	}

}
