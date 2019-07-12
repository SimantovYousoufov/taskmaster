package testing

import (
	"reflect"
	"testing"
	"runtime"
)

//
// Testing helpers
///
func AssertNilError(t *testing.T, err error) {
	if err != nil {
		file, no := GetCaller()
		t.Fatalf("%s:%d Failed due to error %s", file, no, err)
	}
}

func AssertErrorValue(t *testing.T, err error, expect string) {
	if err == nil || err.Error() != expect {
		file, no := GetCaller()
		t.Fatalf("%s:%d expected error with value %s, received %s", file, no, expect, err)
	}
}

func AssertNotHasKey(t *testing.T, key string, m map[string]string) {
	if _, ok := m[key]; ok {
		file, no := GetCaller()
		t.Fatalf("%s:%d Expected map to not have key %s but it is set", file, no, key)
	}
}

func AssertNotHasKey2D(t *testing.T, key string, m map[string][]string) {
	if _, ok := m[key]; ok {
		file, no := GetCaller()
		t.Fatalf("%s:%d Expected map to not have key %s but it is set", file, no, key)
	}
}

func AssertHasKey(t *testing.T, key string, m map[string][]string) {
	if _, ok := m[key]; ! ok {
		file, no := GetCaller()
		t.Fatalf("%s:%d Expected map to have key %v but it is not set", file, no, key)
	}
}

func AssertNil(t *testing.T, val interface{}) {
	if val != nil {
		file, no := GetCaller()
		t.Fatalf("%s:%d Expected %v to be nil", file, no, val)
	}
}

func AssertNotNil(t *testing.T, val interface{}) {
	if val == nil {
		file, no := GetCaller()
		t.Fatalf("%s:%d Expected %v to not be nil", file, no, val)
	}
}

func AssertSame(t *testing.T, expect interface{}, actual interface{}) {
	expectedType := reflect.TypeOf(expect)
	actualType := reflect.TypeOf(actual)

	if expectedType != actualType {
		file, no := GetCaller()

		t.Fatalf("%s:%d Expected %v to be of type %v", file, no, actualType, expectedType)
	}

	if expect != actual {
		file, no := GetCaller()

		t.Fatalf("%s:%d Expected \n %v but got \n %v", file, no, expect, actual)
	}
}

func AssertFalse(t *testing.T, value bool) {
	if value {
		file, no := GetCaller()

		t.Fatalf("%s:%d Expected false, received true", file, no)
	}
}

func AssertTrue(t *testing.T, value bool) {
	if ! value {
		file, no := GetCaller()

		t.Fatalf("%s:%d Expected true, received false", file, no)
	}
}

func GetCaller() (string, int) {
	// Get caller of caller
	_, file, no, ok := runtime.Caller(2)

	if ! ok {
		panic("could not get caller")
	}

	return file, no
}