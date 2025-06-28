package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	port := flag.String("port", "8080", "Port to run the web server on")

	// Customize usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of triple-s\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintln(os.Stderr, "  go run main.go -port=3000")
	}

	flag.Parse()

	addr := fmt.Sprintf(":%s", *port)
	http.HandleFunc("/", handler)
	fmt.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}