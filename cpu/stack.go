package cpu

const StackSize = 256

type Stack struct {
	Data []uint16
}

func NewStack() *Stack {
	return &Stack{
		Data: []uint16{},
	}
}

func (s *Stack) Push(value uint16) {
	s.Data = append(s.Data, value)
	if len(s.Data) >= StackSize {
		panic("Stack overflow")
	}
}

func (s *Stack) Pop() uint16 {
	if len(s.Data) == 0 {
		panic("Stack underflow")
	}
	value := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return value
}
