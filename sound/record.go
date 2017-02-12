package sound

// Recorder holds the recorded StereoSamples
type Recorder struct {
	Samples []StereoSamples
}

// Process writes the Audio samples into Memory
func (rec *Recorder) Process(in interface{}) {

	data := in.(StereoSamples)

	buf := make(StereoSamples, len(data))
	for i, v := range data {
		buf[i] = append([]Sample(nil), v...)
	}

	rec.Samples = append(rec.Samples, buf)
}

// GetData returns the slice of recorded StereoSamples
func (rec *Recorder) GetData() []StereoSamples {
	return rec.Samples
}
