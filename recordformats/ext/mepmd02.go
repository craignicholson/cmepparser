package mepmd02

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// MEPMD02: Metering Data Type 2 – TOU Data, Net Metering

type MEPMD02 struct {
	RecordType          string // MEPMD02
	RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
	SenderID            string // Fixed value, text
	SenderCustomerID    string // Sensus code - customer:Key for flexible fields
														 // ACME:012000
	ReceiverID          string // Flexible field – see Table 9 for options
	ReceiverCustomerID  string // Flexible field – see Table 9 for options
	TimeStamp           string // Date&time this record was created YYYYMMDDHHMM
														 // 200801310855
	RecordID            string // Optional field in MEPEC01, blank
  OperationType       string // “CFG” used for current configuration,
                             // no change has occurred.

  Purpose             string // Text,‘OK’ for normal transmission.
	Comment             string // Text, Optional field in MEPEC01
	Commodity           string // Text, Values {E,W,G,S : Electric, Water, Gas Steam}
	Activity            string // <Blank> used to indicate an automatic
                             // transmission caused this record to be sent.
  EquipmentType       string // The kind of equipment in the record. 'METER'
  Manufacture         string // Manufacturers name, not needed for this extract.
                             // blank
  Model               string // Device’s model name, not needed for this extract.
                             // blank
  SerialNumber        string // Flexible field – see Table 9 for options
  Identifier          string // Optional field in MEPEC01. Not needed for this extract
                             // blank
  DateOfPurchase      string // The meter inventory date. <CCYYMMDD>
  DateOfInstallation  string // The meter installation date.
                             // Blank installation date indicates a meter
                             // that has not been installed.  <CCYYMMDD>
  Owner               string // float32, Owner of the device.
                             // Not needed in this extract. blank
	Count    string 					 // int32 Number of triples to follow.

	Data []Interval 			 	   // Interval Data, Maximum of 48 allowed per record. ????
                             // [n][0] Text, Parameter name
                             // [n][1] Text, Parameter value
                             // Parameter names/values: Table 10.


}

// Interval Block of Data, can up Register and Interval Data
// TODO(cn): Parse out the Bit for the Data Quality Flag, this value
// holds a text value and a bit value, for two values with seperate meanings
type Data struct {
	TouLable        string  // OnPeak: Time-of-use label. Table 5
	DataQualityFlag string  // Data Quality Flag: see Table 6.
	MeasuredValue   float32 // float64, 345.6789 : Measured Value
}
