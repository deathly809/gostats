package prob

import (
	"math/big"
	"testing"
)

func TestNChooseK(t *testing.T) {
	N := 1000
	for k := 0; k <= N; k++ {
		if NChooseK(N, k).Cmp(NChooseK(N, N-k)) != 0 {
			t.Fail()
			return
		}
	}

	expected := NChooseK(N, 0)
	if expected.Cmp(big.NewInt(1)) != 0 {
		t.Logf("Expected %d choose 0 = 1, found %s", N, expected.String())
		t.Fail()
		return
	}
}

func TestNChooseKKnownAnswers(t *testing.T) {
	Correct := []int64{
		1,   // 0
		10,  // 1
		45,  // 2
		120, // 3
		210, // 4
		252, // 5
		210, // 6
		120, // 7
		45,  // 8
		10,  // 9
		1,   // 10
	}

	N := len(Correct) - 1

	for k := 0; k <= N; k++ {
		result := NChooseK(N, k).Int64()
		if Correct[k] != result {
			t.Logf("For (%d choose %d) expected %d but found %d", N, k, Correct[k], result)
			t.Fail()
			return
		}
	}
}
