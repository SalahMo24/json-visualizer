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

func TestDiffWithArabicAndEnglishValues(t *testing.T) {
	// Initialize a Differ struct with an empty Differece slice
	d := NewDiffer()

	// Define the new and old maps with both Arabic and English values
	mapNew := jsonmodule.Input{
		"signUp_screen_details_city": map[string]string{
			"ar": "المدينة الجديدة", // New city in Arabic
			"en": "New City",        // New city in English
		},
		"welcome_message": "Welcome to our platform", // A new key added in mapNew
	}

	mapOld := jsonmodule.Input{
		"signUp_screen_details_city": map[string]string{
			"ar": "المدينة",  // Old city in Arabic
			"en": "Old City", // Old city in English
		},
		"goodbye_message": "Thank you for visiting", // A key that was deleted in mapNew
	}

	// Call the Diff function to compare mapNew and mapOld
	d.Diff(mapNew, mapOld)

	// Define the expected differences
	expectedDiffs := []Difference{
		{
			New: map[string]interface{}{
				"signUp_screen_details_city": map[string]string{
					"ar": "المدينة الجديدة",
					"en": "New City",
				},
			},
			Old: map[string]interface{}{
				"signUp_screen_details_city": map[string]string{
					"ar": "المدينة",
					"en": "Old City",
				},
			},
			Type: ValueChange,
		},
		{
			New: map[string]interface{}{
				"welcome_message": "Welcome to our platform",
			},
			Old:  nil,
			Type: Addition,
		},
		{
			New: nil,
			Old: map[string]interface{}{
				"goodbye_message": "Thank you for visiting",
			},
			Type: Deletion,
		},
	}

	// Compare the actual differences with the expected differences
	if !reflect.DeepEqual(d.GetDifference(), expectedDiffs) {
		t.Errorf("\ngot\n %v, \nwant\n %v", d.GetDifference(), expectedDiffs)
	}
}

func TestDiffWithNestedMaps(t *testing.T) {
	tests := []struct {
		name     string
		mapNew   jsonmodule.Input
		mapOld   jsonmodule.Input
		expected []Difference
	}{
		{
			name: "ValueChange in nested map",
			mapNew: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
				},
			},
			mapOld: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة قديمة",
						"en": "Old City",
					},
				},
			},
			expected: []Difference{
				{
					New: map[string]interface{}{
						"city": map[string]string{
							"ar": "مدينة جديدة",
							"en": "New City",
						},
					},
					Old: map[string]interface{}{
						"city": map[string]string{
							"ar": "مدينة قديمة",
							"en": "Old City",
						},
					},
					Type: ValueChange,
				},
			},
		},
		{
			name: "Addition in nested map",
			mapNew: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
					"population": "1 million",
				},
			},
			mapOld: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
				},
			},
			expected: []Difference{
				{
					New: map[string]interface{}{
						"population": "1 million",
					},
					Old:  nil,
					Type: Addition,
				},
			},
		},
		{
			name: "Deletion in nested map",
			mapNew: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
				},
			},
			mapOld: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
					"population": "1 million",
				},
			},
			expected: []Difference{
				{
					New: nil,
					Old: map[string]interface{}{
						"population": "1 million",
					},
					Type: Deletion,
				},
			},
		},
		{
			name: "Complex nested map changes",
			mapNew: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة جديدة",
						"en": "New City",
					},
					"population": "2 million",
				},
				"greeting": "Welcome",
			},
			mapOld: jsonmodule.Input{
				"details": map[string]interface{}{
					"city": map[string]string{
						"ar": "مدينة قديمة",
						"en": "Old City",
					},
					"population": "1 million",
				},
				"farewell": "Goodbye",
			},
			expected: []Difference{
				{
					New: map[string]interface{}{
						"city": map[string]string{
							"ar": "مدينة جديدة",
							"en": "New City",
						},
					},
					Old: map[string]interface{}{
						"city": map[string]string{
							"ar": "مدينة قديمة",
							"en": "Old City",
						},
					},
					Type: ValueChange,
				},
				{
					New: map[string]interface{}{
						"population": "2 million",
					},
					Old: map[string]interface{}{
						"population": "1 million",
					},
					Type: ValueChange,
				},
				{
					New: map[string]interface{}{
						"greeting": "Welcome",
					},
					Old:  nil,
					Type: Addition,
				},
				{
					New: nil,
					Old: map[string]interface{}{
						"farewell": "Goodbye",
					},
					Type: Deletion,
				},
			},
		},
	}

	// Iterate over all the test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Initialize the Differ instance
			d := &Differ{
				Differece: []Difference{},
			}

			// Call the Diff function with the current test's input maps
			d.Diff(tt.mapNew, tt.mapOld)

			// Compare the actual differences with the expected ones
			if !reflect.DeepEqual(d.Differece, tt.expected) {
				t.Errorf("Test %s failed:\n got:\n  %v\n want:\n %v", tt.name, d.Differece, tt.expected)
			}
		})
	}
}
