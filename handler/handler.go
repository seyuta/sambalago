package handler

import (
	"math"
	"reflect"
)

// Paginate function to handle pagination logic for any type
func Paginate[T any](items []T, page, pageSize int) (results []T, totalItems, totalPages int) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}

	totalItems = len(items)
	totalPages = int(math.Ceil(float64(totalItems) / float64(pageSize)))

	start := (page - 1) * pageSize
	end := start + pageSize

	if start > totalItems {
		return
	}

	if end > totalItems {
		end = totalItems
	}

	results = items[start:end]
	return
}

// FilterSliceCondition function to check whether a value satisfies a condition
type FilterSliceCondition func(value any) bool

// FilterSliceOfStruct function to filter slice of struct based on conditions (FilterSliceCondition)
// Example:
//
//	conditions := map[string]FilterSliceCondition{
//		"FieldName": func(value any) bool {
//			val, ok := value.(string)
//			return ok && strings.Contains(val, "a")
//		},
//	}
//
// FieldName = field name from struct
func FilterSliceOfStruct[T any](slice []T, conditions map[string]FilterSliceCondition) []T {
	filteredSlice := make([]T, 0)
	for _, elem := range slice {
		match := true
		elemValue := reflect.ValueOf(elem)
		for field, condition := range conditions {
			fieldVal := elemValue.FieldByName(field)
			if !fieldVal.IsValid() {
				match = false
				break
			}
			if !condition(fieldVal.Interface()) {
				match = false
				break
			}
		}
		if match {
			filteredSlice = append(filteredSlice, elem)
		}
	}
	return filteredSlice
}
