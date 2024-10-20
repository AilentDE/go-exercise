package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/custom-types/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving wrong.")
		return
	} else {
		fmt.Println("Saved success.")
	}
}


func getNoteData() (string, string) {
	title := getInput("Note title")
	content := getInput("Note content")

	return title, content
}

func getInput(prompt string) string {
	fmt.Printf("%v: ", prompt)
	// var value string
	// fmt.Scanln(&value)

	// fmt.Scan is not useful when reading long content.
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')// must use ''
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}