package test

import (
	"errors"

	icd "github.com/dh1tw/goSoundbench/icd"
	"github.com/dh1tw/goSoundbench/sound"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

// DeserializeTestCases deserializes a protocol buffer containing
// testcases and returns a slice of Testcases fullfilling the
// TestcaseI interface
func DeserializeTestCases(data []byte) ([]TestcaseI, error) {

	pbTCs := icd.TestCases{} // protocol buffer (pb) TestCases
	if err := proto.Unmarshal(data, &pbTCs); err != nil {
		return nil, err
	}

	var tcs []TestcaseI

	// iterate through the pb sine test cases and copy the data in the
	// corresponing structures within the application
	if pbTCs.SineTestCases != nil {
		// iterate over all sine test cases in the protocol buffer structure
		for _, pbSTC := range pbTCs.GetSineTestCases() {
			STC := SineTestCase{}
			STC.ID = pbSTC.GetId()
			STC.Duration = pbSTC.GetDuration()
			// iterate over all sine wave test elements within the pb sine test case
			if pbSTC.Elements != nil {
				for _, pbSTE := range pbSTC.GetElements() {
					STE, err := createSineTestElement(pbSTE)
					if err != nil {
						return nil, err
					}
					STC.SineTestElements = append(STC.SineTestElements, &STE)
				}
			}
			tcs = append(tcs, &STC)
		}
	}
	// return interface containing all sine test cases
	return tcs, nil
}

// createSineTestElement creates a new SineTestElement from a protocol buffer.
// The function is mainly used during the deserialization
func createSineTestElement(pbSTE *icd.SineTestElement) (SineTestElement, error) {
	var STE = SineTestElement{}
	var direction int
	var samplingrate float64 = 48000 //default value
	var bufferlength int
	var deviceName string
	chs := map[int]sound.Channel{}

	mapping := viper.GetStringMapString("topic-device-mapping")

	if _, ok := mapping[pbSTE.GetChannel()]; ok {
		deviceName = mapping[pbSTE.GetChannel()]
	}

	switch pbSTE.GetDirection() {
	case icd.DIRECTIONS_INPUT:
		direction = sound.INPUT
		bufferlength = 24000
	case icd.DIRECTIONS_OUTPUT:
		direction = sound.OUTPUT
		bufferlength = 1024
	}

	if pbSTE.Samplingrate != nil {
		samplingrate = float64(pbSTE.GetSamplingrate())
	}

	if pbSTE.Bufferlength != nil {
		bufferlength = int(pbSTE.GetBufferlength())
	}

	if pbSTE.AudioChannels != nil {
		pChs := pbSTE.GetAudioChannels()
		for _, pCh := range pChs {
			chID, ch, err := createChannel(pCh)
			if err != nil {
				return STE, err
			}
			chs[chID] = ch
		}
	}

	if direction == sound.OUTPUT {

		// create a muted LEFT channel (writing silence), otherwise
		// portaudio would old use one channel and consider it the
		// left channel
		if _, ok := chs[sound.LEFT]; !ok {
			chs[sound.LEFT] = sound.Channel{
				AudioChId: sound.LEFT,
				Tones:     []sound.Tone{sound.Tone{0.0, 0.0}},
			}
		}

		out := sound.SinusOut{}
		if _, ok := chs[sound.LEFT]; ok {
			out.TonesL = chs[sound.LEFT].Tones
		}
		if _, ok := chs[sound.RIGHT]; ok {
			out.TonesR = chs[sound.RIGHT].Tones
		}
		out.Setup(samplingrate)
		STE.AudioStream.Out = &out
	}

	if direction == sound.INPUT {
		in := sound.Recorder{}
		STE.AudioStream.In = &in
	}

	STE.AudioStream.DeviceName = deviceName
	STE.AudioStream.Direction = direction
	STE.AudioStream.Samplingrate = samplingrate
	STE.AudioStream.FramesPerBuffer = bufferlength
	STE.AudioStream.Channels = chs

	return STE, nil
}

func createChannel(pbCH *icd.AudioChannel) (int, sound.Channel, error) {
	ch := sound.Channel{}
	var chID int

	switch pbCH.GetAchid() {
	case icd.ACHID_LEFT:
		chID = sound.LEFT
	case icd.ACHID_RIGHT:
		chID = sound.RIGHT
	default:
		return 0, ch, errors.New("Unknown Audio Channel Id")
	}

	pTones := pbCH.GetTones()
	for _, pTone := range pTones {
		tone, err := createTone(pTone)
		if err != nil {
			return 0, ch, err
		}
		ch.Tones = append(ch.Tones, tone)
	}

	return chID, ch, nil
}

func createTone(pbTone *icd.Tone) (sound.Tone, error) {

	var tone = sound.Tone{}

	if pbTone.Frequency != nil {
		tone.Frequency = float64(pbTone.GetFrequency())
	}

	if pbTone.Amplitude != nil {
		tone.Amplitude = float64(pbTone.GetAmplitude())
	}

	return tone, nil
}
