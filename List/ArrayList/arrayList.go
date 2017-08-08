package ArrayList

import (
	"fmt"
	"strings"
	"sync"

	"github.com/fighterlyt/ADT/List"
)

type arrayList struct {
	*sync.RWMutex
	data      []interface{}
	iteartors []*arrayListIterator
}

func New() List.List {
	return &arrayList{
		RWMutex: &sync.RWMutex{},
	}
}

func (a arrayList) getData() []interface{} {
	return a.data
}

func (a arrayList) IsEmpty() bool {
	return a.Size() == 0
}

func (a arrayList) Iterator() List.Iterator {
	a.RLock()
	defer a.RUnlock()

	return &arrayListIterator{
		RWMutex: a.RWMutex,
		index:   0,
		base:    &a,
	}
}

func (a *arrayList) Append(values ...interface{}) bool {
	a.Lock()
	defer a.Unlock()

	a.data = append(a.data, values...)
	return true
}

func (a *arrayList) Prepend(value ...interface{}) bool {
	panic("implement me")
}

func (a arrayList) Tail() List.List {
	a.RLock()
	defer a.RUnlock()

	if a.size() == 0 {
		return &arrayList{
			RWMutex: a.RWMutex,
		}
	} else {
		return &arrayList{
			RWMutex: a.RWMutex,
			data:    a.data[1:],
		}
	}
}

func (a arrayList) size() int {
	return len(a.data)
}
func (a arrayList) Size() int {
	a.RLock()
	defer a.RUnlock()

	return a.size()
}

func (a arrayList) String() string {
	a.RLock()
	defer a.RUnlock()

	result := make([]string, 0, len(a.data))

	for _, data := range a.data {
		result = append(result, fmt.Sprint(data))
	}
	return "[ " + strings.Join(result, ",") + " ]"
}

func (a arrayList) Clear() {
	a.Lock()
	defer a.Unlock()

	if len(a.data) != 0 {
		a.data = a.data[:0]
	}
}

type arrayListIterator struct {
	*sync.RWMutex
	base  *arrayList
	index int
}

func (a *arrayListIterator) Head() (interface{}, bool) {

	a.index = 0
	a.RLock()
	defer a.RUnlock()

	if a.base.size() == 0 {
		return nil, false
	} else {
		return a.getData()[0], true
	}
}

func (a arrayListIterator) End() (interface{}, bool) {
	a.RLock()
	defer a.RUnlock()

	if size := a.base.size(); size == 0 {
		a.index = 0
		return nil, false
	} else {
		a.index = size - 1
		return a.getData()[size-1], true
	}
}
func (a arrayListIterator) Next() (interface{}, bool) {
	a.RLock()
	defer a.RUnlock()

	size := a.base.size()

	if a.index < size-1 {
		a.index++
		return a.getData()[a.index], true
	} else {
		return nil, false
	}

}
func (a arrayListIterator) Prev() (interface{}, bool) {
	a.RLock()
	defer a.RUnlock()

	size := a.base.size()

	if a.index > 0 && a.index < size-2 {
		a.index--
		return a.getData()[a.index], true
	} else {
		return nil, false
	}
}

func (a *arrayListIterator) getData() []interface{} {
	return a.base.getData()
}
