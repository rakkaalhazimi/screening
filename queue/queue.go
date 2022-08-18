package queue

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type MyQueue []interface{}

func (q *MyQueue) Push(item interface{}) {
	*q = append(*q, item)
}

func (q MyQueue) Len() int {
	return len(q)
}

func (q MyQueue) Contains(key interface{}) bool {
	var found bool
	for _, value := range q {
		if value == key {
			found = true
		}
	}
	return found
}

func (q MyQueue) Keys() []interface{} {
	return q
}

func (q *MyQueue) Pop() interface{} {
	lastItem := (*q)[q.Len()-1]
	*q = (*q)[:q.Len()-1]
	return lastItem
}

func New(size int) Queue {
	my_queue := MyQueue{}

	return my_queue
}
