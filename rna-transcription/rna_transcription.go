package strand

func ToRNA(dna string) string {
	dnsToRna := map[byte]byte{
		'G': 'C',
		'C': 'G',
		'T': 'A',
		'A': 'U',
	}

	dnsbs := []byte(dna)
	rnabs := make([]byte, len(dnsbs))
	for i, d := range dnsbs {
		rnabs[i] = dnsToRna[d]
	}

	return string(rnabs)
}
