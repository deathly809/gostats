package discrete

/*
   Current information is from
   https://en.wikipedia.org/wiki/Bernoulli_distribution
*/

import (
	"math"

	"github.com/deathly809/gostats"
)

// Bernoulli contains information about the distribution
type Bernoulli struct {
	p                float64
	median           float64
	variance, stddev float64
}

// PDF computes the probability density function
func (b *Bernoulli) PDF(isHead ...float64) float64 {
	result := 0.0
	switch int(isHead[0]) {
	case 0:
		result = 1 - b.p
	case 1:
		result = b.p
	}
	return result
}

// CDF computes the cumulative density function
func (b *Bernoulli) CDF(flips ...float64) float64 {
	result := 0.0
	for i := 0.0; i < flips[0]; i++ {
		result += b.PDF(i)
	}
	return result
}

// PMF probability mass function
func (b *Bernoulli) PMF(isHead ...float64) float64 {
	return b.PDF(isHead...)
}

// Mean is the most commong occuring result
func (b *Bernoulli) Mean() float64 {
	return b.p
}

// Median is the middle element
func (b *Bernoulli) Median() float64 {
	return b.median
}

// Mode returns 0, .5, or 1 depending on p
func (b *Bernoulli) Mode() float64 {
	if b.p < .5 {
		return 0
	}
	if b.p == .5 {
		return .5
	}
	return 1
}

// Variance is the skew of the data
func (b *Bernoulli) Variance() float64 {
	return b.variance
}

// StdDev is the standard deviation
func (b *Bernoulli) StdDev() float64 {
	return b.stddev
}

// Dimension is 1
func (b *Bernoulli) Dimension() int {
	return 1
}

//Sample this distribution
func (b *Bernoulli) Sample() []float64 {
	return nil
}

// NewBernoulli returns a new Bernoulli distribution
func NewBernoulli(p float64) gostats.Distribution {
	med := 0.0

	switch {
	case p == 0.5:
		med = 0.5
	case p > 0.5:
		med = 1.0
	}

	return &Bernoulli{
		p:        p,
		median:   med,
		variance: p * (1 - p),
		stddev:   math.Sqrt(p * (1 - p)),
	}
}
