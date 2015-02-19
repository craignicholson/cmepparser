package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//TODO:  Create struct for each RecordFormat, a few vendors have
//extended the RecordFormat and made the CMEP format into a
//propritary format.  The user will have to identify which format to
//load the data into by passing in a flag for the format, Original or extended
//MEPMD01x: Metering Data Type 1 – Interval Data, Pulse Data, Reference Register Reads
//MEPMD02x: Metering Data Type 2 – TOU Data, Net Metering
//MLA01x: Meter Level Alarms
//MEPEC01x: Equipment Configuration Type 1 – Meter configuration information

//TODO:  Original CMEP Specifications
//MEPAD01" - Administrative Data Type 1 - DASR
//MEPAD02" - Administrative Data Type 2 - Credit Data
//MEPMD01" - Metering Data Type 1 - Interval Data
//MEPMD02" - Metering Data Type 2 - TOU Data
//MEPBD01" - Billing Data Type 1 - Billed Dollars
//MEPBD02" - Billing Data Type 2 - Interval Pricing Plan
//MEPBD03" - Billing Data Type 3 - TOU Pricing Plan
//MEPLF01" - Distribution Loss Factors - Electric
//MEPEC01" - Equipment Configuration Type 1
//MEPRR01" - Record Reject Type 1

//TODO: For each file loaded return a [] struct  of the results over the wire
//to be logged in a database.
//FileName, Meter, Errors [Missing REad]

//TODO: Emit to REST Service, a message when a file fails to Parse.
//TODO: Emit to REST Service, a message when a file loads with the Analysis.

//TODO: Mark the file as done, complete, finished so a filemover can
//move the file somewhere

//TODO: Write to another file with the extension .err for error when we
//have an error in the data which failed to load.


func main() {
	//Open the file
	file, err := os.Open("sensus_cmep_extended.dat")

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
		ProcessLine(text,recordformats)
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
func ProcessLine(line string, recordformats map[string]int){

	//We only want the record format which is the first value in this slice
	values := strings.Split(line, ",")
	recordformats = RecordLocatorCount(values[0], recordformats)
}

//Analysis scans the files and generates a set of keys and values
//which describes the content of the file for historical review
//and reporting.  This data can be used for ...
func Analysis(){

}

func ErrorCollection(){

}

//Create map containing the RecordLocator and the number
//of instances the RecordLocator has occured.
func RecordLocatorCount(rl string, counts map[string]int) map[string]int {
	counts[rl]++
	return counts
}
