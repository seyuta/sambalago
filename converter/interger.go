package converter

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// ToInt converts any value to int
func ToInt(value interface{}) (result int) {
	switch v := value.(type) {
	case string:
		str := strings.TrimSpace(v)
		r, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert value ToInt: [%#v] with error: [%v]\n", v, err)
			return
		}

		return r
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case []byte:
		r, err := strconv.Atoi(string(v))
		if err != nil {
			log.Printf("failed to convert value ToInt: [%#v] with error: [%v]\n", v, err)
			return
		}

		return r
	default:
		return
	}
}

// ToInt16 converts any value to int16
func ToInt16(value interface{}) (result int16) {
	switch v := value.(type) {
	case string:
		str := strings.TrimSpace(v)
		r, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert value ToInt16: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int16(r)
	case int:
		return int16(v)
	case int16:
		return v
	case int32:
		return int16(v)
	case float32:
		return int16(v)
	case float64:
		return int16(v)
	case []byte:
		r, err := strconv.Atoi(string(v))
		if err != nil {
			log.Printf("failed to convert value ToInt16: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int16(r)
	default:
		return
	}
}

// ToInt32 converts any value to int32
func ToInt32(value interface{}) (result int32) {
	switch v := value.(type) {
	case string:
		str := strings.TrimSpace(v)
		r, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert value ToInt32: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int32(r)
	case int:
		return int32(v)
	case int32:
		return v
	case float32:
		return int32(v)
	case float64:
		return int32(v)
	case []byte:
		r, err := strconv.Atoi(string(v))
		if err != nil {
			log.Printf("failed to convert value ToInt32: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int32(r)
	default:
		return
	}
}

// ToInt64 converts any value to int64
func ToInt64(value interface{}) (result int64) {
	switch v := value.(type) {
	case string:
		str := strings.TrimSpace(v)
		r, err := strconv.Atoi(str)
		if err != nil {
			log.Printf("failed to convert value ToInt64: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int64(r)
	case int:
		return int64(v)
	case int32:
		return int64(v)
	case int64:
		return v
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	case []byte:
		r, err := strconv.Atoi(string(v))
		if err != nil {
			log.Printf("failed to convert value ToInt64: [%#v] with error: [%v]\n", v, err)
			return
		}

		return int64(r)
	default:
		return
	}
}

// ToArrayOfInt convert any value to []int
func ToArrayOfInt(v interface{}) (result []int) {
	switch v := v.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			log.Printf("failed to convert value ToArrayOfInt: [%#v] with error: [%v]\n", v, err)
			return
		}

		return
	case []string:
		b := v
		for _, vv := range b {
			return append(result, ToInt(vv))
		}

		return
	case [][]byte:
		b := v
		for _, vv := range b {
			return append(result, ToInt(vv))
		}

		return
	default:
		return
	}
}

// ToArrayOfInt32 convert any value to []int32
func ToArrayOfInt32(v interface{}) (result []int32) {
	switch v := v.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			log.Printf("failed to convert value ToArrayOfInt32: [%#v] with error: [%v]\n", v, err)
			return
		}

		return
	case []string:
		b := v
		for _, vv := range b {
			result = append(result, ToInt32(vv))
		}

		return
	case [][]byte:
		b := v
		for _, vv := range b {
			result = append(result, ToInt32(vv))
		}

		return
	default:
		return
	}
}

// ToArrayOfInt64 convert any value to []int64
func ToArrayOfInt64(v interface{}) (result []int64) {
	switch v := v.(type) {
	case string:
		err := json.Unmarshal([]byte(v), &result)
		if err != nil {
			log.Printf("failed to convert value ToArrayOfInt64: [%#v] with error: [%v]\n", v, err)
			return
		}

		return
	case []string:
		b := v
		for _, vv := range b {
			result = append(result, ToInt64(vv))
		}

		return
	case [][]byte:
		b := v
		for _, vv := range b {
			result = append(result, ToInt64(vv))
		}

		return
	default:
		return
	}
}
