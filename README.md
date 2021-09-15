# github_stars_guessing_game

## Design a “Guess the Stars” CLI game based on the GitHub API

- [x] Allows users to choose language (optionally).

    - [x] Build default query option.
    - [x] Build query based on input provided.

- [x] Shows trending repositories.

    - [x] Add func to display repos from array of repos.
    - [x] Add hint (forks_count).
    - [x] Randomize it using rand.Seed

- [x] Allows you to guess the number of stars the repo has.

    - [x] Use Scan fmt.Scan.

- [x] Total 5 rounds.

    - [x] Iterate above steps for MaxRounds.

- [x] Each round presents you with a new repo.

    - [x] Have a banner showing round number.

- [x] You win a round by guessing the stars within 10% tolerance.

- [x] You win the game if you get 4 or more correct answers out of 5.

    - [x] Have a banner show round result.

- [x] Write tests for github_stars.go