package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

/* Based on what you’ve learned, write a command-line application that
counts the total number of lines of Go in a repo and prints the result to
the console. No 3rd-party libraries are allowed, only the Golang core
library. The application will be run against a Targeting repo. Any
file that does not end in .go should be excluded, as should empty
lines and comment lines in .go files. The main function should not
include all logic and functionality of the application, i.e., you
should break logical portions of the application into their functions.
Your functions should have accompanying tests, all passing. Time to
complete is 20 hours. After 20 hours, we’ll walk through the solutions and
compare the outputs. As you’re developing, you’ll run your code against
a single Targeting repo, but after completion we’ll run against multiple
repos and compare outputs with others. ADDITIONAL CHALLENGE FOR THE OVER
ACHIEVER: From the script, create a CSV file in your home directory
(~/ on Linux and Mac) with two columns, filename and numLines. For
each .go file, write a row in this document with the file name and the
number of code lines in that file. If the application isn’t complete
in time, that’s ok. We’ll work through the sticking points.*/

func go_file_check(fp string) bool {
	return filepath.Ext(fp) == ".go"
}

func recurse_through_repo(dir string) []string {

	fileList := make([]string, 0)
	e := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if go_file_check(path) {
			fileList = append(fileList, path)
		}
		return err
	})

	if e != nil {
		panic(e)
	}

	for _, file := range fileList {
		fmt.Println(file)
	}

	return fileList
}

func count_lines_in_files(fileList []string) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for _, file_path := range fileList {
		file, _ := os.Open(file_path)
		c, err := file.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			continue
		}
	}

	return count

}

func make_csv(files []string) {
	csv_file, err := os.Create(os.Getenv("HOME") + "/internship.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csv_file.Close()

	writer := csv.NewWriter(csv_file)
	defer writer.Flush()

	writer.Write([]string{"filename", "numLines"})

	for _, file := range files {
		count := count_lines_in_files([]string{file})
		record := []string{filepath.Base(file), strconv.Itoa(count)}
		fmt.Println(record)
		writer.Write(record)
	}
}

func main() {
	files := recurse_through_repo("/home/alexdripchak/go/src/")
	make_csv(files)
}
