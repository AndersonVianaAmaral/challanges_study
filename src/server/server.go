package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	Routes()
	MountServer()
}

func MountServer() {
	fmt.Printf("Starting server at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Code Education Rocks!")
}

func Routes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, Greeting("Code Education Rocks !"))
	})
}

func Greeting(text string) string {
	return "<b>" + text + "</b>"
}
