package test_utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

func AssertStringEquals(t *testing.T, have, want string) {
	if have != want {
		t.Errorf("Have '%s' != '%s' Want", have, want)
	}
}

func AssertIntEquals(t *testing.T, have, want int) {
	if have != want {
		t.Errorf("Have '%d' != '%d' Want", have, want)
	}
}

func ClearSlash(in string) string {
	return strings.Replace(strings.Replace(in, "//", "/", -1), `\\`, `\`, -1)
}

func CheckFor500(t *testing.T, statusCode int) {
	if status := statusCode; status != http.StatusInternalServerError {
		t.Fatalf("wrong status code: got '%d' want '%d'", status, http.StatusInternalServerError)
	}
}

// jsonEqual tries to compare s1 and s2 as json, return true
// if the content is the same, false otherwise
func JsonEqual(s1, s2 string) bool {
	if s1 == s2 {
		return true
	}
	var i1, i2 interface{}
	err := json.Unmarshal([]byte(s1), &i1)
	if err != nil {
		return false
	}
	err = json.Unmarshal([]byte(s2), &i2)
	if err != nil {
		return false
	}
	b1, err := json.Marshal(i1)
	if err != nil {
		return false
	}
	b2, err := json.Marshal(i2)
	if err != nil {
		return false
	}
	// fmt.Printf("%s\n\n%s\n\n", string(b1), string(b2))
	if string(b1) == string(b2) {
		return true
	} else {
		return false
	}
}

func AssertErrorEquals(t *testing.T, have, want error) {
	if want == nil && have != nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}
	if want != nil && have == nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}

	if want == nil && have == nil {
		return
	}

	if have.Error() != want.Error() {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}
}

func AsserErrorEqualsAny(t *testing.T, have error, want []error) {
	if want == nil && have != nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}
	if want != nil && have == nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}

	if want == nil && have == nil {
		return
	}
	for _, v := range want {
		if have.Error() == v.Error() {
			return
		}
	}

	t.Errorf("'%v' not in '%v'", have, want)
}

func AssertErrorContains(t *testing.T, have, want error) {
	if want == nil && have != nil {
		t.Errorf("'%v' != '%v'", have, want)
		return
	}
	if want != nil && have == nil {
		t.Errorf("'%v' != '%v'", have, want)
		return
	}

	if want == nil && have == nil {
		return
	}

	if !strings.Contains(have.Error(), want.Error()) {
		t.Errorf("'%v' was not found in '%v'", want, have)
		return
	}
}

func CheckError(t *testing.T, err error, errorMessage string) {
	if err != nil {
		t.Error(errorMessage)
	}
}

func ExpectError(t *testing.T, err error, errorMessage string) {
	if err == nil {
		t.Error(errorMessage)
	}
}

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

func AssertStringContainsSubstringsNoOrder(t *testing.T, body string, expectedStrings []string) {
	for _, v := range expectedStrings {
		if !strings.Contains(body, v) {
			t.Errorf("'%s' not found.", v)
		}
	}
}

func AssertStringContainsNoneOfTheSubstrings(t *testing.T, body string, notExpectedLogEntries []string) {
	print := false
	for _, v := range notExpectedLogEntries {
		if strings.Contains(body, v) {
			t.Errorf("We did not expect '%s' but found it in '%s'", v, body)
			print = true
		}

	}

	if print {
		t.Log(body)
	}
}

// AssertMapsEqual checks if two maps have the exact same content
// Attention: This function changes the value of the first map!
func AssertMapsEqual(t *testing.T, got, want map[string]interface{}) {

	for k, v := range want {
		if got[k] != v {
			t.Errorf("[%s] Got '%v' != '%v' Want", k, got[k], v)
		} else {
			delete(got, k)
		}
	}

	for k, v := range got {
		t.Errorf("[%s] Got '%v' != '' Want", k, v)

	}
}
func AssertStringArraysEqualNoOrder(t *testing.T, have, want []string) {
	wantInner := make([]string, len(want))
	copy(wantInner, want)

	if want == nil && have != nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}
	if want != nil && have == nil {
		t.Errorf("Have '%v' != '%v' Want", have, want)
		return
	}

	if want == nil && have == nil {
		return
	}

	for _, v := range have {
		for ik, iv := range wantInner {
			if v == iv {
				wantInner = append(wantInner[:ik], wantInner[ik+1:]...)
				break
			}
		}
	}

	if len(wantInner) > 0 {
		for _, v := range wantInner {
			t.Errorf("'%s' not found", v)
		}
	}

}
