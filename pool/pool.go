package pool

import "fmt"

type Task struct {
	taskId int
	f func() error
}

func NewTask(id int, f func() error) *Task {
	return &Task{
		taskId: id,
		f: f,
	}
}

func (t *Task) execute() {
	t.f()
}

type Pool struct {
	workerNum int
	EntryChan chan *Task
}

func NewPool(num int) *Pool {
	return &Pool{
		workerNum: num,
		EntryChan: make(chan *Task,num),
	}
}

func (p *Pool) worker(id int) {
	for task := range p.EntryChan {
		task.execute()
		fmt.Println("workerId:", id, "taskId:", task.taskId, "is done")
	}
}

func (p *Pool) AddTask(task *Task) {
	p.EntryChan <- task
}

func (p *Pool) Run() {
	for i := 0; i < p.workerNum; i++ {
		go p.worker(i)
	}
}

func (p *Pool) Close() {
	close(p.EntryChan)
}