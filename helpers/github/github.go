package github

import (
	"github.com/google/go-github/v56/github"
	"context"
)


func GetPullRequests() ([]*github.PullRequest, error) {
	client := github.NewClient(nil).WithAuthToken("")


	prs, _, err := client.PullRequests.List(context.Background(), "", "", nil)

	return prs, err
}
