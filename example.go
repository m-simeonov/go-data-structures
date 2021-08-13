package example

import (
	"fmt"
	"go-data-structures/list"
	"go-data-structures/stack"
)

func RunExampleStack() {
	lifoStack := stack.New()
	// Push Element to LifoStack
	lifoStack.Push("one")
	lifoStack.Push("two")

	fmt.Println(lifoStack.GetLength()) // will print `2`

	element, err := lifoStack.Pop()
	if err == nil {
		fmt.Println(element) // will print `two`
	}
}

func RunExampleList() {
	linkedList := list.New()
	linkedList.Add("one")
	linkedList.Add("two")
	linkedList.Add("three")

	data := linkedList.ToSlice()
	fmt.Println(data) // will print [one two three]

	len := linkedList.GetLength()
	fmt.Println(len) // will print 3

	head := linkedList.GetHead()
	fmt.Println(head.Data) // will print `one`
	linkedList.Delete(head)

	data = linkedList.ToSlice()
	fmt.Println(data) // will print [two three]
}
