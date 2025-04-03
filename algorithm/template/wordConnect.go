package template

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func connectWord(words []string, startWord string, endWord string) int {
	if words == nil || len(words) == 0 {
		return 0
	}
	var allWord = make(map[string]bool)
	for _, word := range words {
		allWord[word] = true
	}
	visitedWord := make(map[string]bool)
	queue := make([]string, 0)
	result := 1
	queue = append(queue, startWord)
	queue = append(queue, "")
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		if word == "" {
			if len(queue) > 0 {
				result++
				queue = append(queue, "")
			}
			continue
		}
		arr := []rune(word)
		for i := 0; i < len(arr); i++ {
			old := arr[i]
			for j := 'a'; j <= 'z'; j++ {
				arr[i] = j
				newWord := string(arr)
				if allWord[newWord] && !visitedWord[newWord] {
					visitedWord[newWord] = true
					queue = append(queue, newWord)
					if newWord == endWord {
						return result + 1
					}
				}
			}
			arr[i] = old
		}
	}
	return 0
}

func TestConnectWord() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	input := strings.Split(scanner.Text(), " ")
	beginStr := input[0]
	endStr := input[1]

	wordList := []string{beginStr, endStr}
	for i := 0; i < n; i++ {
		scanner.Scan()
		wordList = append(wordList, scanner.Text())
	}
	count := connectWord(wordList, beginStr, endStr)
	fmt.Println(count)
}
