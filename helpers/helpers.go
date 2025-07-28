package helpers

import (
	"encoding/csv"
	"log"
	"net"
	"os"
	flp "path/filepath"
	"strings"
	"time"
)

var Directory string

// Don't forget to close files after creating or reading them

func CreateBucketsCSV() {
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err2 := file.WriteString("Name,CreationTime,LastModifiedTime,Status\n")
	if err2 != nil {
		log.Fatal(err)
	}

}

func CreateObjectsCSV(filepath string) {
	file, err := os.OpenFile(filepath+"/objects.csv", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
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
	file, err := os.OpenFile(Directory+"/buckets.csv", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	temp := []string{filename, time.Now().Format(time.UnixDate), time.Now().Format(time.UnixDate), "Active\n"}

	_, err2 := file.WriteString(strings.Join(temp, ","))

	if err2 != nil {
		log.Fatal("Error writing to the file: ", filename)
	}
}

func AppendObjects(objectKey, size, contentType, filepath string) {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	temp := []string{objectKey, size, contentType, time.Now().Format(time.UnixDate) + "\n"}

	_, err2 := file.WriteString(strings.Join(temp, ","))

	if err2 != nil {
		log.Fatal("Error writing to the file: ", objectKey)
	}
}

func ReadCSV(filepath string) *[][]string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	records = records[1:]
	return &records
}

func WriteCSV(filepath string, records *[][]string) error {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	if flp.Base(filepath) == "buckets.csv" {
		_, err2 := file.WriteString("Name,CreationTime,LastModifiedTime,Status\n")
		if err2 != nil {
			log.Fatal(err)
		}
	} else {
		_, err2 := file.WriteString("ObjectKey,Size,ContentType,LastModified\n")
		if err2 != nil {
			log.Fatal(err)
		}
	}

	return writer.WriteAll(*records)
}

func DeleteRecord(name, filepath string) {
	records := ReadCSV(filepath)
	for i, v := range *records {
		if v[0] == name {
			*records = append((*records)[:i], (*records)[i+1:]...)
			break
		}
	}
	if err := WriteCSV(filepath, records); err != nil {
		log.Fatal(err)
	}
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
		if (v < 'a' || v > 'z') && !(v >= '0' && v <= '9') && v != '-' && v != '.'{
			return false
		} else if v == '.' && ((i != len(name)-1 && rune(name[i+1]) == '.') || (i != 0 && rune(name[i-1]) == '.')) {
			return false
		}
	}

	return net.ParseIP(name) == nil
}

func IsUniqueName(name, filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return !strings.Contains(string(data), name)
}

func IsEmptyCSV(filepath string) bool {
	records := ReadCSV(filepath)
	return len(*records) == 0
}
