package programs

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func PrintOneToTenWithTenGoRoutines() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {

		go func(j int) {
			ch <- j
		}(i)

	}

	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
}

func ProducerConsumerOfIntegers() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(j int) {
			ch <- j
		}(i)
	}
	receiver(ch)
	close(ch)
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

func FibonacciSeries(n int) {
	ch := make(chan int, n)

	generateFib(ch, n)

	for val := range ch {
		fmt.Println(val)
	}

}

func generateFib(ch chan int, n int) {
	a := 0
	b := 1
	ch <- a
	ch <- b
	c := a + b
	for i := 2; i < n; i++ {
		ch <- c
		a = b
		b = c
		c = a + b
	}
	close(ch)
}

func SumOfArrayWithFourGoRoutines() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	l1 := len(arr)
	iter := l1 / 4
	ch := make(chan int)
	for i := 0; i < l1; i = i + iter {
		go sumArr(ch, arr[i:i+iter])
	}
	finalSum := 0
	for i := 0; i < 4; i++ {
		finalSum += <-ch
	}
	close(ch)
	fmt.Println(finalSum)
}

func sumArr(ch chan int, arr []int) {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	ch <- sum
}

type Person struct {
	Label        string
	RightFork    bool
	LeftFork     bool
	LastConsumed bool
	Last         *Person
	Next         *Person
}

func createCircularArrangment(n int) *Person {
	var p1 *Person
	var lp *Person
	var lastNode *Person
	for i := 0; i < n; i++ {
		obj := Person{
			Label:     strconv.Itoa(i + 1),
			RightFork: true,
			LeftFork:  true,
			Last:      lp,
			Next:      nil,
		}
		if lastNode != nil {
			lastNode.Next = &obj
		}
		lp = &obj
		if i == 0 {
			p1 = &obj
		}
		lastNode = &obj
		if i == n-1 {
			p1.Last = &obj
			lastNode.Next = p1
		}

	}
	return p1
}

func printCircularList(head *Person) {
	p1 := head
	for p1.Next != head {
		fmt.Print(p1.Label, " -> ")
		p1 = p1.Next
	}
	fmt.Print(p1.Label, " -> Cyclic point \n")

}
func DiningPhilosophersProblem() {
	p1 := createCircularArrangment(5)
	printCircularList(p1)
	ch := make(chan struct{})
	go foodProducer(ch)
	go foodConsumer(ch, p1)
	go foodConsumer(ch, p1.Next.Next)
	time.Sleep(10 * time.Second)
}

func foodConsumer(ch <-chan struct{}, head *Person) {
	p1 := head
	for {
		if p1.LeftFork && p1.RightFork {
			p1.LeftFork = false
			p1.RightFork = false
			p1.Last.RightFork = false
			p1.Next.LeftFork = false
			time.Sleep(10)
			<-ch
			fmt.Println("Consumed by : ", p1.Label)
			p1.Last.RightFork = true
			p1.Next.LeftFork = true
			p1.LeftFork = true
			p1.RightFork = true
			p1 = p1.Next
		} else {
			p1 = p1.Next
		}
	}

}

func foodProducer(ch chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
	}
}

func CheckpointSynchronization() {
	var wg sync.WaitGroup
	workers := []string{"sagar", "sanket", "shweta", "snehal"}
	noOfTasks := 3
	for i := 1; i <= noOfTasks; i++ {
		fmt.Printf("Task number: %v starts \n", i)
		wg.Add(len(workers))
		for _, worker := range workers {
			go work(worker, &wg)
		}
		wg.Wait()
		fmt.Printf("Task number: %v completed \n", i)
	}

}

func work(worker string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(worker, " begins his/her work")
	time.Sleep(10)
	fmt.Println(worker, " completed his/her work")
}

func ProducerConsumerProblem() {
	msgs := make(chan int, 5)
	done := make(chan bool)
	go produce(msgs, done)
	go consume(msgs)
	<-done
	fmt.Println("all items are produced and consumed")
}
func produce(msgs chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		msgs <- i
		fmt.Println("Produced: ", i)

	}
	done <- true
}

func consume(msgs chan int) {
	for {
		val := <-msgs
		fmt.Println("consumed: ", val)
	}

}
