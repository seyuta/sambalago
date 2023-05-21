package converter

import (
	"encoding/json"
	"log"
	"strconv"
	"time"
)

// ToString converts any value to string
func ToString(value interface{}) (result string) {
	if value == nil {
		return
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case bool:
		return strconv.FormatBool(v)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []uint8:
		return string(v)
	case time.Time:
		return v.Format(time.RFC3339)
	default:
		resultJSON, err := json.Marshal(v)
		if err == nil {
			return string(resultJSON)
		} else {
			log.Printf("failed to convert data ToString: [%#v] with error: [%v]\n", v, err)
		}
	}

	return
}

// ArrIntToArrString converts an array of integers to an array of strings.
func ArrIntToArrString(arrInt []int) []string {
	result := make([]string, len(arrInt))
	for i, x := range arrInt {
		result[i] = strconv.Itoa(x)
	}
	return result
}

// ToArrayOfString convert any value to []string
func ToArrayOfString(value interface{}) (result []string) {
	r := make([]string, 0)
	switch v := value.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &r)
		if err != nil {
			log.Printf("failed to convert value ToArrayOfString: [%#v] with error: [%v]\n", v, err)
			return
		}

		return r
	case [][]byte:
		b := v
		for _, vv := range b {
			return append(result, string(vv))
		}
	default:
		return
	}

	return
}
