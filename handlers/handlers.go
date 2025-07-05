package handlers

import (
	"encoding/xml"
	"net/http"
	"path"

	"github.com/LatievA/triple-s/helpers"
)

type ListAllMyBucketsResult struct {
	Buckets []Bucket
	Owner Owner
}

type Bucket struct {
	CreationDate string
	Name string
}

type Owner struct {
	DisplayName string
	ID string
}

var NewOwner *Owner = &Owner{DisplayName: "Abylay", ID: "1"}

func RooterWays() *http.ServeMux {
	mux := http.NewServeMux()

	// TO DO:
	mux.HandleFunc("PUT /{BucketName}", CreateBucket) // Done
	mux.HandleFunc("GET /", ListBuckets)
	mux.HandleFunc("DELETE /{BucketName}", DeleteBucket)
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", PutObject)
	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", GetObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", DeleteObject)

	return mux
}

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(r.URL.Path)
	if !helpers.IsValidName(bucketName) {
		http.Error(w, "Invalid bucket name", http.StatusBadRequest)
		return
	}

	if !helpers.IsUniqueName(bucketName, helpers.Directory + "/buckets.csv") {
		http.Error(w, "Bucket name already exists", http.StatusConflict)
		return
	}
	
	helpers.CreateDir(helpers.Directory + "/" + bucketName)
	helpers.AppendBuckets(bucketName)
	helpers.CreateObjectsCSV(helpers.Directory + "/" + bucketName)
}

func ListBuckets(w http.ResponseWriter, r *http.Request) {
	result := new(ListAllMyBucketsResult)
	result.Owner = *NewOwner
	records := helpers.ReadCSV(helpers.Directory + "/buckets.csv")
	buckets := make([]Bucket, len(records))
	for i, v := range records{
		buckets[i] = Bucket{
			Name:         v[0],
			CreationDate: v[1],
		}
	}
	result.Buckets = buckets

	w.Header().Set("Content-Type", "application/xml")
	data, err := xml.Marshal(result) 
	if err != nil {
		http.Error(w, "error marshaling xml", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(r.URL.Path)
	if !helpers.IsValidName(bucketName) {
		http.Error(w, "Invalid bucket name", http.StatusBadRequest)
		return
	}

	if helpers.IsUniqueName(bucketName, helpers.Directory + "/buckets.csv") {
		http.Error(w, "Bucket doesn't exist", http.StatusBadRequest)
		return
	}
}

func PutObject(w http.ResponseWriter, r *http.Request) {
	// objectName := path.Base(r.URL.Path)
}

func GetObject(w http.ResponseWriter, r *http.Request) {}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	// objectName := path.Base(r.URL.Path)
}
