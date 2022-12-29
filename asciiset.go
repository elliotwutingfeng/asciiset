// Package asciiset is an ASCII character bitset
package asciiset

import (
	"math/bits"
	"unicode/utf8"
)

// ASCIISet is a 32-byte value, where each bit represents the presence of a
// given ASCII character in the set. The 128-bits of the lower 16 bytes,
// starting with the least-significant bit of the lowest word to the
// most-significant bit of the highest word, map to the full range of all
// 128 ASCII characters. The 128-bits of the upper 16 bytes will be zeroed,
// ensuring that any non-ASCII character will be reported as not in the set.
// This allocates a total of 32 bytes even though the upper half
// is unused to avoid bounds checks in ASCIISet.Contains.
type ASCIISet [8]uint32

// MakeASCIISet creates a set of ASCII characters and reports whether all
// characters in chars are ASCII.
func MakeASCIISet(chars string) (as ASCIISet, ok bool) {
	for i := 0; i < len(chars); i++ {
		c := chars[i]
		if c >= utf8.RuneSelf {
			return as, false
		}
		as[c/32] |= 1 << (c % 32)
	}
	return as, true
}

// Add inserts character c into the set.
func (as *ASCIISet) Add(c byte) {
	if c < utf8.RuneSelf { // ensure that c is an ASCII byte
		as[c/32] |= 1 << (c % 32)
	}
}

// Contains reports whether c is inside the set.
func (as *ASCIISet) Contains(c byte) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}

// Remove removes c from the set
//
// if c is not in the set, the set contents will remain unchanged.
func (as *ASCIISet) Remove(c byte) {
	as[c/32] &= ^(1 << (c % 32))
}

// Size returns the number of characters in the set.
func (as *ASCIISet) Size() int {
	return bits.OnesCount(uint(as[0])) + bits.OnesCount(uint(as[1])) + bits.OnesCount(uint(as[2])) + bits.OnesCount(uint(as[3]))
}

// Union returns a set containing all characters that belong to either as and as2.
func (as *ASCIISet) Union(as2 ASCIISet) (as3 ASCIISet) {
	as3[0] = as[0] | as2[0]
	as3[1] = as[1] | as2[1]
	as3[2] = as[2] | as2[2]
	as3[3] = as[3] | as2[3]
	return
}

// Intersection creates a new set that contain all characters that belong to both as and as2.
func (as *ASCIISet) Intersection(as2 ASCIISet) (as3 ASCIISet) {
	as3[0] = as[0] & as2[0]
	as3[1] = as[1] & as2[1]
	as3[2] = as[2] & as2[2]
	as3[3] = as[3] & as2[3]
	return
}

// Subtract creates a new set that contain all characters that belong to as but not as2.
func (as *ASCIISet) Subtract(as2 ASCIISet) (as3 ASCIISet) {
	as3[0] = as[0] &^ as2[0]
	as3[1] = as[1] &^ as2[1]
	as3[2] = as[2] &^ as2[2]
	as3[3] = as[3] &^ as2[3]
	return
}

// Equals returns whether as contains the same characters as as2.
func (as *ASCIISet) Equals(as2 ASCIISet) bool {
	return as[0] == as2[0] && as[1] == as2[1] && as[2] == as2[2] && as[3] == as2[3]
}
