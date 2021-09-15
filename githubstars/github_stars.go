package githubstars

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RepoInfo struct {
	Id          int    `json:"id"`
	Name        string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	Language    string `json:"language"`
}

type SearchData struct {
	Items []RepoInfo
}

func responseToJson(data []byte) SearchData {
	var searchResult SearchData
	_ = json.Unmarshal(data, &searchResult)
	return searchResult
}

func customQueryParameters(language string) string {
	var queryParameter string
	if language != "" {
		queryParameter = "?q=language:" + language + "&sort=stars&per_page=50" //language query
	} else {
		queryParameter = "?q=is:public&sort=stars&per_page=50" //Default
	}
	return queryParameter
}

func GetTrendingRepos(language string) (SearchData, int) {
	url := "https://api.github.com/search/repositories"

	queryParams := customQueryParameters(language)

	request, _ := http.NewRequest("GET", url+queryParams, nil)

	request.Header.Add("accept", "application/vnd.github.v3+json")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err, resp.StatusCode)
	}

	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	searchResult := responseToJson(data)

	return searchResult, resp.StatusCode
}
