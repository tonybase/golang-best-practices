package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Map map[string]interface{}

func (m Map) Get(key string) (interface{}, error) {
	var (
		keys = strings.Split(key, ".")
		next = m
	)
	for _, k := range keys {
		v, ok := next[k].(map[string]interface{})
		if ok {
			next = v
		} else {
			if raw, ok := next[k]; ok {
				return raw, nil
			}
		}
	}
	return nil, fmt.Errorf("key not found: %s", key)
}

func main() {
	text := `
		{
			"struct": {
				"string_key": "string_value",
				"int_key": 100,
				"float_key": 1000.1
			}
		}
	`
	m := &Map{}
	if err := json.Unmarshal([]byte(text), &m); err != nil {
		panic(err)
	}
	fmt.Println(m)
	fmt.Println(m.Get("struct.string_key"))
	fmt.Println(m.Get("struct.int_key"))
	fmt.Println(m.Get("struct.float_key"))
	fmt.Println(m.Get("struct.key"))
}
