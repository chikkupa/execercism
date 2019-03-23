package strand

// ToRNA convert DNA strand to RNA
func ToRNA(dna string) string {
	rna := []byte(dna)

	for i := 0; i < len(rna); i++ {
		switch rna[i] {
		case 'G':
			rna[i] = 'C'
		case 'C':
			rna[i] = 'G'
		case 'T':
			rna[i] = 'A'
		case 'A':
			rna[i] = 'U'
		}
	}

	return string(rna)
}
