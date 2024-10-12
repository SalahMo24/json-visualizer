package localizationparser

import (
	jsonmodule "json-visualizer/pkg/json-module"
	"reflect"
	"testing"
)

func TestMergeMaps(t *testing.T) {
	tests := []struct {
		name     string
		map1     jsonmodule.Input
		map2     jsonmodule.Input
		lang1    string
		lang2    string
		expected jsonmodule.Input
	}{
		{
			name: "Test simple merge",
			map1: jsonmodule.Input{
				"name":  "Apple",
				"price": 5000,
			},
			map2: jsonmodule.Input{
				"name":  "تفاح",
				"price": 5000,
			},
			lang1: "en",
			lang2: "ar",
			expected: jsonmodule.Input{
				"name":  jsonmodule.Input{"en": "Apple", "ar": "تفاح"},
				"price": jsonmodule.Input{"en": 5000, "ar": 5000},
			},
		},
		{
			name: "Test nested merge",
			map1: jsonmodule.Input{
				"description": jsonmodule.Input{
					"short": "A sweet fruit",
					"long":  "A sweet and crunchy fruit.",
				},
			},
			map2: jsonmodule.Input{
				"description": jsonmodule.Input{
					"short": "فاكهة حلوة",
					"long":  "فاكهة حلوة ومقرمشة.",
				},
			},
			lang1: "en",
			lang2: "ar",
			expected: jsonmodule.Input{
				"description": jsonmodule.Input{
					"short": jsonmodule.Input{"en": "A sweet fruit", "ar": "فاكهة حلوة"},
					"long":  jsonmodule.Input{"en": "A sweet and crunchy fruit.", "ar": "فاكهة حلوة ومقرمشة."},
				},
			},
		},
		{
			name: "Test merge with missing keys",
			map1: jsonmodule.Input{
				"name": "Apple",
			},
			map2: jsonmodule.Input{
				"description": jsonmodule.Input{
					"short": "فاكهة حلوة",
				},
			},
			lang1: "en",
			lang2: "ar",
			expected: jsonmodule.Input{
				"name": jsonmodule.Input{"en": "Apple", "ar": nil},
				"description": jsonmodule.Input{
					"short": jsonmodule.Input{"en": nil, "ar": "فاكهة حلوة"},
				},
			},
		},
		{
			name: "Test deeply nested merge",
			map1: jsonmodule.Input{
				"product": jsonmodule.Input{
					"info": jsonmodule.Input{
						"details": jsonmodule.Input{
							"material": "Metal",
							"usage":    "Outdoor",
						},
						"origin": "USA",
					},
				},
			},
			map2: jsonmodule.Input{
				"product": jsonmodule.Input{
					"info": jsonmodule.Input{
						"details": jsonmodule.Input{
							"material": "معدن",
							"usage":    "خارجي",
						},
						"origin": "الولايات المتحدة الأمريكية",
					},
				},
			},
			lang1: "en",
			lang2: "ar",
			expected: jsonmodule.Input{
				"product": jsonmodule.Input{
					"info": jsonmodule.Input{
						"details": jsonmodule.Input{
							"material": jsonmodule.Input{"en": "Metal", "ar": "معدن"},
							"usage":    jsonmodule.Input{"en": "Outdoor", "ar": "خارجي"},
						},
						"origin": jsonmodule.Input{"en": "USA", "ar": "الولايات المتحدة الأمريكية"},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Merge(tt.map1, tt.map2, tt.lang1, tt.lang2)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
