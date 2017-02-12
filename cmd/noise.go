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
)

// noiseCmd represents the noise command
var noiseCmd = &cobra.Command{
	Use:   "noise",
	Short: "Generate random noise",
	Long:  `This function generates random noise on an audio output device`,
	Run: func(cmd *cobra.Command, args []string) {
		generateNoise()
	},
}

func init() {
	generateCmd.AddCommand(noiseCmd)
}

func generateNoise() {

	as := sound.AudioStream{}
	as.DeviceName = deviceName
	as.Samplingrate = samplingrate
	as.FramesPerBuffer = frames
	as.Direction = sound.OUTPUT

	// get actual default audio output device
	if deviceName == "default" {
		d, err := portaudio.DefaultOutputDevice()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		as.DeviceName = d.Name
	} else {
		as.DeviceName = deviceName
	}

	out := sound.NoiseOut{}
	as.Out = &out
	switch strings.ToUpper(channels) {
	case "LEFT":
		out.Left = true
		as.Channels = sound.STEREO
	case "RIGHT":
		out.Right = true
		as.Channels = sound.STEREO
	case "BOTH":
		out.Left = true
		as.Channels = sound.MONO
	}

	if err := as.Initialize(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if err := as.Start(); err != nil {
		fmt.Println(err)
	}

	t1 := time.Now()
	for time.Now().Sub(t1) < duration {
		time.Sleep(100 * time.Millisecond)
	}

	if err := as.Stop(); err != nil {
		fmt.Println(err)
	}

	as.Close()
}
