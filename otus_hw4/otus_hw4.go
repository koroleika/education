package main

import (
	"container/list"
	"fmt"
)
/*
func PushFrontFunc(list *list.List , i int) *list.List  {
	//fmt.Println(list.Len())
	list.PushFront(i)
	return list
	//fmt.Println(list.Len())
	//-----------------
	//PushFrontFunc(list)
	//list.PushFront(1)
	for i:=10; i<=0;i--{
		PushFrontFunc(list,i)
	}
	fmt.Println(list.Len())
}
*/
func main(){
	list2 := list.New()
	/*
	element1 := list.PushFront(2)
	element2 := list.PushBack(3)
	list.InsertBefore(1,element1)
	list.InsertAfter(4,element2)
*/
	for i:=10; i<=0;i--{
		list2.PushFront(i)
	}

	fmt.Println(list2.Len())
	//fmt.Println(list.Back())
	for i:= list2.Front(); i!=nil; i = i.Prev(){
		fmt.Println(i.Value)
	}



}
