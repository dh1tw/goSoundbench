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

	"github.com/dh1tw/goSoundbench/mqtt"
	"github.com/dh1tw/goSoundbench/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// mqttCmd represents the mqtt command
var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "MQTT Server which can execute sound test cases",
	Long:  `MQTT Server which can execute sound test cases`,
	Run: func(cmd *cobra.Command, args []string) {
		MqttServer()
	},
}

func init() {
	serverCmd.AddCommand(mqttCmd)
	mqttCmd.Flags().StringP("url", "u", "localhost", "URL of the MQTT broker")
	mqttCmd.Flags().IntP("port", "p", 1883, "Port of the MQTT broker")
	mqttCmd.Flags().Float64P("samplerate", "s", 48000, "Sampling rate")
	mqttCmd.Flags().IntP("frame-length-input", "i", 24000, "Buffer frame length for recording audio devices")
	mqttCmd.Flags().IntP("frame-length-output", "o", 1024, "Buffer frame length for playing audio devices")
	viper.BindPFlag("mqtt.broker-url", mqttCmd.Flags().Lookup("url"))
	viper.BindPFlag("mqtt.broker-port", mqttCmd.Flags().Lookup("port"))
	viper.BindPFlag("audio.samplingrate", mqttCmd.Flags().Lookup("samplingrate"))
	viper.BindPFlag("audio.frame-length-input", mqttCmd.Flags().Lookup("frame-length-input"))
	viper.BindPFlag("audio.frame-length-output", mqttCmd.Flags().Lookup("frame-length-output"))
}

func MqttServer() {

	mapping := viper.GetStringMapString("topic-device-mapping")
	if len(mapping) == 0 {
		fmt.Println("topic - device map is empty; Check config file")
		os.Exit(-1)
	}

	topics := make([]string, 0, len(mapping))

	for k := range mapping {
		topics = append(topics, k)
	}

	ms := mqttserver.MqttSettings{
		Transport:  "tcp",
		BrokerURL:  viper.GetString("mqtt.broker-url"),
		BrokerPort: viper.GetInt("mqtt.broker-port"),
		ClientID:   utils.RandSeq(10),
		Topics:     topics,
	}

	mqttserver.MqttServer(ms)
}
