package ds

import "fmt"

type Queue struct {
	vals []int
}

func (q *Queue) Enqueue(v int) {
	q.vals = append(q.vals, v)
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.vals) == 0 {
		return -1, fmt.Errorf("Queue is empty.")
	}

	v := q.vals[0]
	q.vals = q.vals[1:]

	return v, nil
}
