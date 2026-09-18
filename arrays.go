package main
import "fmt"

func main(){
	//initiallyif array is int->0 all zero values,string->"",bool->false
	// var nums[4]int
	// nums[0]=1
	// fmt.Println(nums)
	//fmt.Println(len(nums))

	// nums:=[3]int{1,2,3}
	// fmt.Println(nums)

	//2d arrays
	nums:=[2][2]int{{3,4},{5,6}}
	fmt.Println(nums)

	//benefits of arrays
	//fixed size that is predictable
	//memory optimatization
	//constant time access
}
