// 支持线程安全的链表实现

package goq

import "sync"

type lister interface {
	length() int
	empty() bool
	pushBack(em interface{})
	popFront() (em interface{})
}

// 切片实现法
type slist struct {
	m      *sync.RWMutex
	Values []interface{}
}

func newSList() *slist {
	return &slist{m: &sync.RWMutex{}}
}

func (l *slist) length() int {
	l.m.RLock()
	defer l.m.RUnlock()

	return len(l.Values)
}

func (l *slist) empty() bool {
	return (l.length() == 0)
}

func (l *slist) pushBack(em interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	l.Values = append(l.Values, em)
}

func (l *slist) popFront() (em interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	if len(l.Values) > 0 {
		em = l.Values[0]
		l.Values = l.Values[1:]
	}
	return em
}

// 链表实现法
type listSNode struct {
	v interface{}
	n *listSNode
}

func newNode(v interface{}) *listSNode {
	return &listSNode{v: v}
}

type clist struct {
	l int
	h *listSNode
	t *listSNode
	m *sync.RWMutex
}

func newCList() *clist {
	return &clist{m: &sync.RWMutex{}}
}

func (l *clist) length() int {
	l.m.RLock()
	defer l.m.RUnlock()

	return l.l
}

func (l *clist) empty() bool {
	l.m.RLock()
	defer l.m.RUnlock()

	return (l.l == 0)
}

func (l *clist) pushBack(em interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.h == nil {
		l.h = newNode(em)
		l.t = l.h
	} else {
		l.t.n = newNode(em)
		l.t = l.t.n
	}
	l.l++
}

func (l *clist) popFront() (em interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.h != nil {
		em = l.h.v
		l.h = l.h.n
		l.l--
	}
	return em
}
