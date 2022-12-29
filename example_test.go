package asciiset

import "fmt"

// Basic asciiset operations
func Example_basics() {
	// Making an ASCIISet
	// No need to ensure that all characters in chars are unique
	chars := "3gqZ1mAhVcA#Z7eKvwPN8J@D"
	as, ok := MakeASCIISet(chars)
	if ok {
		fmt.Println("as created")
	}

	// Character not in ASCIISet
	if !as.Contains('n') {
		fmt.Println("as does not contain 'n'")
	}

	// Adding character to ASCIISet
	as.Add('b')

	// Character is in ASCIISet
	if as.Contains('b') {
		fmt.Println("as contains 'b'")
	}

	// Attempting to add same character to ASCIISet again will not return an error
	// and the ASCIISet contents will remain unchanged
	as.Add('b')

	// Adding non-ASCII byte characters will fail silently
	britishPound := byte('£') // this is not an ASCII character
	as.Add(britishPound)
	if !as.Contains(britishPound) {
		fmt.Printf("as does not contain %s\n", string(britishPound))
	}

	// Removing character from ASCIISet
	as.Remove('3')
	if !as.Contains('3') {
		fmt.Println("as does not contain 3")
	}

	// Attempting to remove same character from ASCIISet again will not return an error
	// and the ASCIISet contents will remain unchanged
	as.Remove('3')

	// Getting size of ASCIISet
	fmt.Println(as.Size())

	// Output: as created
	// as does not contain 'n'
	// as contains 'b'
	// as does not contain £
	// as does not contain 3
	// 22
}

// Operations involving multiple sets
func Example_multiple_sets() {
	as, _ := MakeASCIISet("ABCD")
	as2, _ := MakeASCIISet("CDEF")

	expectedUnion, _ := MakeASCIISet("ABCDEF")
	expectedIntersection, _ := MakeASCIISet("CD")
	expectedSubtract, _ := MakeASCIISet("AB")

	union := as.Union(as2)
	if union.Equals(expectedUnion) {
		fmt.Println(`Union of as and as2 is "ABCDEF"`)
	}
	intersection := as.Intersection(as2)
	if intersection.Equals(expectedIntersection) {
		fmt.Println(`Intersection of as and as2 is "CD"`)
	}
	subtract := as.Subtract(as2)
	if subtract.Equals(expectedSubtract) {
		fmt.Println(`Subtraction of as2 from as is "AB"`)
	}
	// Output: Union of as and as2 is "ABCDEF"
	// Intersection of as and as2 is "CD"
	// Subtraction of as2 from as is "AB"
}
