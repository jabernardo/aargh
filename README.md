# Aargh! 

A simple skeleton for Command-line Applications in Go

## Usage

Get a copy first!
```sh
go get github.com/jabernardo/aargh
```


A simple hello to everyone.
```go
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
  
  // Run application
  app.Run()
}

```
[![Run on Repl.it](https://repl.it/badge/github/jabernardo/aargh)](https://repl.it/github/jabernardo/aargh)

## Test it out!
```sh
# No parameters
go run main.go hello

# With sample parameter
go run main.go hello --name="Your Name"
```

## License

The `aargh` is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT).
