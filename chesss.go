package main

import "fmt"

func main() {

	var c int
	fmt.Println("Введите размерность шахматной доски: ")
	fmt.Scan(&c)

	for i := 0; i < c; i++ {
		for j := 0; j < c; j++ {
			if (i+j)%2 == 0 {
				fmt.Print("   ")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
