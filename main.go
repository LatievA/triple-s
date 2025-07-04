package main

// TO DO
// Create buckets.csv when starting server if it doesn't exist +
// Create a dir and and it to csv file, also create objects.csv inside
// List all buckets from csv file as xml response
// If given bucket exists in buckets csv and objects.csv of this bucket is empty delete dir and this bucket from buckets.csv

/* bucket naming
unique
3-63 char long
only lowercase letter, numbers, '-' and '.'
not ip address
not begin and end with '-'
not contain two consecutive dots
if not return 400 bad request
*/

/*
	Don't forget to close files after creating or reading them
*/

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LatievA/triple-s/handlers"
	"github.com/LatievA/triple-s/helpers"
)

func main() {
	port := flag.String("port", "8080", "Port to run the web server on")
	dir := flag.String("dir", "./data", "Directory to serve files from")

	// Customize usage message
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of triple-s\n")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nExample:")
		fmt.Fprintln(os.Stderr, "  go run main.go -port=3000")
		fmt.Fprintln(os.Stderr, "  go run main.go -dir=/path/to/data")
	}

	flag.Parse()

	// Set the directory to serve files from
	if *dir != "./data" {
		helpers.Directory = "./data" + *dir
	} else {
		helpers.Directory = *dir
	}

	helpers.CreateDir(helpers.Directory)
	helpers.CreateBucketsCSV(helpers.Directory)

	addr := fmt.Sprintf(":%s", *port)
	log.Printf("Server is running on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, handlers.RooterWays()); err != nil {
		log.Printf("Failed to start server: %v\n", err)
	}
}
