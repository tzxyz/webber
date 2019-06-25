package webber

import (
	"container/list"
	"sync"
)

type Scheduler interface {
	Push(request *Request)
	Poll() *Request
}

type QueueScheduler struct {
	queue *list.List
	lock  sync.Mutex
}

func (s *QueueScheduler) Push(request *Request) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.queue.PushBack(request)
	logger.Debug("QueueScheduler Push Url: ", request.url)
}

func (s *QueueScheduler) Poll() *Request {
	s.lock.Lock()
	defer s.lock.Unlock()
	e := s.queue.Front()
	if e == nil {
		return nil
	}
	s.queue.Remove(e)
	request := e.Value.(*Request)
	logger.Debug("QueueScheduler Poll Url: ", request.url)
	return request
}

var InMemoryScheduler = &QueueScheduler{queue: list.New()}
