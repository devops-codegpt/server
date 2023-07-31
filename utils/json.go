package utils

import (
	"encoding/json"
	"fmt"
)

func Struct2Json(obj any) string {
	str, err := json.Marshal(obj)
	if err != nil {
		fmt.Printf("struct covert to json failed: %v", err)
	}
	return string(str)
}

func Json2Struct(str string, obj any) {
	if err := json.Unmarshal([]byte(str), obj); err != nil {
		fmt.Printf("json convert to struct failed: %v", err)
	}
}

// Struct2Struct  is an intermediate bridge, dist must be passed as a pointer
func Struct2Struct(src, dist any) {
	jsonStr := Struct2Json(src)
	Json2Struct(jsonStr, dist)
}
