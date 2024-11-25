package main

import (
	"bufio"
	"example.com/note/note"
	todo "example.com/note/todo"
	"fmt"
	"os"
	"strings"
)

type saver interface {
	Save() error
}

// this is an example of embedded interface. one interface having another interface
type outputable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todoItem, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(todoItem)

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputData(userNote)
}

func outputData(data outputable) {
	data.Display()
	err := data.Save()
	if err != nil {
		fmt.Println("error saving the data %v", data)
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Printf(err.Error())
		return err
	}

	fmt.Println("\nsaving data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
