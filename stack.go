package goraph

type NodeStack struct {
	stack []int
}

func NewNodeStack() *NodeStack {

	s := new(NodeStack)
	s.stack = make([]int, 0)
	return s
}

func (s *NodeStack) Push(nodeID int) {
	s.stack = append(s.stack, nodeID)
}

func (s *NodeStack) Pop() int {
	if len(s.stack) > 0 {
		result := s.stack[len(s.stack)-1]
		if len(s.stack) > 1 {
			s.stack = s.stack[0 : len(s.stack)-1]
		} else {
			s.stack = make([]int, 0)
		}
		return result
	}
	panic("stack is empty")
}

func (s *NodeStack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *NodeStack) Len() int {
	return len(s.stack)
}
