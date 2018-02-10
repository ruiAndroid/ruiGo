package goPackage


import (
	"container/list"
	"fmt"
)

/**
利用container 中的list(双向链表) 放入101,102,103并打印出来
 */
func ListMethod() {
	//创建双向链表
	lst:=list.New()
	lst.PushBack(101)
	lst.PushBack(102)
	lst.PushBack(103)

	//打印list
	for e:=lst.Front();e!=nil;e=e.Next(){
		fmt.Println(e.Value)
	}
}
