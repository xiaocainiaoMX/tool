package sqlStruct

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// custom datetime struct
type MyDate sql.NullTime

const dateFormat = "2006-01-02"

func (v MyDate) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time.UTC().Format(dateFormat))
	} else {
		return json.Marshal("")
	}
}
func (v *MyDate) UnmarshalJSON(data []byte) error {
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
func (t MyDate) Value() (driver.Value, error) {
	if t.Valid {
		return t.Time.UTC().Format(dateFormat), nil
	} else {
		return "", nil
	}
}
func (t *MyDate) Scan(v interface{}) error {
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
func (t *MyDate) String() string {
	if t.Valid {
		return t.Time.UTC().Format(dateFormat)
	} else {
		return ""
	}
}
