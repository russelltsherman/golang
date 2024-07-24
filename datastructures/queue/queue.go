package queue

// Queue structure
type Queue struct {
	items []int
}

// Enqueue add an item to the queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue remove an item from the queue
func (q *Queue) Dequeue() int {
	if len(q.items) == 0 {
		return -1
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Queue structure using go channel
type QueueC struct {
	items chan int
}

// Enqueue add an item to the queue
func (q *QueueC) Enqueue(i int) {
	q.items <- i
}

// Dequeue remove an item from the queue
func (q *QueueC) Dequeue() int {
	// no provision for empty channel
	return <-q.items
}
