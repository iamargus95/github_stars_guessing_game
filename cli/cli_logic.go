package cli

import (
	"fmt"
	"iamargus95/githubGuessStars/githubstars"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

const (
	MaxRounds   = 5
	Threshold   = 10
	WinScenario = 4
)

func Start() {

	var (
		language     string
		repositories githubstars.SearchData
		statusCode   int
		roundsWon    int
	)

	printToCLI()

	for {
		language = getLanguage()
		repositories, statusCode = githubstars.GetTrendingRepos(language)
		if statusCode == 442 {
			fmt.Println("Enter a valid language. Eg. Golang, Python, Java.....")
		} else if statusCode == 200 {
			break
		} else {
			log.Fatalf("API Call failed with status Code - %v", statusCode)
		}
	}

	for round := 1; round <= MaxRounds; round++ {
		roundsWon = roundsWon + playRound(round, MaxRounds, repositories.Items, Threshold)
	}

	displayResult(roundsWon, MaxRounds, WinScenario)
}

func printToCLI() {
	fmt.Println("\n\n\t--------------------------- Welcome to the StarGazer Academy CLI Game ----------------------------")
	fmt.Println("\n\tYou will be given a Public repository from github, Your goal is to guess the number of stars it has.")
	fmt.Println("\n\t\tA total of 5 rounds will be played in a game. You win the game if you win 4 rounds")
	fmt.Println("\n\t\t\tYou can guess stars within +-10`%` of the actual number of stars.")
	fmt.Println("\n\t\t\t\t\t\tGOODLUCK!!")
}

func getLanguage() string {
	var language string
	fmt.Printf("\n\n\tEnter a language of your choice (Just press enter if no preference):")
	_, err := fmt.Scanf("%s", language)
	if err != nil && (err.Error() != "unexpected newline") {
		log.Fatal(err)
	}

	return language
}

func displayRoundHeader(roundnumber, totalrounds int) {
	fmt.Println("\n\t**********************************")
	fmt.Printf("\n\t*******  Round No: %v / %v  ********\n", roundnumber, MaxRounds)
	fmt.Println("\n\t**********************************")
}

func randomRepo(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func displayRepoInfo(repo githubstars.RepoInfo) {
	fmt.Printf("\n\n\tRepository Name: %v (%v)\n", repo.Name, repo.Language)
	fmt.Printf("\n\tDescription: %v\n", repo.Description)
	fmt.Printf("\n\tHint - Number of Forks : %v\n", repo.Forks)
}

func getStarsinput() int {
	var guessInput int
	var starsInput string

	for {
		fmt.Printf("\n\n\tGuess the stars of this Repository : ")
		_, err := fmt.Scan(&starsInput)
		if err != nil {
			log.Fatal(err)
		}
		guessInput, err = strconv.Atoi(starsInput)
		if err != nil {
			fmt.Println("\tNumber of lines should be a valid integer")
		} else {
			break
		}
	}

	return guessInput
}

func computeRoundResult(stargazersCount, guessInput int, Threshold float64) int {
	var wins int

	deviation := math.Abs(float64(stargazersCount) - float64(guessInput))
	deviationPercent := (deviation / float64(stargazersCount)) * 100

	if deviationPercent <= Threshold {
		wins = 1
		fmt.Printf("\n\n\tRound WON!!. Actual stars in this repository: %v\n", stargazersCount)
	} else {
		wins = 0
		fmt.Printf("\n\n\tRound LOST!!. Actual stars in this repository: %v\n", stargazersCount)
	}
	fmt.Println("")

	return wins
}

func playRound(roundnumber, MaxRounds int, repositories []githubstars.RepoInfo, Threshold float64) int {

	displayRoundHeader(roundnumber, MaxRounds)

	randomRepoIndex := randomRepo(len(repositories))
	chosenRepo := repositories[randomRepoIndex]

	displayRepoInfo(chosenRepo)
	guessInt := getStarsinput()
	roundsWon := computeRoundResult(chosenRepo.Stars, guessInt, Threshold)

	return roundsWon
}

func displayResult(roundsWon, MaxRounds, WinScenario int) {
	if roundsWon >= WinScenario {
		fmt.Println("\t********* Congratulations! You have won GitHub's Guess the Stars CLI game. *********")
	} else {
		fmt.Println("\t******** Better luck next time. ********")
		fmt.Printf("\tYou only won %v out of %v rounds. Win %v rounds to win the game.\n", roundsWon, MaxRounds, WinScenario)
	}
}
