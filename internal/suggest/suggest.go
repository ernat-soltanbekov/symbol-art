package suggest

import (
	"fmt"
	"strings"
	"symbol-art/internal/banner"
	"unicode"
)

func GetSuggestions(input, font string) []string { // Главная функция рекомендаций. Получает текст пользователя и формирует список полезных рекомендаций.
	var suggestion []string // Здесь постепенно будет собираться список рекомендаций.

	suggestion = append(suggestion, "\n--- AI Suggestions ---")

	hasLetters := false // Флаг показывает, встретились ли в тексте вообще буквы.
	isAllLower := true  // Предполагаем, что все буквы строчные, пока не встретим обратное.
	isAllUpper := true  // Предполагаем, что все буквы заглавные, пока не встретим обратное.

	for _, r := range input { // Последовательно просматриваем каждый символ входной строки.
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsLower(r) {
				isAllLower = false
			}
			if !unicode.IsUpper(r) {
				isAllUpper = false
			}
		}
	}

	if hasLetters && isAllLower { // Если все найденные буквы оказались строчными, предлагаем вариант с верхним регистром.
		suggestion = append(suggestion, fmt.Sprintf("- Input is all lowercase. Try %q for emphasis.", strings.ToUpper(input)))
	} else if hasLetters && isAllUpper { // Если текст полностью состоит из заглавных букв, рекомендуем смешанный регистр.
		suggestion = append(suggestion, "- Input is all uppercase. Consider using mixed case for better display.")
	}

	words := strings.Fields(input) // Разбиваем строку на отдельные слова для дальнейшего анализа.
	hasPunctuation := false        // Подготавливаем флаг, который покажет, содержит ли текст знаки препинания.

	for _, r := range input { // Еще раз проходим по каждому символу и ищем знаки препинания.
		if unicode.IsPunct(r) {
			hasPunctuation = true
			break
		}
	}

	if len(words) == 1 && !hasPunctuation { // Если введено только одно слово без пунктуации, предлагаем добавить знак препинания.
		suggestion = append(suggestion, "- Single word detected. Consider adding punctuation.")
	}

	if strings.Contains(input, "  ") { // Проверяем, нет ли подряд нескольких пробелов.
		suggestion = append(suggestion, "- Multiple consecutive spaces detected. Consider normalizing spaces.")
	}

	totalLines, maxWidth := calculateDimensions(input, font) // Вычисляем предполагаемые размеры будущего ASCII-арта.
	suggestion = append(suggestion, fmt.Sprintf("- Output dimensions: %d lines × %d characters.", totalLines, maxWidth))

	return suggestion // Возвращаем полный список найденных рекомендаций.
}

func calculateDimensions(input, font string) (int, int) { // Рассчитывает высоту и максимальную ширину будущего ASCII-арта.
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

	return totalLines, maxWidth // Возвращаем рассчитанную высоту и максимальную ширину результата.
}
