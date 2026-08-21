package suggest

import (
	"slices"
	"testing"
)

func TestGetSuggestions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		font  string
		want  []string
	}{
		{
			name:  "строчные буквы",
			input: "hello",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				`- Input is all lowercase. Try "HELLO" for emphasis.`,
				"- Single word detected. Consider adding punctuation.",
				"- Output dimensions: 8 lines × 31 characters.",
			},
		},
		{
			name:  "заглавные буквы",
			input: "HELLO",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Input is all uppercase. Consider using mixed case for better display.",
				"- Single word detected. Consider adding punctuation.",
				"- Output dimensions: 8 lines × 45 characters.",
			},
		},
		{
			name:  "смешанный регистр",
			input: "HeLlo",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Single word detected. Consider adding punctuation.",
				"- Output dimensions: 8 lines × 37 characters.",
			},
		},
		{
			name:  "одно слово со знаком препинания",
			input: "Hello!",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 8 lines × 36 characters.",
			},
		},
		{
			name:  "несколько слов",
			input: "Hello There",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 8 lines × 77 characters.",
			},
		},
		{
			name:  "несколько пробелов",
			input: "Hello  World",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Multiple consecutive spaces detected. Consider normalizing spaces.",
				"- Output dimensions: 8 lines × 86 characters.",
			},
		},
		{
			name:  "строчные буквы и несколько слов",
			input: "hello world",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				`- Input is all lowercase. Try "HELLO WORLD" for emphasis.`,
				"- Output dimensions: 8 lines × 75 characters.",
			},
		},
		{
			name:  "заглавные буквы и несколько слов",
			input: "HELLO WORLD",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Input is all uppercase. Consider using mixed case for better display.",
				"- Output dimensions: 8 lines × 102 characters.",
			},
		},
		{
			name:  "цифры",
			input: "1234",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Single word detected. Consider adding punctuation.",
				"- Output dimensions: 8 lines × 29 characters.",
			},
		},
		{
			name:  "пустая строка",
			input: "",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 1 lines × 0 characters.",
			},
		},
		{
			name:  "текст с буквальной последовательностью переноса строки",
			input: `Hello\nThere`,
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 16 lines × 39 characters.",
			},
		},
		{
			name:  "текст с двумя буквальными последовательностями переноса строки",
			input: `Hello\n\nThere`,
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 17 lines × 39 characters.",
			},
		},
		{
			name:  "специальные символы",
			input: "{Hello & There #}",
			font:  "../../standard.txt",
			want: []string{
				"\n--- AI Suggestions ---",
				"- Output dimensions: 8 lines × 121 characters.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetSuggestions(tt.input, tt.font)

			if !slices.Equal(got, tt.want) {
				t.Errorf("GetSuggestions() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCalculateDimensions(t *testing.T) {
	tests := []struct {
		name  string
		input string
		font  string
		want  int
		want2 int
	}{
		{
			name:  "слово Hello",
			input: "Hello",
			font:  "../../standard.txt",
			want:  8,
			want2: 32,
		},
		{
			name:  "слово hello",
			input: "hello",
			font:  "../../standard.txt",
			want:  8,
			want2: 31,
		},
		{
			name:  "слово HELLO",
			input: "HELLO",
			font:  "../../standard.txt",
			want:  8,
			want2: 45,
		},
		{
			name:  "смешанный регистр",
			input: "HeLlo",
			font:  "../../standard.txt",
			want:  8,
			want2: 37,
		},
		{
			name:  "цифры",
			input: "1234",
			font:  "../../standard.txt",
			want:  8,
			want2: 29,
		},
		{
			name:  "Hello 2024",
			input: "Hello 2024",
			font:  "../../standard.txt",
			want:  8,
			want2: 71,
		},
		{
			name:  "Hello There",
			input: "Hello There",
			font:  "../../standard.txt",
			want:  8,
			want2: 77,
		},
		{
			name:  "несколько пробелов",
			input: "Hello  World",
			font:  "../../standard.txt",
			want:  8,
			want2: 86,
		},
		{
			name:  "строчные буквы и несколько слов",
			input: "hello world",
			font:  "../../standard.txt",
			want:  8,
			want2: 75,
		},
		{
			name:  "заглавные буквы и несколько слов",
			input: "HELLO WORLD",
			font:  "../../standard.txt",
			want:  8,
			want2: 102,
		},
		{
			name:  "перенос строки",
			input: `Hello\nThere`,
			font:  "../../standard.txt",
			want:  16,
			want2: 39,
		},
		{
			name:  "две пустые строки между текстом",
			input: `Hello\n\nThere`,
			font:  "../../standard.txt",
			want:  17,
			want2: 39,
		},
		{
			name:  "пустая строка",
			input: "",
			font:  "../../standard.txt",
			want:  1,
			want2: 0,
		},
		{
			name:  "специальные символы",
			input: "{Hello & There #}",
			font:  "../../standard.txt",
			want:  8,
			want2: 121,
		},
		{
			name:  "шрифт shadow",
			input: "Hello",
			font:  "../../shadow.txt",
			want:  8,
			want2: 33,
		},
		{
			name:  "шрифт thinkertoy",
			input: "Hello",
			font:  "../../thinkertoy.txt",
			want:  8,
			want2: 17,
		},
		{
			name:  "несуществующий шрифт",
			input: "Hello",
			font:  "../../missing.txt",
			want:  0,
			want2: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := calculateDimensions(tt.input, tt.font)

			if got != tt.want {
				t.Errorf(
					"calculateDimensions() высота = %v, want %v",
					got,
					tt.want,
				)
			}

			if got2 != tt.want2 {
				t.Errorf(
					"calculateDimensions() ширина = %v, want %v",
					got2,
					tt.want2,
				)
			}
		})
	}
}
