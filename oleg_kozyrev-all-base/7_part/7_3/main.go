/*
Вы разрабатываете сервис, который обрабатывает изображения.
Каждое изображение проходит дорогостоящую обработку, например, наложение водяного знака.
Поскольку обработка каждого изображения занимает значительное время, необходимо обрабатывать их параллельно, чтобы
ускорить процесс. Однако, чтобы избежать излишней нагрузки на систему, вы хотите ограничить количество одновременно
работающих горутин.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	numOfWorker = 3
	numOfTasks  = 10
)

type Task struct {
	ID       int64
	FileName string
}

func ProcessImage(task Task) string {
	time.Sleep(1 * time.Second)
	return fmt.Sprintf("File processed %s (Task ID %d)", task.FileName, task.ID)
}

func RunWorker(id int64, taskCh <-chan Task, resultCh chan<- string) {
	for task := range taskCh {
		fmt.Printf("Worker %d started task %d\n", id, task.ID)
		resultCh <- ProcessImage(task)
		fmt.Printf("Worker %d finished task %d\n", id, task.ID)
	}

}

func main() {
	taskCh := make(chan Task, numOfTasks)
	resultCh := make(chan string, numOfTasks)

	wg := sync.WaitGroup{}

	for i := 0; i < numOfWorker; i++ {
		wg.Add(1)

		// start worker
		go func() {
			defer wg.Done()
			RunWorker(int64(i), taskCh, resultCh)
		}()
	}

	go func() {
		for i := 0; i < numOfTasks; i++ {
			// add task
			taskCh <- Task{ID: int64(i), FileName: fmt.Sprintf("Image_%d.png", i)}
		}

		close(taskCh)
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for res := range resultCh {
		fmt.Println(res)
	}

	fmt.Println("All tasks processed!")
}
