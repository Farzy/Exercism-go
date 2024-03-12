package isogram

import "testing"

func TestIsIsogram(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogram(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogram(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogram(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogram(c.input)
		}
	}
}

func TestIsIsogramBitfield(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogramBitfield(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogramBitfield(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogramBitfield(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramBitfield(c.input)
		}
	}
}

func TestIsIsogram2(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogram2(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogram2(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogram2(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogram2(c.input)
		}
	}
}

func TestIsIsogramInt(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogramInt(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogramInt(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogramInt(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramInt(c.input)
		}
	}
}

func TestIsIsogramMap(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := IsIsogramMap(tc.input); actual != tc.expected {
				t.Fatalf("IsIsogramMap(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkIsIsogramMap(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, c := range testCases {
			IsIsogramMap(c.input)
		}
	}
}
