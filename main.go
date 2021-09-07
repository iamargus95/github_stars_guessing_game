package main

import "fmt"

func main() {
	var language string
	fmt.Scanln(&language)

	data, statusCode := githubstars.GetTrendingRepos(language)
}
