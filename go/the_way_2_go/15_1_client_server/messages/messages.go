// handling messages on the server side
package messages

import (
    "bytes"
    "fmt"
    "os"
    "net"
)

// struct to integrate message data & methods on the server side
type Message struct {
    rawStream []byte
    connection net.Conn
    listener net.Listener
    users *map[string]int
}

// stringificator
func (m *Message)String() string {
    return string(m.rawStream)
}

// factory function for creating new instances of Message
func NewMessage(r []byte, conn net.Conn, list net.Listener, users *map[string]int) (message *Message) {
    message = new(Message)
    message.rawStream = r
    message.connection = conn
    message.listener = list
    message.users = users
    return
}

// handler function for `SH' command
func (m *Message)shHandler() {
    fmt.Println("Server recieved kill command.")
    m.connection.Close()
    m.listener.Close()
    os.Exit(1)
}

// handler function for `WHO' command
func (m *Message)whoHandler() {
    activeUsers := ""
    for i, v := range *m.users {
        activeUsers += fmt.Sprintf("user: %v\tactive: %v.\n", i, v)
    }
    _, _ = m.connection.Write([]byte(activeUsers))  // TODO: check for len(activeUsers)
}

// extracts user name from clients's message stream
func (m *Message)ExtractName() (name []byte) {
    stop := bytes.IndexByte(m.rawStream, byte(':')) + 1 // index of the first msg byte
    if stop != -1 {
        name = m.rawStream[:stop - 6]        // to remove ` says' from the name
    }                                                         // may return nil
    return
}

// extracts message content from clients's message stream
func (m *Message)ExtractMsg() (msg []byte) {
    start := bytes.IndexByte(m.rawStream, byte(':')) + 1 // index of the first msg byte
    stop := bytes.IndexByte(m.rawStream, byte(0))        // index of the first empty byte
    if stop != -1 || start != -1 {
        msg = m.rawStream[start:stop]
    }                                                         // may return nil
    return
}

// run handler on demand
func (m *Message)Execute() {
    switch command := string(m.ExtractMsg()); command {
        case "SH": m.shHandler()
        case "WHO": m.whoHandler()
        default: fmt.Println(m)
    }
}
