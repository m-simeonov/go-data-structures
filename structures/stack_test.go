package structures_test

import (
	"data_structures/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	t.Parallel()
	stack := structures.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

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
