# Data Structures

####install package
<code>go get github.com/m-simeonov/go-data-structures </code>

##LIFO Stack (Last In First Out)

Stack is a linear data structure which follows a particular order in which the operations are performed. The order is Last In First Out.

#### Init Stack
<code>stack := structures.NewStack()</code>

#### Push Element
<code>stack.Push("one")</code>

#### Pop Element
<code>val := stack.Pop()</code>

#### Get Length
<code>len := stack.GetLength()</code>

#### Example LIFO Stack
```` 
stack := structures.NewLifoStack()
//Push Element to LifoStack
stack.Push("one")
stack.Push("two")

fmt.Println(stack.GetLength()) //will print `2`

element, err := stack.Pop()
if err == nil {
    fmt.Println(element) //will print `two`
}
```` 

## Linked List
A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. A linked list consists of nodes where each node contains a data field and a reference(link) to the next node in the list

#### Init linked list
<code>list := structures.NewLinkedList()</code>

#### Add to linked list
<code>list.Add(&structures.Node{Data: "one"})</code>

#### Delete node
<code>list.Delete(head)</code>

#### Convert linked list to slice
<code>data := list.ToSlice()</code>

#### Get Head
<code>tail := list.GetHead()</code>

#### Get Tail
<code>tail := list.GetTail()</code>

#### Get Length
<code>len := list.GetLength()</code>

####Example linked list
```` 
list := structures.NewLinkedList()
list.Add(&structures.Node{Data: "one"})
list.Add(&structures.Node{Data: "two"})
list.Add(&structures.Node{Data: "three"})

data := list.ToSlice()
fmt.Println(data) // will print [one two three]

len := list.GetLength()
fmt.Println(len) // will print 3

head := list.GetHead()
fmt.Println(head.Data) // will print `one`
list.Delete(head)

data = list.ToSlice()
fmt.Println(data) // will print [two three]
````