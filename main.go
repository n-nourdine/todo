package main

import (
	"os"

	h "github.com/n-nourdine/todo/handlers"
)

func main() {
	h.Start(os.Getenv("PORT"))
}
