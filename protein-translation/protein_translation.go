package protein

// FromCodon doc here
func FromCodon(codon string) string {
	var codonMap = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine",
		"UUC": "Phenylalanine",
		"UUA": "Leucine",
		"UUG": "Leucine",
		"UCU": "Serine",
		"UCC": "Serine",
		"UCA": "Serine",
		"UCG": "Serine",
		"UAU": "Tyrosine",
		"UAC": "Tyrosine",
		"UGU": "Cysteine",
		"UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP",
		"UAG": "STOP",
		"UGA": "STOP",
	}
	return codonMap[codon]
}

// FromRNA doc here
func FromRNA(rna string) []string {
	ret := make([]string, 0)
	for i := 0; i < len(rna); i += 3 {
		r := rna[i : i+3]
		codon := FromCodon(r)
		if codon == "STOP" {
			break
		}
		ret = append(ret, codon)
	}
	return ret
}
