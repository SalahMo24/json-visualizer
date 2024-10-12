package parser

import (
	"errors"
)

func (parser Parser) EntryParser() (map[string]interface{}, error) {

	if parser.IsEmpty() {
		return nil, errors.New("empty array provided")
	}
	vals := make(map[string]interface{}, len(parser.Entries))

	for _, entry := range parser.Entries {
		vals[entry.Key] = entry.Val

	}

	return vals, nil
}

func (parser Parser) IsEmpty() bool {

	return len(parser.Entries) == 0
}

func NewParser(entries []Entry) *Parser {
	return &Parser{Entries: entries}
}
