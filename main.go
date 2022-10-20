package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Repository struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `jsonL:"language"`
	CreatedAt   string `jsonL:"created_at"`
	UpdatedAt   string `jsonL:"updated_at"`
}

func main() {
	repositories, err := fetchGithubRepositoriesInformation()
	if err != nil {
		log.Println(err)
	}

	for _, v := range repositories {
		fmt.Printf("Id: %d Nome: %s Linguagem: %s\n", v.Id, v.Name, v.Language)
	}

}

func fetchGithubRepositoriesInformation() ([]Repository, error) {
	token := os.Getenv("GITHUB_TOKEN")
	bearer := "Bearer " + token

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/gdguesser/repos", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", bearer)
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var repositories []Repository
	err = json.Unmarshal([]byte(bytes), &repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}
