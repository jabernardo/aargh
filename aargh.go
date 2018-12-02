package aargh

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// aargh.App Structure
type App struct {
	Name    string // Application Name
	Version string // Application Version
	Author  string // Application Authow

	CommandActive string              // Active Command
	Commands      map[string]callback // Commands available

	Flags     map[string]bool   // Flags mapped throught command-line arguments (-f)
	Options   map[string]string // key=value options mapped
	Arguments []string          // Extra arguments
}

// Create new instance
func New() *App {
	return &App{}
}

// Initialize Application
// Reads os.Args and parse command, flags, options and arguments  for our
// application
func (app *App) Init() {
	if len(os.Args) == 1 {
		fmt.Println("No command found.")
		os.Exit(1) // No command
	}

	// Application Defaults
	app.CommandActive = os.Args[1] // Command

	// Parse Arguments
	flags := make(map[string]bool)
	options := make(map[string]string)
	var arguments []string

	for _, arg := range os.Args[2:] {
		// Options
		if strings.Index(arg, "--") == 0 {
			if strings.Index(arg, "=") == -1 {
				fmt.Printf("No value given for option: %s\n", arg)
				os.Exit(3)
			}

			option := strings.TrimLeft(arg, "--")
			equals_index := strings.Index(option, "=")
			key := option[:equals_index]
			value := option[equals_index+1:]

			options[key] = value

			// Flags
		} else if strings.Index(arg, "-") == 0 {
			key := strings.TrimLeft(arg, "-")
			flags[key] = true

			// Arguments
		} else {
			arguments = append(arguments, arg)
		}
	}

	app.Flags = flags
	app.Options = options
	app.Arguments = arguments
}

// Run Application
// Executes command from command map
func (app *App) Run() {
	app.Init()

	if !app.Call(app.CommandActive) {
		fmt.Printf("%s. Command not found\n", app.CommandActive)
		os.Exit(2) // Command not found
	}
}

// Call command
// Returns `true` if the callback was executed else `false`
func (app *App) Call(name string) bool {
	if command, ok := app.Commands[app.CommandActive]; ok {
		if reflect.TypeOf(command).Kind() != reflect.Func {
			os.Exit(3) // Invalid callback
		}

		command(app)

		return true
	}

	return false
}

// Register Command
// First parameter is the command name
// Second parameter is callback function
func (app *App) Command(name string, fn callback) {
	command_existing := make(map[string]callback)

	for k, v := range app.Commands {
		command_existing[k] = v
	}

	command_existing[name] = fn

	app.Commands = command_existing
}
