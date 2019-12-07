// TODO describe program
package main

import (
	"fmt"
	"os"

	"github.com/artyom/autoflags"
)

func main() {
	args := struct {
		Name string `flag:"name"`
	}{}
	autoflags.Parse(&args)
	if err := run(args.Name); err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(1)
	}
}

func run(name string) error {
	if name == "" {
		name = "world"
	}
	Hello(name)
	return nil
}

func Hello(name string) { fmt.Println("Hello,", name) }
