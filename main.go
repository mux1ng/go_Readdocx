package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

type ZipFile struct {
	data *zip.ReadCloser
}

func main() {
	file := flag.String("f", "", "the file to read")
	flag.Parse()
	readDocxFile(*file)
}

func readDocxFile(file string) {
	reader, _ := zip.OpenReader(file)
	zipData := ZipFile{data: reader}
	text := readText(zipData.data.File)

	re := regexp.MustCompile("<.*?>")
	result := re.ReplaceAllString(text, "")
	fmt.Println(result)
}

func readText(files []*zip.File) string {
	var documentFile *zip.File
	for _, f := range files {
		if f.Name == "word/document.xml" {
			documentFile = f
		}
	}

	documentReader, _ := documentFile.Open()
	content, _ := ioutil.ReadAll(documentReader)
	return string(content)
}
