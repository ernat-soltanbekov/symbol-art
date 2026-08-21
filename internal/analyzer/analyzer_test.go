package analyzer

import (
	"slices"
	"testing"
)

func TestGetAnalyze(t *testing.T) {
	tests := []struct {
		name  string
		input string
		font  string
		want  []string
	}{
		{
			name:  "анализ слова Hello",
			input: "Hello",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Analysis ---",
				"Character Breakdown:",
				"  Uppercase: 1",
				"  Lowercase: 4",
				"  Digits: 0",
				"  Special: 0",
				"  Spaces: 0",
				"",
				"Patterns Detected:",
				"  - Mixed case detected",
				"  - Repeated characters:",
				`       "Hello" -> 'l' (x2)`,
				"",
				"Complexity Score: 80.00%",
				"Art Dimensions: 8 lines × 32 characters",
			},
		},
		{
			name:  "анализ слова в верхнем регистре",
			input: "HELLO",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Analysis ---",
				"Character Breakdown:",
				"  Uppercase: 5",
				"  Lowercase: 0",
				"  Digits: 0",
				"  Special: 0",
				"  Spaces: 0",
				"",
				"Patterns Detected:",
				"  - Repeated characters:",
				`       "HELLO" -> 'L' (x2)`,
				"",
				"Complexity Score: 80.00%",
				"Art Dimensions: 8 lines × 45 characters",
			},
		},
		{
			name:  "анализ числовой последовательности",
			input: "1234",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Analysis ---",
				"Character Breakdown:",
				"  Uppercase: 0",
				"  Lowercase: 0",
				"  Digits: 4",
				"  Special: 0",
				"  Spaces: 0",
				"",
				"Patterns Detected:",
				`  - Numeric sequence: "1234"`,
				"",
				"Complexity Score: 100.00%",
				"Art Dimensions: 8 lines × 29 characters",
			},
		},
		{
			name:  "анализ слова с числовой последовательностью",
			input: "Hello 2024",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Analysis ---",
				"Character Breakdown:",
				"  Uppercase: 1",
				"  Lowercase: 4",
				"  Digits: 4",
				"  Special: 0",
				"  Spaces: 1",
				"",
				"Patterns Detected:",
				"  - Mixed case detected",
				"  - Repeated characters:",
				`       "Hello" -> 'l' (x2)`,
				`  - Numeric sequence: "2024"`,
				"",
				"Complexity Score: 80.00%",
				"Art Dimensions: 8 lines × 71 characters",
			},
		},
		{
			name:  "анализ слова со специальным символом",
			input: "Hello!",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Analysis ---",
				"Character Breakdown:",
				"  Uppercase: 1",
				"  Lowercase: 4",
				"  Digits: 0",
				"  Special: 1",
				"  Spaces: 0",
				"",
				"Patterns Detected:",
				"  - Mixed case detected",
				"  - Repeated characters:",
				`       "Hello!" -> 'l' (x2)`,
				"",
				"Complexity Score: 83.33%",
				"Art Dimensions: 8 lines × 36 characters",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetAnalyze(tt.input, tt.font)

			if !slices.Equal(got, tt.want) {
				t.Errorf("GetAnalyze() = %q, want %q", got, tt.want)
			}
		})
	}
}
