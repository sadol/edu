package main

import (
    "./github"
    "fmt"
    "log"
    "os"
//    "sort"
    "time"
)

const (
    DAY = 24                                         // how many hours in a day
    MONTH = DAY * 31                                 // ... in a month
    YEAR = MONTH * 365                               // ... in a year
)

func main() {
    result, err := github.SearchIssues(os.Args[1:])
    if err != nil {
        log.Fatal(err)
    }

    latestIssues := make([]*github.Issue, 0)
    moderateIssues := make([]*github.Issue, 0)
    oldestIssues:= make([]*github.Issue, 0)
    now := time.Now()

    for _, issue := range result.Items {
        diff := int(now.Sub(issue.CreatedAt).Hours())
        switch {
        case diff < MONTH :
            latestIssues = append(latestIssues, issue)
        case diff > MONTH && diff < YEAR :
            moderateIssues = append(moderateIssues, issue)
        default :
            oldestIssues = append(oldestIssues, issue)
        }
    }

    fmt.Printf("%d issues (less than 1 mth old):\n", len(latestIssues))
    for _, issue := range latestIssues {
        fmt.Printf("#%-5d %v %9.9s %.55s\n",
                   issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
    }
    fmt.Println("=======================================")
    fmt.Printf("%d issues (more than 1 mth old && less than 1 year old):\n",
               len(moderateIssues))
    for _, issue := range moderateIssues {
        fmt.Printf("#%-5d %v %9.9s %.55s\n",
                   issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
    }
    fmt.Println("=======================================")
    fmt.Printf("%d issues (more than 1 year old):\n", len(oldestIssues))
    for _, issue := range oldestIssues {
        fmt.Printf("#%-5d %v %9.9s %.55s\n",
                   issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
    }
}
