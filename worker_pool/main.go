package main

import (
	"fmt"
	"sync"
	"time"
)

// Task Types of this interface are handled by worker pool
type Task interface {
	Process(id int)
}

type IdTask struct {
	id int
}

func (task *IdTask) Process(id int) {
	fmt.Printf("thread {%d} is processing task with id {%d}\n", id, task.id)
	time.Sleep(2 * time.Second)
}

type WorkerPool struct {
	concurrency int
	taskQueue   chan Task
	wg          sync.WaitGroup
}

func NewWorkerPool(concurrency int) *WorkerPool {
	return &WorkerPool{
		concurrency: concurrency,
		taskQueue:   make(chan Task),
		wg:          sync.WaitGroup{},
	}
}

func (p *WorkerPool) Start() {
	for i := 0; i < p.concurrency; i++ {
		go func(i int) {
			defer p.wg.Done()
			for task := range p.taskQueue {
				task.Process(i)
			}
		}(i)
	}
	p.wg.Add(p.concurrency)
}

func (p *WorkerPool) Submit(task Task) {
	p.taskQueue <- task
}

func (p *WorkerPool) Wait() {
	close(p.taskQueue)
	p.wg.Wait()
}

func main() {
	pool := NewWorkerPool(3)
	pool.Start()
	for i := 1; i <= 15; i++ {
		task := &IdTask{i}
		pool.Submit(task)
	}
	pool.Wait()
}
