package programs

import "fmt"

//PairsOfSum ...
func PairsOfSum(arr []int, sum int) {
	type Pair struct {
		a, b int
	}
	var pairs []Pair
	myMap := make(map[int]bool)
	for _, value := range arr {
		if _, ok := myMap[value]; ok {
			pairs = append(pairs, Pair{value, sum - value})
		}
		val := sum - value
		myMap[val] = true
	}
	fmt.Println(pairs)
}
