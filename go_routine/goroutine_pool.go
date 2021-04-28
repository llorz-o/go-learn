package go_routine

import (
	"fmt"
	"math/rand"
)

type Job struct {
	Id      int
	RandNum int
}

type Result struct {
	job *Job
	sum int
}

func FuncGoroutinePool() {
	var id = 0
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	for {
		id++
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
	}
}

func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan <-chan *Job, resultChan chan<- *Result) {
			for job := range jobChan {
				rNum := job.RandNum
				var sum int
				for rNum != 0 {
					temp := rNum % 10
					sum += temp
					rNum /= 10
				}

				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}
