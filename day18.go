package main

import (
	"os"
	"panic"
	"readers"
	"strconv"
)

type snailFish struct {
	left   *snailFish
	right  *snailFish
	parent *snailFish
	value  int
}

func Day18Part1() int {
	file, err := os.Open("assets/day18.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	snailFishes := make([]*snailFish, len(lines))

	for i, line := range lines {
		snailFishes[i] = parseSnailFish(line)
	}

	for len(snailFishes) > 1 {
		snailFishes[1] = snailFishes[0].add(snailFishes[1])
		snailFishes = snailFishes[1:]
	}

	return snailFishes[0].magnitude()
}

func Day18Part2() int {
	file, err := os.Open("assets/day18.txt")
	panic.Check(err)

	lines, err := readers.ReadStrings(file)
	panic.Check(err)

	snailFishes := make([]*snailFish, len(lines))

	for i, line := range lines {
		snailFishes[i] = parseSnailFish(line)
	}
	return maxSumMagnitude(snailFishes)
}

func maxSumMagnitude(numbers []*snailFish) int {
	maxFound := 0
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if i == j {
				continue
			}
			sum := numbers[i].clone().add(numbers[j].clone()).magnitude()
			if sum > maxFound {
				maxFound = sum
			}
		}
	}
	return maxFound
}

func (s *snailFish) clone() *snailFish {
	clone := &snailFish{}

	if s.left != nil {
		clone.left = s.left.clone()
		clone.left.parent = clone
	}
	if s.right != nil {
		clone.right = s.right.clone()
		clone.right.parent = clone
	}

	clone.value = s.value

	return clone
}

func (s *snailFish) magnitude() int {
	if s.left == nil && s.right == nil {
		return s.value
	}

	sum := 0

	if s.left != nil {
		sum += s.left.magnitude() * 3
	}
	if s.right != nil {
		sum += s.right.magnitude() * 2
	}

	return sum
}

func (s *snailFish) add(other *snailFish) *snailFish {
	sum := &snailFish{
		left:  s,
		right: other,
	}
	sum.left.parent = sum
	sum.right.parent = sum

	for {
		if sum.reduceExplode(sum, 0) {
			continue
		}
		if !sum.reduceSplit(sum) {
			break
		}
	}

	return sum
}

func (s *snailFish) reduceExplode(head *snailFish, depth int) bool {
	if depth == 5 {
		s.parent.explode()
		return true
	}

	if s.parent == nil && depth != 0 {
		return false
	}
	if s.left == nil {
		return false
	}

	if s.right == nil || s.left.reduceExplode(head, depth+1) {
		return true
	}

	return s.right.reduceExplode(head, depth+1)
}

func (s *snailFish) reduceSplit(head *snailFish) bool {
	if s.left == nil && s.right == nil {
		if s.value > 9 {
			s.split()
			return true
		}
		return false
	}
	if s.left == nil {
		return false
	}

	if s.right == nil || s.left.reduceSplit(head) {
		return true
	}

	return s.right.reduceSplit(head)
}

func (s *snailFish) split() {
	s.left = &snailFish{
		parent: s,
		value:  s.value / 2,
	}
	s.right = &snailFish{
		parent: s,
		value:  (s.value / 2) + (s.value % 2),
	}
	s.value = 0
}

func (s *snailFish) explode() {
	left := s.findAnyLeftValue()
	if left != nil {
		left.value += s.left.value
	}

	right := s.findAnyRightValue()
	if right != nil {
		right.value += s.right.value
	}

	s.left.parent = nil
	s.left = nil
	s.right.parent = nil
	s.right = nil
	s.value = 0
}

func (s *snailFish) findAnyLeftValue() *snailFish {
	prev := s
	for number := s.parent; number != nil; {
		if number.right == prev {
			prev = number
			number = number.left
			continue
		}
		if number.left == prev {
			if s.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.right
	}

	return nil
}

func (s *snailFish) findAnyRightValue() *snailFish {
	prev := s
	for number := s.parent; number != nil; {
		if number.left == prev {
			prev = number
			number = number.right
			continue
		}
		if number.right == prev {
			if s.parent == nil {
				break
			}
			prev = number
			number = number.parent
			continue
		}
		if number.left == nil && number.right == nil {
			return number
		}
		number = number.left
	}

	return nil
}

func parseSnailFish(line string) *snailFish {
	var err error
	currentSnailFish := &snailFish{}

	for _, symbol := range line {
		switch symbol {
		case '[':
			currentSnailFish.left = &snailFish{parent: currentSnailFish}
			currentSnailFish.right = &snailFish{parent: currentSnailFish}
			currentSnailFish = currentSnailFish.left
		case ',':
			currentSnailFish = currentSnailFish.parent.right
		case ']':
			currentSnailFish = currentSnailFish.parent
		default:
			currentSnailFish.value, err = strconv.Atoi(string(symbol))
			panic.Check(err)
		}
	}

	return currentSnailFish
}
