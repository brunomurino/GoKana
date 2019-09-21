package main

import (
	"hiragana"
	"katakana"
	"fmt"
	"os"
	"os/exec"

	"github.com/common-nighthawk/go-figure"
)

// clearScreen clears the terminal in which this program is running.
func clearScreen() {
	clear := exec.Command("clear")
	clear.Stdout = os.Stdout
	clear.Run()
}

// printMenu takes the available options and prints the program header.
func printMenu(options []string) {

	myFigure := figure.NewFigure("GoKana", "", true)
	myFigure.Print()

	fmt.Println("")
	for key, value := range options {
		fmt.Printf("\t%d. %s\n", key+1, value)
	}
	fmt.Println("")
}

// getChoice read user inputs for option and difficulty.
func getChoice(options []string) (choice int, difficulty int) {

	difficulty = 3
	validChoice := false
	for !validChoice {
		fmt.Print("\tChoose option:\t")
		fmt.Scanf("%d", &choice)
		if choice > 0 && choice <= len(options) {
			validChoice = true
			if choice == 3 {
				fmt.Print("\n\tSetting difficulty:\t")
				fmt.Scanf("%d", &difficulty)
				validChoice = false
			} else {
				fmt.Println("\t\tThe difficulty is", difficulty)
				break
			}
		} else {
			fmt.Println("\t\tChoose a valid option.")
		}
	}
	return
}

// preprareRound randomly creates the questions/truths pairs.
func preprareRound(m map[string]string, difficulty int) ([]string, []string) {

	var questions [10]string
	var truths [10]string

	for i := 0; i < difficulty; i++ {
		for key, value := range m {
			questions[i] = value
			truths[i] = key
			break
		}
	}

	return questions[:difficulty], truths[:difficulty]
}

// play formats the screen and runs rounds of practice.
func play(kana map[string]string, difficulty int, kanaName string) {

	fmt.Println("\n\tPracticing", kanaName, "!")
	fmt.Println("\t    Press 'q' to go back to the menu.")

	var roundQuestions, roundTruths []string

	roundQuestions, roundTruths = preprareRound(kana, difficulty)

	fmt.Print("\n\t")
	for i := range roundQuestions {
		fmt.Printf("%s ", roundQuestions[i])
	}
	fmt.Println("")

	answers := make([]string, difficulty)

	fmt.Print("\n\t")
	for i := range roundQuestions {
		fmt.Scanf("%s", &answers[i])
		if answers[i] == "q" {
			main()
		}
	}

	fmt.Print("\n\t")
	for i := range roundQuestions {
		fmt.Printf("%s ", roundTruths[i])
	}

	score := 0
	for i := 0; i < difficulty; i++ {
		if answers[i] == roundTruths[i] {
			score++
		}
	}

	fmt.Print("\n\t\tTotal score:\t", score, "\t")
	fmt.Scanf("%s")

}

// resetScreen prints the program header without taking inputs.
func resetScreen(options []string, choice int, difficulty int) {
	fmt.Println("\n\tKana Practice in Go")
	clearScreen()
	printMenu(options)
	fmt.Print("\tChoose option:\t", choice)
	fmt.Println("\n\t\tThe difficulty is", difficulty)
}

func main() {

	clearScreen()

	options := make([]string, 4)

	options[0] = ("Play Hiragana")
	options[1] = ("Play Katakana")
	options[2] = ("Set difficulty")
	options[3] = ("Quit")

	printMenu(options)

	switch choice, difficulty := getChoice(options); choice {
	case 1:
		resetScreen(options, choice, difficulty)
		for {
			play(hiragana.Hiragana(), difficulty, "Hiragana")
			resetScreen(options, choice, difficulty)
		}
	case 2:
		resetScreen(options, choice, difficulty)
		for {
			play(katakana.Katakana(), difficulty, "Katakana")
			resetScreen(options, choice, difficulty)
		}
	case 4:
		os.Exit(0)
	}

}
