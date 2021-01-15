package utils

import (
	"os"
	"reflect"

	"github.com/google/uuid"
)

type Reply int

type RegisterComponentArgs struct {
	ComponentType string
	ID            uuid.UUID
}

// RunningInDocker Returns if the code is running within a Docker container
func RunningInDocker() bool {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		return true
	}

	return false
}

// MatchStructureWithSchema Matches the data with the given schema
func MatchStructureWithSchema(data map[string]interface{}, schema map[string]interface{}) map[string]interface{} {
	var i interface{}

	for key, value := range data {
		if schema[key] == nil {
			delete(data, key)
		} else {
			if reflect.TypeOf(value) == reflect.MapOf(reflect.TypeOf("string"), reflect.TypeOf(&i).Elem()) &&
				reflect.TypeOf(schema[key]) == reflect.MapOf(reflect.TypeOf("string"), reflect.TypeOf(&i).Elem()) {
				data[key] = MatchStructureWithSchema(data[key].(map[string]interface{}), schema[key].(map[string]interface{}))
			}
		}
	}

	return data
}
