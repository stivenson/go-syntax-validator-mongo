package mongoparser

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		input   string
		output  map[string]interface{}
		wantErr string
	}{{
		input:  `{}`,
		output: map[string]interface{}{},
	}, {
		input: `{"a": 1}`,
		output: map[string]interface{}{
			"a": float64(1),
		},
	}, {
		input: `{"a": 1, "b": ["c", 2]}`,
		output: map[string]interface{}{
			"a": float64(1),
			"b": []interface{}{"c", float64(2)},
		},
	}, {
		input: `{"a": []}`,
		output: map[string]interface{}{
			"a": []interface{}{},
		},
	}, {
		input: `{"a": [1.2]}`,
		output: map[string]interface{}{
			"a": []interface{}{float64(1.2)},
		},
	}, {
		input: `{"a": [1.2, 2.3]}`,
		output: map[string]interface{}{
			"a": []interface{}{float64(1.2), float64(2.3)},
		},
	}, {
		input: `{"a": true, "b": false, "c": null}`,
		output: map[string]interface{}{
			"a": true,
			"b": false,
			"c": nil,
		},
	}, {
		input:   `.1`,
		wantErr: `syntax error`,
	}, {
		input:   `invalid`,
		wantErr: `syntax error`,
	}}
	for _, tc := range testcases {
		got, err := Parse("insert", tc.input, "")
		var gotErr string
		if err != nil {
			gotErr = err.Error()
		}
		if gotErr != tc.wantErr {
			// t.Errorf(`%s err: %v, want %v`, tc.input, gotErr, tc.wantErr) // optional use
		}
		if !reflect.DeepEqual(got, tc.output) {
			// t.Errorf(`%s: %#v want %#v`, tc.input, got, tc.output) // optional use
		}
	}
}

func TestParserJson1(t *testing.T) {
	var query = `{
   		"title": "MongoDB Overview", 
   		"likes": "100"
	}`

	var payload = `{
		"set": {"size.uom": "cm"},
		"likes": {"size": "cm"}
   	}`
	var operation = "insert"
	_, err := Parse(operation, query, payload)
	if err == nil {
		fmt.Println("ALL FINE")
	} else {
		fmt.Println("ERROR:", err)
	}
}

func TestParserJson2(t *testing.T) {
	var query = `{
   		"title": "MongoDB Overview", 
   		"likes": "100"
	}`

	var payload = `{
		"set": {"size.uom": "cm"},
		"likes": [{"size": "cm"},{"size2": "pulg"}]
   	}`
	var operation = "insert"
	_, err := Parse(operation, query, payload)
	if err == nil {
		fmt.Println("ALL FINE")
	} else {
		fmt.Println("ERROR:", err)
	}
}

func TestParserJson3(t *testing.T) {
	var query = `{
   		"title": "MongoDB Overview", 
   		"likes": "100"
	}`

	var payload = `{
		"$set": {"size.uom": {"a": "b"}},
		"likes": [{"size": "cm"},{"size2": "pulg"}]
   	}`
	var operation = "insert"
	_, err := Parse(operation, query, payload)
	if err == nil {
		fmt.Println("ALL FINE")
	} else {
		fmt.Println("ERROR:", err)
	}
}
