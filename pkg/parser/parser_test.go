package parser

import (
	"testing"
)

func TestEntryParser(t *testing.T) {
	var tests = []struct {
		name  string
		input []Entry
		want  map[string]interface{}
	}{
		{"parse one value", []Entry{{Key: "key1", Val: "val1"}}, map[string]interface{}{"key1": "val1"}},
		{"parse values", []Entry{{Key: "key1", Val: "val1"}, {Key: "key2", Val: "val2"}}, map[string]interface{}{"key1": "val1", "key2": "val2"}},
	}

	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pasrser := NewParser(tt.input)
			parsed, err := pasrser.EntryParser()
			if err != nil {
				t.Error(err)
				panic("test failed")
			}
			for index, parsedVals := range parsed {
				if parsedVals != tt.want[index] {
					t.Errorf("got %s, want %s", parsedVals, tt.want)
				}

			}
		})
	}
}
