package handlers

import (
	"fmt"
	"net/http"
	"path"

	"github.com/LatievA/triple-s/helpers"
)

func RooterWays() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Handler)
	mux.HandleFunc("PUT /{BucketName}", CreateBucket)
	mux.HandleFunc("GET /", ListBuckets)
	mux.HandleFunc("DELETE /{BucketName}", DeleteBucket)
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", PutObject)
	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", GetObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", DeleteObject)

	return mux
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(r.URL.Path)
	helpers.CreateDir(helpers.Directory + "/" + bucketName)
	helpers.AppendBuckets(bucketName)
	helpers.CreateObjectsCSV(helpers.Directory + "/" + bucketName)
}

func ListBuckets(w http.ResponseWriter, r *http.Request) {}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {}

func PutObject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PutObject called with method: %s", r.Method)
}

func GetObject(w http.ResponseWriter, r *http.Request) {}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "DeleteObject called with method: %s", r.Method)
}
