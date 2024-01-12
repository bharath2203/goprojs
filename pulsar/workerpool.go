package pulsar

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type WorkerPool interface {
	HandleMessage(msg string)
}

type workerPool struct {
	totalWorkers int
	tickets      chan struct{}
	counter      int64
	mu           sync.RWMutex
}

func NewWorkerPool(workerCount int) WorkerPool {
	return &workerPool{
		totalWorkers: workerCount,
		tickets:      make(chan struct{}, workerCount),
		mu:           sync.RWMutex{},
	}
}

func (w *workerPool) HandleMessage(msg string) {
	w.tickets <- struct{}{}
	go func() {

		w.mu.Lock()
		fmt.Printf("Received Message: %s, count=%d\n", msg, w.counter)
		w.counter++
		w.mu.Unlock()

		sleepWindow := rand.Int() % 10
		fmt.Printf("Sleeping for %d seconds\n", sleepWindow)
		time.Sleep(time.Duration(sleepWindow) * time.Second)

		<-w.tickets
	}()
}
