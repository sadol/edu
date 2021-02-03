// basic client program to mock network traffic
package main

import (
    "fmt"
    "os"
    "net"
    "bufio"
    "strings"
    "../neterr"
)

const (
    HOST = "localhost"
    PORT = "50000"
)

func main() {
    // initializing connection with the server
    conn, err := net.Dial("tcp", net.JoinHostPort(HOST, PORT))
    if neterr.CheckError(err) == false { return }

    //run server response handler
    //go whoPrinter(conn)

    inputReader := bufio.NewReader(os.Stdin)
    fmt.Println("Name?")  // WARNING: there is no duplicate check on the server side
    clientName, _ := inputReader.ReadString('\n')
    trimmedClient := strings.Trim(clientName, "\n")
    for {
        fmt.Println("Send sth to the server (Q to quit).")
        input, _ := inputReader.ReadString('\n')
        trimmedInput := strings.Trim(input, "\n")

        switch trimmedInput {
        case "Q":
            neterr.CloseAndInform(conn, "Closing client.")
            return
        case "SH":
            _, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
            if neterr.CheckErrorConn(err, conn) == false { return }
            neterr.CloseAndInform(conn, "Closing server (and client by implication).")
            return
        case "WHO":
            _, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
            if neterr.CheckErrorConn(err, conn) == false { return }
            fmt.Println("List of active users:")
            whoPrinter(conn)
        default:
            _, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
            if neterr.CheckErrorConn(err, conn) == false { return }
        }
    }
}

// client function to handle response from the server
func whoPrinter(conn net.Conn) {
    buf := make([]byte, 512)
    _, err := conn.Read(buf)

    if neterr.CheckError(err) == false {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(string(buf))
}
