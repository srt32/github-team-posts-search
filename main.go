package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

// define the github post struct
type Post struct {
	URL       string `json:"html_url"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	ID        int    `json:"id"`
	CreatedAt string `json:"created_at"`
}

// Usage: GITHUB_TOKEN=<token> go run main.go -org=github -team=engineering -query=mysql1
func main() {
	orgName := flag.String("org", "", "The name of the organization")
	teamSlug := flag.String("team", "", "The slug of the team")
	query := flag.String("query", "", "The query to search for")

	flag.Parse()

	if *orgName == "" || *teamSlug == "" {
		log.Fatal("You must specify an organization and team name")
	}

	// fetch the team posts from the github rest api
	posts, err := GetTeamPosts(*orgName, *teamSlug)

	// check the error
	if err != nil {
		log.Fatal(err)
	}

	matches := []Post{}

	// search the post bodies for the query
	for _, post := range posts {
		if strings.Contains(post.Body, *query) || strings.Contains(post.Title, *query) {
			matches = append(matches, post)
		}
	}

	// print the matches
	for _, match := range matches {
		log.Printf("%s - %s - %s", match.CreatedAt, match.Title, match.URL)
	}
}

// GetTeamPosts returns the most recent 100 posts from the team
func GetTeamPosts(orgName string, teamSlug string) ([]Post, error) {
	// create an http client
	client := &http.Client{}
	// make a get request to the github api
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/orgs/%s/teams/%s/discussions?per_page=100", orgName, teamSlug), nil)

	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("You must set the GITHUB_TOKEN environment variable")
	}

	// set the authorization header
	req.Header.Set("Authorization", "token "+token)

	// make the request
	resp, err := client.Do(req)
	// check the error
	if err != nil {
		// log the error
		log.Fatal("Error making request: ", err)
	}

	// check the resp status code
	if resp.StatusCode != 200 {
		log.Fatal("Error: ", resp.StatusCode)
	}

	// close the response body
	defer resp.Body.Close()

	// parse the response body
	var posts []Post
	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		log.Fatal("Error parsing response: ", err)
	}

	// return the posts
	return posts, nil
}
