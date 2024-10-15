package diff

import (
	jsonmodule "json-visualizer/pkg/json-module"
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		mapNew   jsonmodule.Input
		mapOld   jsonmodule.Input
		expected []Difference
	}{
		{
			name: "Basic test with additions, deletions, and changes",
			mapNew: jsonmodule.Input{
				"key1": "value1",      // unchanged
				"key2": "new_value",   // changed value
				"key3": "added_value", // added key
			},
			mapOld: jsonmodule.Input{
				"key1": "value1",        // unchanged
				"key2": "old_value",     // old value
				"key4": "deleted_value", // deleted key
			},
			expected: []Difference{
				{New: jsonmodule.Input{"key2": "new_value"}, Old: jsonmodule.Input{"key2": "old_value"}, Type: ValueChange},
				{New: jsonmodule.Input{"key3": "added_value"}, Old: nil, Type: Addition},
				{New: nil, Old: jsonmodule.Input{"key4": "deleted_value"}, Type: Deletion},
			},
		},
		{
			name: "Nested map with changes",
			mapNew: jsonmodule.Input{
				"key1": "value1", // unchanged
				"key2": jsonmodule.Input{
					"nested_key1": "nested_value1",    // unchanged
					"nested_key2": "new_nested_value", // changed value
					"nested_key3": "new_nested_added", // new nested key
				},
			},
			mapOld: jsonmodule.Input{
				"key1": "value1", // unchanged
				"key2": jsonmodule.Input{
					"nested_key1": "nested_value1",    // unchanged
					"nested_key2": "old_nested_value", // old nested value
					"nested_key4": "nested_deleted",   // deleted nested key
				},
			},
			expected: []Difference{
				{New: jsonmodule.Input{"nested_key2": "new_nested_value"}, Old: jsonmodule.Input{"nested_key2": "old_nested_value"}, Type: ValueChange},
				{New: jsonmodule.Input{"nested_key3": "new_nested_added"}, Old: nil, Type: Addition},
				{New: nil, Old: jsonmodule.Input{"nested_key4": "nested_deleted"}, Type: Deletion},
			},
		},
		{
			name: "Nested map with no changes",
			mapNew: jsonmodule.Input{
				"key1": "value1",
				"key2": jsonmodule.Input{
					"nested_key1": "nested_value1",
					"nested_key2": "nested_value2",
				},
			},
			mapOld: jsonmodule.Input{
				"key1": "value1",
				"key2": jsonmodule.Input{
					"nested_key1": "nested_value1",
					"nested_key2": "nested_value2",
				},
			},
			expected: []Difference{}, // No differences expected
		},
		{
			name:   "All values deleted",
			mapNew: jsonmodule.Input{},
			mapOld: jsonmodule.Input{
				"key1": "value1",
				"key2": "value2",
			},
			expected: []Difference{
				{New: nil, Old: jsonmodule.Input{"key1": "value1"}, Type: Deletion},
				{New: nil, Old: jsonmodule.Input{"key2": "value2"}, Type: Deletion},
			},
		},
	}

	// Loop through all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the Differ struct
			differ := NewDiffer()

			// Run the Diff function
			differ.Diff(tt.mapNew, tt.mapOld)
			if !reflect.DeepEqual(differ.GetDifference(), tt.expected) {
				t.Errorf("got %v, want %v", differ.GetDifference(), tt.expected)
			}

		})
	}
}
