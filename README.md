# Data Structures

### install package
```` 
go get github.com/m-simeonov/go-data-structures
````

## LIFO Stack (Last In First Out)

Stack is a linear data structure which follows a particular order in which the operations are performed. The order is Last In First Out.

#### Init Stack
````
stack := stack.New()
````

#### Push Element
````
stack.Push("one")
````

#### Pop Element
````
val := stack.Pop()
````

#### Get Length
````
len := stack.GetLength()
````

#### Example LIFO Stack
```` 
lifoStack := stack.New()
// Push Element to LifoStack
lifoStack.Push("one")
lifoStack.Push("two")

fmt.Println(lifoStack.GetLength()) // will print `2`

element, err := lifoStack.Pop()
if err == nil {
    fmt.Println(element) // will print `two`
}
```` 

## Linked List
A linked list is a linear data structure, in which the elements are not stored at contiguous memory locations. A linked list consists of nodes where each node contains a data field and a reference(link) to the next node in the list

#### Init linked list
````
linkedList := list.New()
````

#### Add Value to linked list
````
linkedList.Add("one")
````

#### Delete node
````
linkedList.Delete(head)
````

#### Convert linked list to slice
````
data := linkedList.ToSlice()
````

#### Get Head
````
tail := linkedList.GetHead()
````

#### Get Tail
````
tail := linkedList.GetTail()
````

#### Get Length
````
len := linkedList.GetLength()
````

#### Example linked list
```` 
linkedList := list.New()
linkedList.Add("one")
linkedList.Add("two")
linkedList.Add("three")

data := linkedList.ToSlice()
fmt.Println(data) // will print [one two three]

len := linkedList.GetLength()
fmt.Println(len) // will print 3

head := linkedList.GetHead()
fmt.Println(head.Data) // will print `one`
linkedList.Delete(head)

data = linkedList.ToSlice()
fmt.Println(data) // will print [two three]
````