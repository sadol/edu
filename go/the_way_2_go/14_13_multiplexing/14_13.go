// Multiplexing in golang using idioms
package main

import (
    "fmt"
    "strconv"
)

// Request object to ask server AND store server's reply
type Request struct {
    a, b int
    chReply chan int // no need to store the channel's value in the dedicated struct (instead of `int') in this simple case
}

// Stringificator for Request object
func (r *Request) String() string {
    return "Request{a:" + strconv.Itoa(r.a) + ", b:" + strconv.Itoa(r.b) + "}"
}

// Operator function header for server to use
type binOp func(a, b int) int

// Operator function for server to use
func run(operation binOp, req *Request) {
    req.chReply <-operation(req.a, req.b)
}

// Server function(with on demand termination functionality)
func server(operation binOp, chService chan *Request, chTerminate chan bool) {
    for {
        select {
        case req := <-chService:
            go run(operation, req)
        case <-chTerminate:
            return
        }
    }
}

// Start server is an IDIOMATIC HELPER, very nice function in golang in which
// channels are used as convinient communication medium between main thread &
// goroutines.
func startServer(operation binOp) (chService chan *Request, chTerminate chan bool) {
    chService = make(chan *Request)
    chTerminate = make(chan bool)
    go server(operation, chService, chTerminate)
    return
}

// and driver, finally
func main () {
    chAdder, chTerminator := startServer(func (a, b int) int { return a + b })  // inline

    // clients mockery
    const N = 100
    var reqs [N]Request
    for i := 0; i < N; i++ {
        req := &reqs[i]
        req.a = i
        req.b = i + N
        req.chReply = make(chan int)
        chAdder <- req
    }

    // do something with results; note there is NO GUARANTEE that all elements
    // are properly services by server at this point (no additional
    // synchronization)
    for i := N - 1; i >= 0; i-- {
        if <-reqs[i].chReply != N + 2 * i {
            fmt.Printf("Illegal alien at: %v.\n", i)
        } else {
            fmt.Println(&reqs[i])
        }
    }
    chTerminator <-true
    fmt.Println("Done.")
}
