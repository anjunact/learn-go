package stack

import "testing"

func TestStack_Pop(t *testing.T) {
	c:= new (Stack)
	c.Pop()
	if c.Pop() !=5{
		t.Log("pop desn't give 5")
		t.Fail()
	}
}