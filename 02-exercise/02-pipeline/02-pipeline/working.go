//// generator() -> square() -> print
//
package main

//
//import (
//	"fmt"
//)
//
//func generator(nums ...int) <-chan int {
//	out := make(chan int)
//
//	go func() {
//		for _, n := range nums {
//			out <- n
//		}
//		close(out)
//	}()
//	return out
//}
//
//func square(in <-chan int) <-chan int {
//	out := make(chan int)
//	go func() {
//		for n := range in {
//			out <- n * n
//		}
//		close(out)
//	}()
//	return out
//}
//
//func merge(cs ...<-chan int) <-chan int {
//	ch := make(chan int, 2)
//	ch <- 1
//	ch <- 2
//	fmt.Println(<-ch)
//	fmt.Println(<-ch)
//
//	return ch
//}
//
//func main() {
//	in := generator(2,3)
//
//	res := merge(square(in), square(in), square(in))
//
//	fmt.Println(<- res)
//}
