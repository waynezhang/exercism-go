package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[byte]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Count counts number of occurrences of given nucleotide in given DNA.
//
// This is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Count method has a receiver of type DNA named d.
func (d DNA) Count(nucleotide byte) (int, error) {
	count := 0
	for _, ch := range d {
		switch byte(ch) {
		case 'A', 'T', 'C', 'G':
			if byte(ch) == nucleotide {
				count++
			}
		default:
			return 0, errors.New("")

		}
	}
	return count, nil
}

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = make(map[byte]int)
	for _, ch := range []byte{'A', 'C', 'T', 'G'} {
		if count, err := d.Count(ch); err == nil {
			h[ch] = count
		} else {
			return h, err
		}
	}
	return h, nil
}
