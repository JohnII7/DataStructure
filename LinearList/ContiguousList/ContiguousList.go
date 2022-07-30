package main

// 顺序表
import (
	"errors"
	"fmt"
)

const MaxSize int = 10

type SeqList struct {
	data   []int
	length int
}

//NewSeqList 声明并初始化顺序表，为顺序表分配内存空间
func NewSeqList() *SeqList {
	if MaxSize == 0 {
		return nil
	}
	return &SeqList{
		data:   make([]int, MaxSize, MaxSize),
		length: 0,
	}
}

//CreateList 给表里添加初始元素
func (s *SeqList) CreateList(data []int, n int) error {
	if n > MaxSize {
		return errors.New("表长不足")
	}
	for i, value := range data {
		s.data[i] = value
	}
	s.length = n
	return nil
}

// InsertList 把value插入第i位置，i-1为index；i从1开始
func (s *SeqList) InsertList(value, i int) error {
	if i < 1 || i > s.length+1 {
		return errors.New("error insert location")
	}
	for j := s.length; j >= i-1; j-- {
		s.data[j] = s.data[j-1]
	}
	s.data[i-1] = value
	s.length++
	return nil
}

//DeleteList 删除第i个
func (s *SeqList) DeleteList(i int) error {
	if i < 1 || i > s.length {
		return errors.New("delete error location")
	}
	for j := i; j < s.length; j++ {
		s.data[j-1] = s.data[j]
	}
	s.length--
	return nil
}

//SeqListPrint 打印顺序表
func (s *SeqList) SeqListPrint() {
	fmt.Println("========================")
	fmt.Print("data: ")
	for i := 0; i < s.length; i++ {
		fmt.Printf("%v", s.data[i])
	}
	fmt.Println()
	fmt.Println("length:", s.length)
}

func main() {
	seqList := NewSeqList()
	initData := []int{0, 1, 2, 3, 4, 5, 6, 7}
	err := seqList.CreateList(initData, len(initData))
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 1; i++ {
		seqList.InsertList(90, 2)
	}
	//seqList.DeleteList(2)
	seqList.SeqListPrint()
}
