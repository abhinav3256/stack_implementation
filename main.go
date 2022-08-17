package main

import "fmt"

type stack struct {
	size int
	top  int
	data []int
}

func main() {
	s := stack{}
	empty := s.isEmpty()
	fmt.Println(empty)
	s.pop()
	s.size = 4
	s.isFull()
	s.push(10)
	s.push(15)
	s.push(19)
	s.pop()
	s.push(12)
	empty = s.isEmpty()
	fmt.Println(empty)

	fmt.Println(s.data)

}

func (s *stack) isFull() bool {

	//fmt.Println(s.top)
	if s.size == s.top {
		return true
	} else {
		return false
	}
}

func (s *stack) isEmpty() bool {
	if s.top == 0 {
		return true
	} else {
		return false
	}
}

func (s *stack) push(x int) {

	if !s.isFull() {
		s.data = append(s.data, x)
		s.top++
	} else {
		fmt.Println("overflow condition")
	}

}

func (s *stack) pop() {
	fmt.Println(s.data)
	if !s.isEmpty() {
		s.data = append(s.data[:s.top-1], s.data[s.top-1+1:]...)
		s.top--
	} else {
		fmt.Println("underflow condition")
	}

	fmt.Println(s.data)
}
