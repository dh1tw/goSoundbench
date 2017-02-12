// Code generated by protoc-gen-go.
// source: soundbench.proto
// DO NOT EDIT!

/*
Package Soundbench is a generated protocol buffer package.

It is generated from these files:
	soundbench.proto

It has these top-level messages:
	TestCases
	SineTestCase
	SineTestElement
	TestResults
	SineTestResult
	SineTestElementResults
	AudioChannel
	Tone
	Gpio
*/
package Soundbench

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type DIRECTIONS int32

const (
	DIRECTIONS_INPUT  DIRECTIONS = 1
	DIRECTIONS_OUTPUT DIRECTIONS = 2
)

var DIRECTIONS_name = map[int32]string{
	1: "INPUT",
	2: "OUTPUT",
}
var DIRECTIONS_value = map[string]int32{
	"INPUT":  1,
	"OUTPUT": 2,
}

func (x DIRECTIONS) Enum() *DIRECTIONS {
	p := new(DIRECTIONS)
	*p = x
	return p
}
func (x DIRECTIONS) String() string {
	return proto.EnumName(DIRECTIONS_name, int32(x))
}
func (x *DIRECTIONS) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(DIRECTIONS_value, data, "DIRECTIONS")
	if err != nil {
		return err
	}
	*x = DIRECTIONS(value)
	return nil
}
func (DIRECTIONS) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ACHID int32

const (
	ACHID_LEFT  ACHID = 1
	ACHID_RIGHT ACHID = 2
)

var ACHID_name = map[int32]string{
	1: "LEFT",
	2: "RIGHT",
}
var ACHID_value = map[string]int32{
	"LEFT":  1,
	"RIGHT": 2,
}

func (x ACHID) Enum() *ACHID {
	p := new(ACHID)
	*p = x
	return p
}
func (x ACHID) String() string {
	return proto.EnumName(ACHID_name, int32(x))
}
func (x *ACHID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(ACHID_value, data, "ACHID")
	if err != nil {
		return err
	}
	*x = ACHID(value)
	return nil
}
func (ACHID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type TestCases struct {
	SineTestCases    []*SineTestCase `protobuf:"bytes,1,rep,name=sineTestCases" json:"sineTestCases,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *TestCases) Reset()                    { *m = TestCases{} }
func (m *TestCases) String() string            { return proto.CompactTextString(m) }
func (*TestCases) ProtoMessage()               {}
func (*TestCases) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TestCases) GetSineTestCases() []*SineTestCase {
	if m != nil {
		return m.SineTestCases
	}
	return nil
}

type SineTestCase struct {
	Id               *string            `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Duration         *uint32            `protobuf:"varint,2,req,name=duration" json:"duration,omitempty"`
	Elements         []*SineTestElement `protobuf:"bytes,3,rep,name=elements" json:"elements,omitempty"`
	Gpio             []*Gpio            `protobuf:"bytes,4,rep,name=gpio" json:"gpio,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *SineTestCase) Reset()                    { *m = SineTestCase{} }
func (m *SineTestCase) String() string            { return proto.CompactTextString(m) }
func (*SineTestCase) ProtoMessage()               {}
func (*SineTestCase) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SineTestCase) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *SineTestCase) GetDuration() uint32 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *SineTestCase) GetElements() []*SineTestElement {
	if m != nil {
		return m.Elements
	}
	return nil
}

func (m *SineTestCase) GetGpio() []*Gpio {
	if m != nil {
		return m.Gpio
	}
	return nil
}

type SineTestElement struct {
	Channel          *string         `protobuf:"bytes,1,req,name=channel" json:"channel,omitempty"`
	Direction        *DIRECTIONS     `protobuf:"varint,2,req,name=direction,enum=Soundbench.DIRECTIONS" json:"direction,omitempty"`
	Samplingrate     *int32          `protobuf:"varint,3,opt,name=samplingrate" json:"samplingrate,omitempty"`
	Bufferlength     *int32          `protobuf:"varint,4,opt,name=bufferlength" json:"bufferlength,omitempty"`
	AudioChannels    []*AudioChannel `protobuf:"bytes,5,rep,name=audioChannels" json:"audioChannels,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SineTestElement) Reset()                    { *m = SineTestElement{} }
func (m *SineTestElement) String() string            { return proto.CompactTextString(m) }
func (*SineTestElement) ProtoMessage()               {}
func (*SineTestElement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SineTestElement) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *SineTestElement) GetDirection() DIRECTIONS {
	if m != nil && m.Direction != nil {
		return *m.Direction
	}
	return DIRECTIONS_INPUT
}

func (m *SineTestElement) GetSamplingrate() int32 {
	if m != nil && m.Samplingrate != nil {
		return *m.Samplingrate
	}
	return 0
}

func (m *SineTestElement) GetBufferlength() int32 {
	if m != nil && m.Bufferlength != nil {
		return *m.Bufferlength
	}
	return 0
}

func (m *SineTestElement) GetAudioChannels() []*AudioChannel {
	if m != nil {
		return m.AudioChannels
	}
	return nil
}

type TestResults struct {
	SineTestResult   *SineTestResult `protobuf:"bytes,1,opt,name=sineTestResult" json:"sineTestResult,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *TestResults) Reset()                    { *m = TestResults{} }
func (m *TestResults) String() string            { return proto.CompactTextString(m) }
func (*TestResults) ProtoMessage()               {}
func (*TestResults) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TestResults) GetSineTestResult() *SineTestResult {
	if m != nil {
		return m.SineTestResult
	}
	return nil
}

type SineTestResult struct {
	Id               *string                   `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Results          []*SineTestElementResults `protobuf:"bytes,2,rep,name=results" json:"results,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *SineTestResult) Reset()                    { *m = SineTestResult{} }
func (m *SineTestResult) String() string            { return proto.CompactTextString(m) }
func (*SineTestResult) ProtoMessage()               {}
func (*SineTestResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SineTestResult) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *SineTestResult) GetResults() []*SineTestElementResults {
	if m != nil {
		return m.Results
	}
	return nil
}

type SineTestElementResults struct {
	Channel          *string         `protobuf:"bytes,1,req,name=channel" json:"channel,omitempty"`
	AudioChannels    []*AudioChannel `protobuf:"bytes,2,rep,name=audioChannels" json:"audioChannels,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *SineTestElementResults) Reset()                    { *m = SineTestElementResults{} }
func (m *SineTestElementResults) String() string            { return proto.CompactTextString(m) }
func (*SineTestElementResults) ProtoMessage()               {}
func (*SineTestElementResults) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *SineTestElementResults) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *SineTestElementResults) GetAudioChannels() []*AudioChannel {
	if m != nil {
		return m.AudioChannels
	}
	return nil
}

type AudioChannel struct {
	Achid            *ACHID  `protobuf:"varint,1,req,name=achid,enum=Soundbench.ACHID" json:"achid,omitempty"`
	Tones            []*Tone `protobuf:"bytes,2,rep,name=tones" json:"tones,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AudioChannel) Reset()                    { *m = AudioChannel{} }
func (m *AudioChannel) String() string            { return proto.CompactTextString(m) }
func (*AudioChannel) ProtoMessage()               {}
func (*AudioChannel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *AudioChannel) GetAchid() ACHID {
	if m != nil && m.Achid != nil {
		return *m.Achid
	}
	return ACHID_LEFT
}

func (m *AudioChannel) GetTones() []*Tone {
	if m != nil {
		return m.Tones
	}
	return nil
}

type Tone struct {
	Frequency        *float32 `protobuf:"fixed32,1,opt,name=frequency" json:"frequency,omitempty"`
	Amplitude        *float32 `protobuf:"fixed32,2,opt,name=amplitude" json:"amplitude,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Tone) Reset()                    { *m = Tone{} }
func (m *Tone) String() string            { return proto.CompactTextString(m) }
func (*Tone) ProtoMessage()               {}
func (*Tone) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Tone) GetFrequency() float32 {
	if m != nil && m.Frequency != nil {
		return *m.Frequency
	}
	return 0
}

func (m *Tone) GetAmplitude() float32 {
	if m != nil && m.Amplitude != nil {
		return *m.Amplitude
	}
	return 0
}

type Gpio struct {
	Pin              []bool `protobuf:"varint,1,rep,name=pin" json:"pin,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Gpio) Reset()                    { *m = Gpio{} }
func (m *Gpio) String() string            { return proto.CompactTextString(m) }
func (*Gpio) ProtoMessage()               {}
func (*Gpio) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Gpio) GetPin() []bool {
	if m != nil {
		return m.Pin
	}
	return nil
}

func init() {
	proto.RegisterType((*TestCases)(nil), "Soundbench.TestCases")
	proto.RegisterType((*SineTestCase)(nil), "Soundbench.SineTestCase")
	proto.RegisterType((*SineTestElement)(nil), "Soundbench.SineTestElement")
	proto.RegisterType((*TestResults)(nil), "Soundbench.TestResults")
	proto.RegisterType((*SineTestResult)(nil), "Soundbench.SineTestResult")
	proto.RegisterType((*SineTestElementResults)(nil), "Soundbench.SineTestElementResults")
	proto.RegisterType((*AudioChannel)(nil), "Soundbench.AudioChannel")
	proto.RegisterType((*Tone)(nil), "Soundbench.Tone")
	proto.RegisterType((*Gpio)(nil), "Soundbench.Gpio")
	proto.RegisterEnum("Soundbench.DIRECTIONS", DIRECTIONS_name, DIRECTIONS_value)
	proto.RegisterEnum("Soundbench.ACHID", ACHID_name, ACHID_value)
}

var fileDescriptor0 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x1d, 0x9b, 0xc6, 0x93, 0x34, 0x4d, 0x17, 0x54, 0x59, 0x80, 0xa0, 0x5a, 0x2e, 0xa1,
	0x82, 0x20, 0x85, 0x2b, 0x97, 0x2a, 0x0d, 0xad, 0x25, 0xd4, 0xd2, 0xc4, 0xbd, 0x70, 0x73, 0xe3,
	0x49, 0xbc, 0x92, 0xbb, 0x36, 0xde, 0xf5, 0x01, 0xf1, 0x47, 0x7c, 0x25, 0xb3, 0x76, 0x1a, 0x6c,
	0x12, 0xc1, 0xc9, 0xde, 0x37, 0x6f, 0xde, 0xbe, 0x79, 0xb3, 0x30, 0x54, 0x59, 0x29, 0xe3, 0x7b,
	0x94, 0xcb, 0x64, 0x9c, 0x17, 0x99, 0xce, 0x18, 0x2c, 0xb6, 0x08, 0xff, 0x04, 0x5e, 0x88, 0x4a,
	0x4f, 0x23, 0x85, 0x8a, 0x7d, 0x80, 0x43, 0x25, 0x24, 0x6e, 0x01, 0xdf, 0x3a, 0xed, 0x8c, 0x7a,
	0x13, 0x7f, 0xfc, 0xa7, 0x61, 0xbc, 0x68, 0x10, 0xf8, 0x4f, 0xe8, 0x37, 0xcf, 0x0c, 0xc0, 0x16,
	0x31, 0x75, 0xd9, 0x23, 0x8f, 0x0d, 0xa1, 0x1b, 0x97, 0x45, 0xa4, 0x45, 0x26, 0x7d, 0x9b, 0x90,
	0x43, 0xf6, 0x1e, 0xba, 0x98, 0xe2, 0x03, 0x4a, 0xad, 0xfc, 0x4e, 0xa5, 0xfc, 0x62, 0x9f, 0xf2,
	0xac, 0xe6, 0xb0, 0x57, 0xe0, 0xac, 0x73, 0x91, 0xf9, 0x4e, 0x45, 0x1d, 0x36, 0xa9, 0x97, 0x84,
	0xf3, 0x5f, 0x16, 0x1c, 0xfd, 0xdd, 0x73, 0x04, 0x07, 0xcb, 0x24, 0x92, 0x12, 0xd3, 0x8d, 0x8b,
	0xb7, 0xe0, 0xc5, 0xa2, 0xc0, 0xe5, 0xd6, 0xc6, 0x60, 0x72, 0xd2, 0x54, 0xba, 0x08, 0xe6, 0xb3,
	0x69, 0x18, 0xdc, 0x5c, 0x2f, 0xd8, 0x33, 0xe8, 0xab, 0xe8, 0x21, 0x4f, 0x85, 0x5c, 0x93, 0x6d,
	0x24, 0x8b, 0xd6, 0xc8, 0x35, 0xe8, 0x7d, 0xb9, 0x5a, 0x61, 0x91, 0xa2, 0x5c, 0xeb, 0x84, 0xdc,
	0x18, 0x94, 0x92, 0x8a, 0xca, 0x58, 0x64, 0xd3, 0xfa, 0x32, 0xe5, 0xbb, 0xbb, 0x49, 0x9d, 0x37,
	0x08, 0xfc, 0x1c, 0x7a, 0xc6, 0xe7, 0x1c, 0x55, 0x99, 0x6a, 0xc5, 0x26, 0x30, 0x78, 0x4c, 0xba,
	0x86, 0xc8, 0xae, 0x45, 0x02, 0xcf, 0xf7, 0x05, 0x52, 0x33, 0xf8, 0x2d, 0x0c, 0xda, 0x48, 0x2b,
	0xee, 0x8f, 0x70, 0x50, 0xd4, 0xe2, 0x34, 0xa6, 0xf1, 0xc2, 0xff, 0x91, 0xed, 0xc6, 0x06, 0xff,
	0x06, 0x27, 0xfb, 0x2b, 0xbb, 0x41, 0xee, 0x4c, 0x6c, 0xff, 0x67, 0xe2, 0x5b, 0xe8, 0x37, 0xcf,
	0xec, 0x14, 0xdc, 0x68, 0x99, 0x6c, 0xfc, 0x0e, 0x26, 0xc7, 0xad, 0xc6, 0xe9, 0x55, 0x70, 0xc1,
	0x5e, 0x83, 0xab, 0x33, 0x89, 0x8f, 0xd2, 0xad, 0x8d, 0x87, 0x54, 0xe0, 0xef, 0xc0, 0x31, 0x5f,
	0x76, 0x0c, 0xde, 0xaa, 0xc0, 0xef, 0x25, 0x55, 0x7e, 0x54, 0xc1, 0xd9, 0x06, 0xaa, 0x76, 0xa7,
	0xcb, 0x18, 0xa9, 0x9f, 0x20, 0xfe, 0x14, 0x1c, 0xf3, 0x4e, 0x58, 0x0f, 0x3a, 0xb9, 0x90, 0xd5,
	0x5b, 0xee, 0x9e, 0xbd, 0x01, 0x68, 0xac, 0xdc, 0x03, 0x37, 0xb8, 0xfe, 0x7a, 0x17, 0x0e, 0x2d,
	0xca, 0xf2, 0xc9, 0xcd, 0x5d, 0x68, 0xfe, 0xed, 0xb3, 0x97, 0xe0, 0xd6, 0x8e, 0xba, 0xe0, 0x7c,
	0x99, 0x7d, 0x36, 0x65, 0x62, 0xce, 0x83, 0xcb, 0x2b, 0xaa, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0xa1, 0xce, 0x80, 0x34, 0x51, 0x03, 0x00, 0x00,
}