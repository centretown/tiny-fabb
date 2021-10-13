package forms

import (
	"fmt"

	"github.com/centretown/tiny-fabb/docs"
	"github.com/golang/glog"
)

// Entry -
type Entry struct {
	ID    string  `json:"entryId"`
	Name  string  `json:"name"`
	Code  string  `json:"code"`
	Label string  `json:"label"`
	URL   string  `json:"url"`
	Type  string  `json:"type"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Step  float32 `json:"step"`
	Item  uint    `json:"item"`
	Val   func(interface{}) interface{}
	Scan  func(string, interface{}) error
}

type Entries []*Entry

func (ent *Entry) Documents() docs.Docs {
	return Documents()
}

func (ent *Entry) FindDoc(code string) *docs.Doc {
	return findDoc(code)
}

func (ent *Entry) Value(v interface{}) (r interface{}) {
	if ent.Val != nil {
		r = ent.Val(v)
	} else {
		switch t := v.(type) {
		case *uint:
			r = *t
		case *float64:
			r = *t
		case *float32:
			r = *t
		case *bool:
			r = *t
		case *string:
			r = *t
		case string:
			r = t
		default:
			r = v
		}
	}
	return
}

func (ent *Entry) ResponseList(v interface{}) []string {
	l, ok := v.([]string)
	if ok && len(l) > 0 {
		return l
	}
	return []string{"empty"}
}

func (ent *Entry) Checked(v interface{}) (b bool) {
	switch e := v.(type) {
	case *bool:
		b = *e
	case bool:
		b = e
	case uint8:
		b = e != 0
	case *uint8:
		b = *e != 0
	}
	return
}

func (ent *Entry) InfoURL(base string) (s string) {
	return base + ent.URL
}

type InputFormat struct {
	Entry      *Entry
	Name       string
	Class      string
	Value      string
	HasChecked bool
	HasRange   bool
	HasStep    bool
	ReadOnly   bool
}

func (ent *Entry) FormatInput(value interface{}, first *Entry) (f *InputFormat) {
	eval := ent.Value(value)
	f = &InputFormat{
		Name:  ent.ID,
		Class: "w3-input",
		Entry: ent,
		Value: fmt.Sprint(eval),
	}
	switch ent.Type {
	case "text":
	case "checkbox":
		f.Class = "w3-check"
		f.HasChecked = ent.Checked(eval)
	case "radio":
		f.Class = "w3-" + ent.Type
		f.HasChecked = ent.Checked(eval)
		f.Name = first.ID
		f.Value = fmt.Sprint(ent.Item)
	case "number":
		f.HasRange = ent.Min != ent.Max
		f.HasStep = ent.Step != 0
	case "readonly":
		f.ReadOnly = true
	}
	return
}

func (ent *Entry) ScanInput(s string, value interface{}) (err error) {
	if ent.Scan != nil {
		err = ent.Scan(s, value)
	} else {
		t, ok := value.(string)
		if ok {
			value = t
			return
		}

		_, ok = value.([]string)
		if ok {
			return
		}

		_, err = fmt.Sscan(s, value)
		if err != nil {
			glog.Warningln(ent.Code, err)
		}
	}
	return
}

func Mask(m interface{}, mask uint) (b bool) {
	pVal, ok := m.(*uint)
	if ok {
		b = (*pVal&mask != 0)
	}
	return
}

func UnMasks(m interface{}, mask uint, s string) (err error) {
	return UnMask(m, mask, s == "true")
}

func UnMask(m interface{}, mask uint, isTrue bool) (err error) {
	v, ok := m.(*uint)
	if !ok {
		err = fmt.Errorf("unmask wrong type %v", m)
		return
	}
	if isTrue {
		*v |= mask
	} else {
		*v &= ^mask
	}
	return
}

func (ent *Entry) InBounds(s string) bool {
	var val float32
	if ent.Min >= ent.Max {
		return true
	}
	fmt.Sscan(s, &val)
	return val >= ent.Min && val <= ent.Max
}
