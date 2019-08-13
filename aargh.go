package aargh

import (
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

// aargh.App Structure
type App struct {
	// Application Name
	Name string

	// Application Version
	Version string

	// Application Author
	Author string

	// Active Command
	CommandActive string

	// Commands available
	Commands map[string]callback

	// Flags mapped throught command-line arguments (-f)
	Flags map[string]bool

	// key=value options mapped
	Options map[string]string

	// Extra arguments
	Arguments []string

	// Console Logging
	ConsoleLogging bool
}

// Create new instance of (*Aargh) App
//
// Returns:
//  - (*App) New *Aargh.App{}
func New() *App {
	log.SetFlags(0)

	return &App{}
}

// (*App).init() -  Initialize Application
// Reads os.Args and parse command, flags, options and arguments  for our
// application
//
// Returns:
//  - (error) error
func (app *App) init() error {
	// Application Defaults
	app.CommandActive = "default" // Command
	var index_start int = 1

	if len(os.Args) > 1 &&
		!strings.HasPrefix(os.Args[1], "-") &&
		!strings.HasPrefix(os.Args[1], "--") {
		app.CommandActive = os.Args[1]
		index_start = 2
	}

	if _, ok := app.Commands[app.CommandActive]; !ok {
		return app.handleError(100)
	}

	// Parse Arguments
	flags := make(map[string]bool)
	options := make(map[string]string)
	var arguments []string

	if len(os.Args) > index_start {
		for _, arg := range os.Args[index_start:] {
			// Options
			if strings.Index(arg, "--") == 0 {
				if len(arg) == 2 {
					return app.handleError(101)
				}

				option := strings.TrimLeft(arg, "--")

				if strings.Index(arg, "=") == -1 {
					return app.handleError(102, option)
				}

				equals_index := strings.Index(option, "=")
				key := option[:equals_index]
				value := option[equals_index+1:]

				if len(value) == 0 {
					return app.handleError(102, key)
				}

				options[key] = value

				// Flags
			} else if strings.Index(arg, "-") == 0 {
				if len(arg) == 1 {
					return app.handleError(103)
				}

				key := strings.TrimLeft(arg, "-")
				flags[key] = true

				// Arguments
			} else {
				arguments = append(arguments, arg)
			}
		}
	}

	app.Flags = flags
	app.Options = options
	app.Arguments = arguments

	return nil
}

// (*App).Run - Run Application
// Executes command from command map
//
// Arguments:
//  - logging (bool) Console logging?
//
// Returns:
//  - (error) error
func (app *App) Run(logging bool) error {
	// Set logging option
	app.ConsoleLogging = logging

	if err := app.init(); err != nil {
		return err
	}

	if err := app.Call(app.CommandActive); err != nil {
		if errd := app.Call("default"); errd != nil {
			return err
		}
	}

	return nil
}

// (*App).Call - Call command function
//
// Arguments:
//  - name (string) Command name
//
// Returns:
//  -  (error) If the callback was not executed
func (app *App) Call(name string) error {
	err_msg := app.getError(104, name)

	if command, ok := app.Commands[name]; ok {
		if reflect.TypeOf(command).Kind() != reflect.Func {
			return errors.New(err_msg)
		}

		command(app)

		return nil
	}

	return errors.New(err_msg)
}

// (*App).Register - Add/register new command to application
//
// Arguments:
//  - name (string) Command name
//  - fn (callback) Function
func (app *App) Command(name string, fn callback) {
	command_existing := make(map[string]callback)

	for k, v := range app.Commands {
		command_existing[k] = v
	}

	command_existing[name] = fn

	app.Commands = command_existing
}

// (*App).HasFlag - Check if name is flagged upon running the application
//
// Arguments:
//  - name (string) Flag name
//
// Returns:
//  - (bool)
func (app *App) HasFlag(name string) bool {
	if app.Flags[name] {
		return true
	}

	return false
}

// (*App).HasOption - Check if option is passed upon running the application
//
// Arguments:
//  - name (string) Option name
//
// Returns:
//  - (bool)
func (app *App) HasOption(name string) bool {
	if _, ok := app.Options[name]; ok {
		return true
	}

	return false
}

// (*App).GetOption - Get Option value
//
// Arguments:
//  - name (string) Option
//  - default_value (string) Default value
//
// Returns:
//  - (string) Option value
func (app *App) GetOption(name string, default_value ...string) string {
	if option, ok := app.Options[name]; ok {
		return option
	}

	return strings.Join(default_value, "")
}

// (*App).handleError - Handle error messages
//
// Arguments:
//  - code (int) Error code
//  - args (interface) Interface
//
// Returns:
//  - (error) Error Message
func (app *App) handleError(code int, args ...interface{}) error {
	err_msg := app.getError(code, args...)

	if app.ConsoleLogging {
		log.Fatalln(err_msg)
	}

	return errors.New(err_msg)
}

// (*App).getError - Get error message from `Error`
//
// Arguments:
//  - code (int) Error code
//  - args (interface) Interface
//
// Returns:
//  - (string) Error Message
func (app *App) getError(code int, args ...interface{}) string {
	err_msg := ""

	if _, ok := ERROR[code]; !ok {
		return err_msg
	}

	if len(args) > 0 {
		err_msg = fmt.Sprintf(ERROR[code], args...)
	} else {
		err_msg = ERROR[code]
	}

	return err_msg
}
