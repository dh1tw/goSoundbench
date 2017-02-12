package mqttserver

// func DeserializeTestResults(data []byte) (sb.TestResult, error) {

// 	var tr sb.TestResult

// 	p := icd.TestResults{}
// 	if err := proto.Unmarshal(data, &p); err != nil {
// 		return tr, err
// 	}

// 	if p.SineTestResult != nil {
// 		pStr := p.GetSineTestResult()
// 		sTr := sb.SineTestResult{}
// 		sTr.Id = pStr.GetId()

// 		for _, pSter := range pStr.GetResults() {
// 			sTer := sb.SineTestElementResult{}
// 			if pSter.Channel != nil {
// 				sTer.DeviceName = pSter.GetChannel()
// 			}

// 			if len(pSter.AudioChannels) > 0 {
// 				for _, pAudioCh := range pSter.GetAudioChannels() {
// 					audioCh := sb.Channel{}

// 					if pAudioCh.Achid != nil {
// 						if pAudioCh.GetAchid() == icd.ACHID_LEFT {
// 							audioCh.AudioChId = sb.LEFT
// 						} else if pAudioCh.GetAchid() == icd.ACHID_RIGHT {
// 							audioCh.AudioChId = sb.RIGHT
// 						} else {
// 							return tr, errors.New("unknown Audio Channel")
// 						}
// 					}

// 					for _, pTone := range pAudioCh.GetTones() {
// 						tone := sb.Tone{}
// 						if pTone.Frequency != nil {
// 							tone.Frequency = float64(pTone.GetFrequency())
// 						}
// 						if pTone.Amplitude != nil {
// 							tone.Amplitude = float64(pTone.GetAmplitude())
// 						}
// 						audioCh.Tones = append(audioCh.Tones, tone)
// 					}

// 					sTer.Channels = append(sTer.Channels, audioCh)
// 				}
// 			}
// 			sTr.Results = append(sTr.Results, sTer)
// 		}

// 		tr = &sTr
// 	}
// 	return tr, nil
// }
