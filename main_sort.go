package main
import (
	"fmt"
	"./algorithms"
)
func main() {
	fmt.Println("Trying bouble sorting...")
	aList := []int{8,4,2,7,1,3,5}
	fmt.Println("Original: ", aList)
	newList := algorithms.Bubble_sort_modified(aList)
	fmt.Println("Sorted: ", newList)
}