package diff

import jsonmodule "json-visualizer/pkg/json-module"

func NewDifference(newVal, oldVal interface{}, chnageType ChangeType) Difference {
	return Difference{
		New:  newVal,
		Old:  oldVal,
		Type: chnageType,
	}
}

func NewDiffer() IDiffer {
	return &Differ{Differece: make([]Difference, 0)}
}

func (d *Differ) GetDifference() []Difference {

	return d.Differece
}
func (d *Differ) Addition(newVal interface{}) {
	diff := NewDifference(newVal, nil, Addition)
	d.Differece = append(d.Differece, diff)
}

func (d *Differ) Deletion(deletedVal interface{}) {
	diff := NewDifference(nil, deletedVal, Deletion)
	d.Differece = append(d.Differece, diff)

}
func (d *Differ) ValueChange(newVal, oldVal interface{}) {
	diff := NewDifference(newVal, oldVal, ValueChange)
	d.Differece = append(d.Differece, diff)

}
func (d *Differ) KeyChange(newVal, oldVal interface{}) {
	diff := NewDifference(newVal, oldVal, KeyChange)
	d.Differece = append(d.Differece, diff)

}
func (d *Differ) Diff(mapNew, mapOld jsonmodule.Input) {
	if d.Differece == nil {
		panic("difference array should not be nil")
	}
	for k, v := range mapNew {
		vOld, ok := mapOld[k]

		if ok {
			switch vType := v.(type) {
			case jsonmodule.Input:
				{
					vOldTyped, _ := vOld.(map[string]interface{})
					d.Diff(vType, vOldTyped)
					delete(mapOld, k)

				}
			default:
				{
					if v != vOld {
						d.ValueChange(jsonmodule.Input{k: v}, jsonmodule.Input{k: vOld})
					}
					delete(mapOld, k)

				}
			}
		} else {
			d.Addition(jsonmodule.Input{k: v})
		}
	}

	for k, v := range mapOld {
		d.Deletion(jsonmodule.Input{k: v})
	}
}
