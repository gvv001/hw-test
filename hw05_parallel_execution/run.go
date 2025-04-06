package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var errorsCount int32
	var wg sync.WaitGroup
	taskChan := make(chan Task)

	worker := func() {
		defer wg.Done()
		for task := range taskChan {
			if err := task(); err != nil {
				atomic.AddInt32(&errorsCount, 1)
			}
		}
	}

	// Запускаем воркеры
	for i := 0; i < n; i++ {
		wg.Add(1)
		go worker()
	}

	// передаём таски воркерам если не привысили лимит по ошибкам
	for _, task := range tasks {
		if atomic.LoadInt32(&errorsCount) >= int32(m) {
			break
		}
		taskChan <- task
	}

	close(taskChan)
	wg.Wait()

	// Если счётчик ошибок больше M, возвращаем ErrErrorsLimitExceeded
	if atomic.LoadInt32(&errorsCount) >= int32(m) {
		return ErrErrorsLimitExceeded
	}

	return nil
}
