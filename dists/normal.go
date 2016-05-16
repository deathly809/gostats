package dists

/*
   Created using
   https://en.wikipedia.org/wiki/Normal_distribution
*/

import (
	"math"

	"github.com/deathly809/gostats"
)

const (
	// SqrtTwoPi is the square root of 2 times Pi
	SqrtTwoPi = 2.50662827463
)

// Normal holds all information related to a normal distribution
type Normal struct {
	mean     float64
	stddev   float64
	variance float64
}

// PDF is the probability density function.  If the distributions
// is discrete then this is the probability mass function
func (n *Normal) PDF(x ...float64) float64 {
	exp := -math.Pow(x[0]-n.mean, 2) / (2 * n.variance)
	return math.Exp(exp) / (n.stddev * SqrtTwoPi)
}

// CDF is the cumulative density function
func (n *Normal) CDF(x ...float64) float64 {
	return 0.5 * (1 + math.Erf((x[0]-n.mean)/(n.stddev*math.Sqrt2)))
}

// PMF calls PDF
func (n *Normal) PMF(x ...float64) float64 {
	return n.PDF(x...)
}

// Mean is the mean of the distribution
func (n *Normal) Mean() float64 {
	return n.mean
}

// Median is the median of the distribution
func (n *Normal) Median() float64 {
	return n.mean
}

// Mode is the mean
func (n *Normal) Mode() float64 {
	return n.mean
}

// Variance is the variance of the distribution
func (n *Normal) Variance() float64 {
	return n.variance
}

// StdDev is the variance of the distribution
func (n *Normal) StdDev() float64 {
	return n.stddev
}

// Dimension is 1
func (n *Normal) Dimension() int {
	return 1
}

// Sample the distribution
func (n *Normal) Sample() []float64 {
	return nil
}

// NewNormal creates a new normal distribution
func NewNormal(mean, variance float64) gostats.Distribution {
	return &Normal{
		mean:     mean,
		variance: variance,
		stddev:   math.Sqrt(variance),
	}
}
