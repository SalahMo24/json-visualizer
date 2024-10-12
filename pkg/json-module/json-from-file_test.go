package jsonmodule

import (
	"testing"
)

func TestReader(t *testing.T) {
	file, err := NewFile("../../sample.json")
	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		name string
		want string
	}{
		// the table itself
		{"read first value in file", "val1"},
	}
	parsed, err := file.Reader()
	if err != nil {
		t.Error(err)
		panic("test failed")
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := parsed["key1"]
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
func TestWriter(t *testing.T) {
	file, err := NewFile("../../test.json")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		name  string
		input Input
		want  string
	}{
		// the table itself
		{"write a value to a file", map[string]interface{}{"key1": "val11", "key2": "val2", "key3": map[string]string{"key4": "val4"}}, "val11"},
		{"append a value to a file", map[string]interface{}{"key5": "val5"}, "val5"},
	}

	err = file.Writer(tests[0].input)
	if err != nil {
		t.Error(err)
		panic("writing failed")
	}
	parsed, err := file.Reader()
	if err != nil {
		t.Error(err)
		panic("reader failed")
	}
	// The execution loop

	t.Run(tests[0].name, func(t *testing.T) {
		ans := parsed["key1"]
		if ans != tests[0].want {
			t.Errorf("got %s, want %s", ans, tests[0].want)
		}
	})
	err = file.Writer(tests[1].input)
	if err != nil {
		t.Error(err)
		panic("writing failed")
	}
	parsed, err = file.Reader()
	if err != nil {
		t.Error(err)
		panic("reader failed")
	}
	t.Run(tests[1].name, func(t *testing.T) {
		ans := parsed["key5"]
		if ans != tests[1].want {
			t.Errorf("got %s, want %s", ans, tests[1].want)
		}
	})

}
func TestAppender(t *testing.T) {
	file, err := NewFile("../../test.json")

	if err != nil {
		t.Error(err)
	}

	var tests = []struct {
		name  string
		input Input
		want  Input
	}{
		// the table itself
		{"append one new value to file",
			map[string]interface{}{"key3": "val3"},
			map[string]interface{}{"key1": "val1", "key2": "val2", "key3": "val3"}},
	}

	res, err := file.Append(tests[0].input)
	if err != nil {
		t.Error(err)

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for index, parsedVals := range res {

				if parsedVals != tt.want[index] {
					t.Errorf("got %s, want %s", parsedVals, tt.want[index])
				}

			}
		})
	}

}
