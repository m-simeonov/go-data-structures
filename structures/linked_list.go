package structures

type LinkedList interface {
	Add(node *Node)
	ToSlice() []interface{}
	Delete(node *Node) error
	GetLength() int
	GetHead() *Node
	GetTail() *Node
}

type Node struct {
	Data interface{}
	Next *Node
}

type linkedList struct {
	length int
	tail   *Node
	head   *Node
}

func NewLinkedList() LinkedList {
	return &linkedList{}
}

func (l *linkedList) Add(node *Node) {
	if l.head == nil {
		l.head = node
	} else {
		l.tail.Next = node
	}

	l.tail = node
	l.length++
}

func (l *linkedList) Delete(nodeDel *Node) error {
	i := 0
	n := l.head
	if n == nodeDel {
		l.head = nodeDel.Next
		l.length--
		return nil
	}

	for {
		if n == nil {
			break
		} else if n.Next == nodeDel {
			l.length--
			n.Next = nodeDel.Next
		}

		i++
		n = n.Next
	}

	return nil
}

func (l linkedList) ToSlice() []interface{} {
	result := make([]interface{}, l.length)
	node := l.head
	i := 0
	for {
		if node == nil {
			break
		}
		result[i] = node.Data
		i++
		node = node.Next
	}

	return result
}

func (l linkedList) GetHead() *Node {
	return l.head
}

func (l linkedList) GetTail() *Node {
	return l.tail
}

func (l linkedList) GetLength() int {
	return l.length
}
