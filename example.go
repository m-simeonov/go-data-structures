package example

import (
	"fmt"
	"go-data-structures/structures"
)

func RunExampleStack() {
	stack := structures.NewLifoStack()
	//Push Element to LifoStack
	stack.Push("one")
	stack.Push("two")

	fmt.Println(stack.GetLength()) //will print `2`

	element, err := stack.Pop()
	if err == nil {
		fmt.Println(element) //will print `two`
	}
}

func RunExampleList() {
	list := structures.NewLinkedList()
	list.Add("one")
	list.Add("two")
	list.Add("three")

	data := list.ToSlice()
	fmt.Println(data) // will print [one two three]

	len := list.GetLength()
	fmt.Println(len) // will print 3

	head := list.GetHead()
	fmt.Println(head.Data) // will print `one`
	list.Delete(head)

	data = list.ToSlice()
	fmt.Println(data) // will print [two three]
}
