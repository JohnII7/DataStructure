package main

import (
	"fmt"
)

type LinkNode struct {
	Data interface{}
	Next *LinkNode
}

type LList = LinkNode

type LLister interface {
	Creat_head()                  //头插法创建链表
	Creat_tail()                  //尾插法创建链表
	Insert(v interface{}, i int)  //插入v到第ige
	Delete(index int) interface{} //删除第i+1（i从0开始）个，返回value
	GetLength() int               //统计节点个数（不包括头节点）
	Search(v interface{}) int     //查找值为v的元素，返回所在位置
	Display()
}

func newLList() *LList {
	return &LList{Next: nil}
}

func (head *LList) Creat_head() {
	for i := 0; i < 10; i++ {
		s := &LinkNode{
			Data: i,
		}
		s.Next = head.Next
		head.Next = s
	}
}

func (head *LList) Creat_tail() {
	p := head
	for i := 0; i < 10; i++ {
		s := &LinkNode{
			Data: i,
		}
		p.Next = s
		p = s
	}
}

func (head *LList) Insert(v interface{}, index int) {
	length := head.GetLength()       //链表的长度，如3个就是3
	if index < 0 || index > length { //i=0时相当于头插法；i=length相当于尾插法；i=1时是在第一个节点之后插入
		fmt.Printf("index out of range %d\n", length)
		return
	}
	p := head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	newNode := &LinkNode{Data: v}
	newNode.Next = p.Next
	p.Next = newNode
}

func (head *LList) Delete(index int) interface{} {
	length := head.GetLength()
	if index < 0 || index >= length {
		fmt.Printf("index out of range %d\n", length)
		return "index out of range, please check."
	}
	p := head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	deleteNode := p.Next
	p.Next = p.Next.Next
	return deleteNode.Data
}

func (head *LList) GetLength() int {
	p := head.Next
	var length int = 0
	for p != nil {
		length++
		p = p.Next
	}
	return length
}

func (head *LList) Search(value int) interface{} {
	p := head.Next
	index := 1
	for p != nil {
		if p.Data == value {
			return index
		}
		index++
		p = p.Next
	}
	return fmt.Sprintf("the value %v is not in the llist;", value)
}

func (head *LList) Display() {
	p := head.Next
	fmt.Print("head")
	for p != nil {
		fmt.Printf("->%v", p.Data)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	head := newLList()
	head.Creat_tail()
	head.Display()
	head.Insert(99, 10)
	head.Display()
	result := head.Search(99)
	fmt.Printf("search result is: %v\n", result)
	deleteValue := head.Delete(10)
	fmt.Printf("delete node value is %v\n", deleteValue)
	head.Display()
	fmt.Printf("link list length is: %v", head.GetLength())
}
