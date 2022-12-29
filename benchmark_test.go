package asciiset

import (
	"testing"
	"unicode/utf8"
)

// prevent compiler from erroneously optimising benchmark computations
var globalExists bool
var globalSize int

func BenchmarkASCIISet(b *testing.B) {
	sets := []ASCIISet{}
	for i := 0; i < 100000; i++ {
		var as ASCIISet
		sets = append(sets, as)
	}
	b.Run("ASCIISet Add()", func(b *testing.B) {
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c += 2 {
					// add every 2nd ASCII character
					as.Add(c)
				}
			}
		}
	})
	b.Run("ASCIISet Contains()", func(b *testing.B) {
		var exists bool
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					exists = as.Contains(c)
				}
			}
		}
		globalExists = exists
	})
	b.Run("ASCIISet Remove()", func(b *testing.B) {
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					as.Remove(c)
				}
			}
		}
	})
	b.Run("ASCIISet Size()", func(b *testing.B) {
		var c byte
		var size int
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					size = as.Size()
				}
			}
		}
		globalSize = size
	})
}

func BenchmarkMapSet(b *testing.B) {
	sets := []map[byte]struct{}{}
	for i := 0; i < 100000; i++ {
		sets = append(sets, make(map[byte]struct{}))
	}
	b.Run("map Add()", func(b *testing.B) {
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c += 2 {
					// add every 2nd ASCII character
					as[c] = struct{}{}
				}
			}
		}
	})
	b.Run("map Contains()", func(b *testing.B) {
		var exists bool
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					_, exists = as[c]
				}
			}
		}
		globalExists = exists
	})
	b.Run("map Remove()", func(b *testing.B) {
		var c byte
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					delete(as, c)
				}
			}
		}
	})
	b.Run("map Size()", func(b *testing.B) {
		var c byte
		var size int
		for i := 0; i < b.N; i++ {
			for _, as := range sets {
				for c = 0; c < utf8.RuneSelf; c++ {
					size = len(as)
				}
			}
		}
		globalSize = size
	})
}
