package converter

import (
	"encoding/json"
	"log"
)

// ToJson convert any value to json
func ToJson(data interface{}) (jsonData []byte) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to convert data ToJson: [%#v] with error: [%v]\n", data, err)
		return
	}

	return
}

// ToJsonString convert any value to json string
func ToJsonString(data interface{}) (jsonString string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("failed to convert data ToJsonString: [%#v] with error: [%v]\n", data, err)
		return
	}
	jsonString = string(jsonData)

	return
}

// JsonToMapStrInterface convert json to map[string]interface{}
func JsonToMapStrInterface(jsonData []byte) (result map[string]interface{}) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		log.Printf("failed to convert data JsonToMapStrInterface: [%#v] with error: [%v]\n", jsonData, err)
		return
	}

	return
}

// JsonToArrayOfInterface convert json to []interface{}
func JsonToArrayOfInterface(jsonData []byte) (result []interface{}) {
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		log.Printf("failed to convert data JsonToArrayOfInterface: [%#v] with error: [%v]\n", jsonData, err)
		return
	}

	return
}
