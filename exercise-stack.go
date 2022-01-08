// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Stack struct {
	i    int
	data []int
}

func NewStack(d []int) *Stack {
	return &Stack{data: d, i: len(d) - 1}
}

func (s *Stack) push(n int) {
	s.data = append([]int{n}, s.data...)
	s.i++
}

func (s *Stack) pop() (n int) {
	n = s.data[s.i]
	s.data = s.data[:s.i]
	s.i--
	return n
}

func main() {
	s := NewStack([]int{1, 2, 3})
	fmt.Println(s)
	s.push(5)
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s)
}
