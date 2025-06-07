/*
Вы разрабатываете систему для загрузки файлов с удалённого сервера. Каждый файл загружается через отдельное соединение.
Однако, сервер ограничивает количество одновременно активных соединений, и ваше приложение не должно превышать этот
лимит. В программе нельзя просто запускать фиксированное количество горутин. Нужно обеспечить динамически изменяемое
ограничение количество одновременно работающих соединений.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	connectionsLimit = 3
)

func downloadFile(url string) {
	fmt.Printf("Downloading file from %s\n", url)
	time.Sleep(1 * time.Second)
	fmt.Printf("File from %s downloaded\n", url)
}

func main() {
	files := []string{
		"https://www.google.com",
		"https://www.yandex.ru",
		"https://www.amazon.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.twitter.com",
		"https://www.linkedin.com",
		"https://www.reddit.com",
		"https://www.tiktok.com",
		"https://www.netflix.com",
	}

	wg := sync.WaitGroup{}

	semaphore := make(chan struct{}, connectionsLimit)

	wg.Add(len(files))

	for _, file := range files {
		semaphore <- struct{}{}
		go func() {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			downloadFile(file)
		}()
	}

	wg.Wait()
}
