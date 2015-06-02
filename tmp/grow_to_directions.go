package transposablematrix

import (
	"fmt"
	"math/rand"
)

func max(x ...int) (ret int) {
	for i := 0; i < len(x); i++ {
		if x[i] > ret {
			ret = x[i]
		}
	}
	return
}

func grow() {

	mW, mE, mN := rand.Intn(9)+1, rand.Intn(9)+1, rand.Intn(9)+1
	maxG := max(mW, mE, mN)
	fmt.Printf("%v %v %v %v\n", mW, mE, mN, maxG)

	W, E, N := 0, 0, 0
	for i := 0; i <= maxG; i++ {
		W = incUpto(W, mW)
		searchMatch(W, E, N)
		E = incUpto(E, mE)
		searchMatch(W, E, N)
		N = incUpto(N, mN)
		searchMatch(W, E, N)
		fmt.Printf("%v %v %v\n", W, E, N)
	}
}

func searchMatch(w, e, n int) {
}

func incUpto(cur, max int) int {
	cur++
	if cur > max {
		cur = max
	}
	return cur
}
