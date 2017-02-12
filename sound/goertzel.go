package sound

import (
	"errors"
	"math"
)

// CalcGoertzel calculates the power for a given frequency in a MonoSample. The Goertzel
// Algorithm requires much less CPU cycles than calculating the sprectrum power density
// through an FFT. For more details on the Goertzel Filter check out:
// https://courses.cs.washington.edu/courses/cse466/12au/calendar/Goertzel-EETimes.pdf
func CalcGoertzel(freq float64, samplingRate float64, samples MonoSamples) float64 {
	var k int
	var floatnumSamples float64
	var omega, sine, cosine, coeff, q0, q1, q2, magnitude, real, imag float64

	data := samples.toFloat64()
	numSamples := len(data)
	scalingFactor := float64(numSamples / 2.0)

	floatnumSamples = float64(numSamples)

	k = int(0.5 + (floatnumSamples * freq / samplingRate))

	omega = (2.0 * math.Pi * float64(k)) / floatnumSamples

	sine = math.Sin(float64(omega))
	cosine = math.Cos(float64(omega))
	coeff = 2.0 * cosine
	q0 = 0
	q1 = 0
	q2 = 0

	for _, d := range data {
		q0 = coeff*q1 - q2 + float64(d)
		q2 = q1
		q1 = q0
	}

	real = (q1 - q2*cosine) / scalingFactor
	imag = (q2 * sine) / scalingFactor

	magnitude = 10 * math.Log10(math.Sqrt(float64(real*real+imag*imag)))

	return magnitude
}

// findMaxValuePosition iterates over a []float64 slice an returns
// the maximum value and the corresponding position in the slice
func findMaxValuePosition(in []float64) (int, float64) {

	var maxValue float64
	var maxPosition int

	maxValue = in[0]
	maxPosition = 0
	for i, el := range in {
		if el > maxValue {
			maxValue = el
			maxPosition = i
		}
	}
	return maxPosition, 10 * math.Log10(maxValue)
}

// CalculateAverage calculates the mean value over a slice of float64 values
func CalculateAverage(slice []float64) (float64, error) {

	if len(slice) == 0 {
		return 0.0, errors.New("Result array empty")
	}

	var result float64

	for _, el := range slice {
		result += el
	}

	avg := result / float64(len(slice))

	return avg, nil
}

// toFloat64 converts a []Sample to []float64
func (in *MonoSamples) toFloat64() []float64 {

	out := make([]float64, len(*in))
	for i, el := range *in {
		out[i] = float64(el)
	}

	return out
}
