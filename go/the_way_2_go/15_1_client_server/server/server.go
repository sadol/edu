// simple tcp server in golang
package main

import (
    "fmt"
    "net"
    "../messages"
    "../neterr"
)

// struct to use in updates of statuses
type Request struct {
    name string
    status int
}

const (
    protocol = "tcp"
    host = "localhost"
    port = "50000"
)

func main () {
    users := make(map[string]int)                         // bottle neck object
    chRequest := make(chan Request)                // activity status semaphore
    fmt.Println("Starting server...")
    // listennig socket creation
    listener, err := net.Listen(protocol, net.JoinHostPort(host, port))
    if neterr.CheckError(err) == false { return }
    go userHandler(users, chRequest)             // updating active users list
    for {                 // accepting connections via socket & delegating work
        conn, err := listener.Accept()
        if neterr.CheckError(err) == false { return }
        go commandHandler(conn, listener, users, chRequest)
    }
}

// server request handler
func commandHandler(conn net.Conn, list net.Listener, users map[string]int, chRequest chan Request) {
    var name string = ""

    for {
        buf := make([]byte, 512)
        _, err := conn.Read(buf)

        if neterr.CheckError(err) == false {
            if name != "" {
                // request to change activation status
                chRequest <- Request{name: name, status: 0}           // blocks
            }
            return
        }

        msg := messages.NewMessage(buf, conn, list, &users)

        // fill this variable only the first time
        if name == "" {
            name = string(msg.ExtractName())
        }
        // and try to put it into `users' map
        chRequest <- Request{name: name, status: 1}                   // blocks

        msg.Execute()
    }
}

// update function for user's map ---> aggregate function for subgoroutines
func userHandler(users map[string]int, chRequest chan Request) {
    for {
        someRequest := <-chRequest                                    // blocks
        someName, someStatus := someRequest.name, someRequest.status
        users[someName] = someStatus
    }
}
