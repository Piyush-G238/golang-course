package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/models"
)

func main() {
	// fmt.Print("Hello, world");

	note, err := GetNoteData()

	if err != nil {
		fmt.Print(err.Error())
		return
	}

	note.SaveNote()
}

func GetNoteData() (models.Note, error) {
	title := GetUserInput("Note title: ")
	content := GetUserInput("Note content: ")

	return models.NewNotes(title, content)
}

func GetUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, error := reader.ReadString('\n')
	if error != nil {
		return ""
	}
	value = strings.TrimSuffix(value, "\n")
	value = strings.TrimSuffix(value, "\r")
	return value
}
