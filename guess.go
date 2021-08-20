package main // the first statement in a go source file must start with a package "name". All the functions declared in this program becomes part of the declared package. Go programs start running in the main package. This tells the Go compiler to compile the package as an executable program.
//An Executable simply means a file which contains a program which is capable of being executed or run as a program in your computer.

import (
	"fmt"       //  import the fmt package which allows you to use text formatting,reading input & printing output fucntions
	"math/rand" //import the rand package which allows you to generate random numbers
	"time"      // import the package which will provide time functionality to measure time
)

func main() {
	fmt.Println("Game: Guess a number between 0 and 10") // This informs the player about how to play the game.
	fmt.Println("You have three(3) tries ")

	// generate a random number
	source := rand.NewSource(time.Now().UnixNano()) //The default number generator is predictable, so it will produce the same sequence of numbers each time. To produce varying range of numbers, give it a seed that changes (in this case: time would ensure it changes ). Note that this is not safe to use for random numbers you want to be secret; use crpyto/rand for those.
	randomizer := rand.New(source)
	secretNumber := randomizer.Intn(10) // generate numbers between 0 and 10 only. If you want to change the range change the value 10 to a higher or lower value

	var guess int // this is one form of declaration in go; you have to add the type of the variable been declared. "var guess" wont work

	for try := 1; try <= 3; try++ { // declaring the conditions for the for loop ; the shorthand form of declaring a variable was used here. Declare and Initialize ' := ' you declare and assign a value upon declaration. Go will automatically infer the type of the variable since you already assigned a value to it.
		fmt.Printf("TRIAL %d\n", try)           // print out the number of times the player has made a guess
		fmt.Println("Please enter your number") // the program will prompt the player to make a guess and enter a number
		fmt.Scan(&guess)                        // this function makes it possible for the program to receive the input

		if guess < secretNumber { // if the guessed number is less than or greater than the correct number; give the player a hint
			fmt.Printf("Sorry, wrong guess ; number is too small\n ")
		} else if guess > secretNumber {
			fmt.Printf("Sorry, wrong guess ; number is too large\n ")
		} else {
			fmt.Printf("You win!\n") // Notify the player that he has won when he guesses the correct number
			break
		}

		if try == 3 { // if the number of tries is equal to 3, print game over and also the correct number
			fmt.Printf("Game over!!\n ")
			fmt.Printf("The correct number is %d\n", secretNumber)
			break
		}
	}
}
