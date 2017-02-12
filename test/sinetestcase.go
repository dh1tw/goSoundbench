package test

import (
	"log"
	"time"

	"github.com/dh1tw/goSoundbench/sound"
)

type SineTestCase struct {
	ID               string
	Duration         uint32
	SineTestElements []*SineTestElement
}

type SineTestElement struct {
	sound.AudioStream
	results []SineTestElementResult
}

// Setup initializes the PortAudio System and creates the Streams
// for all SineTestElements within a SineTestCase.
func (tc *SineTestCase) Setup() error {
	// portaudio.Initialize()

	// Identify devices and load the Device Infos
	for _, te := range tc.SineTestElements {

		err := te.AudioStream.Initialize()
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute the TestCase with all its test elements
func (tc *SineTestCase) Execute() error {

	// Start the streams in the testcase
	for _, te := range tc.SineTestElements {
		if err := te.AudioStream.Start(); err != nil {
			return err
		}
	}

	// Wait until testcase has finished
	t1 := time.Now()
	for time.Now().Sub(t1) < time.Second*time.Duration(tc.Duration) {
		time.Sleep(100 * time.Millisecond)
	}

	for _, te := range tc.SineTestElements {
		if err := te.AudioStream.Stop(); err != nil {
			return err
		}
	}

	if err := tc.getGoertzelResults(); err != nil {
		return err
	}

	return nil
}

// GetID returns the ID of the Test Case
func (tc *SineTestCase) GetID() string {
	return tc.ID
}

// GetResult returns an interface variable containing the test results
func (sTC *SineTestCase) GetResult() (TestResultI, error) {

	str := SineTestResult{}
	str.ID = sTC.GetID()
	for _, te := range sTC.SineTestElements {
		str.Results = append(str.Results, te.results...)
	}
	var testResults TestResultI

	if len(str.Results) > 0 {
		testResults = &str
		return testResults, nil
	}

	// in case it does not contain Results, return a nil value for the interface
	return nil, nil
}

// Cleanup closes the Audio streams and terminates portaudio
func (tc *SineTestCase) Cleanup() {
	for _, te := range tc.SineTestElements {
		te.AudioStream.Close()
	}
	// portaudio.Terminate()
}

// getGoertzelResults calculates the results for all TestElements in the TestCase, using the
// Goertzel Algorithm.
func (tc *SineTestCase) getGoertzelResults() error {

	for _, te := range tc.SineTestElements {

		if te.Direction == sound.INPUT {
			recordedData := te.AudioStream.In.GetData()
			if len(recordedData) == 0 {
				return &sound.SoundbenchError{sound.EINTERNAL, te.DeviceName, "No data recorded"}
			}

			if err := te.getGoertzelPerSineTestElement(); err != nil {
				return err
			}
		}
	}
	return nil
}

// getGoertzelPerSineTestElement iterates over raw the recorded samples of a SineTestElement for either the
// LEFT or RIGHT Audio Channel. The function calculates the signal magnitude for the Tones of the SineTestElement.
// The results are stored as a SineTestElementResult within the SineTestElement.
func (te *SineTestElement) getGoertzelPerSineTestElement() error {

	sTer := SineTestElementResult{}
	sTer.DeviceName = te.DeviceName

	for audioChId, audioCh := range te.AudioStream.Channels {

		tones := sound.Tones{}

		for _, tone := range audioCh.Tones {
			var magnitudes []float64
			var magnitudesCleaned []float64
			// calculate the signal magnitude for the given tone
			for _, data := range te.AudioStream.In.GetData() {
				if audioChId == sound.LEFT {
					magnitudes = append(magnitudes, sound.CalcGoertzel(tone.Frequency, te.AudioStream.Samplingrate, sound.MonoSamples(data[0])))
				} else {
					magnitudes = append(magnitudes, sound.CalcGoertzel(tone.Frequency, te.AudioStream.Samplingrate, sound.MonoSamples(data[1])))
				}
			}
			if len(magnitudes) > 5 {
				// remove the values at the beginning and end as they might not be correct
				magnitudesCleaned = magnitudes[3 : len(magnitudes)-2]
			} else {
				log.Printf("Warning: Too few recorded samples (%d)", len(magnitudes))
				magnitudesCleaned = magnitudes
			}
			avgMagnitude, err := sound.CalculateAverage(magnitudesCleaned)
			if err != nil {
				return &sound.SoundbenchError{sound.EINTERNAL, te.DeviceName, err.Error()}
			}

			tones = append(tones, sound.Tone{
				Frequency: tone.Frequency,
				Amplitude: avgMagnitude,
			})
		}

		// ignore the dummy audioCh in case just the right channel was ment
		// to be recorded
		if len(audioCh.Tones) > 0 {
			sTer.Channels = append(sTer.Channels, sound.Channel{AudioChId: audioChId, Tones: tones})
		}
	}
	te.results = append(te.results, sTer)
	return nil
}
