// double linked lists in golang

package main

import (
	"container/list"
	"fmt"
)

func main() {
	// double linked list experiment:
	var dblLinkList *list.List = list.New()
	dblLinkList.PushFront(101)
	dblLinkList.PushFront(102)
	dblLinkList.PushFront(103)
	for element := dblLinkList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
	fmt.Println()
	fmt.Println(*dblLinkList)
}
