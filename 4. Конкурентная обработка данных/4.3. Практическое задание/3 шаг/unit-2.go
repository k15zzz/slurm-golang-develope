package main

import (
	"fmt"
	"sync"
	"time"
)

// WorkerPool имитирует пул воркеров, опрашивающих внешнюю систему.
type WorkerPool struct {
	cond    *sync.Cond
	active  bool
	workers int
}

// NewWorkerPool создает новый пул воркеров.
func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		cond:    sync.NewCond(&sync.Mutex{}),
		active:  false,
		workers: workers,
	}
}

// Start запускает воркеры.
func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		go wp.worker(i)
	}
}

// worker представляет воркера, опрашивающего внешнюю систему.
func (wp *WorkerPool) worker(id int) {
	fmt.Printf("Воркер %d начал работу\n", id)
	for {
		wp.cond.L.Lock()
		for !wp.active {
			wp.cond.Wait()
		}
		wp.cond.L.Unlock()

		// Имитация выполнения задачи.
		fmt.Printf("Воркер %d выполняет задачу\n", id)
		time.Sleep(time.Second)

		// Имитация проверки доступности внешней системы.
		if !wp.active {
			fmt.Printf("Воркер %d приостановлен\n", id)
		}
	}
}

// CheckExternalSystem имитирует проверку доступности внешней системы.
func (wp *WorkerPool) CheckExternalSystem() {
	for {
		time.Sleep(5 * time.Second) // Имитация времени проверки.

		wp.cond.L.Lock()
		wp.active = !wp.active // Имитация изменения состояния системы.
		if wp.active {
			fmt.Println("Внешняя система доступна, уведомляем воркеров...")
			wp.cond.Broadcast()
		} else {
			fmt.Println("Внешняя система недоступна, воркеры приостановлены...")
		}
		wp.cond.L.Unlock()
	}
}

func main() {
	workers := 3
	pool := NewWorkerPool(workers)
	pool.Start()

	// Запускаем горутину для проверки доступности внешней системы.
	go pool.CheckExternalSystem()

	// Даем программе поработать некоторое время.
	time.Sleep(20 * time.Second)
}
