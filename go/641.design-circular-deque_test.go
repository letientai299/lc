package main

import "testing"

func TestMyCircularDeque(t *testing.T) {
	q := Constructor(4)
	t.Log(q.InsertFront(9))
	t.Log(q.DeleteLast())
	t.Log(q.GetRear())
}
