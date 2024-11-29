package util

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const choiceCharacters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// QuestionC will ask a question to the user, and present them with a list of options
func QuestionC(question string, choices ...string) (int, error) {
	result := -1

	fmt.Printf("%s:\n", question)

	for i, choice := range choices {
		fmt.Printf("%*s: %s\n", 1, string(choiceCharacters[i]), choice)
	}

	reader := bufio.NewReader(os.Stdin)

	for result < 0 {
		fmt.Print("Choose an option: ")
		input, _, err := reader.ReadRune()
		if err != nil {
			return -1, err
		}

		c := byte(input)
		result = strings.IndexByte(choiceCharacters, c) % 26
		if result >= len(choices) || result < 0 {
			fmt.Println("Invalid choice, please choose again.")
			result = -1 // Reset result to keep the loop going for a valid choice
		}
	}

	return result, nil
}
