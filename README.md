![](https://img.shields.io/twitter/follow/programmfabrik.svg?label=Follow&style=social)  [![](http://img.shields.io/badge/docs-GoDoc-red.svg)](https://godoc.org/github.com/programmfabrik/go-test-utils)
             
                                                                                                         
# Go Test Utils

This small packages gives you some handy testing functions, to supercharge your go unit testing workflow.

**Full documentation?** Use the source, Luke (or go to [GoDoc](https://godoc.org/github.com/programmfabrik/go-test-utils))

**Do you have your own small functions, that help you with you unit test? Do not hesitate to open a pull request so we can build a great tool chain and make unit go testing even more effective.**

## Assert

`AssertStringEquals(t testing.TB, expected, got string)` checks if two strings are equal

`AssertIntEquals(t testing.TB, expected, got int)` checks if two strings are equal

`AssertIsError(t testing.TB, err error)` checks if there is an error

`AssertErrorEquals(t *testing.T, expected, got error)` checks if the actually error is equal the expected one, by doing a string compare

`AsserErrorEqualsAny(t *testing.T, got error, expectAnyIn []error)` checks if the actually error is equal to one in your error slice.

`AssertErrorContains(t *testing.T, err error, shouldContain string)` checks if the given error contains the expected substring

`AssertStringContainsSubstringsInOrder(t *testing.T, body string, expectedStrings []string)`  checks if the given string contains all the expected strings in the right order.

`AssertStringContainsSubstringsNoOrder(t *testing.T, body string, expectedStrings []string)` checks if the given string contains all the expected strings, we do not care about the order

`AssertStringContainsNoneOfTheSubstrings(t *testing.T, body string, nonExpectedStrings []string)`  checks if the given string contains any of the nonExpectedStrings

`AssertMapsEqual(t *testing.T, got, expected map[string]interface{})` checks if two maps have the exact same content

`AssertStringArraysEqualNoOrder(t *testing.T, got, expected []string)` checks if two string slices are the same, but do not have the same order

## Server


`type HandleFunc func(*http.ResponseWriter, *http.Request)` the type of func that a typical go http handler has

`type Routes map[string]HandleFunc` is a map of functions, that help you to define your testserver

`NewTestServer(routes Routes) *httptest.Server` creates a new go testing server with the given routes, so you can define more complex test server setups in no time

## Util


`ClearSlash(in string) string` removes double forward and backward slashes from your string

`CheckFor500(t *testing.T, statusCode int)` checks if the given http status code equals StatusInternalServerError

`JsonEqual(s1, s2 string) bool` is an easy json compare function that tries to tell you if the given json strings are equal or not


---

We use this toolset internally in our backend team @[Programmfabrik](https://www.programmfabrik.de) (FYI: [We are hiring ;)](https://www.programmfabrik.de/career/?lang=en))
