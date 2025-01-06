package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Text    string
	Options []string
	Answer  int
}

func main() {
	questions := []Question{
		{"What is the location of NexTurn 2025 interns?", []string{"1. ICRISAT", "2. BALVIKAS", "3. AMMERPET", "4. KOTI"}, 2},
		{"What is the name of BoardInfinity trainer?", []string{"1. Vaibhav", "2. Raghu", "3. Bhawana", "4. Lavanya"}, 3},
		{"What is the value of 2+3?", []string{"1. 5", "2. 8", "3. 7", "4. 3"}, 1},
	}

	reader := bufio.NewReader(os.Stdin)
	score := 0
	var input string

	fmt.Println("Welcome to the Online Examination System!")
	fmt.Println("Enter 'exit' at any time to quit the quiz.")

	for i, q := range questions {
		fmt.Printf("\nQuestion %d: %s\n", i+1, q.Text)
		for _, option := range q.Options {
			fmt.Println(option)
		}

		for {
			fmt.Print("Your answer: ")
			input, _ = reader.ReadString('\n')
			input = strings.TrimSpace(input)

			if strings.ToLower(input) == "exit" {
				fmt.Println("\nExiting the quiz. Thank you!")
				break
			}

			answer, err := strconv.Atoi(input)
			if err != nil || answer < 1 || answer > len(q.Options) {
				fmt.Println("Invalid input. Please enter a valid option number.")
				continue
			}

			if answer == q.Answer {
				fmt.Println("Correct!")
				score++
			} else {
				fmt.Println("Incorrect!")
			}
			break
		}

		if strings.ToLower(input) == "exit" {
			break
		}
	}

	fmt.Printf("\nQuiz Completed! Your Score: %d/%d\n", score, len(questions))
	switch {
	case score == len(questions):
		fmt.Println("Performance: Excellent")
	case score >= len(questions)/2:
		fmt.Println("Performance: Good")
	default:
		fmt.Println("Performance: Needs Improvement")
	}
}
