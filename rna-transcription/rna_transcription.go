// RNA transcription package
package strand

import "strings"

// ToRNA converts a DNA string to an RNA string
func ToRNA(dna string) string {
	var str strings.Builder
	str.Grow(len(dna))

	for _, c := range dna {
		switch c {
		case 'G':
			str.WriteByte('C')
		case 'C':
			str.WriteByte('G')
		case 'T':
			str.WriteByte('A')
		case 'A':
			str.WriteByte('U')
		}
	}

	return str.String()
}
