package helpers

import (
	"log"
	"os"
	"strings"
	"time"
)

var Directory string

func CreateFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("Error creating file: ", err)
	}
	defer file.Close()
	log.Println("File created or truncated: ", filename)
}

func CreateBucketsCSV(filename string) {
	CreateFile(filename + "/buckets.csv")
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err2 := file.WriteString("Name,CreationTime,LastModifiedTime,Status\n")
	if err2 != nil {
		log.Fatalln(err)
	}

}

func CreateObjectsCSV(filename string) {
	CreateFile(filename + "/objects.csv")
	file, err := os.OpenFile(filename+"/objects.csv", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	_, err2 := file.WriteString("ObjectKey,Size,ContentType,LastModified\n")
	if err2 != nil {
		log.Fatalln(err)
	}
}

func CreateDir(filepath string) {
	err := os.MkdirAll(filepath, 0755)
	if err != nil {
		if !os.IsExist(err) {
			log.Println("Directory already exists, using it:", filepath)
		} else {
			log.Fatalln("Failed to create directory:", err)
		}
	}
}

func AppendBuckets(filename string) {
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	temp := []string{filename, time.Now().Format(time.UnixDate), time.Now().Format(time.UnixDate), "Active"}

	_, err2 := file.WriteString(strings.Join(temp, ","))

	if err2 != nil {
		log.Fatalln("Error writing to the file: ", filename)
	}
}
