package leetcodetop100

import (
	"fmt"
	"testing"
)

func dailyTemperatures(temperature []int) []int {
	var result = make([]int, len(temperature))
	// result[len(temperature) - 1] = 0
	for i := len(temperature) - 2; i >= 0; i-- {
		println("the num of i is ", i)
		for j := i + 1; j < len(temperature); {
			println("the num of j is ", j)
			if temperature[j] > temperature[i] {
				result[i] = j - i
				break
			} else {
				if result[j] == 0 {
					println("no temperature higher than day ", j)
					break
				}
				j = j + result[j]
				println("the next day choose to compare temperature is ", j)
			}
		}
	}
	return result
}

func TestDailyTemperatures(t *testing.T) {
	arr := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures(arr)
	str := fmt.Sprintf("%v", result)
	println(str)
}
