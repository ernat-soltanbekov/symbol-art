package main

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// Функция преобразует команду из input-файла в аргументы командной строки.
func parseCommand(input string) []string {
	var args []string
	var current strings.Builder
	inQuotes := false

	for _, char := range strings.TrimSpace(input) {
		switch char {
		case '"':
			inQuotes = !inQuotes

		case ' ', '\t':
			if inQuotes {
				current.WriteRune(char)
			} else if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

func TestGolden(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{
			name:  "первый golden",
			input: "input1.txt",
			want:  "output1.txt",
		},
		{
			name:  "второй golden",
			input: "input2.txt",
			want:  "output2.txt",
		},
		{
			name:  "анализ",
			input: "input3.txt",
			want:  "output3.txt",
		},
		{
			name:  "рекомендации",
			input: "input4.txt",
			want:  "output4.txt",
		},
		{
			name:  "шрифт shadow",
			input: "input5.txt",
			want:  "output5.txt",
		},
		{
			name:  "шрифт thinkertoy",
			input: "input6.txt",
			want:  "output6.txt",
		},
		{
			name:  "шрифт shadow с анализом и рекомендациями",
			input: "input7.txt",
			want:  "output7.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputPath := filepath.Join("samples", tt.input)
			outputPath := filepath.Join("samples", tt.want)

			input, err := os.ReadFile(inputPath)
			if err != nil {
				t.Fatalf("не удалось прочитать %s: %v", inputPath, err)
			}

			args := parseCommand(string(input))

			if len(args) == 0 {
				t.Fatal("input-файл пустой")
			}

			cmd := exec.Command(
				"go",
				append([]string{"run", "."}, args...)...,
			)

			var got bytes.Buffer
			var stderr bytes.Buffer

			cmd.Stdout = &got
			cmd.Stderr = &stderr

			if err := cmd.Run(); err != nil {
				t.Fatalf(
					"команда завершилась с ошибкой: %v\nstderr:\n%s",
					err,
					stderr.String(),
				)
			}

			want, err := os.ReadFile(outputPath)
			if err != nil {
				t.Fatalf("не удалось прочитать %s: %v", outputPath, err)
			}

			if got.String() != string(want) {
				t.Errorf(
					"вывод не совпадает с golden-файлом\n\nПолучено:\n%s\n\nОжидалось:\n%s",
					got.String(),
					string(want),
				)
			}
		})
	}
}
