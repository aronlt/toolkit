package demo

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func tempDir() {
	// Create our Temp File
	tmpFile, err := ioutil.TempFile(os.TempDir(), "prefix-")
	if err != nil {
		log.Fatal("Cannot create temporary file", err)
	}

	fmt.Println("Created File: " + tmpFile.Name())

	// Example writing to the file
	_, err = tmpFile.Write([]byte("This is a www.twle.cn example!"))
	if err != nil {
		log.Fatal("Failed to write to temporary file", err)
	}

	// Remember to clean up the file afterwards
	defer os.Remove(tmpFile.Name())
}
