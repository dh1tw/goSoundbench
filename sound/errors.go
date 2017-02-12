package sound

import "fmt"

// Errorcodes
const (
	EDEVICE    = 1
	ECHANNEL   = 2
	EFREQ      = 3
	EMONO      = 4
	ESTEREO    = 5
	EAMPLITUDE = 6
	EFUNCTION  = 7
	EINTERNAL  = 8
)

// map containing the string representation of the Errorcode
var ErrorName = map[int]string{
	EDEVICE:    "EDEVICE",
	ECHANNEL:   "ECHANNEL",
	EFREQ:      "EFREQ",
	EMONO:      "EMONO",
	ESTEREO:    "ESTEREO",
	EAMPLITUDE: "EAMPLITUDE",
	EFUNCTION:  "EFUNCTION",
	EINTERNAL:  "EINTERNAL",
}

// map containing the description of the Errorcode
var ErrorDescription = map[int]string{
	EDEVICE:    "Unknown Device",
	ECHANNEL:   "Channel unknown, not defined or contains invalid data",
	EFREQ:      "Invalid Frequency",
	EMONO:      "A Mono Stream takes exactly one Channel",
	ESTEREO:    "A Stereo Stream takes exactly two Channels",
	EAMPLITUDE: "Invalid Amplitude",
	EFUNCTION:  "Test Function does not exist",
	EINTERNAL:  "Internal Error",
}

type SoundbenchError struct {
	Errorcode      int
	Unit           string
	AdditionalInfo string
}

func (e *SoundbenchError) Error() string {
	return fmt.Sprintf("ERROR %s: %s - %s (%s)", ErrorName[e.Errorcode], e.Unit, ErrorDescription[e.Errorcode], e.AdditionalInfo)
}
