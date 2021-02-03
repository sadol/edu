// simple web app with forms and stuff, VERY crude
package main

import (
    "net/http"
    "io"
    "log"
    "net"
    "sort"
    "strings"
    "strconv"
    "math"
    "errors"
)

const (
    form1 = `<html><body><h1>Statistics</h1><h1></h1>
            <h2>Computes basic statistics for a given list of numbers</h2>
            <h1></h1><form action="#" method="post" name="stats">
            <label for="in_numbers">Numbers (comma or space separated)</label><br>
            <input type="text" id="in_numbers" name="in_numbers"><br>
            <input type="submit" value="Submit"></form><h1></h1>
            <table border ="1" style="width:50%"><th colspan="2">Results</th>
            <tr><td>Numbers</td><td>`
    form2 = `</td></tr><tr><td>Count</td><td>`
    form3 = `</td></tr><tr><td>Mean</td><td>`
    form4 = `</td></tr><tr><td>Median</td><td>`
    form5 = `</td></tr></table></body></html>`
    HOST = "localhost"
    PORT = "9001"
    FORMDIR = "/test2"
    MAINDIR = "/test1"
)

//GET handler for some url
func SimpleServer(w http.ResponseWriter, request *http.Request) {
    io.WriteString(w, "<h1>hello, world</h1>")
}

// POST & GET handler of some url
func FormServer(w http.ResponseWriter, request *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    switch request.Method {
    case "GET":                              // initial interaction with a user
        io.WriteString(w, form1 + form2 + form3 + form4 + form5)
    case "POST":
        request.ParseForm()
        rawInputs := request.FormValue("in_numbers")
        inputs, err := processRawInput(rawInputs)
        var output string
        if err != nil {
            output = form1 + err.Error() + form2 + form3 + form4 + form5
        } else {
            //output = form1 + rawInputs + form2 + string(len(inputs)) + form3
            output = form1 + rawInputs + form2 + strconv.Itoa(len(inputs)) + form3
            output += myMean(inputs) + form4
            output += myMedian(inputs) + form5
        }

        io.WriteString(w, output)
    }
}

// crude implementation of the mean
func myMean(inputs []float64) string {
    var mean float64 = 0.00
    for _, v := range inputs {
        mean += v
    }
    mean = mean / float64(len(inputs))
    return strconv.FormatFloat(mean, 'f', 6, 64)
}

// crude implementation of the median
func myMedian(inputs []float64) string {
    sort.Float64s(inputs)
    var median float64
    l := len(inputs)
    if l == 0 {
        median = 0.00
    } else if l == 1 {
        median = inputs[0]
    } else if int(math.Mod(float64(l), 2)) == 0 {
        median = (inputs[l / 2] + inputs[(l / 2) - 1]) / 2.00
    } else if int(math.Mod(float64(l), 2)) == 1 { median = inputs[(l - 1) / 2]
    }
    return strconv.FormatFloat(median, 'f', 6, 64)
}

// splits raw input string accordingly
func processRawInput(rawInput string) (processed []float64, err error) {
    // at first determine type of the separator
    separators := []string{",", " "}                        // legal separators
    var separator string
    containsComa, containsSpace := strings.Contains(rawInput, separators[0]),
        strings.Contains(rawInput, separators[1])
    almostProcessed := make([]string, 0)
    var justOne float64

    if containsComa && containsSpace == true {
        err = errors.New("Ambiguous separators in one input string!!!")
        return
    } else if containsComa == true && containsSpace == false {
        separator = separators[0]
    } else if containsComa == false && containsSpace == true {
        separator = separators[1]
    } else {
        // check if there is only one value
        if justOne, err = strconv.ParseFloat(rawInput, 64); err == nil {
            processed = append(processed, justOne)
        }
        return                                                // may return err
    }

    almostProcessed = strings.Split(rawInput, separator)
    for _, v := range almostProcessed {
        if len(v) > 0 {                                  // only non-empty vals
            if justOne, err = strconv.ParseFloat(v, 64); err == nil {
                processed = append(processed, justOne)
            } else {                                          // silent zeroing
                processed = make([]float64, 0)
                return
            }
        }
    }
    return
}

func main() {
    // css loading
    //fs := http.FileServer(http.Dir("./stylesheets"))
    //http.Handle("/", fs)

    http.HandleFunc(MAINDIR, SimpleServer)
    http.HandleFunc(FORMDIR, FormServer)
    if err := http.ListenAndServe(net.JoinHostPort(HOST, PORT), nil); err != nil {
        log.Fatalf("ListenAndServe: %v.", err)
    }
}
