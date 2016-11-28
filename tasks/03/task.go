package main

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Execute(int) (int, error)
}

// 1. Pipeline

type Pipe struct {
	tasks []Task
}

func (p *Pipe) Execute(i int) (int, error) {
	if len(p.tasks) == 0 {
		return 0, fmt.Errorf("Pipelord is not pleased with your lack of tasks!")
	}
	result := i
	var err error
	for _, task := range p.tasks {
		if result, err = task.Execute(result); err != nil {
			return 0, err
		}
	}
	return result, err
}

func Pipeline(tasks ...Task) Task {
	t := &Pipe{}
	t.tasks = tasks
	return t
}

// 2. Fast

type Fast struct {
	tasks []Task
}

type channelResult struct {
	err    error
	result int
}

func (f *Fast) Execute(i int) (int, error) {
	if len(f.tasks) == 0 {
		return 0, fmt.Errorf("Fastlord is not pleased with your lack of tasks!")
	}
	c := make(chan *channelResult)
	var once sync.Once

	for _, task := range f.tasks {
		go func(task Task) {
			result, err := task.Execute(i)
			once.Do(func() {
				c <- &channelResult{result: result, err: err}
			})
		}(task)
	}
	cr := <-c
	return cr.result, cr.err
}

func Fastest(tasks ...Task) Task {
	return &Fast{tasks}
}

// 3. Timed

type Timey struct {
	task    Task
	timeout time.Duration
}

func (t *Timey) Execute(i int) (int, error) {
	c := make(chan *channelResult)
	go func(task Task) {
		result, err := task.Execute(i)
		c <- &channelResult{result: result, err: err}
	}(t.task)

	select {
	case cr := <-c:
		return cr.result, cr.err
	case <-time.After(t.timeout):
		return 0, fmt.Errorf("Timelord is not pleased with your dallying!")
	}
}

func Timed(task Task, timeout time.Duration) Task {
	return &Timey{task, timeout}
}

// 4. ConcurentMapReduce

type MapReduce struct {
	tasks  []Task
	reduce func(results []int) int
}

func (mr *MapReduce) Execute(i int) (int, error) {
	if len(mr.tasks) == 0 {
		return 0, fmt.Errorf("Maplord is not reduced by your lack of tasks!")
	}
	return 0, nil
}

func ConcurrentMapReduce(reduce func(results []int) int, tasks ...Task) Task {
	return &MapReduce{tasks, reduce}
}

type GreatestSearch struct {
	errorLimit int
	tasks      <-chan Task
}

func GreatestSearcher(errorLimit int, tasks <-chan Task) Task {
	return &GreatestSearch{errorLimit, tasks}
}

func (gs *GreatestSearch) Execute(i int) (int, error) {
	return 0, nil
}
