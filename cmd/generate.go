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

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a sound on an audio device",
	Long:  `Generate a sound on an audio device`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please select a command (run --help for a list of available commands)")
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.PersistentFlags().StringVarP(&deviceName, "device", "D", "default", "Output device")
	generateCmd.PersistentFlags().Float64VarP(&samplingrate, "samplingrate", "s", 48000, "Sampling rate for the output device")
	generateCmd.PersistentFlags().IntVarP(&frames, "frames", "F", 1024, "Frames per Buffer")
	generateCmd.PersistentFlags().StringVarP(&channels, "channel", "c", "both", "Output channel (left, right, both)")
	generateCmd.PersistentFlags().DurationVarP(&duration, "duration", "d", 5e9, "Duration in seconds")
}
