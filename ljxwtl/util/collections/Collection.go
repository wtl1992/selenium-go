package collections

type Collection interface {
	Size() int64
	IsEmpty() bool
	Contains(object interface{}) bool
	Iterator() Iterator
	ToArray() []interface{}
	ADD(ele interface{}) bool
	Remove(ele interface{}) bool
	ContainsAll(collect []interface{}) bool
	AddAll(collect []interface{}) int
	RemoveAll() bool
	Clear() bool
	Equals() bool
	HashCode() int
	Stream() interface{}
	ParallelStream() interface{}
}
