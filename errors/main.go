package main

import (
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	if err := raise(); err != nil {
		log.Fatal(err)
	}
}

func raise() error {
	_, err := os.Open("a.txt")
	if err != nil {
		return errors.Wrap(err, "open failed")
	}

	return nil
}
