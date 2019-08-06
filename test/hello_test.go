package tests

import (
	"fmt"
	"reflect"
	"testing"

	".."
)

func TestHello(t *testing.T) {
	app := aargh.New()

	if reflect.TypeOf(app) != reflect.TypeOf(&aargh.App{}) {
		t.Error("Invalid object")
	}

	app.Command("default", func(app *aargh.App) {
		fmt.Println("Hello World")
	})

	if err := app.Call("default"); err != nil {
		t.Error("No callback registered")
	}
}
