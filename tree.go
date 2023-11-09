// The package tree provides resources for building and manipulating a tree data structure, which consists of nodes containing objects.
// A Node is a container that holds data and references to other nodes, and can have one parent and multiple children.
// It can store any type of data and has a key/value field for storing additional data.
package tree

import (
	"errors"
)

// A Node is the fundamental building block of a tree data structure.
// It is a container that holds data and references to other nodes.
// A node may have zero or one parent, and it may have multiple children.
// The root node is the top-level node in a tree, and it is the only node that does not have a parent.
// A node can store any type of data.
// It also has a key/value field, which is a map that can be used to store additional data that may be used for processing or rendering.
type Node struct {
	parent   *Node
	children []*Node
	Object   interface{}
	values   map[string]interface{}
}

// Parent returns a pointer of the parent Node. A node can only have one parent.
// If a node has no parent, it is the root of the tree and Parent() returns nil.
func (o *Node) Parent() *Node { return o.parent }

// SetParent sets or changes the parent Node.
func (o *Node) SetParent(parent *Node) { o.parent = parent }

// Children return a slice of the children Nodes.
func (o *Node) Children() []*Node { return o.children }

func (o *Node) SetChildren(children []*Node) { o.children = children }

// GetChild returns the child node at the given index.
// If the index is out of range, it returns nil.
func (o *Node) GetChild(index int) *Node {
	if o.HasChild(index) {
		return o.children[index]
	}
	return nil
}

func (o *Node) HasChild(index int) bool {
	return index >= 0 && index < len(o.children)
}

type Filter func(*Node) bool

type SearchFunc func(*Node, interface{}) bool

func NewNode() *Node {
	o := new(Node)
	return o
}

func (o *Node) Adopt(child *Node) error {
	if child == nil {
		return errors.New("Child is nil")
	}
	if child.parent != nil {
		return errors.New("Child already has a parent")
	}
	o.children = append(o.children, child)
	child.parent = o
	return nil
}

// Attach attaches the given child node to the first node found by the searchFunc function
// using the given searchCriterion value. If no node is found, an error is returned.
func (o *Node) Attach(child *Node, searchFunc SearchFunc, searchCriterion interface{}) error {
	if child == nil {
		return errors.New("Child is nil")
	}
	if searchFunc == nil {
		return errors.New("Invalid search function")
	}
	// Search for the node using the provided search function and value.
	foundNode, ok := o.Search(searchFunc, searchCriterion)
	if !ok {
		return errors.New("Node not found")
	}
	if err := foundNode.Adopt(child); err != nil {
		return err
	}
	return nil
}

// RecursiveProcessing allows to define a pre and post processing fun
func (o *Node) RecursiveProcessing(preProcessing func(*Node), postProcessing func(*Node)) {
	if preProcessing != nil {
		preProcessing(o)
	}
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			child.RecursiveProcessing(preProcessing, postProcessing)
		}
	}
	if postProcessing != nil {
		postProcessing(o)
	}
}

func (o *Node) StoppableRecursiveProcessing(preProcessing func(*Node) bool, postProcessing func(*Node) bool) (stop bool) {
	if preProcessing != nil {
		stop = preProcessing(o)
		if stop {
			return true
		}
	}
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			stop = child.StoppableRecursiveProcessing(preProcessing, postProcessing)
			if stop {
				return true
			}
		}
	}
	if postProcessing != nil {
		stop = postProcessing(o)
		if stop {
			return true
		}
	}
	return false
}

func (o *Node) Search(searchFunc SearchFunc, searchCriterion interface{}) (*Node, bool) {
	if searchFunc(o, searchCriterion) {
		return o, true
	}
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			if searchFunc(child, searchCriterion) {
				return child, true
			}
			searchedVal, ok := child.Search(searchFunc, searchCriterion)
			if ok {
				return searchedVal, true
			}
		}
	}
	return nil, false
}

func (o *Node) Deepness() int {
	var deep int
	o.deepness(0, &deep)
	return deep
}

func (o *Node) deepness(d int, deep *int) {
	d = d + 1
	if d > *deep {
		*deep = d
	}
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			child.deepness(d, deep)
		}
	}
}

func (o *Node) PreOrderIter() (nodes []*Node) {
	nodes = append(nodes, o)
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			for _, cn := range child.PreOrderIter() {
				nodes = append(nodes, cn)
			}
		}
	}
	return nodes
}

func (o *Node) PreOrderIterAdvanced(filter Filter) (nodes []*Node) {
	if filter == nil {
		filter = func(*Node) bool { return true }
	}
	if filter(o) {
		nodes = append(nodes, o)
	}
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			for _, cn := range child.PreOrderIterAdvanced(filter) {
				if filter(cn) {
					nodes = append(nodes, cn)
				}
			}
		}
	}
	return nodes
}

func (o *Node) PostOrderIter() (nodes []*Node) {
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			for _, cn := range child.PostOrderIter() {
				nodes = append(nodes, cn)
			}
		}
	}
	nodes = append(nodes, o)
	return nodes
}

func (o *Node) LevelOrderIter() (nodes [][]*Node) {
	nodes = make([][]*Node, 0)
	o.levelOrderIter(0, &nodes)
	return nodes
}

func (o *Node) levelOrderIter(deepness int, nodes *[][]*Node) {
	if len(*nodes) <= deepness {
		*nodes = append(*nodes, make([]*Node, 0))
	}
	(*nodes)[deepness] = append((*nodes)[deepness], o)
	deepness = deepness + 1
	childrenCount := len(o.children)
	if childrenCount > 0 {
		for _, child := range o.children {
			child.levelOrderIter(deepness, nodes)
		}
	}
}

func (o *Node) IsRoot() bool {
	return (o.parent == nil)
}

func (o *Node) GetRoot() *Node {
	for {
		if o.parent == nil {
			return o
		}
		o = o.parent
	}
}

func (o *Node) HasChildren() bool {
	return (len(o.children) != 0)
}

// ChildIndex returns the index of the node in it's parent's children list.
func (o *Node) ChildIndex() (int, error) {
	if o.IsRoot() {
		return 0, errors.New("Root node has no parent and therefore no child index")
	}
	parent := o.Parent()
	for index, child := range parent.children {
		if o == child {
			return index, nil
		}
	}
	// This error should never be reached
	return 0, errors.New("Root not found")
}

func (o *Node) NextSibling() (*Node, error) {
	if o.IsRoot() {
		return nil, errors.New("Root node has no sibling")
	}
	childIndex, err := o.ChildIndex()
	if err != nil {
		return nil, err
	}
	parent := o.Parent()
	if childIndex == len(parent.children)-1 {
		return nil, errors.New("This node is the last of the siblings")
	}
	return parent.GetChild(childIndex + 1), nil
}

func (o *Node) PreviousSibling() (*Node, error) {
	if o.IsRoot() {
		return nil, errors.New("Root node has no sibling")
	}
	childIndex, err := o.ChildIndex()
	if err != nil {
		return nil, err
	}
	parent := o.Parent()
	if childIndex == 0 {
		return nil, errors.New("This node is the first of the siblings")
	}
	return parent.GetChild(childIndex - 1), nil
}

func (o *Node) ProcessNodeAndAscendants(processing func(*Node) bool) (stop bool) {
	stop = processing(o)
	if stop {
		return true
	}
	parent := o.Parent()
	if parent == nil {
		return true
	}
	stop = parent.ProcessNodeAndAscendants(processing)
	if stop {
		return true
	}
	return false
}

// ChildrenCount returns the number of children Nodes.
func (o *Node) ChildrenCount() int {
	return len(o.children)
}

func (o *Node) SwapChildren(childIndex1, childIndex2 int) error {
	childrenCount := len(o.children)
	if childIndex1 < 0 || childIndex1 >= childrenCount || childIndex2 < 0 || childIndex2 >= childrenCount {
		return errors.New("bad index for child1 or child2")
	}
	if childIndex1 == childIndex2 {
		return errors.New("childIndex1 and childIndex2 must be different")
	}
	child := o.children[childIndex1]
	o.children[childIndex1] = o.children[childIndex2]
	o.children[childIndex2] = child
	return nil
}
