package rbtree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNodeInsert(t *testing.T) {
	for _, test := range []struct {
		name  string
		input []int
		in    []int
		pre   []int
		post  []int
	}{
		{
			name:  "simple",
			input: []int{10, 0, 20, 40, 100, 0},
			in:    []int{0, 10, 20, 40, 100},
			pre:   []int{},
			post:  []int{},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var root *Node
			for _, key := range test.input {
				root, _ = root.Insert(key)
			}
			testInOrder(t, root, test.in)
			testPreOrder(t, root, test.pre)
			testPostOrder(t, root, test.post)
		})
	}
}

func TestNodeRotate(t *testing.T) {
	for _, test := range []struct {
		name  string
		dir   rune
		input []int
		in    [2][]int
		pre   [2][]int
		post  [2][]int
	}{
		{
			name:  "left",
			dir:   'l',
			input: []int{3, 2, 5, 4, 7},
			in: [2][]int{
				{2, 3, 4, 5, 7},
				{2, 3, 4, 5, 7},
			},
			pre: [2][]int{
				{3, 2, 5, 4, 7},
				{5, 3, 2, 4, 7},
			},
			post: [2][]int{
				{2, 4, 7, 5, 3},
				{2, 4, 3, 7, 5},
			},
		},
		{
			name:  "right",
			dir:   'r',
			input: []int{5, 3, 2, 4, 7},
			in: [2][]int{
				{2, 3, 4, 5, 7},
				{2, 3, 4, 5, 7},
			},
			pre: [2][]int{
				{5, 3, 2, 4, 7},
				{3, 2, 5, 4, 7},
			},
			post: [2][]int{
				{2, 4, 3, 7, 5},
				{2, 4, 7, 5, 3},
			},
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			var root *Node
			for _, key := range test.input {
				root, _ = root.Insert(key)
			}

			// Pre-rotation checks.
			testInOrder(ErrorPrefix{t, "before: "}, root, test.in[0])
			testPreOrder(ErrorPrefix{t, "before: "}, root, test.pre[0])
			testPostOrder(ErrorPrefix{t, "before: "}, root, test.post[0])

			switch test.dir {
			case 'l':
				root = root.RotateLeft()
			case 'r':
				root = root.RotateRight()
			default:
				t.Fatalf("unknown rotate direction: %q", test.dir)
			}

			// Post-rotation checks.
			testInOrder(ErrorPrefix{t, "after: "}, root, test.in[1])
			testPreOrder(ErrorPrefix{t, "after: "}, root, test.pre[1])
			testPostOrder(ErrorPrefix{t, "after: "}, root, test.post[1])
		})
	}
}

func inOrder(n *Node) (keys []int) {
	n.InOrder(func(key int) {
		keys = append(keys, key)
	})
	return
}

func preOrder(n *Node) (keys []int) {
	n.PreOrder(func(key int) {
		keys = append(keys, key)
	})
	return
}

func postOrder(n *Node) (keys []int) {
	n.PostOrder(func(key int) {
		keys = append(keys, key)
	})
	return
}

type ErrReporter interface {
	Errorf(string, ...interface{})
}

type ErrorPrefix struct {
	T      *testing.T
	Prefix string
}

func (ep ErrorPrefix) Errorf(f string, args ...interface{}) {
	ep.T.Error(fmt.Sprint(ep.Prefix, fmt.Sprintf(f, args...)))
}

func testInOrder(t ErrReporter, root *Node, exp []int) {
	if act := inOrder(root); !reflect.DeepEqual(act, exp) {
		t.Errorf("unexpected in order sequence: %v; want %v", act, exp)
	}
}
func testPreOrder(t ErrReporter, root *Node, exp []int) {
	if act := preOrder(root); !reflect.DeepEqual(act, exp) {
		t.Errorf("unexpected pre order sequence: %v; want %v", act, exp)
	}
}
func testPostOrder(t ErrReporter, root *Node, exp []int) {
	if act := postOrder(root); !reflect.DeepEqual(act, exp) {
		t.Errorf("unexpected post order sequence: %v; want %v", act, exp)
	}
}
