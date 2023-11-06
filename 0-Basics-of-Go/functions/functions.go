package functions

import (
	"bufio"
	"fmt"
	"os"
)

var K int = 0
var ans int = -1

func mergeSort(A *[]int, p int, r int) {
	if p < r {
		q := (p + r) >> 1
		mergeSort(A, p, q)
		mergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
}

func merge(A *[]int, p int, q int, r int) {
	i := p
	j := q + 1
	t := 0

	tmp := make([]int, r-p+1)

	for i <= q && j <= r {
		if (*A)[i] <= (*A)[j] {
			tmp[t] = (*A)[i]
			t++
			i++
		} else {
			tmp[t] = (*A)[j]
			t++
			j++
		}
	}

	for i <= q {
		tmp[t] = (*A)[i]
		t++
		i++
	}

	for j <= r {
		tmp[t] = (*A)[j]
		t++
		j++
	}

	i = p
	t = 0

	for i <= r {
		(*A)[i] = tmp[t]
		K--
		if K == 0 {
			ans = (*A)[i]
		}
		i++
		t++
	}
}

// Solves problem of BOJ 24060
func main() {
	br := bufio.NewReader(os.Stdin)
	bw := bufio.NewWriter(os.Stdout)
	defer bw.Flush()

	var N int
	fmt.Fscanln(br, &N, &K)

	var A = make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Fscan(br, &A[i])
	}

	mergeSort(&A, 0, N-1)

	fmt.Fprintln(bw, ans)
}
