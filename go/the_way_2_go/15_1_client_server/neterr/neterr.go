// package of (unnecessary)helper error hanlding functions for the `net' package
package neterr

import (
    "fmt"
    "net"
    "io"
)

// error checking without established connection
func CheckError(err error) bool {
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

// error checking WITH established connection (there are no optional args in
// golang thus exsistance of thi function is necessary)
func CheckErrorConn(err error, conn net.Conn) bool {
    if err != nil {
        CloseAndInform(conn, err.Error())
        return false
    }
    return true
}

// normal closing procedure for client thread
func CloseAndInform(conn net.Conn, msg string) {
    fmt.Println(msg)
    conn.Close()
}

// check if other side of the connection gracely turned off
func CheckTurnOff(err error) bool {
    if err == io.EOF {                               // check other errors also
        return true
    }
    return false
}
