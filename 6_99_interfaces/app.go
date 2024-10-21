package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}
type outputtable interface {
	Display()
	saver
}

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func main() {
	printAnything(1024)
	printAnythingAnother("For interface test.")


	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userNote)
	if err != nil {
		return
	}
	fmt.Println("Saving the note succeeded!")

	text := getUserInput("Todo content: ")
	userTodo, err := todo.New(text)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(userTodo)
	if err != nil {
		return
	}
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func printAnything(data interface{}) {
// 	switch data.(type) {
// 	case int:
// 		fmt.Println("Prepare int Note.")
// 	case float64:
// 		fmt.Println("Prepare float Note.")
// 	case string:
// 		fmt.Println("Prepare string Note.")
// 	default:
// 		fmt.Println(data)
// 	}

// or
	intData, ok := data.(int)
	if ok {
		fmt.Println("Prepare int Note.", intData)
		return
	}
	float64Data, ok := data.(float64)
	if ok {
		fmt.Println("Prepare float Note.", float64Data)
		return
	}
	stringData, ok := data.(string)
	if ok {
		fmt.Println("Prepare string Note.", stringData)
		return
	}
	fmt.Println(data)
}
func printAnythingAnother(data any) {
	fmt.Println(data)
}