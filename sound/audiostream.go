package sound

import (
	"fmt"

	"github.com/gordonklaus/portaudio"
)

const (
	INPUT  = 1
	OUTPUT = 2
)

const (
	MONO   = 1
	STEREO = 2
)

const (
	LEFT  = 1
	RIGHT = 2
)

type Tones []Tone

type Tone struct {
	Frequency float64
	Amplitude float64
}

type Channel struct {
	AudioChId int
	Tones     Tones
}

var ChName = map[int]string{
	LEFT:  "LEFT",
	RIGHT: "RIGHT",
}

var ChValue = map[string]int{
	"LEFT":  LEFT,
	"RIGHT": RIGHT,
}

type Sample float32
type MonoSamples []Sample
type StereoSamples [][]Sample

// AudioStream contains all data necessary to play/record data from an
// audio device (sound card)
type AudioStream struct {
	DeviceName      string
	Direction       int
	Samplingrate    float64
	Channels        map[int]Channel
	FramesPerBuffer int
	device          *portaudio.DeviceInfo
	stream          *portaudio.Stream
	Out             PlayI
	In              RecordI
}

// PlayI is the interface that must be implemented to play audio on a local
// soundcard
type PlayI interface {
	// Process(StereoSamples)
	Process(interface{})
}

// RecordI is the interface that must be implemented to record from a local
// soundcard
type RecordI interface {
	Process(interface{})
	GetData() []StereoSamples
}

// Initialize the Audiostream and open it
func (as *AudioStream) Initialize() error {
	if err := as.identifyDevice(); err != nil {
		return err
	}
	if err := as.createStream(); err != nil {
		return err
	}

	return nil
}

// Start the Audiostream
func (as *AudioStream) Start() error {
	if err := as.stream.Start(); err != nil {
		return err
	}

	return nil
}

// Stop the Audiostream
func (as *AudioStream) Stop() error {
	if err := as.stream.Stop(); err != nil {
		return err
	}

	return nil
}

// Close the Audiostream
func (as *AudioStream) Close() {
	if as.stream != nil {
		as.stream.Close()
	}
}

// identifyDevice checks if the Audio Devices actually exist
func (as *AudioStream) identifyDevice() error {
	devices, _ := portaudio.Devices()
	for _, device := range devices {
		if device.Name == as.DeviceName {
			as.device = device
			return nil
		}
	}
	return &SoundbenchError{EDEVICE, as.DeviceName, ""}
}

// createStream creates an input or output audio stream with the attributes
// set in the *AudioStream struct.
func (as *AudioStream) createStream() error {
	var err error

	// accept only 1 or 2 channels
	if len(as.Channels) > 2 || len(as.Channels) <= 0 {
		fmt.Println("as.Channels", as.Channels)
		return &SoundbenchError{ECHANNEL, as.DeviceName, ""}
	}

	// setup AudioStream for Play
	if as.Direction == OUTPUT {

		streamDeviceParm := portaudio.StreamDeviceParameters{Device: as.device, Channels: len(as.Channels)}

		var sp portaudio.StreamParameters

		sp = portaudio.StreamParameters{
			Output:          streamDeviceParm,
			FramesPerBuffer: as.FramesPerBuffer,
			SampleRate:      as.Samplingrate,
		}

		as.stream, err = portaudio.OpenStream(sp, as.writeAudioStreamCb)
		if err != nil {
			return err
		}

		// setup AudioStream for Record
	} else if as.Direction == INPUT {

		streamDeviceParm := portaudio.StreamDeviceParameters{Device: as.device, Channels: len(as.Channels)}
		sp := portaudio.StreamParameters{
			Input:           streamDeviceParm,
			FramesPerBuffer: as.FramesPerBuffer,
			SampleRate:      as.Samplingrate,
		}

		as.stream, err = portaudio.OpenStream(sp, as.readAudioStreamCb)
		if err != nil {
			return err
		}
	}
	return nil
}

// readAudioStreamCb is the callback to handle Data which is being read from
// an input audio stream (Record)
func (as *AudioStream) readAudioStreamCb(in StereoSamples) {
	as.In.Process(in)
}

// writeAudioStreamCb is the callback to handle Data which will be written
// to an output audio stream (Play)
func (as *AudioStream) writeAudioStreamCb(out StereoSamples) {
	as.Out.Process(out)
}

// record copies the stereo samples read from the sound card into the data slice of the TestElement
func (as *AudioStream) record(in StereoSamples) {
	// a deep copy is necessary, since portaudio reuses the slice "in"
	buf := make(StereoSamples, len(in))
	for i, v := range in {
		buf[i] = append([]Sample(nil), v...)
	}

	as.In.Process(buf)
}
