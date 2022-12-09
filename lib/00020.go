package lib

type stack struct {
	s []rune
	c int
}

func (s *stack) push(v rune) {
	s.s = append(s.s, v)
	s.c++
}

func (s *stack) pop() rune {
	v := s.s[s.c-1]
	s.c--
	s.s = s.s[:s.c]
	return v
}

func (s *stack) empty() bool {
	return s.c == 0
}

func initStack() *stack {
	return &stack{
		s: make([]rune, 0),
		c: 0,
	}
}

func pair(s *stack, c rune) bool {
	if s.empty() {
		return false
	}
	return s.pop() == c
}

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	ps := map[rune]rune{
		rune(')'): rune('('),
		rune(']'): rune('['),
		rune('}'): rune('{'),
	}
	preMap := make(map[rune]struct{})
	for _, v := range ps {
		preMap[v] = struct{}{}
	}
	stack := initStack()
	for _, sub := range s {
		if _, ok := preMap[sub]; ok {
			stack.push(sub)
			continue
		}
		if pre, ok := ps[sub]; ok {
			if pair(stack, pre) {
				continue
			} else {
				return false
			}
		}
	}
	return stack.empty()
}
