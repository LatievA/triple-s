package helpers

import (
	"log"
	"os"
	"strings"
	"time"
	"encoding/csv"
	"net"
)

var Directory string

// Don't forget to close files after creating or reading them

func CreateBucketsCSV(filename string) {
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString("Name,CreationTime,LastModifiedTime,Status\n")
	if err2 != nil {
		log.Fatal(err)
	}

}

func CreateObjectsCSV(filename string) {
	file, err := os.OpenFile(filename+"/objects.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString("ObjectKey,Size,ContentType,LastModified\n")
	if err2 != nil {
		log.Fatal(err)
	}
}

func CreateDir(filepath string) {
	err := os.MkdirAll(filepath, 0755)
	if err != nil {
		log.Fatal("Failed to create directory:", err)
	}
}

func AppendBuckets(filename string) {
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	temp := []string{filename, time.Now().Format(time.UnixDate), time.Now().Format(time.UnixDate), "Active"}

	_, err2 := file.WriteString(strings.Join(temp, ","))

	if err2 != nil {
		log.Fatal("Error writing to the file: ", filename)
	}
}

func ReadCSV(filename string) [][]string{
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return records[1:]
}

// Validation

func IsValidName(name string) bool {
	if len(name) < 3 || len(name) > 63 {
		return false
	}

	if name[0] == '-' || name[len(name)-1] == '-' {
		return false
	}

	for i, v := range name {
		if (v < 'a' || v > 'z') && v != '-' && v != '.' {
			return false
		} else if v == '.' && ((i != len(name)-1 && rune(name[i+1]) == '.') || (i != 0 && rune(name[i-1]) == '.')) {
			return false
		}
	}

	if net.ParseIP(name) != nil {
		return false
	}

	return true
}

func IsUniqueName(name, filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return !strings.Contains(string(data), name)
}
