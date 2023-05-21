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
