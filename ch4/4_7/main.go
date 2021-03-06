package main

import "fmt"

func main() {
	s := "abcdef"
	fmt.Println(s)
	fmt.Println(string(reverse([]byte(s))))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
