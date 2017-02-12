package test

// var mapping = map[string]string{
// 	"Built-in Output":  "ftd/pilot_out",
// 	"Built-in Microph": "ftd/pilot_in",
// 	// "ftd/iosinstructor":   "CH2",
// 	// "ftd/cabin":           "CH3",
// 	// "ftd/cabininstructor": "CH4",
// }

// TestcaseI is the interface which all Testcase have to implement
type TestcaseI interface {
	Setup() error
	Execute() error
	GetID() string
	GetResult() (TestResultI, error)
	Cleanup()
}
