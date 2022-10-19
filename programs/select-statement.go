package programs

import "fmt"

func NonBlockingChannelOperations() {
	ch1 := make(chan string)
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("Default case executed")
	}

	msg := "Hi there.."
	select {
	case ch1 <- msg:
		fmt.Println(msg)

	default:
		fmt.Println("Default case executed 2")
	}

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		fmt.Println("Default case executed")
	}

}
