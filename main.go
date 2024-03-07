package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

/*
	how do we find the best move?

		first move, guess audio
		any move after, eliminate impossible words, and pick a random valid, choice (if I cared enough, I would choose
																				     the word with common characters)
*/

func getWords(dir string) []string {
	body, err := os.ReadFile(dir)
	if err != nil {
		fmt.Println("error with file opening")
		os.Exit(1)
	}

	words := strings.Split(string(body), " ")
	return words
}

func removeIncorrect(words []string, curWord string) []string {

	for i := 0; i < 5; i++ {
		for q := 0; q < len(words); q++ {
			if words[q][i] != curWord[i] && curWord[i] != '-' {
				words = append(words[:q], words[q+1:]...)
				q--
				continue
			}
		}
	}
	return words
}

func removeNonYellows(words []string, yellows []string, yellowIndexes []int) []string {
	for i := 0; i < len(words); i++ {
		for q := 0; q < len(yellows); q++ {
			if !strings.Contains(words[i], yellows[q]) {
				words = append(words[:i], words[i+1:]...)
				i--
				break
			}

			if words[i][yellowIndexes[q]] == yellows[q][0] {
				words = append(words[:i], words[i+1:]...)
				i--
				break
			}
		}
	}
	return words
}

func removeGreyLetters(words []string, greys []string) []string {
	for i := 0; i < len(words); i++ {
		for q := 0; q < len(greys); q++ {
			if strings.Contains(words[i], greys[q]) {
				words = append(words[:i], words[i+1:]...)
				i--
				break
			}
		}
	}
	return words
}

func getGuessedWord() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What was the guessed word?  ")
	text, _ := reader.ReadString('\n')
	fmt.Println()
	return text
}

func getGuessStatus() string {
	// a '.' is a grey letter, a '?' is a yellow letter and if the letter is green, then just put the letter
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is the status of the word? (or enter ! if the word has been guessed)  ")
	text, _ := reader.ReadString('\n')
	return text
}

func nonRepeatedLetters(words []string) string {
	text := words[0]
	var unique string

	for i := 0; i < len(words); i++ {
		unique = ""
		for q := 0; q < len(words[i]); q++ {
			if !strings.Contains(unique, string(words[i][q])) {
				unique += string(words[i][q])
			}
		}
		if len(unique) == len(words[i]) {
			return words[i]
		}
	}
	return text
}

func runWordle() {

	words := getWords("5_letter_words.txt")
	curWord := "-----"
	var yellows []string
	var yellowIndexes []int
	var greyLetters []string

	text := getGuessedWord()
	for len(words) > 0 {
		status := getGuessStatus()
		if len(strings.TrimSpace(status)) != 5 {
			break
		}

		for i := 0; i < 5; i++ {
			if status[i] != '?' && status[i] != '.' {
				curWord = curWord[:i] + string(status[i]) + curWord[i+1:]
			}

			if status[i] == '?' {
				yellows = append(yellows, string(text[i]))
				yellowIndexes = append(yellowIndexes, i)
			}

			if status[i] == '.' {

				// in the case where a letter occurs more than once
				// we need to ensure that we aren't accidentally setting
				// the letter as grey
				// to fix this, we just add the letter as yellow
				if strings.Count(curWord, string(text[i])) > 0 || slices.Contains(yellows, string(text[i])) {
					yellows = append(yellows, string(text[i]))
					yellowIndexes = append(yellowIndexes, i)
				} else {
					greyLetters = append(greyLetters, string(text[i]))
				}
			}
		}

		words = removeNonYellows(words, yellows, yellowIndexes)
		words = removeIncorrect(words, curWord)
		words = removeGreyLetters(words, greyLetters)

		if len(words) == 0 {
			fmt.Println("Can't find the matching word")
			break
		}

		text = nonRepeatedLetters(words)
		// fmt.Println(words)
		fmt.Println("The best word is : ", text)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	again := "!"
	for again == "!" {
		runWordle()
		fmt.Print("enter '!' if you would like to play again  ")
		again, _ = reader.ReadString('\n')
		again = strings.TrimSpace(again)
	}
}
