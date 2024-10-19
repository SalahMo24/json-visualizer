package handler

import (
	"encoding/json"
	"fmt"
	"json-visualizer/pkg/diff"
	jsonmodule "json-visualizer/pkg/json-module"
	localizationparser "json-visualizer/pkg/localization-parser"
	"json-visualizer/pkg/views/user"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func readFile(fileName string) jsonmodule.Input {
	path, err := filepath.Abs(fileName)
	if err != nil {
		panic(err)
	}

	file, err := jsonmodule.NewFile(path)
	if err != nil {
		panic(err)
	}
	start := time.Now()
	entries, err := file.Reader()
	if err != nil {
		panic(err)
	}
	elapsed := time.Since(start)
	fmt.Println("Time elapsed: ", elapsed)
	return entries
}
func (h UserHandler) HandleUser(c echo.Context) error {

	enEntries := readFile("en.json")
	arEntries := readFile("ar.json")
	merged := localizationparser.Merge(enEntries, arEntries, "en", "ar")
	// fmt.Println(merged)

	return render(c, user.Show(merged))
}

func compare(newMap jsonmodule.Input) []diff.Difference {
	enEntries := readFile("en.json")
	arEntries := readFile("ar.json")
	old := localizationparser.Merge(enEntries, arEntries, "en", "ar")

	// println(reflect.ValueOf(newMap).Kind())
	// println(reflect.ValueOf(old).Kind())
	differ := diff.NewDiffer()

	differ.Diff(newMap, old)
	return differ.GetDifference()

}
func parseFormData(form *multipart.Form) map[interface{}]interface{} {
	res := make(map[interface{}]interface{})

	for key, values := range form.Value {
		if len(values) == 1 {
			// Check if the value is JSON-like and try to parse it
			var nestedMap map[string]interface{} // Using map[string]interface{} for nested structures
			if err := json.Unmarshal([]byte(values[0]), &nestedMap); err == nil {
				// If successful, store as a nested map
				res[key] = nestedMap
			} else {
				// Otherwise, treat it as a regular string
				res[key] = values[0]
			}
		} else {
			// Multiple values, just store them as they are
			res[key] = values
		}
	}

	return res
}

// ConvertToMap recursively converts map[interface{}]interface{} into map[string]interface{}
func ConvertToMap(input map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range input {
		// Convert the key to string (JSON keys must be strings)
		key := fmt.Sprintf("%v", k)

		switch value := v.(type) {
		case map[interface{}]interface{}:
			// Recursively process nested maps
			result[key] = ConvertToMap(value)
		case map[string]interface{}:
			// If it's already a map[string]interface{}, process it directly
			result[key] = value
		default:
			// For any other types, just set the value
			result[key] = value
		}
	}
	return result
}

// Your handler that processes the form data
func (h UserHandler) HandleUpdate(c echo.Context) error {
	// Parse the multipart form from the request
	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}

	convertedMap := ConvertToMap(parseFormData(form))

	diff := compare(convertedMap)

	fmt.Println(diff)
	// Respond to the client
	return c.String(200, "Form data processed and diffed")
}
