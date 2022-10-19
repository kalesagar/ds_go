package programs

import "fmt"

//input: k=3, str=deeedbbcccbdaa
//output: str=aa
//1. ddbbbdaa
//2. dddaa
//3. aa
func RemoveKAdjecentchars(str string, k int) string {
	if k >= len(str) {
		fmt.Println(str)
		return str
	}

	return str

}
