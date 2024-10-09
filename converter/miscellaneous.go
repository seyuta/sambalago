package converter

import (
	"log"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ToTime convert any value to time.Time
// other than string type, parseFormat use empty string
func ToTime(value interface{}, parseFormat string) (result time.Time) {
	if value == nil {
		return
	}

	switch v := value.(type) {
	case time.Time:
		return v
	case string:
		t, err := time.Parse(parseFormat, v)
		if err != nil {
			log.Printf("failed to convert value ToTime: [%#v] with error: [%v]\n", v, err)
			return
		}

		return t
	case int:
		return time.Unix(int64(v), 0)
	case int32:
		return time.Unix(int64(v), 0)
	case int64:
		return time.Unix(v, 0)
	case float32:
		return time.Unix(int64(v), 0)
	case float64:
		return time.Unix(int64(v), 0)
	case primitive.DateTime:
		return v.Time()
	default:
		log.Printf("unsuported type: %s", reflect.TypeOf(value))
		return
	}
}

// ToBool convert any value to boolean
func ToBool(value interface{}) (result bool) {
	if value == nil {
		return false
	}

	switch v := value.(type) {
	case string:
		var err error
		str := strings.TrimSpace(v)
		r, err := strconv.ParseBool(str)
		if err != nil {
			log.Printf("failed to convert value ToBool: [%#v] with error: [%v]\n", v, err)
		}

		return r
	case int:
		return v != 0
	default:
		return false
	}
}

// Function to convert int64 to roman numeral
func Int64ToRoman(num int64) string {
	romanMap := []struct {
		value   int64
		numeral string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result := ""

	for _, item := range romanMap {
		for num >= item.value {
			result += item.numeral
			num -= item.value
		}
	}

	return result
}

// Function to convert roman numeral to int64
func RomanToInt64(s string) int64 {
	romanMap := map[byte]int64{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	result := int64(0)
	n := len(s)

	for i := 0; i < n; i++ {
		if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
			result -= romanMap[s[i]]
		} else {
			result += romanMap[s[i]]
		}
	}

	return result
}
