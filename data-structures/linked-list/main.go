package main

import "fmt"

type ListItem struct {
	data string
	next *ListItem // pointer to next
}

type List struct {
	length int
	start  *ListItem
}

// create an append function on the list type
func (l *List) Append(newItem *ListItem) {
	if l.length == 0 {
		l.start = newItem
	} else {
		currentItem := l.start
		for currentItem.next != nil {
			currentItem = currentItem.next
		}
		currentItem.next = newItem
	}
	l.length++
}

func main() {
	l := &List{}

	i1 := ListItem{
		data: "beep boop",
	}

	l.Append(&i1)

	fmt.Printf("Length: %v\n", l.length)
	fmt.Printf("First: %v\n", l.start)

	i2 := ListItem{
		data: "beep boop beep boop",
	}

	l.Append(&i2)

	fmt.Printf("Length: %v\n", l.length)
	fmt.Printf("First: %v\n", l.start)
	fmt.Printf("Second: %v\n", l.start.next)
}
