package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
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
	args := os.Args[1]

	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Language by " + args + " repositories",
		Subtitle: "Usage by language",
	}))

	repositories, err := fetchGithubRepositoriesInformation(args)
	if err != nil {
		log.Println(err)
	}

	var goLangNum int
	var javaLangNum int

	for _, v := range repositories {
		if v.Language == "Go" {
			goLangNum += 1
		} else if v.Language == "Java" {
			javaLangNum += 1
		}
		fmt.Printf("Id: %d Nome: %s Linguagem: %s\n", v.Id, v.Name, v.Language)
	}

	fmt.Println("goLangNum: ", goLangNum)
	fmt.Println("javaLangNum: ", javaLangNum)

	// Put data into instance
	bar.SetXAxis([]string{"Go", "Ts", "Java", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Golang", generateBarItems(goLangNum)).
		AddSeries("Java", generateBarItems(javaLangNum))
	f, _ := os.Create("bar.html")
	bar.Render(f)
}

func fetchGithubRepositoriesInformation(user string) ([]Repository, error) {
	token := os.Getenv("GITHUB_TOKEN")
	bearer := "Bearer " + token

	req, err := http.NewRequest(http.MethodGet, "https://api.github.com/users/"+user+"/repos", nil)
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

// generate random data for bar chart
func generateBarItems(num int) []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.BarData{Value: num})
	}
	return items
}
