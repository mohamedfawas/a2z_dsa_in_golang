package queueusingarray

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(element int) {
	q.items = append(q.items, element)
}

func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	}

	element := q.items[0]
	q.items = q.items[1:]
	return element, true
}
