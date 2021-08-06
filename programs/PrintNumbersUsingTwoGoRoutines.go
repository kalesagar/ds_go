package programs

import (
	"fmt"
	"sync"
	"time"
)

func PrintOddNumbersUsingGoRoutines() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go printEvens(ch, &wg)
	go printOdds(ch, &wg)
	ch <- struct{}{}
	wg.Wait()
	close(ch)
}

func printOdds(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i = i + 2 {
		time.Sleep(1 * time.Millisecond)
		<-ch
		fmt.Println(i)
		ch <- struct{}{}
	}

}

func printEvens(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i = i + 2 {
		<-ch
		fmt.Println(i)
		if i != 10 {
			ch <- struct{}{}
		}

	}
}
