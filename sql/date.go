package sqlStruct

// custom datetime struct
type MyDate struct {
	MyDateTime
}
var dateFormat =  "2006-01-02"

func init() {
	dateTimeFormat = dateFormat
}
