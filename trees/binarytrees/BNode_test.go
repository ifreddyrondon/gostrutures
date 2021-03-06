package binarytrees_test

import (
	"testing"

	"github.com/ifreddyrondon/gostrutures/trees/binarytrees"
)

func TestNewNode(t *testing.T) {
	var value int
	node := binarytrees.NewBNode(value)

	if node.Value != value {
		t.Errorf("Expected NewBNode value to be '%v'. Got '%v'", value, node.Value)
	}

	if node.Left != nil {
		t.Errorf("Expected NewBNode Left to be nil. Got '%v'", node.Left)
	}

	if node.Right != nil {
		t.Errorf("Expected NewBNode Left to be nil. Got '%v'", node.Left)
	}
}
