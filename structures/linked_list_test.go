package structures_test

import (
	"github.com/stretchr/testify/assert"
	"go-data-structures/structures"
	"testing"
)

func TestAddLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()

	list.Add("one")
	assert.Equal(t, list.GetLength(), 1)
	assert.Equal(t, []interface{}{"one"}, list.ToSlice())

	list.Add("two")
	assert.Equal(t, list.GetLength(), 2)
	assert.Equal(t, []interface{}{"one", "two"}, list.ToSlice())

	node := list.GetHead()
	for i := 0; i < list.GetLength(); i++ {
		if i == 0 {
			assert.Equal(t, node.Data, "one")
		} else if i == 1 {
			assert.Equal(t, node.Data, "two")
		}
		node = node.Next
	}

	list.Add("three")
	assert.Equal(t, list.GetLength(), 3)
	assert.Equal(t, []interface{}{"one", "two", "three"}, list.ToSlice())
}

func TestDeleteHeadLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add("one")
	list.Add("two")
	list.Add("three")
	list.Add("four")
	list.Add("five")

	head := list.GetHead()
	list.Delete(head)

	assert.Equal(t, list.ToSlice(), []interface{}{"two", "three", "four", "five"})
}

func TestDeleteElementLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add("one")
	list.Add("two")
	list.Add("three")
	list.Add("four")
	list.Add("five")

	head := list.GetHead()
	list.Delete(head.Next)

	assert.Equal(t, list.ToSlice(), []interface{}{"one", "three", "four", "five"})
}

func TestDeleteTailLinkedList(t *testing.T) {
	t.Parallel()
	list := structures.NewLinkedList()
	list.Add("one")
	list.Add("two")
	list.Add("three")
	list.Add("four")
	list.Add("five")

	tail := list.GetTail()
	list.Delete(tail)

	assert.Equal(t, list.ToSlice(), []interface{}{"one", "two", "three", "four"})
}
