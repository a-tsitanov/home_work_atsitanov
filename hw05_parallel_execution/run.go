package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

type ErrorInfo struct {
	errorMutex   sync.Mutex
	errorMsg     error
	ErrorCounter int
	ErrorStop    int
}

func (e *ErrorInfo) IncrementErrorCount() {
	e.ErrorCounter++
	if e.ErrorCounter == e.ErrorStop {
		e.errorMsg = ErrErrorsLimitExceeded
	}
}

func Worker(queueTask chan Task, wg *sync.WaitGroup, errorMut *ErrorInfo) {
	defer wg.Done()

	for {
		work, ok := <-queueTask

		if ok == false {
			return
		}

		err := work()
		if err != nil {
			errorMut.errorMutex.Lock()
			if errorMut.errorMsg != nil {
				errorMut.errorMutex.Unlock()
				return
			}
			errorMut.IncrementErrorCount()
			errorMut.errorMutex.Unlock()
			continue
		}
	}
}

func InitQueue(queueTask chan Task, tasks []Task) {
	for _, task := range tasks {
		queueTask <- task
	}
	close(queueTask)
}

func Run(tasks []Task, n, m int) error {

	queueTask := make(chan Task, len(tasks))
	e := &ErrorInfo{ErrorCounter: 0, ErrorStop: m}
	var wg sync.WaitGroup

	InitQueue(queueTask, tasks)

	for workerNum := 0; workerNum < n; workerNum++ {
		wg.Add(1)
		go Worker(queueTask, &wg, e)
	}

	wg.Wait()
	return e.errorMsg
}
