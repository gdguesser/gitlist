package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Repository struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	fetchGithubRepositoriesInformation()
}

func fetchGithubRepositoriesInformation() error {
	token := os.Getenv("GITHUB_TOKEN")
	bearer := "Bearer " + token

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/gdguesser/repos", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var response []Repository
	err = json.Unmarshal([]byte(bytes), &response)
	if err != nil {
		return err
	}

	fmt.Println(response)

	return nil
}
