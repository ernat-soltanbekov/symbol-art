package main

import (
	"fmt"
	"os"
	"strings"

	"symbol-art/internal/analyzer"
	"symbol-art/internal/printer"
	"symbol-art/internal/suggest"
)

func main() { // Точка входа программы. Здесь начинается выполнение всего приложения.
	args := os.Args[1:] // Получаем аргументы командной строки, пропуская имя самой программы.

	if len(args) < 1 { // Если пользователь ничего не передал, завершаем программу без выполнения дальнейшей логики.
		fmt.Fprintln(os.Stderr, "Ошибка: не передано ни одного аргумента.")
		fmt.Fprintln(os.Stderr, "Пример: go run \"Hello World\".")
		os.Exit(1)
	}

	if len(args) > 4 {
		fmt.Fprintln(os.Stderr, "Ошибка: передано слишком много аргументов.")
		fmt.Fprintln(os.Stderr, "Пример: go run . \"Hello World\" standard --analyze --suggest")
		os.Exit(1)
	}

	isSuggestMode := false // Флаг показывает, нужно ли после генерации вывести рекомендации.
	isAnalyzeMode := false // Флаг показывает, нужно ли после генерации вывести информацию об анализе.

	textArg := args[0] // Здесь будет храниться текст, который необходимо преобразовать в ASCII-арт.
	font := "standard.txt"
	fontSelected := false

	for _, symbol := range textArg {
		if symbol > 127 {
			fmt.Fprintln(os.Stderr, "Ошибка: программа поддерживает только ASCII символы.")
			os.Exit(1)
		}
	}

	// Обрабатываем дополнительные аргументы.
	if len(args) > 1 {
		for _, arg := range args[1:] {
			switch arg {
			case "--analyze":
				isAnalyzeMode = true

			case "--suggest":
				isSuggestMode = true

			case "standard", "shadow", "thinkertoy":
				if fontSelected {
					fmt.Fprintln(os.Stderr, "Ошибка: можно выбрать только один шрифт: standard, shadow или thinkertoy.")
					os.Exit(1)
				}

				fontSelected = true

				switch arg {
				case "standard":
					font = "standard.txt"
				case "shadow":
					font = "shadow.txt"
				case "thinkertoy":
					font = "thinkertoy.txt"
				}

			default:
				fmt.Fprintf(os.Stderr, "Ошибка: неизвестный аргумент %q\n", arg)
				os.Exit(1)
			}
		}
	}

	onlyLines := true // Предполагаем, что вход состоит только из пустых строк, пока не обнаружим текст.
	textArg = strings.ReplaceAll(textArg, `\n`, "\n")
	normArgs := strings.Split(textArg, "\n") // Разбиваем строку по последовательности "\n", чтобы обработать каждую строку отдельно.

	for _, arg := range normArgs { // Просматриваем каждую полученную строку.
		if arg != "" {
			onlyLines = false
		}
	}

	if onlyLines { // Если пользователь передал только пустые строки, выводим соответствующее количество пустых строк и завершаем работу.
		for i := 0; i < len(normArgs)-1; i++ {
			fmt.Println()
		}
		return
	}

	printer.PrintLines(os.Stdout, normArgs, font) // Передаем подготовленные строки модулю, который выводит ASCII-арт на экран.

	if isAnalyzeMode { // Если флаг был указан, печатаем анализ.
		for _, str := range analyzer.GetAnalyze(textArg, font) {
			fmt.Println(str)
		}
	}

	if isSuggestMode { // Если флаг был указан, печатаем рекомендации.
		for _, suggestion := range suggest.GetSuggestions(textArg, font) {
			fmt.Println(suggestion)
		}
	}

}
