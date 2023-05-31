package asciiset

import (
	"strings"
	"testing"
)

func TestMakeASCIISet(t *testing.T) {
	if _, ok := MakeASCIISet("@J"); !ok {
		t.Errorf(`MakeASCIISet should identify "@J" as all ASCII`)
	}
	if _, ok := MakeASCIISet("@\u00f8"); ok {
		t.Errorf(`MakeASCIISet should identify "@\u00f8" as not all ASCII`)
	}
}

func TestAdd(t *testing.T) {
	chars := "@J"
	var as ASCIISet
	for _, c := range chars {
		as.Add(byte(c))
		if !as.Contains(byte(c)) {
			t.Errorf("ASCIISet should contain %s", string(c))
		}
	}
	britishPound := byte('£')
	as.Add(britishPound)
	if as.Contains(britishPound) {
		t.Errorf("ASCIISet should not contain %s", string(britishPound))
	}
}

func TestContains(t *testing.T) {
	as, _ := MakeASCIISet("@J")
	if as.Contains('5') {
		t.Errorf("ASCIISet should not contain 5")
	}
}

func TestRemove(t *testing.T) {
	chars := "@J"
	as, _ := MakeASCIISet(chars)
	as.Remove('@')
	if as.Contains('@') {
		t.Errorf("ASCIISet should not contain @")
	}
	if !as.Contains('J') {
		t.Errorf("ASCIISet should contain J")
	}
}

func TestSize(t *testing.T) {
	as, _ := MakeASCIISet("ABCD")
	if size := as.Size(); size != 4 {
		t.Errorf("Expected Size 4, got %d", size)
	}
}

func TestUnion(t *testing.T) {
	as, _ := MakeASCIISet("ABCD")
	as2, _ := MakeASCIISet("CDEF")
	as3 := as.Union(as2)
	for _, c := range "ABCDEF" {
		if !as3.Contains(byte(c)) {
			t.Errorf("ASCIISet should contain %s", string(c))
		}
	}
}

func TestIntersection(t *testing.T) {
	as, _ := MakeASCIISet("ABCD")
	as2, _ := MakeASCIISet("CDEF")
	as3 := as.Intersection(as2)
	for _, c := range "CD" {
		if !as3.Contains(byte(c)) {
			t.Errorf("ASCIISet should contain %s", string(c))
		}
	}
	for _, c := range "ABEF" {
		if as3.Contains(byte(c)) {
			t.Errorf("ASCIISet should not contain %s", string(c))
		}
	}
}

func TestSubtract(t *testing.T) {
	as, _ := MakeASCIISet("ABCD")
	as2, _ := MakeASCIISet("CDEF")
	as3 := as.Subtract(as2)
	for _, c := range "AB" {
		if !as3.Contains(byte(c)) {
			t.Errorf("ASCIISet should contain %s", string(c))
		}
	}
	for _, c := range "CDEF" {
		if as3.Contains(byte(c)) {
			t.Errorf("ASCIISet should not contain %s", string(c))
		}
	}
}

func TestEquals(t *testing.T) {
	as, _ := MakeASCIISet("ABCD")
	as2, _ := MakeASCIISet("ABCD")
	if !as.Equals(as2) {
		t.Errorf("as should be equal to as2")
	}
}

func TestVisit(t *testing.T) {
	// chars must be all unique and in ascending order
	chars := "123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	as, _ := MakeASCIISet(chars)

	// scenario: visit every character in the set
	output := make([]byte, 0, len(chars))
	as.Visit(func(n byte) bool {
		if as.Contains(n) {
			output = append(output, n)
		}
		return false
	})
	if len(output) != len(chars) {
		t.Errorf("output length must be %d; visit every character. Got %d", len(chars), len(output))
	}
	for i := 0; i < len(chars); i++ {
		if output[i] != byte(chars[i]) {
			t.Errorf("%d %d", output[i], byte(chars[i]))
		}
	}
	// scenario: stop early at 'T'
	output = make([]byte, 0, strings.Index(chars, "T")+1)
	as.Visit(func(n byte) bool {
		if as.Contains(n) {
			output = append(output, n)
		}
		return n == 'T'
	})
	if len(output) != strings.Index(chars, "T")+1 {
		t.Errorf("output length must be %d; stop early at 'T'. Got %d", strings.Index(chars, "T")+1, len(output))
	}
	for i := 0; i < strings.Index(chars, "T")+1; i++ {
		if output[i] != byte(chars[i]) {
			t.Errorf("%d %d", output[i], byte(chars[i]))
		}
	}
	// scenario: Add extra '\n'
	output = make([]byte, 0, len(chars))
	as.Visit(func(n byte) bool {
		as.Add('\n')
		if as.Contains(n) {
			output = append(output, n)
		}
		return false
	})
	if as.Size() != len(chars)+1 {
		t.Errorf("as.Size() must be %d; visit every character and add extra '\n'. Got %d", len(chars)+1, as.Size())
	}
	if len(output) != len(chars) {
		t.Errorf("output length must be %d; visit every character and add extra '\n'. Got %d", len(chars), len(output))
	}
	for i := 0; i < len(chars); i++ {
		if output[i] != byte(chars[i]) {
			t.Errorf("%d %d", output[i], byte(chars[i]))
		}
	}
	// scenario: Remove extra '\n'
	output = make([]byte, 0, len(chars))
	as.Visit(func(n byte) bool {
		as.Remove('\n')
		if as.Contains(n) {
			output = append(output, n)
		}
		return false
	})
	if as.Size() != len(chars) {
		t.Errorf("as.Size() must be %d; visit every character and remove extra '\n'. Got %d", len(chars), as.Size())
	}
	if len(output) != len(chars) {
		t.Errorf("output length must be %d; visit every character and remove extra '\n'. Got %d", len(chars), len(output))
	}
	for i := 0; i < len(chars); i++ {
		if output[i] != byte(chars[i]) {
			t.Errorf("%d %d", output[i], byte(chars[i]))
		}
	}
}
