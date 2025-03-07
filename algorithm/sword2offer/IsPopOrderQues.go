package sword2offer

func IsPopOrder(pushA, popA []int) bool {
	index := 0

	stack := make([]int, 0)
	for i := 0; i < len(pushA); i++ {
		stack = append(stack, pushA[i])
		for len(stack) > 0 && stack[len(stack)-1] == popA[index] {
			stack = stack[:len(stack)-1]
			index++
		}
	}
	return index == len(popA)
}
