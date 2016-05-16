package dists

/*
   Created using
   https://en.wikipedia.org/wiki/Uniform_distribution_(continuous)
*/

import (
	"math"
	"math/rand"

	"github.com/deathly809/gomath"
	"github.com/deathly809/gostats"
)

// Uniform holds all information about the distribution
type Uniform struct {
	a, b             float64
	pdf              float64
	mean, median     float64
	variance, stddev float64
}

// PDF will assume that the value passed in is an integer
func (u *Uniform) PDF(x ...float64) float64 {
	result := 0.0
	if x[0] >= u.a && x[0] <= u.b {
		result = u.pdf
	}
	return result
}

// CDF will assume that the value passed in is an integer
func (u *Uniform) CDF(x ...float64) float64 {
	clamped := gomath.ClampFloat64(u.a, u.b, x[0])
	return (clamped - u.a) / (u.b - u.a)
}

// PMF calls PDF
func (u *Uniform) PMF(x ...float64) float64 {
	return u.PDF(x...)
}

// Mean of the discrete uniform distribution
func (u *Uniform) Mean() float64 {
	return u.mean
}

// Median of the discrete uniform distribution
func (u *Uniform) Median() float64 {
	return u.median
}

// Mode is any value in the range, we return the mean
func (u *Uniform) Mode() float64 {
	return u.mean
}

// Variance of the discrete uniform distribution
func (u *Uniform) Variance() float64 {
	return u.variance
}

// StdDev of the discrete uniform distribution
func (u *Uniform) StdDev() float64 {
	return u.stddev
}

// Dimension is 1
func (u *Uniform) Dimension() int {
	return 1
}

// Sample the distribution
func (u *Uniform) Sample() []float64 {
	return []float64{
		u.a + rand.Float64()*(u.b-u.a),
	}
}

// NewUniform creates a continuous uniform distribution
func NewUniform(low, high int) gostats.Distribution {
	a := gomath.MinInt(low, high)
	b := gomath.MaxInt(low, high)

	return &Uniform{
		a:        float64(a),
		b:        float64(b),
		pdf:      1.0 / float64(b-a),
		mean:     0.5 * float64(a+b),
		median:   0.5 * float64(a+b),
		variance: (math.Pow(float64(b-a), 2) - 1.0) / 12.0,
		stddev:   math.Pow(float64(b-a+1), 2) / math.Sqrt(12.0),
	}
}
