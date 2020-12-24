package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var ErrNotFound = errors.New("key not found")

type Map map[string]interface{}

func (m Map) Get(key string) (interface{}, error) {
	var (
		next = m
		keys = strings.Split(key, ".")
	)
	for idx, key := range keys {
		value, ok := next[key]
		if !ok {
			return nil, ErrNotFound
		}
		if idx == len(keys)-1 {
			return value, nil
		}
		if next, ok = value.(map[string]interface{}); !ok {
			return nil, ErrNotFound
		}
	}
	return nil, ErrNotFound
}

func (m Map) Scan(key string, v interface{}) error {
	raw, err := m.Get(key)
	if err != nil {
		return err
	}
	b, err := json.Marshal(raw)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, v)
}

func main() {
	text := `
		{
			"test": {
				"settings" : {
					"int_key": 100,
					"float_key": 1000.1, 
					"string_key": "string_value"
				},
				"server": {
					"addr": "127.0.0.1",
					"port": 8000
				}
			}
		}
	`
	m := &Map{}
	if err := json.Unmarshal([]byte(text), &m); err != nil {
		panic(err)
	}
	var settings struct {
		IntKey    int     `json:"int_key"`
		FloatKey  float32 `json:"float_key"`
		StringKey string  `json:"string_key"`
	}
	fmt.Println("source:", m)
	fmt.Println("scan:", m.Scan("test.settings", &settings), settings)
	fmt.Println(m.Get("test.settings.key"))
	fmt.Println(m.Get("test.settings.int_key"))
	fmt.Println(m.Get("test.settings.float_key"))
	fmt.Println(m.Get("test.settings.string_key"))
	fmt.Println(m.Get("test.server.addr"))
	fmt.Println(m.Get("test.server.port"))
}
