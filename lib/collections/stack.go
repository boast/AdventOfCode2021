package collections

type StackString []string

// IsEmpty checks if the stack is empty.
func (s *StackString) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *StackString) Push(str string) {
	*s = append(*s, str)
}

// Pop removes and returns the top element of the stack.
func (s *StackString) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		indexLastElement := len(*s) - 1
		lastElement := (*s)[indexLastElement]
		*s = (*s)[:indexLastElement]
		return lastElement, true
	}
}

// Peek returns the top element of the stack without removing it.
func (s *StackString) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		indexLastElement := len(*s) - 1
		return (*s)[indexLastElement], true
	}
}

type StackInt []int

// IsEmpty checks if the stack is empty.
func (s *StackInt) IsEmpty() bool {
	return len(*s) == 0
}

// Push adds an element to the top of the stack.
func (s *StackInt) Push(str int) {
	*s = append(*s, str)
}

// Pop removes and returns the top element of the stack.
func (s *StackInt) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		indexLastElement := len(*s) - 1
		lastElement := (*s)[indexLastElement]
		*s = (*s)[:indexLastElement]
		return lastElement, true
	}
}

// Peek returns the top element of the stack without removing it.
func (s *StackInt) Peek() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		indexLastElement := len(*s) - 1
		return (*s)[indexLastElement], true
	}
}
