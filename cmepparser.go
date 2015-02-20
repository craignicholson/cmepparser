package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// CMEP MEPMD01 extended record format
// Metering Data Type 1 – Interval Data, Pulse Data, Reference Register Reads
type MEPMD01x struct {
	RecordType          string
	RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
	SenderID            string // Fixed value
	SenderCustomerID    string // Sensus code - customer:Key for flexible fields
	ReceiverID          string // Flexible field – see Table 9 for options
	ReceiverCustomerID  string // Flexible field – see Table 9 for options
	TimeStamp           string // Date&time this record was created YYYYMMDDHHMM
	MeterID             string // Flexible field – see table 9 for options
	Purpose             string // Table 2
	Commodity           string // Table 3
	Units               string // Table 4
	CalculationConstant string // float32 Multiplier to convert data values to
	// engineering units.
	Interval string // Time interval between readings. 00000015
	Count    string // int32 Number of triples to follow.
	// Maximum of 48 allowed per record.
	Triples []Interval // Stores the Interval Data
}

// Interval Block of Data, can up Register and Interval Data
// TODO(cn): Parse out the Bit for the Data Quality Flag, this value
// holds a text value and a bit value, for two values with seperate meanings
type Interval struct {
	EndTime         string // End time of the interval
	DataQualityFlag string // Data Quality Flag: see Table 6.
	MeasuredValue   string // The measured value
}

// Main Function Entry Point
func main() {

	// Open the file
	file, err := os.Open("cmep.dat")

	// On Error Close
	if err != nil {
		log.Fatal(err)
	}

	// defer the close so we don't leave a locked file open during an error
	defer file.Close()

	// Create a map to hold the counts of the record locators
	recordformats := make(map[string]int)

	//TODO we need to process by RecordType so we need a series of If statements

	//Create the slice of the struct
	mepmd01x := make([]MEPMD01x, 0)

	//Start processing the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		// Example Line by Line Processing
		ProcessLine(text, recordformats)

		// Batching the Results into on RecordType
		mepmd01x = ProcessBatchLine(text, mepmd01x)
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

//A CMEP file is comma delimited file where each line contains
//a specific set of values based on the record locator.
//The record locator is the first value of every line in the file.
func ProcessLine(line string, recordformats map[string]int) MEPMD01x {

	// We only want the record format which is the first value in this slice
	// We are only working with one line at a time
	// this will be many thousands of hits to the database
	// another method would be to batch the results and then push

	// split the csv data
	// need to rework and pass the values around since have
	// to split it for both functions, this way it's done once
	values := strings.Split(line, ",")
	recordformats = RecordLocatorCount(values[0], recordformats)

	data := RecordFormatTransform(line)

	// Printing of the results until I have a test setup
	//json, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//Quick and Dirty Print Here
	//fmt.Println(string(json))
	//fmt.Println("\n")

	return data
	// POST Data to Database, here or in another function
}

// We are only working with one line at a time
// this will be many thousands of hits to the database
// another method would be to batch the results and then push
// Here we are assuming process the entire file and then
// allow another function to post the data to the database
func ProcessBatchLine(line string, records []MEPMD01x) []MEPMD01x {

	data := RecordFormatTransform(line)
	// Expand the slice, expansion does decrease performance
	records = append(records, data)

	return records
}

// Transform the CMEP csv values to the MEPMD01x struct
// and return this struct back the to calling function
func RecordFormatTransform(line string) MEPMD01x {

	//split the csv data
	values := strings.Split(line, ",")

	//values[13] contains the count for intervals at the
	//end of the cmep csv, we can have a max of 48
	count, err := strconv.Atoi(values[13])
	if err != nil {
		// handle error
		fmt.Fprintln(os.Stderr, "convert string to int:", err)
	}

	// After the 14th position, is where the interval data is found
	intervalstartposition := 14

	// Create a slice of all the intervals we need to store
	// We can use a slice because CMEP returns the number of intervals
	// expected, if this is wrong we will to err out and recover and
	// keep processing the data.  We could have more intervals than
	// what the data suggests
	intervals := make([]Interval, count)

	for i := 0; i < count; i++ {
		intervals[i].EndTime = values[intervalstartposition+i]
		intervals[i].DataQualityFlag = values[intervalstartposition+i+1]
		intervals[i].MeasuredValue = values[intervalstartposition+i+2]
	}

	// Populate the data
	// TODO: I should rework this to use .Notation for the values[x]
	// is easier to read
	data := MEPMD01x{values[0],
		values[1],values[2],values[3],values[4],values[5],
		values[6],values[7],values[8],values[9],values[10],
		values[11],values[12],values[13], intervals}

	return data

}

// Create map containing the RecordLocator and the number
// of instances the RecordLocator has occured.  This is for
// reporting purposes only.
func RecordLocatorCount(rl string, counts map[string]int) map[string]int {
	counts[rl]++
	return counts
}
