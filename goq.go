// @Title goq
// @Description 支持线程安全的消息队列
// @Author szr 2020-06-23

package goq

import (
	"sync"
)

type Queue struct {
	l lister
	d *sync.Cond
}

func NewSQueue() *Queue {
	return &Queue{l: newSList(), d: sync.NewCond(&sync.Mutex{})}
}

func NewCQueue() *Queue {
	return &Queue{l: newCList(), d: sync.NewCond(&sync.Mutex{})}
}

func (q *Queue) Size() int {
	return q.l.length()
}

func (q *Queue) Empty() bool {
	return q.l.empty()
}

func (q *Queue) PutQ(em interface{}) {
	q.d.L.Lock()
	defer q.d.L.Unlock()

	if q.Empty() {
		q.d.Signal()
	}
	q.l.pushBack(em)
}

func (q *Queue) GetQ() (em interface{}) {
	q.d.L.Lock()
	defer q.d.L.Unlock()

	for {
		if em = q.l.popFront(); em != nil {
			break
		}
		q.d.Wait()
	}
	return em
}
