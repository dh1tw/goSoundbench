package test

// TestResultI is the interface for any specific test result
type TestResultI interface {
	Serialize() ([]byte, error)
	String() string
}
