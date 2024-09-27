package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type Game struct {
	Rounds []*Round
	Total  int
}

// Track the data of each round
type Round struct {
	Round   int   // The current round
	Score   int   // The current score
	Numbers []int // History of numbers, just in case the user wants to know
}

func main() {
	// Let us welcome the user to the game with a brief overview of the rules
	fmt.Println(`
        Welcome to The Game!
    ----------------------------
        
    The game is simple.
	There are 5 rounds.
    A random number is chosen between 2 and 10.
    The number is added to your score.
    You will be given the number.
    
    You can choose to 'continue' or 'stop' the game at any point.
    If you choose to 'continue', a new number will be chosen and added to your score.
    If you choose to 'stop', the game will end and the sum of the numbers you continued with will be your score.
    
    DO NOT let your score in the round exceed 25! A score of more then 25 will cause you to lose the game.
    
    Understand? Great! Let's begin....
    `)

	// Keep going until the user stops
	for {

		// New game
		game := Game{
			Rounds: []*Round{},
			Total:  0,
		}

		// Start the game
		game.Start()

		// Check if the user wants to play again
		fmt.Println("\nPlay again?")
		if Stop() {
			break
		} else {
			fmt.Print("New Game!\n\n")
		}
	}
}

// The game loop for each game
func (game *Game) Start() {
	// Play 5 games total
	for i := 1; i <= 5; i++ {

		// Round setup
		round := Round{
			Round:   i,
			Score:   0,
			Numbers: []int{},
		}

		// Start the round
		round.Start()

		// Append the round to the game history
		game.Rounds = append(game.Rounds, &round)
		// Update the game score
		game.Total += round.Score

		fmt.Printf("Game total after round %d: %d\n\n", len(game.Rounds), game.Total)
	}

	fmt.Printf("After %d rounds, your score is: %d\n", len(game.Rounds), game.Total)
}

// The game loop for each round
func (round *Round) Start() {
	// Loop until we break
	for {
		// Get the new number
		number := round.NewNumber()

		// Show the user the number
		fmt.Println("Chosen number: ", number)
		fmt.Println("Current score: ", round.Score)

		// Add it to the score
		round.Score += number
		// Check the new score to see if it is over 25
		if round.Score > 25 {
			fmt.Println("\nUh oh! Your score exceeded 25!")
			// Throw some statistics in the user's face. They might find it meaningful
			fmt.Println("Current score: ", round.Score)
			fmt.Printf("Score before %d: %d\n\n", number, round.Score-7)
			fmt.Print("Round over. Lets see how you did...\n\n")
			// Since they lost, the round's score is 0
			round.Score = 0
			break
		}

		// Check if the user wants to quit or continue and break the round loop if so
		if Stop() {
			fmt.Println("You have chosen to end the round. Hopefully it was worth it...")
			break
		}
	}

	// Present the round's score
	fmt.Println("Round score: ", round.Score)
	// Present the chosen numbers that round in a pretty way
	fmt.Print("Round numbers: ")
	for i, n := range round.Numbers {
		fmt.Print(n)
		// Separate each number with a comma or go to a new line at the end of the list
		if i != len(round.Numbers)-1 {
			fmt.Print(" + ")
		} else {
			fmt.Println()
		}
	}
}

// Generates a new number
func (round *Round) NewNumber() int {
	// Get a random number between 2 and 10
	number := rand.IntN(10-2) + 2

	// Store the number for historical reasons
	round.Numbers = append(round.Numbers, number)

	return number
}

/*
Find out if the user wants to stop.

Returns true if the user wants to stop, false if they wish to continue.
Handles invalid input.
*/
func Stop() bool {
	// Just loop until we get an answer and return
	for {
		// Prompt the user
		fmt.Print("Continue or stop?: ")

		var input string
		// Read and store the line of input
		_, err := fmt.Scanln(&input)
		if err != nil {
			// Exit if there is an error
			fmt.Printf("Something went wrong: %v\n\nClosing the application gracefully. Sorry for the inconvenience :(", err)
			os.Exit(1)
		}

		// No matter where the loop goes after here, to look good, we need a line break
		fmt.Println()

		// Compare the lowercase input, returning the valid choice
		switch strings.ToLower(input) {
		case "continue":
			return false
		case "stop":
			return true
		// Invalid choice
		default:
			fmt.Printf("Invalid input. You can either 'continue' or 'stop'. Try again\n\n")
		}

	}
}
