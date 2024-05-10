# Company-System-With-Distributed-Database
Company System With Distributed Database using GO
**Project Description: Distributed Database Server and Client in Go**

**Overview:**
This project implements a simple distributed database system using the Go programming language. It consists of a server and client component that communicate over TCP/IP. The server manages a MySQL database, handling incoming commands from clients, executing SQL queries, and returning results. Clients connect to the server, send SQL commands, and receive responses.

**Components:**
1. **Server:**  
   - The server component is responsible for managing the database and handling client connections.
   - It listens for incoming TCP connections on a specified port.
   - Upon receiving a connection, it spawns a new goroutine to handle the client.
   - It interacts with the MySQL database using the `database/sql` package.
   - The server can execute both `SELECT` queries and data modification queries (`INSERT`, `UPDATE`, `DELETE`).

2. **Client:**
   - The client component allows users to interact with the server by sending SQL commands.
   - It connects to the server via TCP/IP.
   - Users input SQL commands via the command-line interface.
   - Depending on the command type (`SELECT` or data modification), the client receives and displays the server's response.

**Features:**
- **Database Interaction:** Users can execute SQL commands supported by the MySQL database.
- **Client-Server Communication:** The server and client communicate over TCP/IP sockets.
- **Concurrent Handling:** The server handles multiple client connections concurrently using goroutines.
- **Error Handling:** Both the server and client implement error handling for robustness.
- **Feedback:** Clients receive feedback on command execution status from the server.

**Setup:**
1. **Server Setup:**  
   - Install MySQL and configure the database connection parameters in the server code.
   - Compile and run the server code.

2. **Client Setup:**
   - Compile and run the client code, providing the server's IP address and port number.

**Usage:**
1. **Server:**
   - After setup, the server listens for incoming connections on the specified port.
   - It handles client connections, executing SQL commands received from clients.

2. **Client:**
   - Users input SQL commands via the command-line interface.
   - The client sends the command to the server and displays the server's response.

**Notes:**
- Ensure the MySQL database is running and accessible to the server.
- Customize database connection parameters and error handling as needed.
- This project serves as a foundation for building more complex distributed database systems in Go.

**Contributions:**
Contributions to the project are welcome. Users can suggest enhancements, report bugs, or contribute new features through pull requests.

**License:**
This project is open-source and licensed under [LICENSE_NAME]. See the license file for details.

**Authors:**
- Eslam Abdel_Hady Hassanien
- Enas Ragab Abdel_Latif
- Mariem Tarek Saad

**Acknowledgments:**
Special thanks to [list of individuals or organizations] for their contributions or inspiration.
