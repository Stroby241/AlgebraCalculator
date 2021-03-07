package V3

import (
	"fmt"
	"strconv"
)

type INode interface {
	setParent(partent INode)
	getParent() INode
	addChild(child INode, onRightSide bool) INode
	setChilds(childs []INode)
	getChilds() []INode
	getMaxChilds() int

	getType() int
	getRank() int
	getDefiner(vaules bool) string
	getDeepDefiner(vaules bool) string

	copy() INode
	solve() bool
	sort() bool
	print()
	printTree(indentation int)
}

const (
	TypNone            = 0
	TypVector          = 1
	TypVariable        = 2
	TypOpperator       = 3
	TypMathFunction    = 4
	TypSubOperation    = 5
	TypTerm            = 6
	TypComplexFunction = 7

	RankNone            = 0
	RankAppend          = 1
	RankAddSub          = 2
	RankMul             = 3
	RankPow             = 4
	RankMathFunction    = 5
	RankSubOpperation   = 6
	RankTerm            = 7
	RankComplexFunction = 9
)

type Node struct {
	parent    INode
	childs    []INode
	typeId    int
	rank      int
	maxChilds int
	definer   string
}

func NewNode(typeId int, rank int, maxChilds int) *Node {
	return &Node{
		typeId:    typeId,
		rank:      rank,
		maxChilds: maxChilds,
		definer:   strconv.Itoa(typeId),
	}
}

func (t *Node) setParent(partent INode) {
	t.parent = partent
}
func (t *Node) getParent() INode {
	return t.parent
}
func (t *Node) addChild(child INode, onRightSide bool) INode {
	if onRightSide {
		t.childs = append(t.childs, child)
	} else {
		t.childs = append([]INode{child}, t.childs...)
	}
	return child
}
func (t *Node) setChilds(childs []INode) {
	t.childs = childs
}
func (t *Node) getChilds() []INode {
	return t.childs
}
func (t *Node) getMaxChilds() int {
	return t.maxChilds
}

func (t Node) getType() int {
	return t.typeId
}
func (t Node) getRank() int {
	return t.rank
}
func (t Node) getDefiner(vaules bool) string {
	return t.definer
}
func (t Node) getDeepDefiner(vaules bool) string {
	var deepDefiner string
	for _, child := range t.childs {
		deepDefiner += child.getDeepDefiner(vaules)
	}
	deepDefiner += t.definer
	return deepDefiner
}

func (t *Node) copy() INode {
	copy := NewNode(t.typeId, t.rank, t.maxChilds)
	copy.childs = make([]INode, len(t.childs))

	for i, child := range t.childs {
		childCopy := child.copy()
		childCopy.setParent(copy)
		copy.childs[i] = childCopy
	}
	return copy
}
func (t *Node) solve() bool {
	solved := false
	for _, child := range t.childs {
		if child.solve() {
			solved = true
		}
	}
	return solved
}
func (t *Node) sort() bool {
	sorted := false
	for _, child := range t.childs {
		if child.sort() {
			sorted = true
		}
	}
	return sorted
}
func (t *Node) print() {
	if len(t.childs) > 0 {
		fmt.Print("(")
		for _, child := range t.childs {
			child.print()
		}
		fmt.Print(")")
	}
}
func (t *Node) printTree(indentation int) {
	fmt.Print("\n")
	indentation++
	if len(t.childs) > 0 {
		for _, child := range t.childs {
			child.printTree(indentation)
		}
	}
	indentation--
}
func printIndentation(indentation int) {
	for i := 0; i < indentation; i++ {
		if i == indentation-1 {
			fmt.Print("|> ")
		} else if i == 0 {
			fmt.Print("|  ")
		} else {
			fmt.Print("   ")
		}

	}
}

// replaceNode replaces old to new and updates the partent and child pointers to new.
func replaceNode(old INode, new INode) {

	// Copy Node Data to new
	new.setChilds(old.getChilds())
	new.setParent(old.getParent())

	// Set the partents of the childs to new
	for _, child := range new.getChilds() {
		child.setParent(new)
	}

	// Set the childs of the partent to new
	if new.getParent() != nil {
		childs := new.getParent().getChilds()
		for i, child := range childs {
			if child == old {
				childs[i] = new
			}
		}
		new.getParent().setChilds(childs)
	}
}

// insertNode replaces old to new but keep the childs of new while conectiong the partent of old.
func insertNode(old INode, new INode) {

	// Copy Partent pointer
	new.setParent(old.getParent())

	// Set new as child of partent of old
	if new.getParent() != nil {
		childs := new.getParent().getChilds()
		for i, child := range childs {
			if child == old {
				childs[i] = new
			}
		}
		new.getParent().setChilds(childs)
	}
}

func deepEqual(node1 INode, node2 INode) bool {
	return node1.getDeepDefiner(true) == node2.getDeepDefiner(true)
}
