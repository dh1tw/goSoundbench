package test

import (
	"errors"

	icd "github.com/dh1tw/goSoundbench/icd"
	"github.com/dh1tw/goSoundbench/sound"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/viper"
)

// Serialize the results into a protocol buffers byte array
func (STR *SineTestResult) Serialize() ([]byte, error) {
	id := STR.ID

	pbSTR := icd.SineTestResult{}
	pbSTR.Id = &id

	for _, res := range STR.Results {

		pbSTER := icd.SineTestElementResults{}

		for _, ch := range res.Channels {

			pbCH := icd.AudioChannel{}
			var pbAchid icd.ACHID

			if ch.AudioChId == sound.LEFT {
				pbAchid = icd.ACHID_LEFT
			} else if ch.AudioChId == sound.RIGHT {
				pbAchid = icd.ACHID_RIGHT
			} else {
				return []byte{}, errors.New("unknown Audio Channel Id")
			}

			pbCH.Achid = &pbAchid

			for _, tone := range ch.Tones {

				frequency := float32(tone.Frequency)
				amplitude := float32(tone.Amplitude)

				pTone := icd.Tone{
					Frequency: &frequency,
					Amplitude: &amplitude,
				}

				pbCH.Tones = append(pbCH.Tones, &pTone)
			}

			pbSTER.AudioChannels = append(pbSTER.AudioChannels, &pbCH)
		}

		mapping := viper.GetStringMapString("topic-device-mapping")

		if len(mapping) == 0 {
			return nil, errors.New("topic device map is empty; Check config file")
		}

		reverseMap := reverseMap(mapping)

		if _, ok := reverseMap[res.DeviceName]; ok {
			deviceName := mapping[res.DeviceName]
			pbSTER.Channel = &deviceName
		} else {
			return []byte{}, errors.New("could not find device mapping")
		}

		pbSTR.Results = append(pbSTR.Results, &pbSTER)
	}

	pbTR := icd.TestResults{
		SineTestResult: &pbSTR,
	}

	data, err := proto.Marshal(&pbTR)

	if err != nil {
		return nil, err
	}

	return data, nil
}

// reverseMap returns a map with switched keys - values
func reverseMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[v] = k
	}
	return n
}
