package main

import "fmt"

/*队列的一系列操作*/
type QueueNode struct {
	PtrReq *Req
	Next   *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func QueueInit() (q *Queue) {
	q = new(Queue)
	q.head = new(QueueNode)
	q.head.Next = nil
	q.tail = q.head
	return
}

func (q *Queue) IsEmpty() bool {
	if q.head.Next == nil {
		return true
	}
	return false
}

func (q *Queue) PushQueueNode(data *QueueNode) {
	q.tail.Next = data
	q.tail = data
}

func (q *Queue) PopQueueNode() (popnode *QueueNode) {
	if !q.IsEmpty() {
		if q.head.Next == q.tail {
			q.tail = q.head
		}
		popnode = q.head.Next
		q.head.Next = q.head.Next.Next
		return
	}
	return nil
}

func (q *Queue) TravalQueue() {
	if q.IsEmpty() {
		fmt.Println("Null Queue!")
		return
	}
	qhead := q.head.Next
	for qhead != nil {
		fmt.Printf("%d ", qhead.PtrReq.Hoster.Grade)
		qhead = qhead.Next
	}
	fmt.Println()
}

func (q *Queue) DeleteAll() {
	q.head.Next = nil
}
