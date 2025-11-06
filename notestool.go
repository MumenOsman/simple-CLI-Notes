package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 || (len(os.Args) == 2 && os.Args[1] == "help") {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Usage: ./notestool [COLLECTION_NAME]")
		return
	}

	collectionName := os.Args[1]
	notes := LoadNotes(collectionName)

	fmt.Println("\033[H\033[2J")
	fmt.Println("Welcome to the notes tool!")

	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choiceStr := strings.TrimSpace(input)

		fmt.Println("\033[H\033[2J")

		switch choiceStr {
		case "1":
			ShowNotes(notes)
		case "2":
			notes = AddNote(notes)
			SaveNotes(collectionName, notes)
		case "3":
			if len(notes) > 0 {
				ShowNotes(notes)
				notes = DeleteNote(notes)
				SaveNotes(collectionName, notes)
			} else {
				fmt.Println("No notes to delete.")
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number from 1 to 4.")
		}
	}
}

func LoadNotes(collectionName string) []string {
	data, err := os.ReadFile(collectionName)
	if err != nil {
		os.WriteFile(collectionName, []byte(""), 0644)
		return []string{}
	}
	content := string(data)
	notes := strings.Split(content, "\n")
	if len(notes) > 0 && notes[len(notes)-1] == "" {
		notes = notes[:len(notes)-1]
	}
	return notes
}

func SaveNotes(collectionName string, notes []string) {
	content := strings.Join(notes, "\n") + "\n"
	os.WriteFile(collectionName, []byte(content), 0644)
}

func ShowNotes(notes []string) {
	if len(notes) == 0 {
		fmt.Println("\nNo notes yet.")
		return
	}
	fmt.Println("\nNotes:")
	for i, note := range notes {
		fmt.Printf("%03d - %s\n", i+1, note)
	}
}

func AddNote(notes []string) []string {
	fmt.Println("\nEnter the note text (or type 'cancel' to return):")
	reader := bufio.NewReader(os.Stdin)
	newNote, _ := reader.ReadString('\n')
	newNote = strings.TrimSpace(newNote)

	if strings.ToLower(newNote) == "cancel" {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Add note canceled.")
		return notes
	}

	if newNote == "" {
		fmt.Println("Cannot add an empty note.")
		return notes
	}

	fmt.Println("Note added.")
	return append(notes, newNote)
}

// For deleting the note
func DeleteNote(notes []string) []string {
	if len(notes) == 0 {
		return notes
	}

	fmt.Println("\nEnter the number of the note to remove or 0 to cancel:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	num, err := strconv.Atoi(input)
	if err != nil || num < 0 || num > len(notes) {
		fmt.Println("Invalid number.")
		return notes
	}

	if num == 0 {
		fmt.Println("\033[H\033[2J")
		fmt.Println("Deletion canceled.")
		return notes
	}

	index := num - 1
	fmt.Println("Note removed.")
	return append(notes[:index], notes[index+1:]...)
}
