package main

import (
	"github.com/fm0803/routine-pool/pool"
	"fmt"
	"time"
)

func task() error {
	fmt.Println(time.Now(), "Do something")
	return nil
}

func main() {
	p := pool.NewPool(3)
	p.Run()

	id := 0
	//go func() {
		for {
			p.AddTask(pool.NewTask(id, task))
			id++
		}
	//}()

	//p.Run()
}