package main

import (
	"bufio"
	"client/functions"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

    client.WriteToFile("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJBY2Nlc3NUb2tlbjMiLCJleHAiOjE2MzAwMDAwMDB9.RS5MZ1l4Ym9Qd2dKZXBTMW5PVHJQZz10", ".\\file.txt")
    client.HttpGetHandlerFIX("http://localhost:8080/health", client.ReadFile(".\\file.txt"))






    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        os.Exit(1)
    }
    defer conn.Close()
    scanner := bufio.NewScanner(os.Stdin)
    // Send a message to the server
    reader := bufio.NewReader(conn) 
    fmt.Fprintf(conn, client.ReadFile(".\\file.txt")) //here the client will send to the server the jwt for validation
    for 0 < 1 {
        fmt.Print(">>")
        scanner.Scan()
        client_input := strings.ReplaceAll(scanner.Text(), " ", "SPACE")
        fmt.Fprintf(conn, client_input+"\n")
        // Read server's response
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
