/*
В качестве ответа на задание приложите ссылку на репозиторий.
Пожелания к программе:

Использовать методы и структуры пакетов ioutils и regexp.
Программа должна принимать на вход 2 аргумента: имя входного файла и имя файла для вывода результатов.
Если не найден вывод, создать.
Если файл вывода существует, очистить перед записью новых результатов.
Использовать буферизированную запись результатов.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите путь к входному файлу:")
	inFile, err := reader.ReadString('\n') // запрашиваем путь к входному файлу у пользователя
	inFile = strings.TrimSpace(inFile)     // удаляем символ переноса строки
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Введите путь к выходному файлу:")
	outFile, err := reader.ReadString('\n') // запрашиваем путь к выходному файлу у пользователя
	outFile = strings.TrimSpace(outFile)    // удаляем символ переноса строки
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(inFile) // читаем содержимое входного файла
	if err != nil {
		log.Fatal(err)
	}

	fileOut, err := os.Create(outFile) // создаем выходной файл
	if err != nil {
		log.Fatal(err)
	}
	defer fileOut.Close()

	writer := bufio.NewWriter(fileOut) // создаем буферизованный writer для записи в файл

	lines := strings.Split(string(content), "\n") // разбиваем содержимое файла на строки
	for _, line := range lines {                  // проходим по каждой строке файла
		result := parseLine(line) // вызываем функцию parseLine() для каждой строки файла
		if result != "" {
			writer.Write([]byte(result)) // записываем результаты вычислений в выходной файл
		}
	}

	writer.Flush() // записываем данные буфера в файл
	fmt.Println("Done!")
}

// parseLine ищет уравнение в строке и возвращает строку с результатом вычислений.
func parseLine(line string) string {
	re := regexp.MustCompile(`(\d+)([+-])(\d+)=\?`) // регулярное выражение для поиска уравнений
	matches := re.FindStringSubmatch(line)          // поиск уравнения в строке и получение групп в срезе

	if len(matches) == 4 { // если уравнение найдено
		operator := matches[2]
		n1, _ := strconv.Atoi(matches[1])
		n2, _ := strconv.Atoi(matches[3])

		if operator == "+" { // если оператор "+"
			return fmt.Sprintf("%d+%d=%d\n", n1, n2, n1+n2)
		}

		return fmt.Sprintf("%d-%d=%d\n", n1, n2, n1-n2)
	}
	return ""
}
