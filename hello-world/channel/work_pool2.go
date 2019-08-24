package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id        int
	randomint int
}
type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func sumofdigits(randomnum int) int { // 每位数字相加
	sum := 0
	num := randomnum
	for num != 0 {
		no := num % 10 // 个位数
		sum = sum + no
		num /= 10 // 移除个位数
	}
	time.Sleep(time.Second * 2)
	return sum
}

func worker(group *sync.WaitGroup) {
	for job := range jobs {
		out := Result{job, sumofdigits(job.randomint)}
		results <- out
	}
	group.Done()
}

func createworkers(worker_num int) {
	var group sync.WaitGroup
	for i := 0; i < worker_num; i++ {
		group.Add(1)
		go worker(&group)
	}
	group.Wait()
}

func createjobs(numofjob int) {
	for i := 0; i < numofjob; i++ {
		var job = Job{i, rand.Intn(999)}
		jobs <- job
		fmt.Printf("create the %d job.\n", i)
	}
	close(jobs) // close the channel, and from now on, it can't write data to this channel, you can read it till the channel is empty.
}

func printresult(isdone chan bool) {
	for result := range results {
		fmt.Printf("job id %d is done, random integer is %d, result is %d\n", result.job.id, result.job.randomint, result.sumofdigits)
	}
	isdone <- true
}

func main() {
	isdone := make(chan bool)
	numofjob := 100
	go createjobs(numofjob) // create 100 jobs
	go printresult(isdone)
	createworkers(10) // create the worker pool and wait the jobs is submitted.
	close(results)
	<-isdone
	fmt.Println("jobs done.")
}
