package luhn

import "testing"

func TestValid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Valid(tc.input); actual != tc.expected {
				t.Fatalf("Valid(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkValid(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Valid(tc.input)
		}
	}
}

func TestValidRune(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ValidRune(tc.input); actual != tc.expected {
				t.Fatalf("ValidRune(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkValidRune(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			ValidRune(tc.input)
		}
	}
}

//func TestValidExercism(t *testing.T) {
//	for _, tc := range testCases {
//		t.Run(tc.description, func(t *testing.T) {
//			if actual := ValidExercism(tc.input); actual != tc.expected {
//				t.Fatalf("ValidExercism(%q) = %t, want: %t", tc.input, actual, tc.expected)
//			}
//		})
//	}
//}

func BenchmarkValidExercism(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			ValidExercism(tc.input)
		}
	}
}
