package stack_test

import (
	"github.com/stretchr/testify/assert"
	"go-data-structures/stack"
	"testing"
)

func TestPushStack(t *testing.T) {
	t.Parallel()
	stack := stack.New()
	stack.Push("one")
	stack.Push("two")

	assert.Equal(t, stack.GetLength(), 2)
}

func TestPopStack(t *testing.T) {
	t.Parallel()
	stack := stack.New()
	stack.Push("one")
	stack.Push("two")

	elem, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, elem, "two")
}

func TestStackIntegration(t *testing.T) {
	t.Parallel()
	stack := stack.New()
	stack.Push(1, 2, 3)

	assert.Equal(t, stack.GetLength(), 3)

	elem, err := stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, elem, 3)

	elem, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, elem, 2)

	elem, err = stack.Pop()
	assert.Nil(t, err)
	assert.Equal(t, elem, 1)

	_, err = stack.Pop()
	assert.NotNil(t, err)
}
