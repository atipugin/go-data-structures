package ds

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	if NewQueue() == nil {
		t.Fail()
	}
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue()
	queue.enqueue(1)
	if queue.size() == 0 {
		t.Fail()
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueue()
	queue.enqueue(1)
	key, _ := queue.dequeue()
	if key != 1  {
		t.Fail()
	}

	_, err := queue.dequeue()
	if err == nil {
		t.Fail()
	}
}
