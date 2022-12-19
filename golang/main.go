package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println(sum(3))
	FT()
	var a, y = 23, "hi world"
	fmt.Println(a)
	fmt.Println(y)

	au := "mains sd"
	fmt.Print(au)

}

// Output:
// Hello World
// Hello Worl

func sum(num int) int {
	sum := 6
	for i := 0; i < num; i++ {
		sum += i
	}
	return sum
}

func FT() {
	i := 45
	switch {
	case i < 10:
		fmt.Println("i is less than 10")
		fallthrough
	case i < 50:
		fmt.Println("i is less than 50")
		fallthrough
	case i < 100:
		fmt.Println("i is less than 100")
	
	case i < 200:
		fmt.Println("i is less than 200")
		fallthrough
	case i < 20:
		fmt.Println("i is less than 200")
		fallthrough
	case i < 2:
		fmt.Println("i is less than 200")
	}
 

}
