package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

// plus the digits of number
func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(time.Second * 2)
	return sum
}

// create a worker
func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

// create a work pool
func createWorkPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

// allocate the job
func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		random := rand.Intn(999)
		job := Job{i, random}
		jobs <- job
	}
	close(jobs)
}

// print the result
func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func main() {
	startime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkPool(noOfWorkers)
	<-done
	endtime := time.Now()
	diff := endtime.Sub(startime)
	fmt.Println("total time taken ", diff, " seconds.")
}
