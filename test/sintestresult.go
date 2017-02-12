package test

import (
	"fmt"

	"github.com/dh1tw/goSoundbench/sound"
)

type SineTestResult struct {
	ID      string
	Results []SineTestElementResult
}

type SineTestElementResult struct {
	DeviceName string
	Channels   []sound.Channel
}

// GetString returns a string with the test results
func (STR *SineTestResult) String() string {
	txt := "\n"
	// only add in case the test element contains results
	if len(STR.Results) > 0 {
		txt = fmt.Sprintf("Results for Testcase: %s\n", STR.ID)
		for _, tr := range STR.Results {
			for _, audioCh := range tr.Channels {
				for _, tone := range audioCh.Tones {
					txt += fmt.Sprintf("%s (%s): Frequency: %.0f Hz, Amplitude: %.2f dB\n",
						tr.DeviceName,
						sound.ChName[audioCh.AudioChId],
						tone.Frequency,
						tone.Amplitude)
				}
			}
		}
	}
	return txt
}
