package leetcodetop100

func findAnagrams(str string, p string) []int {
	var hash [26]int
	for i := 0; i < len(p); i++ {
		hash[p[i]-'a']++
	}
	left := 0
	right := 0
	var result []int
	var tempHash [26]int
	for ; left <= right && right < len(str); right++ {
		index := str[right] - 'a'
		tempHash[index]++

		if tempHash[index] > hash[index] {
			for ; left <= right && tempHash[index] > hash[index]; left++ {
				tempIndex := str[left] - 'a'
				tempHash[tempIndex]--
			}
		}
		if right-left+1 == len(p) {
			result = append(result, left)
		}
	}
	return result
}
