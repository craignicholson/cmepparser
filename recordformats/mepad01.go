// Package implements the MEPAD01: Administrative Data Type 1 - DASR

package mepad01

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

// MEPAD01: Administrative Data Type 1 - DASR
type MEPAD01 struct{
  RecordType          string // MEPAD01
  RecordVersion       string // Fixed value-Release date to production. YYYYMMDD
  SenderID            string // Fixed value, text
  SenderCustomerID    string // Sensus code - customer:Key for flexible fields
                             // ACME:012000
  ReceiverID          string // Flexible field – see Table 9 for options
  ReceiverCustomerID  string // Flexible field – see Table 9 for options
  TimeStamp           string // Date&time this record was created CCYYMMDDHHMM
                             // 200801310855
  RecordID            string // Text, This is an optional field that may be
                             // supplied in a request record. The contents of
                             // this field will be returned unchanged in the
                             // corresponding response record. The length of
                             // this text shall not exceed 16 characters.
  OperationType       string // Text, Protocol Text: What kind of operation triggered
                             // this record to be transmitted. See "MEPAD01 Operations"
                             // below for examples of field use. Defined values are below:
  SvcRealtionshipCount string // Numeric Integer: The number of relationships included
                              // in the following "Type of Service Relationship" field.
  TypeofServiceRelationship string // Protocol Text: The kind of account/entity
                                   //  relationship this record describes.
                                   // Multiple entries may be supplied. Entries are separated
                                   // by ASCII space character (20 Hexadecimal).
                                   // The length of this field shall not exceed 72 characters               string // Text, KWHREG
  Reason	                  string // Protocol Text: Why this transmission is sent.
  Comment                   string //Arbitrary Text: An optional field used to supply
                                   // additional information about the indicated operation.
                                   // This field is typically used in "NAK" transmissions
                                   // to indicate the reason for rejecting a request.
                                   // It is also used to indicate the reason a meter could
                                   // not be read in a "NO-READ" record. The length of this
                                   // text shall not exceed 64 characters.
  UDCID                     string // Arbitrary Text: Identity of the UDC.
                                   // It will typically be an abbreviation of the UDC company name.
                                   // Currently, only the first 16 characters of this field
                                   // will be recognized by PG&E.
  UDCAccountID              string // Arbitrary Text: ESP account information assigned by UDC.
                                   // This field is supplied to facilitate special contract
                                   // accounting. Currently, only the first 16 characters of
                                   // this field will be recognized by PG&E.
  EffectiveStartDate        string // Date/Time ("CCYYMMDDHHMM"): Communicates requested
                                   // effective date when used supplied in requests for change in account status. Communicates actual effective date in responses or update notices.
  EffectiveEndDate          string // Date/Time("CCYYMMDDHHMM"):Thedatethatthis "Type of Service
                                   // Relationship" account was closed.
  AccountStatus             string // Protocol Text: A descriptive abbreviation of the
                                   // status of this account.
  PendingStatus             string // Protocol Text: A status as described in the
                                   // "Metering account status" field above that will
                                   // take effect some time in the future. This is used by
                                   // UDC to notify SP that account status will change soon.
                                   // The effective change time is described in the
                                   // "Effective start date" or "Effective end date"
                                   // fields above.
  PendingSPidentifier       string // Arbitrary Text: This identifies the SP that will
                                   // "Type of Service Relationship" connection to this
                                   // customer at the date and time described in
                                   // "Pending Effective Date". It will typically be an
                                   // abbreviation of the SP's company name. Currently,
                                   // only the first 16 characters of this field will be
                                   // recognized by PG&E.
  ReadingEstimationMethod   string // Protocol Text: This is a description of the estimation rules
                                   // applied to estimate values for missing data.
                                   // "PG&E" - PG&E's internal estimation method.
                                   // "MADAWG01" - Metering and Data Access Working Group method version
                                   // "NONE" - No estimation will be done. Defined values are:
  Commodity                 string // Text, Values {E:Electric,W:Water,G:Gas,S:Steam}
  CustomerName              string // Arbitrary Text: This is the customer's complete name.
                                   // It is used primarily with commercial accounts for company
                                   // name. Currently, only the first 22 characters of this field
                                   // will be recognized by PG&E, 8 by SCE.
  ContactLastName           string // Arbitrary Text: This is the customer's last name.
                                   // Currently, only the first 22 characters of this field will
                                   // be recognized by PG&E.
  ContactFirstName          string // ArbitraryText:Thisisthecustomer'sfirstname.  Currently, only
                                   // the first 22 characters of this field will be recognized by PG&E.
  ContactMiddleInitial      string // Arbitrary Text: This is the customer's middle initial.
                                   // Currently, only 1 character of this field will be recognized by PG&E.
  HouseBldNmbr              string // Arbitrary text: This is part of the street address.
                                   // Currently, only the first 10 characters of this field will
                                   // be recognized by PG&E, 6 by SCE.
  HouseBldNmbrFractionNmbr  string // Arbitrary text: This is part of the street address.
                                   // Currently, only the first 5 characters of this field will be
                                   // recognized by PG&E, 3 by SCE.
  StreetPrefix              string // Arbitrary text: This is part of the street address. Currently,
                                   // only the first 5 characters of this field will be
                                   // recognized by PG&E, 2 by SCE.
  StreetName                string // Arbitrarytext  //Thisispartofthestreetaddress.Currently,
                                   // onlythe first 22 characters of this field will be
                                   // recognized by PG&E, 25 by SCE.
  StreetSuffix              string // Arbitrary text: This is part of the street address.
                                   // Currently, only the first 5 characters of this field will
                                   // be recognized by PG&E, 4 by SCE.
  UnitNumber                string // Arbitrary text: This is part of the street address.
                                   // Currently, only the first 10 characters of this field will
                                   // be recognized by PG&E, 8 by SCE.
  City                      string // Arbitrary text: This is the address city of the meter.
                                   // Currently, only the first 12 characters of this field
                                   // will be recognized by PG&E, 25 by SCE.
  State                     string // Arbitrary text: This is the address state of the meter.
                                   // This is a standard abbreviation for the state or province.
                                   // Currently, only the first 2 characters of this field will
                                   // be recognized by PG&E, 2 by SCE.
  Country                   string // Arbitrary text: This is the address country of the meter.
                                   // Currently, only the first 15 characters of this field will be recognized by PG&E.
  ZIP                       string // Arbitrary text: This is the address zip code of the meter.
                                   // Currently, only the first 5 characters of this field
                                   // will be recognized by PG&E, 5 by SCE.
  ZIPExtension4             string // Arbitrary text: This is an extension to the address zip code
                                   // of the meter. Currently, only the first 4 characters of this
                                   // field will be recognized by PG&E, 4 by SCE.
  USPSCarrierRoute          string // Arbitrary text: This is an extension to the address zip code
                                   // of the meter. Currently, only the first 2 characters of
                                   // this field will be recognized by PG&E, 4 by SCE.
  StandardTimeZone          string // Numeric Integer: (Generated by MA) Time zone for local
                                   // time calculations when daylight savings time is NOT in
                                   // effect. This value is in minutes difference from
                                   // Universal Coordinated Time (UTC) which, for the purposes
                                   // of this document, is the same as GMT without
                                   // daylight savings time applied. Pacific Standard Time
                                   // has the value -480 (negative four hundred eighty),
                                   // Eastern Standard Time -300 (negative three hundred).
  DaylightTimeZone          string // Numeric Integer: (Generated by MA) Time zone for local time
                                   // calculations when daylight savings is in effect.
                                   // If daylight savings time change is not to be used,
                                   // this field is left empty. Standard algorithms are used
                                   // to calculate when standard versus daylight savings time
                                   // is to be applied. This value is in minutes difference
                                   // from UTC. Pacific Daylight Savings Time has the value -420
                                   // (negative four hundred twenty), Eastern Standard Time -240
                                   // (negative two hundred forty).
  ServiceCategory           string // Protocol Text: The category of service. This information
                                   // is used to calculate distribution loss costs.
                                   // Defined values are:
                                   // A. “S” - Secondary (typically service at less than 2KV).
                                   // B. “P” - Primary (typically service at greater than 2KV and less than 60KV).
                                   // C. "PS" - Primary Subtransmission.
                                   // D. “T” - Transmission (typically service at greater than 60 KV)
  MeterCongestionZone       string // Arbitrary Text: (Generated by UDC) The ISO distribution
                                   // congestion zone identifier. This may alternatively
                                   // be used to indicate Load Group, Load Point, Grid Takeout
                                   // Point, or a combination thereof. Currently, only the
                                   // first 20 characters of this field will be recognized by PG&E.
  UsageProfile              string // Arbitrary Text: (Generated by UDC) The description of this
                                   // accounts usage class. Currently, only the first 12
                                   // characters of this field will be recognized by PG&E.
  BillingOption             string // Protocol Text: This is a description of which entity or
                                   // entities perform billing for service. Defined values are:
                                   // A. "DUAL" - SP bills for service, UDC bills additional fees.
                                   // B. "UDC" - UDC bills customer.
                                   // C. "SP" - SP bills customer.
  UDCRateName               string // Arbitrary Text: (Generated by UDC) UDC rate schedule
                                   // is commonly used with bill ready accounts. Currently,
                                   // only the first 12 characters of this field will be
                                   // recognized by PG&E.
  SPRateName                string // Arbitrary Text: (Generated SP) SP rate schedule required
                                   // for BA or UDC-Consolidated billing. For bill ready accounts,
                                   // the text string "BILL-READY” may be used. Currently, only
                                   // the first 12 characters of this field will be recognized by
                                   // PG&E.
  PhoneInternationalAccess  string // Arbitrary Text: Part of customer phone number.
                                   // Currently, only the first 3 characters of this field
                                   // will be recognized by PG&E.
  PhoneAreaCode             string // Arbitrary Text: Part of customer phone number.
                                   // Currently, only the first 3 characters of this field
                                   // will be recognized by PG&E.
  PhoneNumber               string // Arbitrary Text: Part of customer phone number.
                                   // Currently, only the first 7 characters of this field
                                   // will be recognized by PG&E.
  PhoneExtensionNmber       string // Arbitrary Text: Part of customer phone number.
                                   // Currently, only the first 6 characters of this field
                                   // will be recognized by PG&E.
  FAXnumber                 string // Arbitrary Text: Customer's FAX number. Currently,
                                   // only the first 20 characters of this field
                                   // will be recognized by PG&E.
  RenewableEnergyProvider   string // Protocol Text: (Generated by SP) Indicates renewable energy
                                   // provided for this account. Defined values are:
                                   // A. (blank) - Renewable energy not provided for this account.
                                   // B. “N” - Renewable energy not provided for this account.
                                   // C. “Y” - Renewable energy provided for this account.
  MeterCount                string // Numeric Integer: (Generated by MA) The number of "Meter"
                                   // fields to follow. A maximum of 12 is allowed.
  Meter                     string // Arbitrary Text, Time Interval ("MMDDHHMM"),
                                   // Protocol Text triplet: Each data entry is a set of
                                   // three fields. The number of meter entry sets is
                                   // described in the "Meter Count" field above.
                                   // The "Arbitrary Text" entry is the Meter ID.
                                   // This is the placard identifier or faceplate serial number
                                   // to physically identify a meter. This is usually
                                   // some arbitrary combination of letters and numbers
                                   // that make up a meter manufacturer's serial number.
                                   // It may, however, be some other easily found identifying
                                   // label on the metering equipment. This field may optionally
                                   // be used as a channel identifier for situations where
                                   // that information is useful. Currently, only the first 12 characters
                                   // of this entry will be recognized by PG&E.
                                   // The "Time Interval" entry is the Usage reading interval,
                                   // the time interval that meter data is supplied.
                                   // For example, monthly read meters would be encoded
                                   // as "01000000", weekly as "00070000, hourly as "00000100",
                                   // and 30 minute as "00000030".
                                   // Note: "KWHREG", "KVAHREG", KVARHREG", and "GASREG"
                                   // readings are special cases and are always assumed
                                   // to be available in monthly intervals only for
                                   // inclusion on printed bills. The "Protocol Text:"
                                   // contains the Units parameters. This is a list of metering
                                   // units supplied for this account. When multiple data types
                                   // are available, individual abbreviations are separated by
                                   // ASCII Space characters (20 Hexadecimal).
                                   // For example, Meter reading plus kilowatt hours
                                   // is "KWHREG KWH"; meter reading, kilowatt hours,
                                   // and kilovar hours is "KWHREG KWH KVARH".
                                   // A complete list of abbreviations is supplied in the
                                   // Protocol Text Units listing. A maximum of 8 unit
                                   // parameter entries per field is allowed.
                                   // The length of this text shall not exceed 64 characters.




}

/*
Operation Types

"CUST-REQ" - (Customer to UDC) establish/break direct access with SP.
"CUST-ACK" - (UDC to Customer) acknowledge success of access request.
"CUST-NAK" - (UDC to Customer) reject an access request.
"SP-REQ" - (SP to UDC) establish/break direct access with customer.
"SP-ACK" - (UDC to Customer) acknowledge success of access request.
"SP-NAK" - (UDC to SP) reject an access request.
"ACNT-REQ" - (SP to UDC) request resend of account information.
"ACNT-RESP" - (UDC to SP) resend account information.
"MD-REQ" - (ESP or UDC to MA) request resend of metering data.
"MD-ACK" - (MA to UDC or ESP) acknowledge request for resend metering data.
"MD-NAK" - (MA to SP) reject a request for resend.
"BD-REQ" - (SP to UDC or UDC to SP) request resend of billing data.
"BD-ACK" - (UDC to SP or SP to UDC) acknowledge request for resend billing data.
"BD-NAK" - (UDC to SP or SP to UDC) reject a request for resend of billing data.
"SVC" - (UDC to SP) notify of shutoff or turn-on of service.
"CFG" - (MA to UDC or ESP) notify of metering configuration change.
"METER" - (MA to UDC or ESP) notify of meter change out.
"BILL-ADDR" - (UDC to SP or SP to UDC) notify of billing addresschange.
"ACK" - (SP to UDC or UDC to SP) Acknowledge notification.

TypeofServiceRelationship
A. "MTR-OWN" - Owner of meter.
B. "MTR-INST" - Meter installer or maintainer.
C. "MTR-RDR" - Meter reader.
D. "MTR-AGNT" - Meter agent for customer.
E. "ELEC-ESP" - Electric Energy Service Provider for customer.
F. "ELEC-SC" - Electric Scheduling Coordinator for customer.
G. "GAS-ESP" - Gas Energy Service Provider for customer.
H. "BILLER" - Customer account billing agent.
I. "BILL-CAL" - Customer's agent for billing calculations.
J. "UDC" - Customer's Utility Distribution Company.

Reason
A. "UPDATE" - report a change in status.
B. "RESEND" - repeat of a previous configuration.
C. "ADJUST" - an adjustment to a previously sent configuration that may
involve billing corrections.
D. "CORRECT" - a correction to a previously sent configuration that does not
involve billing changes.
E. "CONNECT" - Request direct access.
F. "DISCONNECT" - Request direct access be discontinued.
G. "NO-READ" - This is notice that a meter could not be read. The reason the
meter could not be read should be placed in the comment field below.

Account Status
"NEW" - Defined but not active.
"PEND-SP" - Customer has requested direct access status change. Change is pending, waiting for SP's request.
"PEND-CUS" - SP has requested direct access status change. Change is pending, waiting for Customer's request or contract execution.
"PEND-MTR" - Customer has requested direct access status change. Change is pending; meter changeout required.
"PEND" - Change in direct access status is approved. Waiting for direct access transfer date that will usually be the beginning of the next billing period.
"OK" - Active account.
"OFF" - Shut-Off.
"INACT" - Inactive.
"DEL" - Deleted.
"NO-DATA" - An active account has no metering data available. This
indicates that one or more readings have been missed and estimation is not
allowed by this account's estimation rules.

*/
