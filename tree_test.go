package tree

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func foundId(n *Node, searchedValue interface{}) bool {
	if n.Values()["id"] == searchedValue.(string) {
		return true
	}
	return false
}

func TestNewNode(t *testing.T) {
	assert := assert.New(t)
	n := NewNode()
	assert.Nil(n.Parent())
	assert.Nil(n.Children())
	assert.Nil(n.Values())
	n.SetValue("test", 0)
	assert.NotNil(n.Values())
	assert.Equal(1, len(n.Values()))
}

func TestNewNodeWithValue(t *testing.T) {
	assert := assert.New(t)
	n := NewNodeWithValue("id", "0123")
	assert.Nil(n.Parent())
	assert.Nil(n.Children())
	assert.NotNil(n.Values())
	assert.Equal(1, len(n.Values()))
	assert.Equal("0123", n.Values()["id"])
}

func TestParent(t *testing.T) {
	// Test parent of root node
	root := NewNode()
	if root.Parent() != nil {
		t.Error("Expected parent of root node to be nil")
	}

	// Test parent of non-root node
	child1 := NewNode()
	root.Adopt(child1)
	if child1.Parent() != root {
		t.Error("Expected parent of child1 to be root")
	}

	// Test changing parent of non-root node
	child2 := NewNode()
	root.Adopt(child2)
	child2.SetParent(child1)
	if child2.Parent() != child1 {
		t.Error("Expected parent of child2 to be child1")
	}
}

func TestChildren(t *testing.T) {
	var l int
	assert := assert.New(t)

	n0 := NewNode()
	n1 := NewNode()
	n2 := NewNode()
	n3 := NewNode()

	l = len(n0.Children())
	assert.Equal(l, 0)
	l = len(n1.Children())
	assert.Equal(l, 0)

	n0.Adopt(n1)
	l = len(n0.Children())
	assert.Equal(l, 1)

	n0.Adopt(n2)
	n0.Adopt(n3)
	l = len(n0.Children())
	assert.Equal(l, 3)
}

func TestValues(t *testing.T) {
	assert := assert.New(t)

	n := NewNode()

	n.SetValue("style", 2)
	n.SetValue("text", "Bonjour")
	n.SetValue("action", func(i int) int { return i * 2 })

	f := func(a int, b int) int {
		return a + b
	}
	n.SetValue("addaction", f)

	// TODO vérifier tous les asserts :  expected puis actual
	i, _ := n.Value("style")
	assert.Equal(2, i.(int))
	s, _ := n.Value("text")
	assert.Equal("Bonjour", s.(string))
	f1, _ := n.Value("action")
	assert.Equal(10, f1.(func(i int) int)(5))
	f2, _ := n.Value("addaction")
	assert.Equal(7, f2.(func(a int, b int) int)(3, 4))
}

func TestTree(t *testing.T) {
	node0 := buildTestTree()
	assert.Equal(t, 4, len(node0.Children()))
	assert.Equal(t, 6, node0.Deepness())
}

func concatNodeIds(nodes []*Node) string {
	s := ""
	for _, n := range nodes {
		s = s + " " + n.Values()["id"].(string)
	}
	return s
}

func TestPreOrderIteration(t *testing.T) {
	var s string
	assert := assert.New(t)

	node0 := buildTestTree()
	assert.Equal(len(node0.Children()), 4)

	s = concatNodeIds(node0.PreOrderIter())
	assert.Equal(s, " 000 111 222 333 444 555 666 777 888 999 AAA BBB CCC DDD EEE FFF GGG HHH")

	fltr := func(node *Node) bool {
		matched, _ := regexp.MatchString(`\d+`, node.Values()["id"].(string)) // node.Value.(Val).Text)
		return matched
	}
	s = concatNodeIds(node0.PreOrderIterAdvanced(fltr))
	assert.Equal(s, " 000 111 222 333 444 555 666 777 888 999")

	s = concatNodeIds(node0.PreOrderIterAdvanced(nil))
	assert.Equal(s, " 000 111 222 333 444 555 666 777 888 999 AAA BBB CCC DDD EEE FFF GGG HHH")
}

func TestPostOrderIteration(t *testing.T) {
	var s string
	node0 := buildTestTree()

	s = concatNodeIds(node0.PostOrderIter())
	assert.Equal(t, s, " 222 111 555 444 333 888 999 CCC BBB AAA 777 GGG FFF EEE DDD 666 HHH 000")
}

func buildTestTree() *Node {
	node0 := NewNodeWithValue("id", "000")
	node1 := NewNodeWithValue("id", "111")
	node0.Adopt(node1)
	node2 := NewNodeWithValue("id", "222")
	node1.Adopt(node2)
	node3 := NewNodeWithValue("id", "333")
	node0.Adopt(node3)
	node4 := NewNodeWithValue("id", "444")
	node3.Adopt(node4)
	node5 := NewNodeWithValue("id", "555")
	node4.Adopt(node5)
	node6 := NewNodeWithValue("id", "666")
	node0.Adopt(node6)
	node7 := NewNodeWithValue("id", "777")
	node6.Adopt(node7)
	node8 := NewNodeWithValue("id", "888")
	node7.Adopt(node8)
	node9 := NewNodeWithValue("id", "999")
	node7.Adopt(node9)
	node10 := NewNodeWithValue("id", "AAA")
	node7.Adopt(node10)
	node11 := NewNodeWithValue("id", "BBB")
	node10.Adopt(node11)
	node12 := NewNodeWithValue("id", "CCC")
	node11.Adopt(node12)
	node13 := NewNodeWithValue("id", "DDD")
	node6.Adopt(node13)
	node14 := NewNodeWithValue("id", "EEE")
	node13.Adopt(node14)
	node15 := NewNodeWithValue("id", "FFF")
	node14.Adopt(node15)
	node16 := NewNodeWithValue("id", "GGG")
	node15.Adopt(node16)
	node17 := NewNodeWithValue("id", "HHH")
	node0.Adopt(node17)

	return node0
}

func TestUpdateDeepnessValues(t *testing.T) {
	node := buildTestTree()
	RecursiveSetValue_Deepness(node, "deepness")
	assert.Equal(t, 0, node.Values()["deepness"].(int))
	// The two lines below tests the same value
	assert.Equal(t, 2, node.Children()[0].Children()[0].Values()["deepness"].(int))
	assert.Equal(t, 2, node.GetChild(0).GetChild(0).ValueInt("deepness", 0))
}

func ExampleLevelOrderIter() {
	var s string
	node := buildTestTree()
	nodes := node.LevelOrderIter()
	for i, ln := range nodes {
		s = s + fmt.Sprintf("level %d:", i)
		for _, n := range ln {
			s = s + " " + n.Values()["id"].(string)
		}
		s = s + "\n"
	}
	fmt.Print(s)
	// Output:
	// level 0: 000
	// level 1: 111 333 666 HHH
	// level 2: 222 444 777 DDD
	// level 3: 555 888 999 AAA EEE
	// level 4: BBB FFF
	// level 5: CCC GGG
}
