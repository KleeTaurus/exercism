package protein

import (
	"errors"
)

var (
	ErrStop        = errors.New("Stop")
	ErrInvalidBase = errors.New("Invalid Base")
)

func FromRNA(rna string) ([]string, error) {
	if len(rna)%3 != 0 {
		return nil, ErrInvalidBase
	}

	proteins := []string{}
	rnabs := []byte(rna)
	for i := 0; i < len(rnabs); i += 3 {
		protein, err := FromCodon(string([]byte{rnabs[i], rnabs[i+1], rnabs[i+2]}))
		if err != nil {
			if err == ErrStop {
				break
			}
			return nil, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, nil
}

func FromCodon(codon string) (string, error) {
	var protein string
	var err error

	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return protein, err
}
