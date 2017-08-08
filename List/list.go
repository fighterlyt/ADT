package List

type List interface {
	IsEmpty() bool
	Append(value ...interface{}) bool
	Prepend(value ...interface{}) bool
	Iterator() Iterator

	Tail() List
	Size() int
	String() string
	Clear()
}

type Element interface {
	Less(another Element) int
}

type Iterator interface {
	Head() (interface{}, bool)
	End() (interface{}, bool)
	Next() (interface{}, bool)
	Prev() (interface{}, bool)
}
