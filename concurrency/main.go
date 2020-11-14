package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	var (
		totalWorker int
		outputDir   string
		source      string
	)

	flag.StringVar(&outputDir, "output", "export", "directory for exported csv, default: \"export/\"")
	flag.IntVar(&totalWorker, "concurrent_limit", 2, "number of workers, default: 2")
	flag.StringVar(&source, "jobs", "jobs.json", "job list, default: jobs.json")

	flag.Parse()

	// JobQueue .
	jobQueue := make(chan Job)

	// create workers pool.
	pool := NewWorkerPool(totalWorker, jobQueue, myJobHandler)
	pool.Run()

	outputDir = normalizePath(outputDir)
	dirInfo, err := os.Stat(outputDir)
	if err != nil && os.IsNotExist(err) {
		log.Fatalf("directory %s does not exists", outputDir)
	}

	if !dirInfo.IsDir() {
		log.Fatal(err)
	}

	jobs, err := createJobs(source, outputDir)
	if err != nil {
		log.Fatal(err)
	}

	// send jobs to queue.
	for _, job := range jobs {
		jobQueue <- *job
	}

	// waiting for interrupt or SIGTERM.
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	<-shutdown
	// run graceful shutdown.
	// wait for all workers to finishing their job.
	log.Println("main: shutdown signal received")
	pool.Close()
	log.Println("main: all workers closed")
}

func createJobs(source, outputDir string) ([]*Job, error) {
	var jobs []*Job

	file, err := os.OpenFile(source, os.O_RDONLY, 0666)
	if err != nil {
		return jobs, err
	}

	if err := json.NewDecoder(file).Decode(&jobs); err != nil {
		return jobs, err
	}

	for _, job := range jobs {
		job.Dir = outputDir
	}

	return jobs, nil
}

func normalizePath(path string) string {
	// normalize trailing slash.
	path = strings.TrimRight(path, string(filepath.Separator))

	return path + "/"
}
