// simple web app with forms and stuff, version using TEMPLATES
package main

import (
    "net/http"
    "io"
    "log"
    "net"
    "html/template"
    //"html"
    "./metrics"
)

const (
    FROM = `<html><body><h1>Statistics</h1><h1></h1>
            <h2>Computes basic statistics for a given list of numbers</h2>
            <h1></h1><form action="#" method="post" name="stats">
            <label for="in_numbers">Numbers (comma or space separated)</label><br>
            <input type="text" id="in_numbers" name="in_numbers"><br>
            <input type="submit" value="Submit"></form><h1></h1>
            {{ if . }}
                <table border ="1" style="width:50%"><th colspan="2">Results</th>
                <tr><td>Numbers</td><td>
                {{ .GetRawList }}
                </td></tr><tr><td>Count</td><td>
                {{ .GetLen }}
                </td></tr><tr><td>Mean</td><td>
                {{ .GetMean }}
                </td></tr><tr><td>Median</td><td>
                {{ .GetMedian }}
                </td></tr></table>
            {{ end }}
            </body></html>`

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
    templ := template.Must(template.New("site").Parse(FORM))

    switch request.Method {
    case "GET":                          // initial interaction with a user
        if err := templ.Execute(w, nil); err != nil {
            log.Fatal("Parsing GET response: ", err)
        }
    case "POST":
        request.ParseForm()
        rawInputs := request.FormValue("in_numbers")
        if metr, err := metrics.NewMetrics(rawInputs, []string{",", " "}); err != nil {
            err = templ.Execute(w, nil)                      // just like a GET
        } else {                                                 // OK; proceed
            if err = templ.Execute(w, metr); err != nil {
                log.Fatal("Parsing POST response: ", err)
            }
        }
    }
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
