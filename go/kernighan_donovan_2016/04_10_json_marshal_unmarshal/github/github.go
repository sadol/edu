// requesting Github issue tracker resonses
package github

import (
    "fmt"
    "time"
    "encoding/json"
    "net/http"
    "net/url"
    "strings"
)

const (
    IssuesURL = "https://api.github.com/search/issues"
)

// helper container
type User struct {
    Login string
    HTMLURL string `json:"html_url"`
}

// response store container
type Issue struct {
    Number int
    HTMLURL string `json:"html_url"`
    Title string
    State string
    User *User
    CreatedAt time.Time `json:"created_at"`
    Body string
}

// response presentation container
type IssuesSearchResult struct {
    TotalCount int `json:"total_count"` // annotation for `json' package to utilize
    Items []*Issue
}

// queries Github issues tracker
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesURL + "?q=" + q)
    if err != nil { return nil, err }
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("search query failed: %s.", resp.Status)
    }

    var result IssuesSearchResult
    // this is STREAMING decoder ( unlike json.Unmarshal)
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil { return nil, err }

    return &result, nil
}
