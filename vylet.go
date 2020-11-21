package main

import "fmt"

func max(arr []int) int {
	var max int = arr[0]
	for _, v := range arr[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	var T int
	fmt.Scan(&T)
	for i := 0; i < T; i++ {
		var N int
		var temps []int
		var min int = 60

		fmt.Scan(&N)
		temps = make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&temps[j])
		}
		for j := 0; j <= N-3; j++ {
			if m := max(temps[j : j+3]); m < min {
				min = m
			}
		}
		fmt.Println(min)
	}
}
