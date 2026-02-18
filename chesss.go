package main

import "fmt"

func main() {
	var a [12][12]string

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			if (i+j)%2 == 0 {
				a[i][j] = "  "
			} else {
				a[i][j] = "#"
			}
		}
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 12; j++ {
			fmt.Print(a[i][j])
		}
		fmt.Println()
	}
}
