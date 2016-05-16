package gostats


// Distribution interface for generic distributions  
type Distribution interface {
    
    // Cumulative Density Function
    CDF(...float64) float64
    // Probability Mass Function
    PMF(...float64) float64
    // Probability Density Function
    PDF(...float64) float64
    
    // Mean of the distribution
    Mean() float64
    // Median of the distribution
    Median() float64    
    // Mode of the distribution
    Mode() float64
    // Variance of the distribution
    Variance() float64
    // Standard Deviation of the distribution
    StdDev() float64
    
    // Dimensionality of the data
    Dimension() int
    
    // Sample from this distribution
    Sample() []float64
}
