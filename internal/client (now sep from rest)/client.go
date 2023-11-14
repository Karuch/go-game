package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()
    scanner := bufio.NewScanner(os.Stdin)
    // Send a message to the server
    
    for 0 < 1 {
        fmt.Print(">>")
        scanner.Scan()
        client_input := strings.ReplaceAll(scanner.Text(), " ", "SPACE")
        fmt.Fprintf(conn, client_input+"\n")
        // Read server's response
        reader := bufio.NewReader(conn)
        response_no_enter, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading:", err.Error())
            return
        }
        response := strings.ReplaceAll(response_no_enter, "ENTER", "\n")
        fmt.Print("Message from server: ", response)
        if strings.Contains(response_no_enter, "YOU LOSE!"){
            break
        } else if strings.Contains(response_no_enter, "YOU WON!") {
            break
        }
    }
}
