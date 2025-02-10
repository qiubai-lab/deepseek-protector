package utils

type FixedQueue struct {
	data     []interface{}
	head     int
	tail     int
	size     int
	capacity int
}

// NewFixedQueue 创建一个固定长度的队列
func NewFixedQueue(capacity int) *FixedQueue {
	return &FixedQueue{
		data:     make([]interface{}, capacity),
		capacity: capacity,
	}
}

// Enqueue 向队列添加元素
func (q *FixedQueue) Enqueue(value interface{}) {
	if q.size == q.capacity { // 如果队列已满，移除最早的元素
		q.head = (q.head + 1) % q.capacity
	} else {
		q.size++
	}
	q.data[q.tail] = value
	q.tail = (q.tail + 1) % q.capacity
}

// Dump 获取队列状态
func (q *FixedQueue) Dump() []interface{} {
	result := make([]interface{}, 0, q.size)
	for i := 0; i < q.size; i++ {
		result = append(result, q.data[(q.head+i)%q.capacity])
	}
	return result
}

// GetLastN 获取最近的n个元素
func (q *FixedQueue) GetLastN(n int) []interface{} {
	if n > q.size || n < 0 { // 返回实际存在的元素
		n = q.size
	}
	result := make([]interface{}, n)
	for i := 0; i < n; i++ {
		result[n-i-1] = q.data[(q.tail-1-i+q.capacity)%q.capacity] // 从尾部开始取值
	}
	return result
}
