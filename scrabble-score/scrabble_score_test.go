package scrabble

import "testing"

func TestScore(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Score(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}

func TestScoreToUpperString(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ScoreToUpperString(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScoreToUpperString(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreToUpperString(test.input)
		}
	}
}

func TestScoreMap(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ScoreMap(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScoreMap(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreMap(test.input)
		}
	}
}

func TestScoreSwitch(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ScoreSwitch(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScoreSwitch(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range scrabbleScoreTests {
			ScoreSwitch(test.input)
		}
	}
}
