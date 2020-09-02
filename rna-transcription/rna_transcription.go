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

// Compare to some other solutions

// https://exercism.io/tracks/go/exercises/rna-transcription/solutions/8932a2ee31024ec4a082095b89ff48fd
func ToRNA_bwarren2(dna string) (out string) {
	var RNAMap = map[string]string{"G": "C", "C": "G", "T": "A", "A": "U"}
	for _, value := range dna {
		out += RNAMap[string(value)]
	}
	return
}

// https://exercism.io/tracks/go/exercises/rna-transcription/solutions/5fbbe4d003222e82656c3834
func ToRna_mgood(input string) string {
	return strings.NewReplacer(
		"G", "C",
		"C", "G",
		"T", "A",
		"A", "U",
	).Replace(input)
}

// https://exercism.io/tracks/go/exercises/rna-transcription/solutions/16da121be28d43e89ee1e8ae881b44e8
var translate = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA - return DNA compliment
func ToRNA_jonahaapala(dna string) string {
	rna := make([]byte, len(dna))
	for i, c := range dna {
		rna[i] = translate[c]
	}
	return string(rna)
}

// https://exercism.io/tracks/go/exercises/rna-transcription/solutions/ec2dc55237cd4931840020b6cca847a8
// ToRNA converts a DNA string into an RNA string
// BenchmarkRNATranscription-4      5000000               301 ns/op
func ToRNA_kgibilterra(dna string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}, dna)
}
