//// generator() -> square() -> print
//
package main

//
//import "fmt"
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
//	res := make(chan int)
//	fmt.Println(res)
//	for _, ch := range cs {
//		for num := range ch {
//			fmt.Println(num) // this just works fine
//			res <- num // this leads to deadlock
//		}
//	}
//	close(res)
//
//	return res
//}
//
//func main() {
//	in := generator(10,20,30)
//	in2 := generator(1,2)
//
//	chs := []<-chan int{square(in), square(in2)}
//	res := merge(chs[0], chs[1])
//
//	for v := range res {
//		fmt.Println(v)
//	}
//}
