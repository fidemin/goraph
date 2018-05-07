package goraph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNodeStack(t *testing.T) {
	assert := assert.New(t)
	s := NewNodeStack()
	s.Push(1)
	s.Push(3)
	s.Push(10)

	assert.Equal(10, s.Pop())
	assert.Equal(3, s.Pop())
	assert.Equal(1, s.Pop())

}
