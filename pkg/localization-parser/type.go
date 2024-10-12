package localizationparser

import jsonmodule "json-visualizer/pkg/json-module"

type LocalizationParser interface {
	Merge(map1, map2 jsonmodule.Input) jsonmodule.Input
}
