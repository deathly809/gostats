package prob

import "math/big"

// NChooseK returns n!/(k!*(n-k)!)
func NChooseK(n, k int) *big.Int {

	if k <= 0 || k >= n {
		return big.NewInt(1)
	}

	if k == 1 || k == n-1 {
		return big.NewInt(int64(n))
	}

	if k > (n - k) {
		k = n - k
	}

	if k > n {
		panic("k is greater than n")
	}

	bigC := big.NewInt(1)

	for i := 1; i <= k; i++ {
		bigC.Mul(bigC, big.NewInt(int64(n+1-i)))
		bigC.Div(bigC, big.NewInt(int64(i)))
	}
	return bigC
}
