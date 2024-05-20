package extractor

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

// FromMapStrInterface extract value from map[string]interface{},
// path for define nested json: key, key.key, key.key.key, and so on
// example path: foo for single key or without nested, foo.bar for 1 nested, foo.bar.baz for 2 nested, and so on
func FromMapStrInterface(data interface{}, path string) (result interface{}) {
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

// CodeLocation just call this function and you will get func name, file path, and line number of code
func CodeLocation() string {
	var pc uintptr
	var ok bool
	pc, origFile, origLine, ok := runtime.Caller(1)
	if !ok {
		return ""
	}

	// get func name
	fn := runtime.FuncForPC(pc)
	name := fn.Name()
	ix := strings.LastIndex(name, ".")
	if ix > 0 && (ix+1) < len(name) {
		name = name[ix+1:]
	}
	funcName := name

	// get file path
	nd, nf := filepath.Split(origFile)
	fileName := filepath.Join(filepath.Base(nd), nf)

	return fmt.Sprintf(
		"inside function %s() | line: %v:%v",
		funcName,
		fileName,
		origLine,
	)
}
