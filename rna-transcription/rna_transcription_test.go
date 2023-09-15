package strand

import "testing"

func TestRNATranscription(t *testing.T) {
	for _, test := range testCases {
		if actual := ToRNA(test.input); actual != test.expected {
			t.Fatalf("FAIL: %s - ToRNA(%q): %q, expected %q",
				test.description, test.input, actual, test.expected)
		}
		t.Logf("PASS: %s", test.description)
	}
}

func BenchmarkRNATranscription(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ToRNA(test.input)
		}
	}
}

// Compare to some other solutions
//
// === RUN   TestRNATranscription
//    rna_transcription_test.go:11: PASS: Empty RNA sequence
//    rna_transcription_test.go:11: PASS: RNA complement of cytosine is guanine
//    rna_transcription_test.go:11: PASS: RNA complement of guanine is cytosine
//    rna_transcription_test.go:11: PASS: RNA complement of thymine is adenine
//    rna_transcription_test.go:11: PASS: RNA complement of adenine is uracil
//    rna_transcription_test.go:11: PASS: RNA complement
//--- PASS: TestRNATranscription (0.00s)
//goos: darwin
//goarch: amd64
//pkg: strand
//BenchmarkRNATranscription
//BenchmarkRNATranscription-16                	 8941611	       131 ns/op	      16 B/op	       5 allocs/op
//BenchmarkRNATranscription_bwarren2
//BenchmarkRNATranscription_bwarren2-16       	  795481	      1280 ns/op	      96 B/op	      11 allocs/op
//BenchmarkRNATranscription_mgood
//BenchmarkRNATranscription_mgood-16          	  518358	      2050 ns/op	    2752 B/op	      24 allocs/op
//BenchmarkRNATranscription_jonahaapala
//BenchmarkRNATranscription_jonahaapala-16    	 6086840	       196 ns/op	      32 B/op	       6 allocs/op
//BenchmarkRNATranscription_kgibilterra
//BenchmarkRNATranscription_kgibilterra-16    	 5887526	       216 ns/op	      37 B/op	       5 allocs/op
//PASS
//ok  	strand	6.365s

func BenchmarkRNATranscription_bwarren2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ToRNA_bwarren2(test.input)
		}
	}
}

func BenchmarkRNATranscription_mgood(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ToRna_mgood(test.input)
		}
	}
}

func BenchmarkRNATranscription_jonahaapala(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ToRNA_jonahaapala(test.input)
		}
	}
}

func BenchmarkRNATranscription_kgibilterra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ToRNA_kgibilterra(test.input)
		}
	}
}
