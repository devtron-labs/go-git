package main

import (
	"fmt"
	"os"

	git "github.com/devtron-labs/go-git"
	"github.com/devtron-labs/go-git/plumbing/transport/http"
)

func main() {
	CheckArgs("<url>", "<directory>", "<github_access_token>")
	url, directory, token := os.Args[1], os.Args[2], os.Args[3]

	// Clone the given repository to the given directory
	Info("git clone %s %s", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		// The intended use of a GitHub personal access token is in replace of your password
		// because access tokens can easily be revoked.
		// https://help.github.com/articles/creating-a-personal-access-token-for-the-command-line/
		Auth: &http.BasicAuth{
			Username: "abc123", // yes, this can be anything except an empty string
			Password: token,
		},
		URL:      url,
		Progress: os.Stdout,
	})
	CheckIfError(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	CheckIfError(err)
	// ... retrieving the commit object
	commit, err := r.CommitObject(ref.Hash())
	CheckIfError(err)

	fmt.Println(commit)
}
