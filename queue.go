package queue

import("fmt")

type Element struct {
	Value interface{}
	previous, next *Element
}

type Queue struct {
	Head, Tail *Element
	Size int
}

func (q *Queue) Init() *Queue {
	if q != nil {
		q.Head = nil
		q.Tail = nil
		q.Size = 0
	} else {
		q = new(Queue)
		q.Head = nil
		q.Tail = nil
		q.Size = 0
	}
	return q
}

func (q *Queue) Enqueue(v interface{}) {
	if q != nil {
		elem := new(Element)
		elem.Value = v
		q.Size++;
		if(q.Head == nil && q.Tail == nil) {
			elem.previous = nil
			elem.next = nil
			q.Head = elem
			q.Tail = elem
		} else {
			elem.previous = q.Tail
			elem.next = nil
			q.Tail.next = elem
			q.Tail = elem
		}
	}
}

func (q *Queue) Top() (val interface{}) {
	if q != nil {
		if q.Head != nil {
			return q.Head.Value
		}
		return nil
	}
	return nil
}

func (q *Queue) GetSize() int {
	if(q != nil) {
		if q.Head != nil && q.Tail != nil && q.Size != 0 {
			return q.Size
		}
		return 0
	}
	return -1
}