package main

import (
	"ds_go/programs"
	"fmt"
)

const (
	abc = "ABC"
)

func main() {
	fmt.Println("Started Main go routine....")

	//-------------------Linked List------------------------
	// llist := linkedlist.CreateNode(10)
	// llist.InsertAtLast(15)
	// llist = llist.InsertAtFirst(5)
	// llist.InsertAtLast(20)
	// llist = llist.InsertAtFirst(1)
	// llist.InsertAtLast(30)
	// llist.InsertAtLast(40)
	// llist.PrintList()
	// llist = llist.ReverseLinkedListNonRecursive()
	// llist.PrintList()
	// llist = llist.DeleteNodeFromLinkedList(15)
	// llist.PrintList()
	// llist.PrintLinnkedListInReverseOrder()

	// llist1 := linkedlist.CreateNode(11)
	// llist1.InsertAtLast(17)
	// llist1 = llist1.InsertAtFirst(5)
	// llist1.InsertAtLast(22)
	// llist1 = llist1.InsertAtFirst(1)
	// llist1 = llist1.InsertAtFirst(1)
	// llist1.InsertAtLast(33)
	// llist1.InsertAtLast(5)
	// llist1.InsertAtLast(45)
	// llist1.InsertAtLast(55)
	// llist1.InsertAtLast(55)
	// llist1.PrintList()
	//llist1 = llist1.DeleteDuplicateNodesFromUnSortedLinkedList()
	// llist1 = llist1.DeleteDuplicateNodesFromUnSortedLinkedList()
	// llist1.PrintList()
	// fmt.Println(linkedlist.CompareTwoLinkedList(llist, llist1))
	// fmt.Println()
	// newLL := linkedlist.MergeSortedLinkedLists(llist, llist1)
	// newLL.PrintList()
	// newLL.InsertAtLast(55)
	// newLL.PrintList()
	// newLL.DeleteDuplicateNodesFromSortedLinkedList()
	// newLL.PrintList()
	// *** create cycle in linked list ***
	// tmp := newLL.InsertAtLast(56)
	// tmp.Next = newLL.Next
	// fmt.Println(newLL.CycleDetectionInLinkedList())
	// *** removing cycle from linked list ***
	// tmp.Next = nil
	// fmt.Println(newLL.CycleDetectionInLinkedList())
	// ***
	// fmt.Println(llist.GetLength())
	// fmt.Println(llist.GetMiddleOfList().Data)
	// if kthlast := llist.FindKthLast(2); kthlast != nil {
	// 	fmt.Println(kthlast.Data)
	// }
	//-------------Queue---------------
	// q := queue.CreateQueue()
	// fmt.Println(q.IsEmptyQ())
	// q.Enque(15)
	// q.Enque(20)
	// q.Enque(30)
	// q.Front.PrintList()
	// q.Enque(40)
	// q.Enque(50)
	// if dequed := q.Deque(); dequed != nil {
	// 	fmt.Println(dequed.Data)
	// }
	// if dequed := q.Deque(); dequed != nil {
	// 	fmt.Println(dequed.Data)
	// }
	// front := q.GetFront()
	// if front != nil {
	// 	fmt.Println(front.Data)
	// }
	//---------program FindLargestNumber-----------
	// largeNo := programs.LargeNumber([]int{78,66,57})
	// fmt.Println(largeNo)

	// programs.PairsOfSum([]int{10, 20, 5, 15, 7, 8}, 25)

	//------ print even and odd numbers using two go routines

	//programs.PrintOddNumbersUsingGoRoutines()

	//----- Check string is pangram ----
	// fmt.Println(programs.IsPangram("The quick brown fox jumps over the lazy dog")) //true
	// fmt.Println(programs.IsPangram("The quick brown fox jumps over the dog")) //false
	//--- MissingCharacterToPangramString ---
	// programs.MissingCharacterToPangramString("The quick brown fox jumps over the lazy dog")
	// programs.MissingCharacterToPangramString("The quick brown fox jumps over the dog")
	//--- Remove Punctuations From String ---
	//programs.RemovePunctuationsFromString("%welcome' to @geeksforgeek<s")
	// fmt.Println(programs.FindVowels("sagar pandurang Kale"))
	// fmt.Println(programs.FindVowels("Yograj Shisode"))
	//fmt.Println(programs.MergeSlices([]string{"potatoes","tomatoes","brinjal"}, []string{"oranges","apples"}))
	// fmt.Println(programs.CheckPanlindrome("abacedbce"))
	// fmt.Println(programs.CheckPanlindrome("xxxxzzy"))
	// fmt.Println(programs.CheckPanlindrome("abab"))
	// fmt.Println(programs.CheckPanlindrome("aaa"))
	// fmt.Println(programs.CheckPanlindrome("sdhjk"))
	//fmt.Println(programs.CheckPanlindrome("ghjjkk"))
	//programs.PrintNNumbers(0, 100)

	// root := trees.CreateBinaryTree()
	// fmt.Println("Print inorder traversal tree:")
	// trees.InorderTraversal(root)
	// fmt.Println("\nPreorder traversal tree before inserting node 6:")
	// trees.PreOrderTraversal(root)
	// root = trees.InsertNode(root, 6)
	// fmt.Println("\npreorder traversal tree after inserting node 6:")
	// trees.PreOrderTraversal(root)

	// fmt.Println("\nPostOrder traversal tree:")
	// trees.PostOrderTraversal(root)

	//var arr [5]int
	//s1 := []int{1,2,3,4,5,6}

	//producer-consumer problem
	/*ch := make(chan int)

	result := make(chan int)

	go producer(ch)
	go consumer(ch, result)
	finalSum := <-result
	fmt.Println(finalSum)*/

	//Expose an sample post API with gin framework
	// router := gin.Default()

	// router.Use(CORSHandler())
	// router.GET("/ping", DefaultHandler)
	// router.POST("/makeentry", MakeEntry)
	// router.Run(":8080")

	//Start a web server using net/http
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Server Started....")
	// })

	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//programs.NonBlockingChannelOperations()
	fmt.Println(programs.LeftShiftCharsInString("DAbcefgh", 16))
	//programs.ProducerConsumerOfIntegers()
	//programs.FibonacciSeries(10)
	//programs.SumOfArrayWithFourGoRoutines()
	//programs.DiningPhilosophersProblem()
	//programs.CheckpointSynchronization()
	programs.ProducerConsumerProblem()
}

// func MakeEntry(c *gin.Context) {
// 	type user struct {
// 		Name string `json:"name"`
// 		Age  int    `json:"age"`
// 	}
// 	var userObj user
// 	fmt.Printf("NAME: %v\n", userObj.Name)
// 	err := c.BindJSON(&userObj)
// 	if err != nil {
// 		log.New(os.Stdout, err.Error(), log.Ldate)
// 	}
// 	fmt.Printf("NAME: %v\n", userObj.Name)

// 	if len(userObj.Name) != 0 {
// 		c.IndentedJSON(http.StatusOK, gin.H{"status": "success"})
// 		return
// 	} else {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "failed"})
// 	}

// }

// func DefaultHandler(c *gin.Context) {
// 	c.String(http.StatusOK, "server is running...")
// }

// func CORSHandler() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
// 		// c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// 		// c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
// 		// c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		if c.Request.Method == http.MethodOptions {
// 			c.AbortWithStatus(http.StatusOK)
// 		} else {
// 			c.Next()
// 		}
// 	}
// }

func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch, result chan int) {
	var sum int
	for {
		if val, ok := <-ch; ok {
			sum = sum + val
		} else {
			break
		}
	}
	fmt.Printf("sum: %v \n", sum)
	result <- sum
}
