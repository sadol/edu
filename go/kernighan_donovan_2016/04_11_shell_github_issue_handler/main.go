// githubv4dev is a test program currently being used for developing githubv4 package.
//
// Warning: It performs some queries and mutations against real GitHub API.
//
// It's not meant to be a clean or readable example. But it's functional.
package main

import (
	"context"
	"encoding/json"
	"flag"
	//"fmt"
	"log"
	"os"
    "errors"
    "strings"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

const (
    DEVELOPER = "sadol"                       // name of the github contributor
    REPOSITORY = "brylog"                            // name of the github repo
    GITHUB_TOKEN = "GITHUB_TOKEN"                        // name of the env var
)

var (
    develFlag = flag.String("D", DEVELOPER, "<string flag>, needs argument of github developer name.")
    repoFlag = flag.String("R", REPOSITORY, "<string flag>, needs argument of github repository name.")
    listFlag = flag.Bool("l", false, "<bool flag>, list github issues.")
    addFlag = flag.Bool("a", false, "<bool flag>, add github issue.")
    changeFlag = flag.String("c", "", "<string flag>, change github issues, needs argument of issue ID.")
    deleteFlag = flag.String("d", "", "<string flag>, delete github issue, needs argument of issue ID.")
)

type githubV4Actor struct {
    Login     githubv4.String
    AvatarURL githubv4.URI `graphql:"avatarUrl(size:72)"`
    URL       githubv4.URI
}

type Issue struct {
    Number         githubv4.Int
    Author         githubV4Actor
    PublishedAt    githubv4.DateTime
    LastEditedAt   *githubv4.DateTime
    Editor         *githubV4Actor
    Body           githubv4.String
    State          githubv4.String
}

// ----------------------------------------------------------------------------
// ---------------------------------MAIN---------------------------------------
// ----------------------------------------------------------------------------
func main() {
	flag.Parse()

    // check environment variables set for github token
    githubToken, ok := os.LookupEnv(GITHUB_TOKEN)
    if !ok {
        panic(errors.New("`" + GITHUB_TOKEN + "' environment variable not set."))
    }

    // creating graphql https client
    src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubToken},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	client := githubv4.NewClient(httpClient)


    // check `flag' arguments

}
// ----------------------------------------------------------------------------
// ----------------------------------------------------------------------------

// printJSON prints v as JSON encoded with indent to stdout. It panics on any error.
// Pretty-printing function, there is no need to beautify this more.
func printJSON(v interface{}) {
	w := json.NewEncoder(os.Stdout)
	w.SetIndent("", "\t")
	err := w.Encode(v)
	if err != nil {
		panic(err)
	}
}

func list_issues(cl *githubv4.Client, user string, repo string) error {
    // list of issues for particular user & repo
    var q struct {
        Repository struct {
            Issues struct {
                Nodes []Issue
                PageInfo struct {               // pagination!!! for slices
                    EndCursor githubv4.String
                    HasNextPage bool
                }
            }  `graphql:"issues(states: OPEN, first:$issuesOnPage, after:$issuesCursor)"`
        } `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
    }

    variables := map[string]interface{}{
        "repositoryOwner": githubv4.String(user),
        "repositoryName":  githubv4.String(repo),
        "issuesOnPage":    githubv4.Int(10),
        "issuesCursor":    (*githubv4.String)(nil),    // null to get first page
    }

    allIssues := make([]Issue, 0)
    // and now gather all pages into one neat slice, paging idiom
    for {
        err := cl.Query(context.Background(), &q, variables)
        if err != nil {
            return err
        }

        allIssues = append(allIssues, q.Repository.Issues.Nodes...)
        if !q.Repository.Issues.PageInfo.HasNextPage { break }
        variables["issuesCursor"] = githubv4.NewString(q.Repository.Issues.PageInfo.EndCursor)
    }
    for _, val := range allIssues { printJSON(val) }
	return nil
}


/*
	// Query some details about a repository, an issue in it, and its comments.
	{
		type githubV4Actor struct {
			Login     githubv4.String
			AvatarURL githubv4.URI `graphql:"avatarUrl(size:72)"`
			URL       githubv4.URI
		}

		var q struct {
			Repository struct {
				DatabaseID githubv4.Int
				URL        githubv4.URI

				Issue struct {
					Author         githubV4Actor
					PublishedAt    githubv4.DateTime
					LastEditedAt   *githubv4.DateTime
					Editor         *githubV4Actor
					Body           githubv4.String
				}  `graphql:"issue(number:$issueNumber)"`
			} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
		}
		variables := map[string]interface{}{
			"repositoryOwner": githubv4.String("sadol"),
			"repositoryName":  githubv4.String("brylog"),
			"issueNumber":     githubv4.Int(1),
			//"commentsFirst":   githubv4.NewInt(1),
			//"commentsAfter":   githubv4.NewString("Y3Vyc29yOjE5NTE4NDI1Ng=="),
		}
		err := client.Query(context.Background(), &q, variables)
		if err != nil {
			return err
		}
		printJSON(q)
		//goon.Dump(out)
		//fmt.Println(github.Stringify(out))
	}
*/
    /*
	// Toggle a 👍 reaction on an issue.
	//
	// That involves first doing a query (and determining whether the reaction already exists),
	// then either adding or removing it.
	{
		var q struct {
			Repository struct {
				Issue struct {
					ID        githubv4.ID
					Reactions struct {
						ViewerHasReacted githubv4.Boolean
					} `graphql:"reactions(content:$reactionContent)"`
				} `graphql:"issue(number:$issueNumber)"`
			} `graphql:"repository(owner:$repositoryOwner,name:$repositoryName)"`
		}
		variables := map[string]interface{}{
			"repositoryOwner": githubv4.String("shurcooL-test"),
			"repositoryName":  githubv4.String("test-repo"),
			"issueNumber":     githubv4.Int(2),
			"reactionContent": githubv4.ReactionContentThumbsUp,
		}
		err := client.Query(context.Background(), &q, variables)
		if err != nil {
			return err
		}
		fmt.Println("already reacted:", q.Repository.Issue.Reactions.ViewerHasReacted)
		if !q.Repository.Issue.Reactions.ViewerHasReacted {
			// Add reaction.
			var m struct {
				AddReaction struct {
					Subject struct {
						ReactionGroups []struct {
							Content githubv4.ReactionContent
							Users   struct {
								TotalCount githubv4.Int
							}
						}
					}
				} `graphql:"addReaction(input:$input)"`
			}
			input := githubv4.AddReactionInput{
				SubjectID: q.Repository.Issue.ID,
				Content:   githubv4.ReactionContentThumbsUp,
			}
			err := client.Mutate(context.Background(), &m, input, nil)
			if err != nil {
				return err
			}
			printJSON(m)
			fmt.Println("Successfully added reaction.")
		} else {
			// Remove reaction.
			var m struct {
				RemoveReaction struct {
					Subject struct {
						ReactionGroups []struct {
							Content githubv4.ReactionContent
							Users   struct {
								TotalCount githubv4.Int
							}
						}
					}
				} `graphql:"removeReaction(input:$input)"`
			}
			input := githubv4.RemoveReactionInput{
				SubjectID: q.Repository.Issue.ID,
				Content:   githubv4.ReactionContentThumbsUp,
			}
			err := client.Mutate(context.Background(), &m, input, nil)
			if err != nil {
				return err
			}
			printJSON(m)
			fmt.Println("Successfully removed reaction.")
		}
	}
    */
//	return nil
//}

