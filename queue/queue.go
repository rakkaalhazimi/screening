package queue


type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

func (q Queue) Push() {
	
}

func New(size int) Queue {
	my_queue := Queue{}

	return nil
}
