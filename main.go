package main

import (
	"fmt"
	"sync"
)

type writeRequest struct {
	key   string
	value string
}

type readRequest struct {
	key   string
	value chan string
}

func randomSleep() {
	//time.Sleep(time.Duration(rand.Intn(10)+10) * time.Millisecond)
}

func createMap() map[string]string {
	return make(map[string]string, 1000)
}

func dedicatedWorker(reqw <-chan writeRequest, reqr <-chan readRequest, stop <-chan struct{}) {
	m := createMap()

	for {
		select {
		case w := <-reqw:
			m[w.key] = w.value
		case r := <-reqr:
			r.value <- m[r.key]
			close(r.value)
		case <-stop:
			return
		}
	}
}

func dedicated(wClients, rClients, chBufSize int) {
	write := make(chan writeRequest, chBufSize)
	read := make(chan readRequest, chBufSize)
	stop := make(chan struct{})

	go dedicatedWorker(write, read, stop)

	var wg sync.WaitGroup
	wg.Add(wClients + rClients)

	for i := 0; i < wClients; i++ {
		go func() {
			defer wg.Done()
			randomSleep()
			write <- writeRequest{
				key:   fmt.Sprintf("key%d", i),
				value: fmt.Sprintf("value%d", i),
			}
		}()
	}

	for i := 0; i < rClients; i++ {
		go func() {
			defer wg.Done()
			randomSleep()
			rr := readRequest{
				key:   fmt.Sprintf("key0"),
				value: make(chan string, 1),
			}
			read <- rr
			_ = <-rr.value
		}()
	}

	wg.Wait()
	close(stop)
}

func synchro(wClients, rClients int) {
	m := createMap()
	var mutex sync.Mutex

	var wg sync.WaitGroup
	wg.Add(wClients + rClients)

	for i := 0; i < wClients; i++ {
		go func() {
			defer wg.Done()
			randomSleep()
			mutex.Lock()
			defer mutex.Unlock()
			m[fmt.Sprintf("key%d", i)] = fmt.Sprintf("value%d", i)
		}()
	}

	for i := 0; i < rClients; i++ {
		go func() {
			defer wg.Done()
			randomSleep()
			mutex.Lock()
			defer mutex.Unlock()
			_ = m[fmt.Sprintf("key0")]
		}()
	}

	wg.Wait()
}

func main() {
	fmt.Println("Main doesn't do anything. Run\ngo test -bench-.")
}
