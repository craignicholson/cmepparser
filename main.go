package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/craignicholson/cmepparser/cmep"
	"log"
	"os"
)

// Main Function Entry Point
func main() {

	// Open the file
	file, err := os.Open("cmep.dat")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() //f.Close will run when we're finished.

	// Create a map to hold the counts of the record locators
	recordformats := make(map[string]int)

	//TODO we need to process by RecordType so we need a series of If statements

	//Create the slice of the struct
	mepmd01x := make([]cmep.MEPMD01x, 0)

	//Start processing the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Example Line by Line Processing
		cmep.ProcessLine(text, recordformats)

		// Batching the Results into on RecordType
		mepmd01x = cmep.ProcessBatchLine(text, mepmd01x)
	}

	// Quick output of found types
	fmt.Printf("RecordTypes : %v\n", recordformats)

	// Printing the results for these first commits and then
	// I will write some tests
	json, err := json.Marshal(mepmd01x)
	if err != nil {
		fmt.Println(err)
		return
	}
	//Quick and Dirty Print Here
	fmt.Println("\n")
	fmt.Println(string(json))
	fmt.Println("\n")

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
