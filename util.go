package test_utils

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

// 	ClearSlash removes double forward and backward slashes from your string
// 	"//" => "/" && "\\" => "\"
// 	Input:
// 		in string	The string to clean
//	Output:
//		string		The cleaned up string
func ClearSlash(in string) string {
	return strings.Replace(strings.Replace(in, "//", "/", -1), `\\`, `\`, -1)
}

// CheckFor500 checks if the given http status code equals StatusInternalServerError
// Input:
//		t *testing.T 		The testing object, so we can call the return functions on it
//		statusCode int		The http status code you want to check
func CheckFor500(t *testing.T, statusCode int) {
	if status := statusCode; status != http.StatusInternalServerError {
		t.Fatalf("wrong status code: got '%d' != '%d' expected", status, http.StatusInternalServerError)
	}
}

// 	JsonEqual is an easy json compare function that tries to tell you if the given json strings are equal or not
// 	Input:
//		s1, s2 string	The two json strings you want to compare
//	Output:
//		bool			Bool that indicates if the strings are equal jsons (true if they are equal)
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
	if string(b1) == string(b2) {
		return true
	} else {
		return false
	}
}
