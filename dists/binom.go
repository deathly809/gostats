package dists

/*
   Binomial Distribution

   Used to model number of success in a sample of size n drawn with replacement.
   If non-replacement is needed please use the Hypergeometric distribution.

*/

import (
	"fmt"
	"math"
	"math/big"
	"math/rand"

	"github.com/deathly809/gostats"
	"github.com/deathly809/gostats/prob"
)

type dist struct {
	q, p, n float64
}

func exp(b, e float64) *big.Float {
	base := big.NewFloat(b)
	exponent := int(e)
	result := big.NewFloat(1)
	for exponent > 0 {
		tmp := new(big.Float)
		if exponent%2 == 1 {
			tmp.Mul(result, base)
			result = tmp
		}
		tmp = new(big.Float)
		tmp.Mul(base, base)
		base = tmp
		exponent /= 2
	}
	return result
}

func (d *dist) PMF(val ...float64) float64 {
	k := val[0]

	nChooseK := prob.NChooseK(int(d.n), int(k))
	happens := exp(d.p, k)
	doesntHappen := exp(d.q, d.n-k)

	nCk, _, _ := big.ParseFloat(nChooseK.String(), 10, 0, big.AwayFromZero)

	tmp := big.NewFloat(0)
	tmp.Mul(nCk, happens)
	tmp.Mul(tmp, doesntHappen)
	result, _ := tmp.Float64()

	if result > 1.0 {
		fmtString := "(PMF) Probability is greater than 1.0 : %f, with \n\tn = %f\n\tp=%f\n\tq=%f\n\tk = %f\n\tnChoosek = %s\n\tsuccess = %s\n\tfailure = %s"
		errMsg := fmt.Sprintf(fmtString,
			result,
			d.n,
			d.p,
			d.q,
			k,
			nChooseK.String(),
			happens.String(),
			doesntHappen.String(),
		)

		panic(errMsg)
	}

	return result
}

func (d *dist) PDF(val ...float64) float64 {
	return d.PMF(val...)
}

// Cumulative Density Function
func (d *dist) CDF(val ...float64) float64 {
	k := val[0]

	if k < 0.0 {
		return 0.0
	}

	if k >= d.n {
		return 1.0
	}

	result := 0.0
	for i := 0.0; i <= k; i++ {
		result += d.PMF(i)
	}
	return result
}

// Mean of the distribution
func (d *dist) Mean() float64 {
	return d.p * d.n
}

// Median of the distribution
func (d *dist) Median() float64 {
	return math.Floor(d.n * d.p)
}

// Mode of the distribution
func (d *dist) Mode() float64 {
	return math.Floor((d.n + 1) * d.p)
}

// Variance of the distribution
func (d *dist) Variance() float64 {
	return d.Mean() * (1 - d.p)
}

func (d *dist) StdDev() float64 {
	return math.Sqrt(d.Variance())
}

// Dimensionality of the data
func (d *dist) Dimension() int {
	return 1
}

// Sample from this distribution
func (d *dist) Sample() []float64 {
	k := 0.0
	prob := rand.Float64()

	result := 0.0
	for ; k < d.n; k++ {
		result += d.PMF(k)
		if result > prob {
			break
		}
	}

	return []float64{k}
}

// Binomial returns a binomial distribution with the parameters
// provided.
//
//  @p - probability of success
//  @n - sample size
//
func Binomial(p, n float64) gostats.Distribution {
	return &dist{
		p: p,
		q: 1 - p,
		n: n,
	}
}
