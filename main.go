package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	. "strings"
)

var errorsCounter int
var guesses string

func main() {
	startGame()
}

func startGame() {
	printStartMessage()
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	userAnswer := scanner.Text()

	if userAnswer != "yes" {
		endTheGame()
	}
	randomWord := getRandomWord(initWords())
	printWordState(randomWord)
	guess := bufio.NewScanner(os.Stdin)

	for !isGameEnded(randomWord) {
		fmt.Println("Try to guess a letter")

		guess.Scan()
		letter := guess.Text()

		if len(letter) < 1 {
			fmt.Println("Please enter a letter or word")
			continue
		}

		if Contains(randomWord, letter) {
			guesses += letter
			printWordState(randomWord)
			fmt.Println("You are right")
		} else {
			errorsCounter++
			printWordState(randomWord)
			fmt.Println("You missed, ur misses: ", errorsCounter)
		}
	}
	if errorsCounter >= 7 {
		fmt.Println("You lose")
	} else {
		fmt.Println("You won")
	}

}

func isWordGuessed(word string) bool {
	for _, letter := range word {
		if !ContainsRune(guesses, letter) {
			return false
		}
	}
	return true
}

func printWordState(secretWord string) {
	for _, letter := range secretWord {
		if ContainsRune(guesses, letter) {
			fmt.Print(string(letter))
		} else {
			fmt.Print("*")
		}
	}
	fmt.Println()
}

func endTheGame() {
	fmt.Println("You have ended.")
	os.Exit(0)
}

func getRandomWord(words []string) string {
	num := rand.Intn(len(words))
	return words[num]
}

func initWords() []string {
	data, err := os.ReadFile("words.txt")

	if err != nil {
		panic(err)
	}

	var arrWords []string
	words := Split(string(data), "\n")

	for _, word := range words {
		word = TrimSpace(word)

		if word != "" {
			arrWords = append(arrWords, word)
		}
	}

	return arrWords
}

func isGameEnded(randomWord string) bool {
	return errorsCounter >= 7 || isWordGuessed(randomWord)
}

func printStartMessage() {
	fmt.Println("HangMan Game" + "\n \n" +
		"Please enter yes to start a new game, or any to end the game. " + "\n")
}
