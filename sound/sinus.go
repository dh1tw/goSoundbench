package sound

import "math"

type Tones []Tone

type Tone struct {
	Frequency float64
	Amplitude float64
}

type SinusOut struct {
	TonesL Tones
	TonesR Tones
	StepL  float64
	StepR  float64
	phaseL float64
	phaseR float64
}

// Process generates a sinewave
// func (so *SinusOut) Process(out StereoSamples) {
func (so *SinusOut) Process(out interface{}) {

	data := out.(StereoSamples)

	for i := range data[0] {
		if len(so.TonesL) > 0 {
			data[0][i] = Sample(so.TonesL[0].Amplitude * math.Sin(2*math.Pi*so.phaseL))
			_, so.phaseL = math.Modf(so.phaseL + so.StepL)
		}
		if len(so.TonesR) > 0 {
			data[1][i] = Sample(so.TonesR[0].Amplitude * math.Sin(2*math.Pi*so.phaseR))
			_, so.phaseR = math.Modf(so.phaseR + so.StepR)
		}
	}

	out = data
}

func (so *SinusOut) Setup(samplingrate float64) {
	if len(so.TonesL) > 0 {
		so.StepL = so.TonesL[0].Frequency / samplingrate
	}
	if len(so.TonesR) > 0 {
		so.StepR = so.TonesR[0].Frequency / samplingrate
	}
}
