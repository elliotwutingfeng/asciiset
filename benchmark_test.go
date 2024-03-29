package asciiset

import (
	"testing"
	"unicode/utf8"
)

// prevent compiler from erroneously optimising benchmark computations
var globalExists bool
var globalSize int
var globalVal byte

// Number of sets
const N int = 10

// setupASCIISets returns an ASCIISet slice of size n where
// each ASCIISet is empty
//
// if populate is true, fill each set with every 2nd ASCII character
func setupASCIISets(n int, populate bool) []ASCIISet {
	sets := []ASCIISet{}
	for i := 0; i < n; i++ {
		var as ASCIISet
		if populate {
			for c := 0; c < utf8.RuneSelf; c += 2 {
				// add every 2nd ASCII character
				as.Add(byte(c))
			}
		}
		sets = append(sets, as)
	}
	return sets
}

// setupMapSets returns a map[byte]struct{} slice of size n where
// each map[byte]struct{} is empty
//
// if populate is true, fill each set with every 2nd ASCII character
func setupMapSets(n int, populate bool) []map[byte]struct{} {
	sets := []map[byte]struct{}{}
	for i := 0; i < n; i++ {
		as := make(map[byte]struct{})
		if populate {
			for c := 0; c < utf8.RuneSelf; c += 2 {
				// add every 2nd ASCII character
				as[byte(c)] = struct{}{}
			}
		}
		sets = append(sets, as)
	}
	return sets
}

func BenchmarkASCIISet(b *testing.B) {
	b.Run("ASCIISet Add()", func(b *testing.B) {
		sets := setupASCIISets(N, false)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c += 2 {
					// add every 2nd ASCII character
					as.Add(c)
				}
			}
		}
	})
	b.Run("ASCIISet Contains()", func(b *testing.B) {
		sets := setupASCIISets(N, true)
		var exists bool
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					exists = as.Contains(c)
				}
			}
		}
		globalExists = exists
	})
	b.Run("ASCIISet Remove()", func(b *testing.B) {
		sets := setupASCIISets(N, true)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					as.Remove(c)
				}
			}
		}
	})
	b.Run("ASCIISet Size()", func(b *testing.B) {
		sets := setupASCIISets(N, true)
		var size int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					size = as.Size()
				}
			}
		}
		globalSize = size
	})
	b.Run("ASCIISet Visit()", func(b *testing.B) {
		sets := setupASCIISets(N, true)
		var val byte
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				as.Visit(func(c byte) bool {
					val = c
					return false
				})
			}
		}
		globalVal = val
	})
}

func BenchmarkMapSet(b *testing.B) {
	b.Run("map Add", func(b *testing.B) {
		sets := setupMapSets(N, false)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c += 2 {
					// add every 2nd ASCII character
					as[c] = struct{}{}
				}
			}
		}
	})
	b.Run("map Contains", func(b *testing.B) {
		sets := setupMapSets(N, true)
		var exists bool
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					_, exists = as[c]
				}
			}
		}
		globalExists = exists
	})
	b.Run("map Remove", func(b *testing.B) {
		sets := setupMapSets(N, true)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					delete(as, c)
				}
			}
		}
	})
	b.Run("map Size", func(b *testing.B) {
		sets := setupMapSets(N, true)
		var size int
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := byte(0); c < utf8.RuneSelf; c++ {
					size = len(as)
				}
			}
		}
		globalSize = size
	})
	b.Run("map Visit", func(b *testing.B) {
		sets := setupMapSets(N, true)
		var val byte
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c := range as {
					val = c
				}
			}
		}
		globalVal = val
	})
}
