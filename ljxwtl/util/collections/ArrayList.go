package collections

import "errors"

type AbstractList struct {
	elements []interface{}
}

func (al *AbstractList) Size() int64 {
	if al.elements == nil {
		panic(errors.New("elements is nil"))
	}
	return int64(len(al.elements))
}

func (al *AbstractList) IsEmpty() bool {
	if al.elements == nil || (al.elements != nil && len(al.elements) <= 0) {
		return true
	}
	return false
}

func (al *AbstractList) Contains(object interface{}) bool {
	if object == nil {
		panic(errors.New("object is nil, cannot execute Contains(object interface{}) func"))
	}

	if al.IsEmpty() {
		return false
	}

	// 当元素不为空，且参数不为nil进行判断
	for _, val := range al.elements {
		if val == object {
			return true
		}
	}

	return false
}
