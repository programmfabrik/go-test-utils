package go_test_utils

import "testing"

//	ExpectError checks if the given error is != nil. If it is the msg is written to testing.Error
func ExpectError(t *testing.T, err error, msg string) {
	if err == nil {
		t.Error(msg)
	}
}

//	func checks if the given error is == nil. If it is not the msg is written to the testing.Error
func ExpectNoError(t *testing.T, err error, msg string) {
	if err != nil {
		t.Error(msg)
	}
}
