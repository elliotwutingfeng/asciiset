package asciiset

import (
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
