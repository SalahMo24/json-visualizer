package parser

type Entry struct {
	Key string
	Val string
}

type ParserInterface interface {
	EntryParser() (map[string]interface{}, error)
	IsEmpty() bool
}

type Parser struct {
	Entries []Entry
}
