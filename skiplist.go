package skiplist

import (
	"math/rand"
	"time"
)

const (
	p        = 0.25
	maxLevel = 32
)

var rd = rand.New(rand.NewSource(time.Now().Unix()))

//SkipList define of skiplist
type SkipList struct {
	head [maxLevel]*SkipListNode
}

type SkipListNode struct {
	next  []*SkipListNode
	level int

	data int
}

//New make a new skiplist
func New() *SkipList {
	return &SkipList{}
}

//Add add value to skiplist
func (s *SkipList) Add(value int) {
	level := randomLevel()
	newNode := &SkipListNode{
		level: level,
		data:  value,
		next:  make([]*SkipListNode, level+1),
	}
	for i := 0; i <= level; i++ {
		node := s.head[i]
		if node == nil {
			s.head[i] = newNode
		} else if node.data > value {
			//change head node
			newNode.next[i] = node
			s.head[i] = newNode
		} else {
			for node.next[i] != nil && node.next[i].data < value {
				node = node.next[i]
			}
			newNode.next[i] = node.next[i]
			node.next[i] = newNode
		}
	}
}

//Find find value in skiplist
func (s *SkipList) Find(value int) (node *SkipListNode) {
	level := maxLevel
	for node == nil && level > 0 {
		level--
		if s.head[level] != nil && s.head[level].data <= value {
			node = s.head[level]
		}
	}

walk:
	for level > -1 {
		if node == nil {
			level--
			goto walk
		}

		if node.data == value {
			return node
		}

		if node.next[level] != nil && node.next[level].data <= value {
			node = node.next[level]
			goto walk
		}

		level--
	}
	return nil
}

//Delete delete value from s
func (s *SkipList) Delete(value int) {
	level := maxLevel
	for level > 0 {
		level--
		node := s.head[level]

		//delete head
		if node != nil && node.data == value {
			s.head[level] = s.head[level].next[level]
			continue
		}

		for node != nil && node.next[level] != nil {
			if node.next[level].data == value {
				node.next[level] = node.next[level].next[level]
				break
			}
			node = node.next[level]
		}
	}
}

//randomLevel
func randomLevel() int {
	level := 0
	for rd.Float64() < p && level < maxLevel {
		level++
	}
	return level
}
