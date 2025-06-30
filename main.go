package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LatievA/triple-s/handlers"
)

var Directory string

func main() {
	port := flag.String("port", "8080", "Port to run the web server on")
	dir := flag.String("dir", "./data", "Directory to serve files from")

	// Customize usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of triple-s\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintln(os.Stderr, "  go run main.go -port=3000")
	}

	flag.Parse()

	// Set the directory to serve files from
	if *dir != "./data" {
		Directory = "./data" + *dir
	} else {
		Directory = *dir
	}

	err := os.MkdirAll(Directory, 0755)
	if err != nil {
		if !os.IsExist(err) {
			log.Println("Directory already exists, using it:", Directory)
		} else {
			log.Fatalln("Failed to create directory:", err)
		}
	}

	addr := fmt.Sprintf(":%s", *port)
	log.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, handlers.RooterWays()); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}
