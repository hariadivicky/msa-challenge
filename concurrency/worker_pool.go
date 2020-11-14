package main

import (
	"log"
	"sync"
)

// Job structure.
type Job struct {
	URL      string
	Filename string
	Dir      string
}

// JobHandler defines job handler.
type JobHandler func(job Job)

// Worker .
type Worker struct {
	ID         int
	Pool       chan chan Job
	JobChannel chan Job
	shutdown   chan bool
	Handler    JobHandler
}

// NewWorker constructs a new worker.
func NewWorker(id int, pool chan chan Job, jobHandler JobHandler) Worker {
	return Worker{
		ID:         id,
		Pool:       pool,
		JobChannel: make(chan Job),
		shutdown:   make(chan bool),
		Handler:    jobHandler,
	}
}

// Listen starts worker and listening for jobs.
func (w Worker) Listen() {
	go func() {
		log.Println("worker", w.ID, ": listening")
		for {
			// registers current job channel to pool
			w.Pool <- w.JobChannel

			select {
			// receiving job
			case job := <-w.JobChannel:
				log.Println("worker", w.ID, ": receiving job", job)
				w.Handler(job)
				log.Println("worker", w.ID, ": was finished job", job)
			case <-w.shutdown:
				log.Println("worker", w.ID, ": stopped")
				return
			}
		}
	}()
}

// Close send shutdown signal to worker to stop working.
func (w Worker) Close(wg *sync.WaitGroup) {
	go func() {
		w.shutdown <- true
		wg.Done()
	}()
}

// WorkerPool .
type WorkerPool struct {
	totalWorker int
	Pool        chan chan Job
	JobQueue    chan Job
	Workers     []*Worker
	Handler     JobHandler
}

// NewWorkerPool constructs a new workers pool.
func NewWorkerPool(totalWorker int, jobQueue chan Job, jobHandler JobHandler) *WorkerPool {
	pool := make(chan chan Job, totalWorker)
	return &WorkerPool{
		totalWorker: totalWorker,
		Pool:        pool,
		JobQueue:    jobQueue,
		Workers:     make([]*Worker, 0),
		Handler:     jobHandler,
	}
}

// Run main worker pool. spawning workers, listen for new job, and dispatch to worker.
func (wp *WorkerPool) Run() {
	// spawn workers.
	for i := 0; i < wp.totalWorker; i++ {
		worker := NewWorker(i, wp.Pool, wp.Handler)
		wp.Workers = append(wp.Workers, &worker)

		worker.Listen()
	}

	// waiting for incoming job.
	go func() {
		for {
			select {
			// job received.
			case job := <-wp.JobQueue:
				go func(job Job) {
					// pull idle worker's job channel.
					// if there is no idle worker, this will block the new request and wait until another worker is idle.
					jobChannel := <-wp.Pool

					// send job to worker's job channel.
					jobChannel <- job
				}(job)
			}
		}
	}()
}

// Close workers pool.
func (wp *WorkerPool) Close() {
	var wg sync.WaitGroup
	wg.Add(len(wp.Workers))

	for _, worker := range wp.Workers {
		worker.Close(&wg)
	}

	wg.Wait()
}
