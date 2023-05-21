package extractor

import (
	"log"
	"strings"
)

// FromMap extract value from map[string]interface{},
// path for define nested json: key, key.key, key.key.key, and so on
// example path: foo for single key or without nested, foo.bar for 1 nested, foo.bar.baz for 2 nested, and so on
func FromMap(data interface{}, path string) (result interface{}) {
	if data == nil {
		log.Println("data is nil")
		return
	}

	// check if interface is map string
	data, isMap := data.(map[string]interface{})
	if !isMap {
		log.Printf("data [%#v] is not map[string]interface{}\n", data)
		return
	}

	keys := strings.Split(path, ".")
	nested := data.(map[string]interface{})
	for _, key := range keys {
		value, ok := nested[key]
		if !ok {
			log.Printf("key [%#v] not found\n", key)
			return
		}
		if nestedMap, ok := value.(map[string]interface{}); ok {
			nested = nestedMap
		} else {
			return value
		}
	}

	return
}
