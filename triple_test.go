package rdf2go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var one = NewTriple(NewResource("a"), NewResource("b"), NewResource("c"))

func TestTripleEquals(t *testing.T) {
	assert.True(t, one.Equal(NewTriple(NewResource("a"), NewResource("b"), NewResource("c"))))
}

func TestTripleString(t *testing.T) {
	assert.Equal(t, "<a> <b> <c> .", one.String())
}

func TestBlankNodeTripleString(t *testing.T) {
	one := NewTriple(NewBlankNode("10"), NewResource("b"), NewResource("c"))
	assert.Equal(t, "_:10 <b> <c> .", one.String())
}
