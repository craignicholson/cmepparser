package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//CMEP MEPMD01 extended record format
//Metering Data Type 1 – Interval Data, Pulse Data, Reference Register Reads
type MEPMD01x struct {
	RecordType          string
	RecordVersion       string //Fixed value -Release date of this protocol to production. YYYYMMDD
	SenderID            string //Fixed value
	SenderCustomerID    string //Sensus code for the customer:Key for flexible fields
	ReceiverID          string //Flexible field – see Table 9 for options
	ReceiverCustomerID  string //Flexible field – see Table 9 for options
	TimeStamp           string //Date and time this record was created YYYYMMDDHHMM
	MeterID             string //Flexible field – see table 9 for options
	Purpose             string //Table 2
	Commodity           string //Table 3
	Units               string //Table 4
	CalculationConstant string // float32 Multiplier to convert data values to engineering units.
	Interval            string //Time interval between readings. 00000015
	Count               string //int32 Number of triples to follow.  Maximum of 48 allowed per record.

	//End time of the interval (may be left empty after the first triple if Interval field is provided).
	//Data Quality Flag: see Table 6.
	//The measured value
	Triples [][]string
}

func main() {
	//Open the file
	file, err := os.Open("cmep.dat")

	//On Error Close
	if err != nil {
		log.Fatal(err)
	}
	//defer the close so we don't leave a locked file open during an error
	defer file.Close()

	//Create a map to hold the counts of the record locators
	//Should I make a pointer here? Instead of passing the variable around
	recordformats := make(map[string]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		ProcessLine(text, recordformats)
	}

	//Quick output of found types
	fmt.Printf("map := %v", recordformats)

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

//A CMEP file is comma delimited file where each line contains
//a specific set of values based on the record locator.
//The record locator is the first value of every line in the file.
func ProcessLine(line string, recordformats map[string]int) {

	//We only want the record format which is the first value in this slice
	values := strings.Split(line, ",")
	recordformats = RecordLocatorCount(values[0], recordformats)

	RecordFormatTransform(line)
}

//Map the CMEP csv values to the struct
func RecordFormatTransform(line string) {
	values := strings.Split(line, ",")
	fmt.Println(len(values))

	//values[13] contains the count for intervals at the
	//end of the cmep csv, we can have a max of 48
	count, err := strconv.Atoi(values[13])
	if err != nil {
		// handle error
		fmt.Fprintln(os.Stderr, "convert string to int:", err)
	}

	//The number of triples after the 14th position
	intervalposition := 14

	//End time of the interval (may be left empty after the first triple if Interval field is provided).
	//Data Quality Flag: see Table 6.
	//The measured value
	//Create a slice of all the intervals we need to store
	intervals := make([][]string, count)

	//Populate the slice with the triples of data [EndTime, DataQuality, Value]
	for i := 0; i < count; i++ {
		innerLen := 3
		intervals[i] = make([]string, innerLen)
		for j := 0; j < 3; j++ {
			x := i + j + intervalposition
			intervals[i][j] = values[x]
		}
	}

	//Make a silce of the struct
	data := MEPMD01x{values[0],
		values[1],
		values[2],
		values[3],
		values[4],
		values[5],
		values[6],
		values[7],
		values[8],
		values[9],
		values[10],
		values[11],
		values[12],
		values[13], intervals}

	//TODO: Write a test, for now review the output here.
	fmt.Println(data)

}

//Analysis scans the files and generates a set of keys and values
//which describes the content of the file for historical review
//and reporting.  This data can be used for ...
func Analysis() {

}

func ErrorCollection() {

}

//Create map containing the RecordLocator and the number
//of instances the RecordLocator has occured.
func RecordLocatorCount(rl string, counts map[string]int) map[string]int {
	counts[rl]++
	return counts
}
