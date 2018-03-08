package webber

import (
	"sync"
	"container/list"
	log "github.com/sirupsen/logrus"
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
	log.Debug("QueueScheduler Push Url: ", request.url)
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
	log.Debug("QueueScheduler Poll Url: ", request.url)
	return request
}
