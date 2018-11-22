package main

import (
	"fmt"
	"learn/runner"
	"os"
	"time"
)

const timeout = 1 * time.Second

func main() {
	fmt.Println("Start work.")

	// 为本次执行分配超时时间
	r := runner.New(timeout)

	// 加入要执行的任务
	r.Add(createTask(), createTask(), createTask())

	// 执行任务并处理结果
	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			fmt.Println("Terminating due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			fmt.Println("Terminating due to interrupt.")
			os.Exit(2)
		}
	}
}

func createTask() func(int) {
	return func(id int) {
		fmt.Printf("Processor - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
