package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

// Функция для подсчета количества кратных b в диапазоне [l, r]
func countMultiples(l, r, b int) int {
	if b == 0 {
		return 0 // Если делитель равен нулю, вернуть 0 (на всякий случай)
	}
	if l > r {
		return 0
	}
	// Количество кратных b от 1 до r
	countR := r / b
	// Количество кратных b от 1 до l-1
	countL := (l - 1) / b
	// Разница даст количество кратных b в [l, r]
	return countR - countL
}

func main() {
	var in = bufio.NewReader(os.Stdin)
	var out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t) // Читаем количество тестов

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n) // Читаем размер массивов

		l := make([]int, n)
		r := make([]int, n)

		// Читаем массивы l и r
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &l[j])
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &r[j])
		}

		result := 1

		for j := 0; j < n; j++ {
			b := j + 1 // b[j] = j + 1
			// Получаем количество подходящих значений для b[j]
			count := countMultiples(l[j], r[j], b)
			result = (result * count) % MOD // Умножаем и берем по модулю
		}

		fmt.Fprintln(out, result)
	}
}
