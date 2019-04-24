package go_test_utils

import (
	"github.com/sergi/go-diff/diffmatchpatch"
	"strings"
	"testing"
)

// 	AssertStringEquals checks if two strings are equal
// 	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//   	expected string 	Te string you want to have
//   	got string			The string you actually got from your test result
func AssertStringEquals(t testing.TB, expected, got string) {
	if expected != got {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(expected, got, false)
		t.Error(dmp.DiffPrettyText(diffs))
	}
}

// 	AssertIntEquals checks if two strings are equal
// 	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//   	expected int 		The int you want to have
//   	got int				The int you actually got from your test result
func AssertIntEquals(t testing.TB, expected, got int) {
	if expected != got {
		t.Errorf("expected '%d' != '%d' Want", expected, got)
	}
}

//	AssertIsError checks if there is an error
//	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//		err error	 		The error you want to check for
func AssertIsError(t testing.TB, err error) {
	if err == nil {
		t.Errorf("err == nil")
	}
}

//	AssertErrorEquals checks if the actually error is equal the expected one, by doing a string compare
//	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//		expected error	 	The error you expect
//		got error	 		The error you actually got
func AssertErrorEquals(t *testing.T, expected, got error) {
	if got == nil && expected == nil {
		return
	}

	if got == nil && expected != nil {
		t.Errorf("expected '%v' != '%v' Want", expected, got)
		return
	}
	if got != nil && expected == nil {
		t.Errorf("expected '%v' != '%v' Want", expected, got)
		return
	}

	if expected.Error() != got.Error() {
		t.Errorf("expected '%v' != '%v' Want", expected, got)
		return
	}
}

//	AsserErrorEqualsAny checks if the actually error is equal to one in your error slice.
//	You can use this function if you expect an error, but you do not exactly know which one
//	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//		got error	 		The error you actually got
//		expectAnyIn []error	All the errors that are a valid result
func AsserErrorEqualsAny(t *testing.T, got error, expectAnyIn []error) {
	if expectAnyIn == nil && got == nil {
		return
	}

	if expectAnyIn == nil && got != nil {
		t.Errorf("got != nil &&  expectAnyIn == nil")
		return
	}

	if expectAnyIn != nil && got == nil {
		t.Errorf("got == nil &&  expectAnyIn != nil")
		return
	}

	for _, v := range expectAnyIn {
		if got.Error() == v.Error() {
			return
		}
	}

	t.Errorf("'%v' not in '%v'", got, expectAnyIn)
}

//	AssertErrorContains checks if the given error contains the expected substring
//	Input:
//		t *testing.T 			The testing object, so we can call the return functions on it
//		err error	 			The error you actually got
//		shouldContain string 	The substring you expect in the error
func AssertErrorContains(t *testing.T, err error, shouldContain string) {
	if shouldContain == "" && err == nil {
		return
	}

	if err == nil && shouldContain != "" {
		t.Errorf("err == nil && shouldContaint != nil")
		return
	}

	if err != nil && shouldContain == "" {
		t.Errorf("err != nil && shouldContaint == nil")
		return
	}

	if !strings.Contains(err.Error(), shouldContain) {
		t.Errorf("'%v' was not found in '%v'", shouldContain, err)
		return
	}
}

//	AssertStringContainsSubstringsInOrder checks if the given string contains all the expected strings in the right order.
//	Prints the complete body if wrong order is present
//	Input:
//		t *testing.T 			The testing object, so we can call the return functions on it
//		body string	 			The string you got
//		expectedStrings string 	The substrings you expect to be in order in the body string
func AssertStringContainsSubstringsInOrder(t *testing.T, body string, expectedStrings []string) {
	cI := 0
	print := false

	for _, v := range expectedStrings {
		i := strings.Index(body, v)
		if i < cI {
			t.Errorf("Wrong order in string. '%s' is at the wrong position", v)
			print = true
		} else {
			cI = i
		}
	}

	if print {
		t.Log(body)
	}
}

//	AssertStringContainsSubstringsNoOrder checks if the given string contains all the expected strings, we do not care about the order
//	Prints the complete body if wrong order is present
//	Input:
//		t *testing.T 			The testing object, so we can call the return functions on it
//		body string	 			The string you got
//		expectedStrings string 	The substrings you expect to be in the body string
func AssertStringContainsSubstringsNoOrder(t *testing.T, body string, expectedStrings []string) {
	for _, v := range expectedStrings {
		if !strings.Contains(body, v) {
			t.Errorf("'%s' not found.", v)
		}
	}
}

//	AssertStringContainsNoneOfTheSubstrings checks if the given string contains any of the nonExpectedStrings
//	If a string is found, the test fails and we print the complete body
//	Prints the complete body if wrong order is present
//	Input:
//		t *testing.T 				The testing object, so we can call the return functions on it
//		body string	 				The string you got
//		nonExpectedStrings string 	The substrings you expect not to be in the body string
func AssertStringContainsNoneOfTheSubstrings(t *testing.T, body string, nonExpectedStrings []string) {
	print := false

	for _, v := range nonExpectedStrings {
		if strings.Contains(body, v) {
			t.Errorf("We did not expect '%s' but found it in '%s'", v, body)
			print = true
		}
	}

	if print {
		t.Log(body)
	}
}

//	AssertMapsEqual checks if two maps have the exact same content
//	If a string is found, the test fails and we print the complete body
//	Prints the complete body if wrong order is present
//	Input:
//		t *testing.T 					The testing object, so we can call the return functions on it
//		got map[string]interface{}	 	The map you got
//		expected map[string]interface{}	The map you expect to have
func AssertMapsEqual(t *testing.T, got, expected map[string]interface{}) {
	// Create new maps
	gotCopy := make(map[string]interface{})
	expectedCopy := make(map[string]interface{})

	// Copy the values over
	for key, value := range expected {
		expectedCopy[key] = value
	}
	for key, value := range got {
		gotCopy[key] = value
	}

	// Iterate over the expected map and check if the values are present and equal in the got map
	for k, v := range expectedCopy {
		if gotCopy[k] != v {
			t.Errorf("[%s] got '%v' != '%v' expected", k, got[k], v)
		} else {
			delete(gotCopy, k)
		}
	}

	// Print if there are extra values in the got map
	for k, v := range gotCopy {
		t.Errorf("[%s] got '%v' != '' expected", k, v)
	}
}

//	AssertStringArraysEqualNoOrder checks if two string slices are the same, but do not have the same order
//	Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//		got []string	 	The slice you got
//		expected []string	The map you expect to have
func AssertStringArraysEqualNoOrder(t *testing.T, got, expected []string) {
	if expected == nil && got == nil {
		return
	}

	if expected == nil && got != nil {
		t.Errorf("got '%v' != '%v' expected", got, expected)
		return
	}

	if expected != nil && got == nil {
		t.Errorf("got '%v' != '%v' expected", got, expected)
		return
	}

	// Copy the expected array over (as me modify it later on)
	expectedInner := make([]string, len(expected))
	copy(expectedInner, expected)

	gotInner := make([]string, 0)

	for _, v := range got {
		notMatched := true
		for ik, iv := range expectedInner {
			if v == iv {
				notMatched = false
				expectedInner = append(expectedInner[:ik], expectedInner[ik+1:]...)
				break
			}
		}

		if notMatched {
			gotInner = append(gotInner, v)
		}

	}

	// Print if there are elements left, that where not found in the got slice
	if len(expectedInner) > 0 {
		for _, v := range expectedInner {
			t.Errorf("'%s' not found", v)
		}
	}

	// Print if there are elements left, that where not found in the want slice
	if len(gotInner) > 0 {
		for _, v := range gotInner {
			t.Errorf("'%s' not expected", v)
		}
	}
}
