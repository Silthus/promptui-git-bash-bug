package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Println("Please run with one of the types: select, prompt, confirm")
		return
	}

	switch args[1] {
	case "select":
		runSelect()
	case "prompt":
		runPrompt()
	case "confirm":
		runConfirm()
	}
}

func runSelect() {
	s := &promptui.Select{
		Items: []string{"One", "Two", "Three"},
	}
	i, result, err := s.Run()
	fmt.Println("Error:", err)
	fmt.Println("Result:", result)
	fmt.Println("Index:", i)
}

func runPrompt() {
	p := &promptui.Prompt{
		Label: "Prompt",
	}
	result, err := p.Run()
	fmt.Println("Error:", err)
	fmt.Println("Result:", result)
}

func runConfirm() {
	p := &promptui.Prompt{
		Label:     "Confirm",
		IsConfirm: true,
	}
	result, err := p.Run()
	fmt.Println("Error:", err)
	fmt.Println("Result:", result)
}
