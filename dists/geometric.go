package dists

/*
   Created using
   https://en.wikipedia.org/wiki/Normal_distribution
*/

import (
	"math"

	"github.com/deathly809/gostats"
)

// Geometric holds all information related to a normal distribution
type Geometric struct {
	p        float64
	mean     float64
	median   float64
	stddev   float64
	variance float64
}

// PDF is the probability density function.  If the distributions
// is discrete then this is the probability mass function
func (n *Geometric) PDF(x ...float64) float64 {
	return n.p * math.Pow(1-n.p, x[0]-1)
}

// CDF is the cumulative density function
func (n *Geometric) CDF(x ...float64) float64 {
	return 1 - math.Pow(1-n.p, x[0])
}

// PMF just calls PDF
func (n *Geometric) PMF(x ...float64) float64 {
	return n.PDF(x...)
}

// Mean is the mean of the distribution
func (n *Geometric) Mean() float64 {
	return n.mean
}

// Median is the median of the distribution
func (n *Geometric) Median() float64 {
	return n.mean
}

// Mode is always 1
func (n *Geometric) Mode() float64 {
	return 1
}

// Variance is the variance of the distribution
func (n *Geometric) Variance() float64 {
	return n.variance
}

// StdDev is the variance of the distribution
func (n *Geometric) StdDev() float64 {
	return n.stddev
}

// Dimension should return 1
func (n *Geometric) Dimension() int {
	return 1
}

// Sample the distribtion
func (n *Geometric) Sample() []float64 {
	return nil
}

// NewGeometric creates a new geometric distribution
func NewGeometric(p float64) gostats.Distribution {
	return &Geometric{
		p:        p,
		mean:     1 / p,
		median:   math.Ceil(-1 / math.Log2(1-p)),
		variance: (1 - p) / (p * p),
		stddev:   math.Sqrt(1-p) / p,
	}
}
