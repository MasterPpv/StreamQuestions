package main

type Element struct {
	Value interface{}
	previous, next *Element
}

type Queue struct {
	head, tail *Element
	Size int
}

func (q *Queue) Init() *Queue {
	if q != nil {
		q.head = nil
		q.tail = nil
		q.Size = 0
	} else {
		q = new(Queue)
		q.head = nil
		q.tail = nil
		q.Size = 0
	}
	return q
}

func (q *Queue) Enqueue(v interface{}) {
	if q != nil {
		elem := new(Element)
		elem.Value = v
		q.Size++;
		if(q.head == nil && q.tail == nil) {
			elem.previous = nil
			elem.next = nil
			q.head = elem
			q.tail = elem
		} else {
			elem.previous = q.tail
			elem.next = nil
			q.tail.next = elem
			q.tail = elem
		}
	}
}

func (q *Queue) Top() (val interface{}) {
	if q != nil {
		if q.head != nil {
			return q.head.Value
		}
		return nil
	}
	return nil
}

func (q *Queue) GetSize() int {
	if q != nil {
		if q.head != nil && q.tail != nil && q.Size != 0 {
			return q.Size
		}
		return 0
	}
	return -1
}

func (q *Queue) IsEmpty() bool {
	if q != nil {
		if q.head != nil && q.tail != nil && q.Size != 0 {
			return false
		}
		return true
	}
	return true
}

func (q *Queue) Dequeue() (val interface{}) {
	if q != nil {
		if q.head != nil && q.tail != nil && q.Size != 0 {
			topval := q.head.Value
			nextup := q.head.next
			if nextup == nil || q.head == q.tail || q.Size == 1 {
				q.head = nil
				q.tail = nil
				q.Size = 0
			} else {
				q.head.next = nil
				q.head = nextup
				nextup.previous = nil
			}
			q.Size--
			return topval
		}
		return nil
	}
	return nil
}