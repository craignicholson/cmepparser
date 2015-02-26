package mepec01

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// MEPEC01: Equipment Configuration Type 1 – Meter configuration information
type MEPEC01 struct {
	RecordType          string // MEPEC01
	RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
	SenderID            string // Fixed value, text
	SenderCustomerID    string // Sensus code - customer:Key for flexible fields
														 // ACME:012000
	ReceiverID          string // Flexible field – see Table 9 for options
	ReceiverCustomerID  string // Flexible field – see Table 9 for options
	TimeStamp           string // Date&time this record was created YYYYMMDDHHMM
														 // 200801310855
	MeterID             string // Text, ABC123, Flexible field – see Table 9 for options
	Purpose             string // Text, "OK"
	Commodity           string // Text, Values {E,W,G,S : Electric, Water, Gas Steam}
	Units               string // Text, KWHREG
  SeasonIdentifier	  string // Text, ProtocolText	<empty>	Flexible field –
                             // see Table 9 for options.
  CalculationConstant string // float32 Multiplier to convert data values to
														 // engineering units.
  DateStartTime              // Date and time that the data interval reported
                             // in this record began.  200801310000
  DataTimeStamp              // Date and time that ends the interval reported
                             // in this record. 200802010000
	Interval string 					 // Time interval between readings. 00000015
	Count    string 					 // int32 Number of triples to follow.

	Triples []Interval 				 // Interval Data, Maximum of 48 allowed per record.
                             // [n][0] Text, OnPeak: Time-of-use label. Table 5
                             // [n][1] Text, R0 : Data quality flag. Table 6
                             // [n][2] float64, 345.6789 : Measured Value

}
