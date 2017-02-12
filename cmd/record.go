// Copyright © 2016 Tobias Wellnitz, DH1TW <Tobias.Wellnitz@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dh1tw/goSoundbench/sound"
	"github.com/gordonklaus/portaudio"
	"github.com/spf13/cobra"
	wav "github.com/youpy/go-wav"
)

// recordCmd represents the record command
var recordCmd = &cobra.Command{
	Use:   "record",
	Short: "Record audio to File",
	Long:  `record audio from a sound card into a file`,
	Run: func(cmd *cobra.Command, args []string) {
		recordToFile()
	},
}
var filename string

func init() {
	RootCmd.AddCommand(recordCmd)
	recordCmd.Flags().StringVarP(&deviceName, "device", "D", "default", "Input device")
	recordCmd.Flags().Float64VarP(&samplingrate, "samplingrate", "s", 48000, "Sampling rate for the input device")
	recordCmd.Flags().StringVarP(&channels, "channels", "c", "both", "Input channels (left, right, both)")
	recordCmd.Flags().DurationVarP(&duration, "duration", "d", 5e9, "Duration in seconds")
	recordCmd.Flags().StringVarP(&filename, "filename", "f", "recording.wav", "Name of the recording file")
	recordCmd.Flags().IntVarP(&frames, "frames", "F", 1024, "Frames per Buffer")
}

func recordToFile() {

	as := sound.AudioStream{}
	as.Channels = make(map[int]sound.Channel)
	as.DeviceName = deviceName
	as.Samplingrate = samplingrate
	as.FramesPerBuffer = frames
	as.Direction = sound.INPUT

	switch strings.ToUpper(channels) {
	case "LEFT":
		as.Channels[sound.LEFT] = sound.Channel{AudioChId: sound.LEFT}
	case "RIGHT":
		as.Channels[sound.RIGHT] = sound.Channel{AudioChId: sound.RIGHT}
	case "BOTH":
		as.Channels[sound.LEFT] = sound.Channel{AudioChId: sound.LEFT}
		as.Channels[sound.RIGHT] = sound.Channel{AudioChId: sound.RIGHT}
	default:
		fmt.Println("Unknown Channel:", channels)
		os.Exit(-1)
	}

	// get actual default audio output device
	if deviceName == "default" {
		d, err := portaudio.DefaultInputDevice()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		as.DeviceName = d.Name
	} else {
		as.DeviceName = deviceName
	}

	in := sound.Recorder{}
	as.In = &in

	if err := as.Initialize(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := as.Start(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// Since we record with an async callback, we have to sleep in the
	// mean routine for the time of the recording
	t1 := time.Now()
	for time.Now().Sub(t1) < duration {
		time.Sleep(100 * time.Millisecond)
	}

	if err := as.Stop(); err != nil {
		fmt.Println(err)
	}

	as.Close()

	// make sure our file ends with .wav
	if !strings.HasSuffix(filename, ".wav") {
		filename += ".wav"
	}

	// create the file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer f.Close()

	recordedData := as.In.GetData()
	recLen := len(recordedData)

	// setup a wav.Writer which will correctly set the wav file header
	writer := wav.NewWriter(f,
		uint32(recLen*as.FramesPerBuffer), // numSamples
		uint16(len(as.Channels)),          // numChannels
		uint32(as.Samplingrate),           // samplingrate
		uint16(16))                        // bits per sample

	wavSamples := make([]wav.Sample, 0, recLen*as.FramesPerBuffer)

	// interlace the data if we recorded stereo
	for i := 0; i < recLen; i++ {
		for j := 0; j < as.FramesPerBuffer; j++ {
			wavSample := wav.Sample{}
			wavSample.Values[0] = int(float32(recordedData[i][0][j]) * 32768)
			if len(as.Channels) == 2 {
				wavSample.Values[1] = int(float32(recordedData[i][1][j]) * 32768)
			}
			wavSamples = append(wavSamples, wavSample)
		}
	}

	// write the interlaced wav samples into the file
	err = writer.WriteSamples(wavSamples)
	if err != nil {
		fmt.Println(err)
	}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
