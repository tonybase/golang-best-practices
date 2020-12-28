package main

import (
	"encoding/json"
	"fmt"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/ghodss/yaml"
)

type Map struct {
	raw *simplejson.Json
}

func NewMap(data []byte) (*Map, error) {
	raw, err := simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return &Map{
		raw: raw,
	}, nil
}

func (m *Map) Value(key string) *simplejson.Json {
	return m.raw.GetPath(strings.Split(key, ".")...)
}

func (m *Map) Scan(key string, v interface{}) error {
	data, err := m.Value(key).MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func main() {
	yml := `
test:
  settings:
    int_key: 100
    float_key: 1000.1
    string_key: string_value
  server:
    addr: 127.0.0.1
    port: 8000
`
	text, err := yaml.YAMLToJSON([]byte(yml))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(text))
	m, err := NewMap([]byte(text))
	if err != nil {
		panic(err)
	}
	var settings struct {
		IntKey    int     `json:"int_key"`
		FloatKey  float32 `json:"float_key"`
		StringKey string  `json:"string_key"`
	}
	fmt.Println("source:", m)
	fmt.Println(m.Scan("test.settings", &settings), settings)
	fmt.Println(m.Value("test.settings.int_key"))
	fmt.Println(m.Value("test.settings.float_key"))
	fmt.Println(m.Value("test.settings.string_key"))
	fmt.Println(m.Value("test.server.addr"))
	fmt.Println(m.Value("test.server.port"))
}
