package tool

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

type JsonCamel struct {
	Value interface{}
}

// 重写 MarshalJSON()
func (c JsonCamel) MarshalJSON() ([]byte, error) {
	var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
	marshalled, err := json.Marshal(c.Value)
	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			matchStr := string(match)
			fmt.Println(matchStr)
			key := matchStr[1 : len(matchStr)-2]
			resKey := Lcfirst(StrToCamel(key))
			return []byte(`"` + resKey + `":`)
		},
	)
	return converted, err
}

// 下划线写法转为驼峰写法 _转驼峰
func StrToCamel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	name = strings.Replace(name, " ", "", -1)
	return name
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
