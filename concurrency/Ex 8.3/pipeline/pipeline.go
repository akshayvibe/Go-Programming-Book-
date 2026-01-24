/*package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)
	go func() {
		for x:=0;x<100;x++{
			natural <- x
		}
		}()
		close(natural)
	go func() {
		// if we do this "<-natural" then it will only accept one data and the pipeline will be closed
			 for x:=range natural{
			squares <- x * x
			 }
		close(squares)
	}()
		for x:=range squares{
			fmt.Println(x)
		}

}
*/

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
