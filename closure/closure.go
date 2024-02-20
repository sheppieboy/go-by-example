package main
import "fmt"

func intSeq() func() int{
	i := 0
	return func() int {
		i++
		return i
	}
}

func adderByTwo() func(num int) int{
	return func(num int) int{
		return num +2
	}
}

func main(){
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	
	newInts := intSeq()
	fmt.Println(newInts())

	add2 := adderByTwo()
	fmt.Println(add2(3))
}