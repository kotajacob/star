// License: GPL-3.0-only
// (c) 2022 Dakota Walsh <kota@nilsu.org>
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	sourcehut "git.sr.ht/~sircmpwn/sourcehut-go"
	"git.sr.ht/~sircmpwn/sourcehut-go/git.sr.ht"
	"github.com/google/go-github/v46/github"
	"github.com/muesli/reflow/wordwrap"
	"golang.org/x/oauth2"
)

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	// Setup API clients for any access tokens we found.
	ghClient := github.NewClient(nil)
	ghToken := os.Getenv("GITHUB_AUTH_TOKEN")
	if ghToken != "" {
		ghTS := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: ghToken},
		)
		ghTC := oauth2.NewClient(context.Background(), ghTS)
		ghClient = github.NewClient(ghTC)
	}
	shClient := sourcehut.NewSrhtClient(
		os.Getenv("SOURCEHUT_AUTH_TOKEN"),
		http.DefaultClient,
	)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		u, err := url.Parse(scanner.Text())
		if err != nil || u.String() == "" {
			continue
		}

		switch u.Hostname() {
		case "github.com":
			s, err := renderGithub(ghClient, u)
			if err != nil {
				if _, ok := err.(*github.RateLimitError); ok {
					log.Fatalln("hit github rate limit")
				}
				// For input errors (such as providing a github user URL instead
				// of a repository URL) we simply skip the line and continue.
				continue
			}
			fmt.Print(s)
		case "git.sr.ht":
			s, err := renderSourcehut(shClient, u)
			if err != nil {
				if _, ok := err.(sourcehut.SrhtErrorResponse); ok {
					log.Fatalf("sourcehut access denied: %v\n", err)
				}
				// For input errors we simply skip the line and continue.
				continue
			}
			fmt.Print(s)
		default:
			fmt.Println(u)
		}
		fmt.Println()
	}
}

type nonRepoError struct {
	u *url.URL
}

func (e *nonRepoError) Error() string {
	return fmt.Sprintf("URL %v: missing username and/or repository", e.u)
}

// repoName parses a git forge URL and returns the username and reponame or an
// error.
func repoName(u *url.URL) (string, string, error) {
	path := strings.TrimPrefix(u.Path, "/")
	split := strings.Split(path, "/")
	if len(split) < 2 {
		return "", "", &nonRepoError{u: u}
	}
	return split[0], split[1], nil
}

func renderGithub(client *github.Client, u *url.URL) (string, error) {
	username, reponame, err := repoName(u)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	repo, _, err := client.Repositories.Get(
		context.Background(),
		username,
		reponame,
	)
	if err != nil {
		return "", err
	}

	b.WriteString(repo.GetHTMLURL() + "\n")

	desc := repo.GetDescription()
	if desc != "" {
		b.WriteString(wordwrap.String("Desc: "+desc, 80) + "\n")
	}

	lang := repo.GetLanguage()
	if lang == "" {
		lang = "Unknown"
	}

	b.WriteString("Lang: " + lang + "\n")
	b.WriteString("Stars: " + strconv.Itoa(repo.GetStargazersCount()) + "\n")

	if len(repo.Topics) != 0 {
		topics := strings.Join(repo.Topics, ", ")
		b.WriteString(wordwrap.String("Topics: "+topics, 80) + "\n")
	}

	return b.String(), nil
}

func renderSourcehut(client *sourcehut.SrhtClient, u *url.URL) (string, error) {
	username, reponame, err := repoName(u)
	if err != nil {
		return "", err
	}

	gitClient := git.NewGitClient(client)
	repo, err := gitClient.GetRepository(username, reponame)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	b.WriteString(u.String() + "\n")
	b.WriteString(wordwrap.String("Desc: "+repo.Description+"\n", 80))

	return b.String(), nil
}
