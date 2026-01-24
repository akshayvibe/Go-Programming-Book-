// /*package main

// import "fmt"

/* NOTE - for x := range order {}
			is actually 
		for {
				x, ok := <-order
				if !ok {
					break
					}
				fmt.Printf("Served coffee no. %d\n", x)
			}


*/

// func main() {
// 	natural := make(chan int)
// 	squares := make(chan int)
// 	go func() {
// 		for x:=0;x<100;x++{
// 			natural <- x
// 		}
// 		}()
// 		close(natural)
// 	go func() {
// 		// if we do this "<-natural" then it will only accept one data and the pipeline will be closed
// 			 for x:=range natural{
// 			squares <- x * x
// 			 }
// 		close(squares)
// 	}()
// 		for x:=range squares{
// 			fmt.Println(x)
// 		}

// }
// */

// ANCHOR - Unidirectional go lang
package main

import "fmt"

func naturals(out chan<- int) {
	for x := range 11{
		out <- x
	}
	close(out)
}
func square(out chan<- int, in <- chan int) {
	// if we do this "<-natural" then it will only accept one data and the pipeline will be closed
	for x := range in {
		out <- x * x
	}
	close(out)
}
func print(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	natural := make(chan int)
	squares := make(chan int)
	go naturals(natural)
	go square(squares, natural)
	print(squares)
}


//ANCHOR - Buffered and Unbuffered Concept
// package main

// import (
// 	"fmt"
// 	// "time"
// )
// func main(){
// 	order:=make(chan int,3)//buffered
//ANCHOR - // order:=make(chan int) //unbuffered
// 	go func ()  {
// 		for x:=range 11{
// 			order<-x
// 			fmt.Printf("Customer order coffe %d\n",x)
// 		}
// 		close(order)
// 	}()
// 		for x:=range order{
// 			fmt.Printf("Served coffee no. %d\n",x)
// 			// fmt.Printf("Coffee no. %d Server\n",x)
// 		}
// }