package strand

func ToRNA(dna string) string {
	rna := ""

	for _, nucleotide := range dna {
		switch string(nucleotide) {
		case "G":
			rna += "C"
		case "C":
			rna += "G"
		case "T":
			rna += "A"
		case "A":
			rna += "U"
		}
	}

	return rna
}
