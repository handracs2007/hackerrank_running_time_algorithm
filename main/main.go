// https://www.hackerrank.com/challenges/runningtime/problem

package main

import "fmt"

func runningTime(arr []int32) int32 {
	var swapCount int32

	for i := 1; i < len(arr); i++ {
		var temp = arr[i]

		for j := i - 1; j >= 0; j-- {
			if temp < arr[j] {
				arr[j+1] = arr[j]

				if j == 0 {
					arr[0] = temp
				}

				swapCount++
			} else {
				if j < i-1 {
					arr[j+1] = temp
				}
				break
			}
		}
	}

	return swapCount
}

func main() {
	fmt.Println(runningTime([]int32{2, 1, 3, 1, 2}))
}
