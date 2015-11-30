package ds

import (
	"errors"
)

type Queue struct {
	list []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (queue *Queue) enqueue(key int) int {
	queue.list = append(queue.list, key)

	return key
}

func (queue *Queue) dequeue() (int, error) {
	if queue.size() == 0 {
		return 0, errors.New("Queue is empty")
	}

	key := queue.list[0]
	queue.list = queue.list[1:]

	return key, nil
}

func (queue *Queue) size() int {
	return len(queue.list)
}
