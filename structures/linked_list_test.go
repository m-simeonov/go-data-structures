package structures_test

import (
	"data_structures/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()

	list.Add(&structures.Node{
		Data: "one", Next: nil,
	})
	assert.Equal(t, list.GetLength(), 1)
	assert.Equal(t, []string{"one"}, list.ToSlice())

	list.Add(&structures.Node{
		Data: "two", Next: nil,
	})
	assert.Equal(t, list.GetLength(), 2)
	assert.Equal(t, []string{"one", "two"}, list.ToSlice())

	node := list.GetHead()
	for i := 0; i < list.GetLength(); i++ {
		if i == 0 {
			assert.Equal(t, node.Data, "one")
		} else if i == 1 {
			assert.Equal(t, node.Data, "two")
		}
		node = node.Next
	}

	list.Add(&structures.Node{
		Data: "three", Next: nil,
	})
	assert.Equal(t, list.GetLength(), 3)
	assert.Equal(t, []string{"one", "two", "three"}, list.ToSlice())
}

func TestDeleteHeadLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add(&structures.Node{Data: "one"})
	list.Add(&structures.Node{Data: "two"})
	list.Add(&structures.Node{Data: "three"})
	list.Add(&structures.Node{Data: "four"})
	list.Add(&structures.Node{Data: "five"})

	head := list.GetHead()
	list.Delete(head)

	assert.Equal(t, list.ToSlice(), []string{"two", "three", "four", "five"})
}

func TestDeleteElementLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add(&structures.Node{Data: "one"})
	list.Add(&structures.Node{Data: "two"})
	list.Add(&structures.Node{Data: "three"})
	list.Add(&structures.Node{Data: "four"})
	list.Add(&structures.Node{Data: "five"})

	head := list.GetHead()
	list.Delete(head.Next)

	assert.Equal(t, list.ToSlice(), []string{"one", "three", "four", "five"})
}

func TestDeleteTailLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add(&structures.Node{Data: "one"})
	list.Add(&structures.Node{Data: "two"})
	list.Add(&structures.Node{Data: "three"})
	list.Add(&structures.Node{Data: "four"})
	list.Add(&structures.Node{Data: "five"})

	tail := list.GetTail()
	list.Delete(tail)

	assert.Equal(t, list.ToSlice(), []string{"one", "two", "three", "four"})
}
