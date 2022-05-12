package string

import (
	"strconv"
	"time"
)

// interface change to string
func ToString(a interface{}) string {
	switch a.(type) {
	case string:
		return a.(string)
	case int:
		return strconv.Itoa(a.(int))
	case uint:
		return strconv.Itoa(int(a.(uint)))
	case uint8:
		return strconv.Itoa(int(a.(uint8)))
	case uint16:
		return strconv.Itoa(int(a.(uint16)))
	case uint32:
		return strconv.FormatUint(uint64(a.(uint32)), 10)
	case uint64:
		return strconv.FormatUint(a.(uint64), 10)
	case int8:
		return strconv.Itoa(int(a.(int8)))
	case int16:
		return strconv.Itoa(int(a.(int16)))
	case int32:
		return strconv.Itoa(int(a.(int32)))
	case int64:
		return strconv.FormatInt(a.(int64), 10) // 在32的机子上，int是32位
	case float32:
		return strconv.FormatFloat(float64(a.(float32)), 'f', -1, 64)
	case float64:
		return strconv.FormatFloat(a.(float64), 'f', -1, 64)
	case time.Time:
		return a.(time.Time).Format("2006-01-02 15:04:05")
	case []uint8:
		var bdata []byte
		for _, b := range a.([]uint8) {
			bdata = append(bdata, byte(b))
		}
		return string(bdata)
	case error:
		return a.(error).Error()
	default:
		panic("This type is not supported!")
	}
}
