package sound

import "math/rand"

type NoiseOut struct {
	Left  bool
	Right bool
}

// Process generates a random noise
// func (no *NoiseOut) Process(out StereoSamples) {
func (no *NoiseOut) Process(out interface{}) {

	data := out.(StereoSamples)

	for i := range data[0] {
		if no.Left {
			data[0][i] = Sample(rand.Float32())
		}
		if no.Right {
			data[1][i] = Sample(rand.Float32())
		}
	}

	out = data
}
