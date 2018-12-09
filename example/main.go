package main

import (
	"fmt"

	"github.com/jabernardo/aargh"
)

func main() {
	// Create a new application
	app := aargh.New()

	// A simple greeting!
	// go run main.go hello
	// go run main.go hello --name="Your Name"
	app.Command("hello", func(app *aargh.App) {
		name := "User"

		// Set visitor name
		if app.HasOption("name") {
			name = app.GetOption("name")
		}

		// Quiet!!!
		if app.HasFlag("q") {
			fmt.Println("Shhh!")
		}

		fmt.Printf("Hello, %s!\n", name)
	})

	// Default command
	app.Command("default", func(app *aargh.App) {
		fmt.Println("A simple greeting!")
		fmt.Println("\tgo run main.go hello --name=\"You Name\"")
	})

	// Help! Help! Help!
	// Calling another commands!
	app.Command("help", func(app *aargh.App) {
		app.Call("default")
	})

	// Run application
	app.Run()
}
