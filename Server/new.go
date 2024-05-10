package main

import (
    "database/sql"
    "fmt"
    "net"
    "io"
    "strings"
    _ "github.com/go-sql-driver/mysql" // Import MySQL driver package
)
func handleConnection(conn net.Conn, db *sql.DB) {
    defer conn.Close()

    for {
        // Read command from client
        cmdBuf := make([]byte, 1024)
        n, err := conn.Read(cmdBuf)
        if err != nil {
            if err == io.EOF {
                fmt.Println("Client closed the connection.")
                return
            }
            fmt.Println("Error reading command:", err)
            return
        }

        // Extract command from buffer
        command := string(cmdBuf[:n])

        fmt.Println("Received command from client:", command)
        if strings.HasPrefix(strings.ToUpper(command), "SELECT") {
            // Execute command on the database
        rows, err := db.Query(command)
        if err != nil {
            fmt.Println("Error executing command:", err)
            return
        }
        defer rows.Close()
        
        // Iterate over the rows returned by the query
        var responseData []byte
        for rows.Next() {
            var id int
            var name string
            err := rows.Scan(&id, &name)
            if err != nil {
                fmt.Println("Error scanning row:", err)
                return
            }
            // Process id and name here as needed
            fmt.Println("Row data - ID:", id, "Name:", name)

            // Append row data to response
            responseData = append(responseData, fmt.Sprintf("ID: %d, Name: %s\n", id, name)...)
        }

        // Write response back to client
        _, err = conn.Write(responseData)
        if err != nil {
            fmt.Println("Error sending response:", err)
            return
        }
		} else {
			command := string(cmdBuf[:n])

        fmt.Println("Received command from client:", command)
            _, err= db.Exec(command)
            if err != nil {
                fmt.Println("Error scanning rows:", err)
                return
            }
            fmt.Println("Data inserted successfully!")
            return
        }
        // Write acknowledgment back to client
        acknowledgment := "Command executed successfully!"
        _, err = conn.Write([]byte(acknowledgment))
        if err != nil {
            fmt.Println("Error sending acknowledgment:", err)
            return
        }
		}
        
    }


func main() {
    // Open a database connection
    db, err := sql.Open("mysql", "root:123456789@tcp(localhost:3306)/company")
    if err != nil {
        fmt.Println("Error connecting to the database:", err)
        return
    }
    defer db.Close() // Defer closing the database connection

    // Ping the database to check if the connection is successful
    err = db.Ping()
    if err != nil {
        fmt.Println("Error pinging the database:", err)
        return
    }

    fmt.Println("Connected to the database successfully!")

    // Create table if not exists
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255))")
    if err != nil {
        fmt.Println("Error creating table:", err)
        return
    }

    // Listen for incoming connections
    listener, err := net.Listen("tcp", ":8081")
    if err != nil {
        fmt.Println("Error listening:", err)
        return
    }
    defer listener.Close()

    fmt.Println("Server listening on port 8080...")

    for {
        // Accept incoming connection
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }

        // Handle connection
        go handleConnection(conn, db)
    }
}