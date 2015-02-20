# cmepparser
California Meter Exchange Protocol (CMEP) Parser

//TODO: Take the CMEP and transform to JSON

//TODO: Take the CMEP and transform to JSON and post to mongodb

//TODO: Take the CMEP and transform and push over the wire to a go rest
//interface which inserts into mongodb


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
