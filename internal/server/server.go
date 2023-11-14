package server

import (
    "bufio"
    "fmt"
    "net"
    "strings"
)

func handleConnection(conn net.Conn) {
    defer conn.Close()

    //initiallize blank and word list
    var word string = "He llo"
    visable_array := strings.Split(word, "")
    inVisable_array := make([]string, len(visable_array))
    copy(inVisable_array[:], visable_array[:])
    for i := 0; i < len(inVisable_array); i++ {
        if inVisable_array[i] != " " {
            inVisable_array[i] = "_"
        }
    }

    for {
        client_input_no_space, err := bufio.NewReader(conn).ReadString('\n')
        if err != nil {
            fmt.Println("Error reading:", err.Error())
            break
        }
        client_input := strings.ReplaceAll(client_input_no_space, "SPACE", " ")
        fmt.Println("Message received:", strings.TrimSpace(client_input))
        doNothurtHM, listState, info := HangmanHandler(inVisable_array, visable_array, word, client_input) //DoNotHurtHM false = hit him
        drawing, life := DrawHM(doNothurtHM, listState, info)
        drawing_for_send := strings.ReplaceAll(drawing, "\n", "ENTER")
        if life < 6 {
            conn.Write([]byte(drawing_for_send+"\n")) // Respond to the client
        } else {
            conn.Write([]byte(fmt.Sprintf("%v YOU LOSE!"+"\n",drawing_for_send))) // Respond to the client
        }
    }
}

func Server() {
    fmt.Println("Starting server...")

    listener, err := net.Listen("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            return
        }
        go handleConnection(conn)
    }
}
