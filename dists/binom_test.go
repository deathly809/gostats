package dists

import (
	"fmt"
	"testing"
)

func TestSample(t *testing.T) {
	N := 20.0
	P := 0.7

	dist := Binomial(P, N)

	counts := make([]float64, int(N)+1)

	Samples := 10000.0
	for i := 0.0; i < Samples; i++ {
		s := dist.Sample()[0]
		counts[int(s)] += 1.0 / Samples
		if s < 0.0 || s > N {
			t.Logf("Should be between 0 and N inclusive, was %f\n", s)
			t.Fail()
			return
		}
	}

	for _, v := range counts {
		fmt.Print(v, " ")
	}
	fmt.Println()

	rows := 100.0
	for r := 0.0; r < rows; r++ {
		for c := 0; c < int(N+1); c++ {
			if counts[c] >= 1.0-(r/rows) {
				fmt.Printf("X   ")
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
	for c := 0; c < int(N+1); c++ {
		fmt.Printf("%-3d ", c)
	}
	fmt.Println()
}

func TestCDF(t *testing.T) {
	N := 100.0
	P := 0.7

	dist := Binomial(P, N)

	found := dist.CDF(-1.0)
	if found != 0.0 {
		t.Logf("CDF(-1.0) should be 0.0, but found %f", found)
		t.Fail()
		return
	}

	found = dist.CDF(N + 1)
	if found != 1.0 {
		t.Logf("CDF(%f) should be 1.0, but found %f", N+1, found)
		t.Fail()
		return
	}
}
