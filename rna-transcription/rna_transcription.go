package strand

// ToRNA doc here
func ToRNA(dna string) string {
	mapping := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}
	rna := make([]byte, len(dna))
	for i, r := range dna {
		rna[i] = mapping[byte(r)]
	}
	return string(rna)
}
