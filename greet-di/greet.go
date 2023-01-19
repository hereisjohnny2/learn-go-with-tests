package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func GreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Marcelo")
}

func main() {
	// Greet(os.Stdout, "Johnny")
	http.ListenAndServe(":8080", http.HandlerFunc(GreetHandler))
}
