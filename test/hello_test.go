package tests

import (
  "testing"
  "reflect"
  ".."
)

func TestHello(t *testing.T) {
  app := aargh.New()

  if reflect.TypeOf(app) != reflect.TypeOf(&aargh.App{}) {
    t.Error("Invalid object")
  }
}
