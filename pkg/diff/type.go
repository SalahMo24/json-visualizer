package diff

import jsonmodule "json-visualizer/pkg/json-module"

type Difference struct {
	Old  interface{}
	New  interface{}
	Type ChangeType
}

type ChangeType int

const (
	Addition ChangeType = iota
	Deletion
	ValueChange
	KeyChange
)

type Differ struct {
	Differece []Difference
}

type IDiffer interface {
	Addition(newVal interface{})
	Deletion(deletedVal interface{})
	ValueChange(newVal, oldVal interface{})
	KeyChange(newVal, oldVal interface{})
	Diff(mapNew, mapOld jsonmodule.Input)
	GetDifference() []Difference
}

func (ct ChangeType) String() string {
	return [...]string{"Addition", "Deletion", "ValueChange", "KeyChange"}[ct]
}

func (ct ChangeType) EnumIndex() int {
	return int(ct)
}
