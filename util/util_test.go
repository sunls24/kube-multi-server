package util

import "testing"

func TestIsZero(t *testing.T) {
	// test int
	if !IsZero(0) {
		t.Error("int 0 is zero value, but return false")
	}
	if IsZero(10) {
		t.Error("int 10 is not zero value, but return true")
	}

	// test string
	if !IsZero("") {
		t.Error("empty string is zero value, but return false")
	}
	if IsZero("not-zero-string") {
		t.Error("'not-zero-string' is not zero value, but return true")
	}

	if !IsZero(nil) {
		t.Error("nil is zero value, but return false")
	}

	// test slice
	if IsZero([]int{0, 1, 2}) {
		t.Error("slice [0,1,2] is not zero value, but return true")
	}

	// test map
	if IsZero(map[string]string{"a": "b"}) {
		t.Error("map {a:b} is not zero value, but return true")
	}

	// test struct
	type s struct {
		id   int
		name string
	}
	if !IsZero(s{}) {
		t.Error("struct {} is zero value, but return false")
	}

	if IsZero(s{id: 1, name: "sls"}) {
		t.Error("{id: 1, name: sls} is not zero value, but return true")
	}
}
