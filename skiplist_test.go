package skiplist

import (
	"fmt"
	"testing"
)

func TestRandomLevel(t *testing.T) {
	max := 0
	for i := 0; i < 1000000; i++ {
		x := randomLevel()
		if x > max {
			max = x
		}
	}
	fmt.Println(max)
}

func TestAdd(t *testing.T) {
	sl := New()
	sl.Add(50)
	sl.Add(20)
	sl.Add(60)
	sl.Add(5)
	sl.Add(55)

	n := sl.head[0]
	for n != nil {
		fmt.Println(n.data)
		n = n.next[0]
	}
}

func TestFind(t *testing.T) {
	sl := New()
	sl.Add(50)
	sl.Add(20)
	sl.Add(60)
	sl.Add(5)
	sl.Add(55)

	if sl.Find(1) != nil {
		t.Error()
	}

	if sl.Find(5).data != 5 {
		t.Error()
	}

	if sl.Find(20).data != 20 {
		t.Error()
	}

	if sl.Find(60).data != 60 {
		t.Error()
	}
}

func TestDelete(t *testing.T) {
	sl := New()
	sl.Add(50)
	sl.Add(20)
	sl.Add(60)
	sl.Add(5)
	sl.Add(55)

	sl.Delete(55)
	sl.Delete(5)
	sl.Delete(60)

	n := sl.head[0]
	for n != nil {
		fmt.Println(n.data)
		n = n.next[0]
	}
}
