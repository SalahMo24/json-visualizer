package jsonmodule

import (
	"encoding/json"
	"io"
	"os"
)

func (file File) Reader() (map[string]interface{}, error) {
	jsonFile, err := os.Open(file.Name)

	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var res map[string]interface{}

	json.Unmarshal(byteValue, &res)

	return res, nil

}

// func (file File) Reader() (map[string]interface{}, error) {
// 	// Open the file
// 	jsonFile, err := os.Open(file.Name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer jsonFile.Close()

// 	// Use a bufio.Scanner to read the file line by line
// 	scanner := bufio.NewScanner(jsonFile)

// 	// String builder to accumulate the entire JSON content
// 	var content string

// 	// Read file line by line and append each line to the content
// 	for scanner.Scan() {
// 		content += scanner.Text()
// 	}

// 	// Check for any errors encountered during reading lines
// 	if err := scanner.Err(); err != nil {
// 		return nil, err
// 	}

// 	// Unmarshal the accumulated JSON content into a map
// 	var res map[string]interface{}
// 	err = json.Unmarshal([]byte(content), &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Return the result
// 	return res, nil
// }

func (file File) Append(newInputs Input) (Input, error) {
	existingData, err := file.Reader()
	var merged = make(Input)
	if err != nil {
		return nil, err
	}
	for key, val := range existingData {
		merged[key] = val
	}
	for key, val := range newInputs {
		merged[key] = val
	}

	return merged, nil

}
func (file File) Writer(inputs map[string]interface{}) error {
	newVals, err := file.Append(inputs)

	if err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(newVals, "", " ")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(file.Name, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func NewFile(name string) (*File, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	stats, err := jsonFile.Stat()

	if err != nil {
		return nil, err
	}
	return &File{Name: name, Stats: stats}, nil
}
