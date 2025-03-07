package sword2offer

func VerifySquenceOfBST(sequence []int) bool {
	length := len(sequence)
	if length == 0 {
		return false
	}
	i := 0
	for length > 0 {
		length--
		for i < length && sequence[i] < sequence[length] {
			i++
		}
		for i < length && sequence[i] > sequence[length] {
			i++
		}
		if i < length {
			return false
		}
		i = 0
	}
	return true
}
