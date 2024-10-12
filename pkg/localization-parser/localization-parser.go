package localizationparser

import (
	jsonmodule "json-visualizer/pkg/json-module"
)

func Merge(map1, map2 jsonmodule.Input, lang1, lang2 string) jsonmodule.Input {
	merged := make(jsonmodule.Input)

	for k, v := range map1 {
		v2, ok := map2[k]
		if !ok {
			v2 = nil
		}

		switch vType := v.(type) {
		case jsonmodule.Input:
			{
				v2Typed, ok := v2.(map[string]interface{})
				if !ok {
					v2Typed = nil
				}
				merged[k] = Merge(vType, v2Typed, lang1, lang2)
			}
		default:
			{
				merged[k] = jsonmodule.Input{
					lang1: v,
					lang2: v2,
				}
			}
		}
	}

	for k, v2 := range map2 {
		switch v2Type := v2.(type) {
		case jsonmodule.Input:
			{
				if _, ok := map1[k]; !ok {
					merged[k] = Merge(v2Type, nil, lang2, lang1)
				}
			}
		default:
			{
				if _, ok := map1[k]; !ok {
					merged[k] = jsonmodule.Input{
						lang1: nil,
						lang2: v2,
					}
				}
			}
		}
	}

	return merged

}
