package banner

import (
	"bufio"
	"os"
)

func Banner(filename string) ([]string, error) { // Загружает файл шрифта в память и возвращает его содержимое в виде среза строк.
	file, mistake := os.Open(filename) // Открываем файл баннера, который будем читать.
	if mistake != nil {                // Если открыть файл не удалось, сразу возвращаем ошибку вызывающему коду.
		return nil, mistake
	}
	
	defer file.Close() // После завершения функции обязательно закрываем файл, чтобы освободить системные ресурсы.

	var lines []string                // Здесь будут постепенно накапливаться все строки из файла.
	scanner := bufio.NewScanner(file) // Создаем сканер, который умеет читать файл построчно.
	for scanner.Scan() {              // Последовательно считываем каждую строку файла.
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil { // После чтения проверяем, не возникло ли ошибок во время работы сканера.
		return nil, err
	}

	return lines, nil // Возвращаем полностью загруженный файл в виде среза строк.
}

func getSymbolLines(lines []string, char rune) []string { // Находит восемь строк ASCII-арта, соответствующих одному символу.
	index := int(char) - int(' ') // Преобразуем символ в порядковый индекс относительно первого печатного символа ASCII.
	start := index*9 + 1          // Вычисляем номер строки, с которой начинается описание нужного символа в файле шрифта.

	return lines[start : start+8] // Возвращаем восемь строк, образующих изображение символа.
}

func renderWord(lines []string, word string) []string { // Собирает ASCII-арт для целого слова, объединяя изображения всех его символов.
	blocks := make([][]string, len(word)) // Подготавливаем место для хранения ASCII-блоков каждого символа слова.
	for i, char := range word {           // Последовательно получаем ASCII-представление каждого символа.
		blocks[i] = getSymbolLines(lines, char)
	}

	result := make([]string, 8) // Будущий результат всегда состоит из восьми строк.

	for stroka := 0; stroka < 8; stroka++ { // Формируем результат построчно, проходя сверху вниз.
		for _, block := range blocks { // На каждой строке последовательно объединяем части всех символов.
			result[stroka] += block[stroka]
		}
	}

	return result // Возвращаем полностью собранный ASCII-арт слова.
}

func Render(filename string, word string) ([]string, error) { // Главная функция модуля: загружает шрифт и строит ASCII-арт для указанного слова.
	lines, mistake := Banner(filename) // Сначала читаем файл шрифта в память.
	if mistake != nil {                // Если файл загрузить не удалось, сразу возвращаем ошибку.
		return nil, mistake
	}
	return renderWord(lines, word), nil // Генерируем ASCII-арт и возвращаем его вызывающему коду.
}
