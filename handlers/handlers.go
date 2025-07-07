package handlers

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/LatievA/triple-s/helpers"
)

type ListAllMyBucketsResult struct {
	Buckets []Bucket
	Owner   Owner
}

type Bucket struct {
	CreationDate string
	Name         string
}

type Owner struct {
	DisplayName string
	ID          string
}

var NewOwner *Owner = &Owner{DisplayName: "Abylay", ID: "1"}

func RooterWays() *http.ServeMux {
	mux := http.NewServeMux()

	// TO DO:
	mux.HandleFunc("PUT /{BucketName}", CreateBucket) // Done
	mux.HandleFunc("GET /", ListBuckets) // Done
	mux.HandleFunc("DELETE /{BucketName}", DeleteBucket) // Done
	mux.HandleFunc("PUT /{BucketName}/{ObjectKey}", PutObject)
	mux.HandleFunc("GET /{BucketName}/{ObjectKey}", GetObject)
	mux.HandleFunc("DELETE /{BucketName}/{ObjectKey}", DeleteObject)

	return mux
}

// change all filepath to r.URL.Path

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(r.URL.Path)
	if !helpers.IsValidName(bucketName) {
		http.Error(w, "Invalid bucket name", http.StatusBadRequest)
		return
	}

	if !helpers.IsUniqueName(bucketName, helpers.Directory+"/buckets.csv") {
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
	buckets := make([]Bucket, len(*records))
	for i, v := range *records {
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

	if helpers.IsUniqueName(bucketName, helpers.Directory+"/buckets.csv") {
		http.Error(w, "Bucket doesn't exist", http.StatusBadRequest)
		return
	}

	if !helpers.IsEmptyCSV(helpers.Directory + "/" + bucketName + "/objects.csv") {
		http.Error(w, "Bucket isn't empty", http.StatusConflict)
		return
	}

	helpers.DeleteRecord(bucketName, helpers.Directory+"/buckets.csv")
	os.RemoveAll(helpers.Directory + "/" + bucketName)
	w.WriteHeader(http.StatusNoContent)
}

func PutObject(w http.ResponseWriter, r *http.Request) {
	bucketName := path.Base(path.Dir(r.URL.Path))
	if !helpers.IsValidName(bucketName) {
		http.Error(w, "bucket name is unvalid", http.StatusBadRequest)
		return
	}
	if helpers.IsUniqueName(bucketName, helpers.Directory + "/buckets.csv") {
		http.Error(w, "bucket not exists", http.StatusBadRequest)
		return
	}

	objectKey := path.Base(r.URL.Path)
	if !helpers.IsValidName(objectKey) {
		http.Error(w, "object name is unvalid", http.StatusBadRequest)
		return
	}

	if !helpers.IsUniqueName(objectKey, helpers.Directory + path.Dir(r.URL.Path)+"/objects.csv") {
		http.Error(w, "objectkey already exists", http.StatusBadRequest)
		return
	}

	contentType := r.Header.Get("Content-Type")
	contentLength := r.Header.Get("Content-Length")
	
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	if err = os.WriteFile(path.Join(helpers.Directory, bucketName, objectKey), body, 0644); err != nil {
		http.Error(w, "error writing object to file", http.StatusInternalServerError)
		return
	}
	

	helpers.AppendObjects(objectKey, contentLength, contentType, helpers.Directory + path.Dir(r.URL.Path)+"/objects.csv")

}

func GetObject(w http.ResponseWriter, r *http.Request) {}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	// objectName := path.Base(r.URL.Path)
}
