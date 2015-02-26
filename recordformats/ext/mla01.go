package mla01

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// MLA01: Meter Level Alarms

type MLA01 struct {
	RecordType          string // MLA01
	RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
	SenderID            string // Fixed value, Text
	SenderCustomerID    string // Sensus code - customer:Key for flexible fields
                             // ACME:012000
	ReceiverID          string // Text, Flexible field – see Table 9 for options
	ReceiverCustomerID  string // Text, Flexible field – see Table 9 for options
	TimeStamp           string // Date&time this record was created YYYYMMDDHHMM
                             // 200801310855
	MeterID             string // Text, ABC123, Flexible field – see table 9 for options
  Purpose             string // Text, "OK"
  Commodity           string // Text, Values {E:Electric,W:Water,G:Gas,S:Steam}
	Units               string // Fixed for this record type. {METERDQ}
	CalculationConstant string // Always blank. Does not apply.
                             // for this customer record type only

	Interval string 					 // Usually the width of time intervals used when
                             // all values are not time stamped. Will not apply
                             // in this customer record type because every
                             // triple will contain a timestamp indicating
                             // the time of event. Always blank.

  Count    string 					 // Number of alarms (triples) reported in the row.

	Triples []Interval 				 // Interval Data, Maximum of 48 allowed per record.
                             // [n][0] Text, 200801301128. Time of the Event
                             // [n][1] Fixed as RO
                             // [n][2] int32, 2 : Meter Alarms: Table 7
                             // meter alarm can be smaller to reduce memory
}

// Interval Block of Data, can up Register and Interval Data
// TODO(cn): Parse out the Bit for the Data Quality Flag, this value
// holds a text value and a bit value, for two values with seperate meanings
type Interval struct {
	TimeOfEvent     string // End time of the interval
	DataQualityFlag string // Data Quality Flag: see Table 6.
	MeterAlarm      int32 // The measured value
}
