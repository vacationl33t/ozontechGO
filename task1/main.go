package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxSalary(s string) string {
	n := len(s)
	if n == 1 {
		return "0"
	}

	removeIndex := -1

	for i := 0; i < n-1; i++ {
		if s[i] < s[i+1] {
			removeIndex = i
			break
		}
	}

	if removeIndex == -1 {
		removeIndex = n - 1
	}

	return s[:removeIndex] + s[removeIndex+1:]
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a int
	fmt.Fscan(in, &a)
	for i := 0; i < a; i++ {
		var s string
		fmt.Fscan(in, &s)
		result := maxSalary(s)
		fmt.Fprintln(out, result)
	}
}
