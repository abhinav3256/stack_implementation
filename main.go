package main

import "fmt"

type stack struct {
	size int
	top  int
	data []int
}

func main() {
	s := stack{}
	s.size = 6
	s.pop()
	s.push(12)
	fmt.Println(s.data)
	s.pop()
	s.push(20)
	s.push(50)

	fmt.Println(s.data)
	fmt.Println(s.top)

	s.pop()

	fmt.Println(s.data)
	fmt.Println(s.top)
}

func (s *stack) push(x int) {

	if !s.isFull() {
		s.data = append(s.data, x)
		s.top++
	} else {
		fmt.Println("overflow")
	}

}

func (s *stack) isFull() bool {

	if s.top == s.size {
		return true
	} else {
		return false
	}
}
func (s *stack) pop() {

	if !s.isEmpty() {
		s.data = append(s.data[:s.top-1], s.data[s.top-1+1:]...)
		s.top--
	} else {
		fmt.Println("underflow")
	}

}

func (s *stack) isEmpty() bool {

	if s.top == 0 {
		return true
	} else {
		return false
	}

}
