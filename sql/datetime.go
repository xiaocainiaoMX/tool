package sqlStruct

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// custom datetime struct
type MyDateTime sql.NullTime

var dateTimeFormat = "2006-01-02 15:04:05"

func (v MyDateTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time.UTC().Format(dateTimeFormat))
	} else {
		return json.Marshal("")
	}
}
func (v *MyDateTime) UnmarshalJSON(data []byte) error {
	var x *sql.NullTime
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
	} else {
		v.Valid = false
	}
	return nil
}
func (t MyDateTime) Value() (driver.Value, error) {
	if t.Valid {
		return t.Time.UTC().Format(dateTimeFormat), nil
	} else {
		return "", nil
	}
}
func (t *MyDateTime) Scan(v interface{}) error {
	if v == nil {
		t.Time, t.Valid = time.Time{}, false
		return nil
	}
	if vtime, ok := v.(time.Time); ok {
		t.Time, t.Valid = vtime, true
		return nil
	} else {
		t.Time, t.Valid = time.Time{}, false
		return errors.New("data type error!")
	}
}
func (t *MyDateTime) String() string {
	if t.Valid {
		return t.Time.UTC().Format(dateTimeFormat)
	} else {
		return ""
	}
}
