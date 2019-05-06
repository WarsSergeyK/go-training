package domain

import "fmt"

type Collection struct {
	lastNodePointer *Node
}

type Node struct {
	value       int
	prevPointer *Node
	nextPointer *Node
}

func (c *Collection) Add(element int) {

	newNode := Node{
		value:       element,
		prevPointer: c.lastNodePointer,
	}

	// Not the first node already
	if c.lastNodePointer != nil {
		newNode.Prev().nextPointer = &newNode
	}

	c.lastNodePointer = &newNode
}

func (c *Collection) Get(index int) *Node {

	getPointer := c.First()

	for i := 0; i < index; i++ {
		getPointer = getPointer.Next()
	}

	return getPointer
}

func (c *Collection) Remove(index int) {

	c.Get(index).Prev().nextPointer = c.Get(index).Next()
	c.Get(index).Next().prevPointer = c.Get(index).Prev()
}

func (c *Collection) First() *Node {
	FirstPointer := c.lastNodePointer

	for FirstPointer.Prev() != nil {
		FirstPointer = FirstPointer.Prev()
	}
	return FirstPointer
}

func (c *Collection) Last() *Node {

	return c.lastNodePointer
}

func (c Collection) Length() int {

	for i := 0; ; i++ {
		if c.Get(i) == nil {
			return i
		}
	}
}

func (c Collection) Print() {

	if c.lastNodePointer == nil {
		fmt.Println("Empty collection")
		return
	}

	toPrint := c.First()

	for toPrint.nextPointer != nil {
		fmt.Printf("%d\t", toPrint.Value())
		toPrint = toPrint.Next()
	}

	fmt.Printf("%d\n", toPrint.Value())
}

func (e *Node) Next() *Node {
	return e.nextPointer
}

func (e *Node) Prev() *Node {

	if e.prevPointer == nil && e.nextPointer == nil {
		fmt.Println("Empty collection")
	}
	return e.prevPointer
}

func (e *Node) Value() int {
	return e.value
}
