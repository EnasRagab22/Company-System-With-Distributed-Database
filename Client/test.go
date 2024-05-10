// package main

// import (
// 	"fmt"
// 	"net"
// )

// func main() {
// 	// Connect to the server
// 	conn, err := net.Dial("tcp", "192.168.205.57:8080")
// 	if err != nil {
// 		fmt.Println("Error connecting to server:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Send the INSERT command to the server
// 	command := "INSERT"
// 	_, err = conn.Write([]byte(command))
// 	if err != nil {
// 		fmt.Println("Error sending command:", err)
// 		return
// 	}

// 	// Read and print the response from the server
// 	response := make([]byte, 1024)
// 	n, err := conn.Read(response)
// 	if err != nil {
// 		fmt.Println("Error reading response:", err)
// 		return
// 	}
// 	fmt.Println("Response from server:", string(response[:n]))
// }
////////////////////////////////////////////////////////////////////////////
// package main

// import (
//     "fmt"
//     "net"
// )

// func main() {
//     // Connect to the server
//     conn, err := net.Dial("tcp", "192.168.205.57:8080")
//     if err != nil {
//         fmt.Println("Error connecting to server:", err)
//         return
//     }
//     defer conn.Close()

//     // Send the INSERT command to the server
//     insertQuery := "INSERT INTO users (name) VALUES ('Mariom')"
//     _, err = conn.Write([]byte(insertQuery))
//     if err != nil {
//         fmt.Println("Error sending INSERT command:", err)
//         return
//     }

//     // Read and print the response from the server
//     response := make([]byte, 1024)
//     n, err := conn.Read(response)
//     if err != nil {
//         fmt.Println("Error reading response:", err)
//         return
//     }
//     fmt.Println("Response from server (after INSERT):", string(response[:n]))

//     // Send the SELECT command to the server
//     selectQuery := "SELECT * FROM users"
//     _, err = conn.Write([]byte(selectQuery))
//     if err != nil {
//         fmt.Println("Error sending SELECT command:", err)
//         return
//     }

//     // Read and print the response from the server
//     n, err = conn.Read(response)
//     if err != nil {
//         fmt.Println("Error reading response:", err)
//         return
//     }
//     fmt.Println("Response from server (after SELECT):", string(response[:n]))
// }
////////////////////////////////////////////////////////////////////////////////////////*
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	// Connect to the server
// 	conn, err := net.Dial("tcp", "192.168.205.57:8080")
// 	if err != nil {
// 		fmt.Println("Error connecting to server:", err)
// 		return
// 	}
// 	defer conn.Close()

// 	// Get user input for the INSERT command
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter name to insert into database: ")
// 	name, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading input:", err)
// 		return
// 	}

// 	// Send the INSERT command to the server
// 	insertCommand := fmt.Sprintf("INSERT INTO users (name) VALUES ('%s')", name)
// 	_, err = conn.Write([]byte(insertCommand))
// 	if err != nil {
// 		fmt.Println("Error sending INSERT command:", err)
// 		return
// 	}

// 	// Read and print the response from the server after INSERT
// 	response := make([]byte, 1024)
// 	n, err := conn.Read(response)
// 	if err != nil {
// 		fmt.Println("Error reading response after INSERT:", err)
// 		return
// 	}
// 	fmt.Println("Response from server after INSERT:", string(response[:n]))

// 	// Send the SELECT command to the server
// 	selectCommand := "SELECT * FROM users"
// 	_, err = conn.Write([]byte(selectCommand))
// 	if err != nil {
// 		fmt.Println("Error sending SELECT command:", err)
// 		return
// 	}

// 	// Read and print the response from the server after SELECT
// 	n, err = conn.Read(response)
// 	if err != nil {
// 		fmt.Println("Error reading response after SELECT:", err)
// 		return
// 	}
// 	fmt.Println("Response from server after SELECT:", string(response[:n]))
// }
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// 	"os"
// )

// func main() {
// 	// Connect to the server
// 	conn, err := net.Dial("tcp", "192.168.205.57:8081")
// 	if err != nil {
// 		fmt.Println("Error connecting to the server:", err)
// 		return
// 	}
// 	defer conn.Close() // Defer closing the connection

// 	// Read command from user
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter SQL command: ")
// 	command, err := reader.ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error reading command:", err)
// 		return
// 	}

// 	// Send command to the server
// 	_, err = conn.Write([]byte(command))
// 	if err != nil {
// 		fmt.Println("Error sending command to the server:", err)
// 		return
// 	}

// 	// Read acknowledgment from the server
// 	ackBuf := make([]byte, 1024)
// 	n, err := conn.Read(ackBuf)
// 	if err != nil {
// 		fmt.Println("Error reading acknowledgment from the server:", err)
// 		return
// 	}

// 	acknowledgment := string(ackBuf[:n])
// 	fmt.Println("Received acknowledgment from the server:", acknowledgment)
// }

/////////////////////////////////////////////////////////////////////

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
