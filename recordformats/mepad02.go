// Package implements the MEPAD02: Administrative Data Type 2 - Credit Data

package mepad02

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// MEPAD02: Administrative Data Type 2 - Credit Data
type MEPAD02 struct{
  RecordType          string // MEPAD02
  RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
  SenderID            string // Fixed value, text
  SenderCustomerID    string // Sensus code - customer:Key for flexible fields
                             // ACME:012000
  ReceiverID          string // Flexible field – see Table 9 for options
  ReceiverCustomerID  string // Flexible field – see Table 9 for options
  TimeStamp           string // Date&time this record was created CCYYMMDDHHMM
                             // 200801310855
  MeterID             string // Arbitrary Text: This is the placard identifier or faceplate serial number
                             // to physically identify a meter. This is usually some arbitrary combination of
                             // letters and numbers that make up a meter manufacturer's serial number. It may,
                             // however, be some other easily found identifying label on the metering equipment.
                             // This field may optionally be used as a channel identifier for situations where that
                             // information is useful. Currently, only the first 12 characters of this entry will be
                             // recognized by PG&E.

  Purpose             string // Protocol Text: Indicates the reason for this data transmission. Defined
                             // values are:
                             // o "OK" - Normal transmission.
                             // o "RESEND" - Retransmission of previously sent data.
                             // o "SUMMARY" - Summary of SP totaled data. Summary data usually
                             // consists of values calculated from metering data such as monthly totals
                             // calculated from daily readings. This data is often supplied on a regular
                             // basis (such as for quarterly reports).
                             // o "HISTORY" - Archival account data. Archival data is retrieved from
                             // long term storage and may be of lesser time resolution than its original
                             // collection period. This data of generally supplied once per request for
                             // analysis purposes.
                             // o "PROFILE" - Account usage profile data.
                             // o "TEMPLATE" - Account usage template data.
                             // o "REJECT" - Data is rejected and is being returned to sender.
Commodity             string // Protocol Text, Values {E:Electric,W:Water,G:Gas,S:Steam}
Units                 string // Protocol Text: Describes the units of the data values. Examples of values
                             // are: "KWHREG", "KWH", and "THERM". A complete list of abbreviations is
                             // supplied in the Protocol Text Units listing. Where multiple unit types and seasons
                             // are transmitted, separate MEPMD02 records are sent for each. Data quality flags
                             // are used to indicate the raw, estimated, valid, etc. status of values transmitted.
SeasonIdentifier      string // Protocol Text: This identifies the season for which the values
                             // apply. Defined values are: "S" - Summer. "W" - Winter. This field may be left
                             // blank for accounts that do not differentiate between seasons. If this field is blank,
                             // it will be interpreted as indicating winter for those accounts that do. A record may
                             // contain data for one season only. Data for different seasons must be sent in
                             // separate records.
CalculationConstant  float32 // Floating-Point: Defines an optional value which
                             // is used as a multiplier to convert data values to engineering units. Typically this
                             // parameter is used with "PULSE" data to allow calculation of equivalent "KWH"
                             // and "THERM" values.
DataStartTime         string // ("CCYYMMDDHHMM"): Describes date and time
                             // that the data interval reported in this record began.
DataTimestamp         string // ("CCYYMMDDHHMM"): Describes date and time
                             // that ends the interval reported in this record.
Count                    int // Numeric Integer: Indicates the number of label- flag-value sets to follow.
                             // A maximum of 6 sets is allowed per record.
Data              []Interval // Protocol Text, Protocol Text, and Numeric Floating-Point
                             // triplet: Each
                             // data entry is a set of three fields.
                             // A maximum of 6 sets is allowed per record.
                             // Each set consists of a
                             // Protocol Text Time-Of-Use component label field, a
                             // Protocol Text data quality flag,
                             // and a Numeric Floating-Point value.
                             // The number of data entry sets is described in the "Count" field above.
                             // Defined values for the  quality flag field are described
                             // in the "MEPMD01" record above.
                             // (An empty indicates that the value is OK.)
                             // Defined values for the label field are:
                             // o "ON-PEAK"
                             // o "OFF-PEAK"
                             // o "PART-PEAK"
                             // o "PEAK-2"
                             // o "PEAK-3"
                             // o "PEAK-4"
                             // o "TOTAL"
}


// Interval Block of Data, can up Register and Interval Data
// TODO(cn): Parse out the Bit for the Data Quality Flag, this value
// holds a text value and a bit value, for two values with seperate meanings
type Data struct {
	TouLable        string  // OnPeak: Time-of-use label. Table 5
	DataQualityFlag string  // Data Quality Flag: see Table 6.
	MeasuredValue   float32 // float64, 345.6789 : Measured Value
}
