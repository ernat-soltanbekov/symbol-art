package analyzer

import (
	"fmt"
	"slices"
	"strings"
	"symbol-art/internal/banner"
)

// func ClassChar(in)
func GetAnalyze(input, font string) []string {
	// words := strings.Split(input, " ") // Делим входную строку на части по пробелам.
	input = strings.ReplaceAll(input, `\n`, "\n")
	lines := strings.Split(input, "\n") // Разбиваем строку по последовательности "\n", чтобы обработать каждую строку отдельно.

	var words []string
	for _, line := range lines {
		words = append(words, strings.Split(line, " ")...)
	}

	var analysisResult []string // Здесь будем собирать результаты анализа.

	// 1. Character Classification
	uppercases := 0
	lowercases := 0
	digits := 0
	special := 0
	spaces := 0

	// Считаем большие/маленькие буквы, цифры, знаки препинания, пробелы.
	for _, line := range lines {
		for _, symbol := range line {
			if symbol >= 'A' && symbol <= 'Z' {
				uppercases++
			} else if symbol >= 'a' && symbol <= 'z' {
				lowercases++
			} else if symbol >= '0' && symbol <= '9' {
				digits++
			} else if symbol == ' ' {
				spaces++
			} else if (symbol >= '!' && symbol <= '/') ||
				(symbol >= ':' && symbol <= '@') ||
				(symbol >= '[' && symbol <= '`') ||
				(symbol >= '{' && symbol <= '~') {
				special++
			}
		}
	}

	// 2. Pattern Detection
	wordStats := make(map[string]map[byte]int)

	// Repeated characters
	for _, word := range words {
		repChars := make(map[byte]int)
		amount := 1
		for i := 1; i < len(word); i++ {
			if word[i] == word[i-1] { // Если буква повторяется
				amount++
			} else if amount > 1 {
				if amount > 1 {
					repChars[word[i-1]] = amount
				}
				amount = 1
			}

			if amount > 1 && i == len(word)-1 {
				repChars[word[i-1]] = amount
			}
		}

		if len(repChars) > 0 {
			wordStats[word] = repChars
		}
	}

	// Numeric sequences
	var numSeqs []string
	for _, word := range words {
		var builder strings.Builder // Строим числовую последовательность
		for i := 0; i <= len(word); i++ {
			if i < len(word) && word[i] >= '0' && word[i] <= '9' {
				builder.WriteByte(word[i])
			} else {
				if builder.Len() >= 2 {
					numSeqs = append(numSeqs, builder.String())
				}
				builder.Reset()
			}
		}
	}

	// 3. Complexity score
	var usedSymbols []rune
	totalCharacters := 0

	for _, line := range lines {
		for _, symbol := range line {
			totalCharacters++
			if !slices.Contains(usedSymbols, symbol) {
				usedSymbols = append(usedSymbols, symbol)
			}
		}
	}

	uniqueCharacters := len(usedSymbols)

	// 4. Art dimensions
	input = strings.ReplaceAll(input, `\n`, "\n")
	normArgs := strings.Split(input, "\n") // Разбиваем строку по последовательности "\n", чтобы обработать каждую строку отдельно.
	totalLines := 0                        // Здесь будем накапливать общую высоту результата.
	maxWidth := 0                          // Здесь будем хранить максимальную ширину среди всех строк арта.

	for _, arg := range normArgs { // Последовательно обрабатываем каждую часть входного текста.
		if arg == "" {
			totalLines++
			continue
		}

		result, err := banner.Render(font, arg) // Генерируем ASCII-арт для текущей строки.
		if err != nil {
			continue
		}

		totalLines += len(result) // Добавляем количество строк, которое получилось после генерации.

		for _, line := range result { // Просматриваем все строки арта и ищем самую длинную.
			if len(line) > maxWidth {
				maxWidth = len(line)
			}
		}
	}

	// Добавляем все результаты в массив с результатами по порядку.

	analysisResult = append(analysisResult, "\n--- AI Analysis ---")
	analysisResult = append(analysisResult, "Character Breakdown:")
	analysisResult = append(analysisResult, fmt.Sprintf("  Uppercase: %d", uppercases))
	analysisResult = append(analysisResult, fmt.Sprintf("  Lowercase: %d", lowercases))
	analysisResult = append(analysisResult, fmt.Sprintf("  Digits: %d", digits))
	analysisResult = append(analysisResult, fmt.Sprintf("  Special: %d", special))
	analysisResult = append(analysisResult, fmt.Sprintf("  Spaces: %d", spaces))
	analysisResult = append(analysisResult, "")

	if uppercases > 0 && lowercases > 0 || len(wordStats) > 0 || len(numSeqs) > 0 {
		analysisResult = append(analysisResult, "Patterns Detected:")
	}

	if uppercases > 0 && lowercases > 0 {
		analysisResult = append(analysisResult, "  - Mixed case detected")
	}

	if len(wordStats) > 0 {
		analysisResult = append(analysisResult, "  - Repeated characters:")

		var wordsWithRepeats []string
		for word := range wordStats {
			wordsWithRepeats = append(wordsWithRepeats, word)
		}
		slices.Sort(wordsWithRepeats)

		for _, word := range wordsWithRepeats {
			repeatsMap := wordStats[word]
			result := fmt.Sprintf("       %q -> ", word)

			var chars []byte
			for char := range repeatsMap {
				chars = append(chars, char)
			}
			slices.Sort(chars)

			for i, char := range chars {
				result += fmt.Sprintf("'%c' (x%d)", char, repeatsMap[char])
				if i != len(chars)-1 {
					result += ", "
				}
			}

			analysisResult = append(analysisResult, result)
		}
	}

	if len(numSeqs) > 0 {
		result := "  - Numeric sequence: "

		for i, num := range numSeqs {
			result += fmt.Sprintf("%q", num)

			if i != len(numSeqs)-1 {
				result += ", "
			}
		}

		analysisResult = append(analysisResult, result)
	}

	analysisResult = append(analysisResult, "")
	analysisResult = append(analysisResult, fmt.Sprintf("Complexity Score: %.2f%%", (float64(uniqueCharacters)/float64(totalCharacters))*100))
	analysisResult = append(analysisResult, fmt.Sprintf("Art Dimensions: %d lines × %d characters", totalLines, maxWidth))

	return analysisResult
}
