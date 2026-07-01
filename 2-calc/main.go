package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
// Принимает неограниченное число чисел через запятую (2, 10, 9)
// Разбивает строку чисел по запятым и затем делает расчёт в зависимости от операции выводя результат
func main() {
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	var operation string
	fmt.Scan(&operation)
	numbersString, err := inputNumbersString()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	nums, err := parseCommaSeparatedInts(numbersString)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	result := calculate(operation, nums)
	fmt.Println("Результат:", result)
}

func inputNumbersString() (string, error) {
	fmt.Print("Введите числа через запятую: ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return "", scanner.Err()
	}

	return scanner.Text(), nil
}

func parseCommaSeparatedInts(s string) ([]int, error) {
	// Если строка пустая — возвращаем пустой слайс без ошибки (частый кейс)
	if s == "" {
		return []int{}, nil
	}

	parts := strings.Split(s, ",")

	result := make([]int, 0, len(parts)) // заранее резервируем ёмкость

	for i, part := range parts {
		part = strings.TrimSpace(part) // убираем пробелы вокруг числа
		num, err := strconv.Atoi(part)
		if err != nil {
			// Возвращаем понятную ошибку с позицией (полезно для логов/мониторинга)
			return nil, fmt.Errorf("некорректное число в позиции %d: %q (ошибка: %v)", i, part, err)
		}
		result = append(result, num)
	}

	return result, nil
}

func calculate(operation string, numbers []int) int {
	switch operation {
	case "AVG":
		return avg(numbers)
	case "SUM":
		return sum(numbers)
	case "MED":
		return med(numbers)
	default:
		return 0
	}
}

// avg возвращает среднее значение списка чисел
func avg(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum / len(numbers)
}

// sum возвращает сумму списка чисел
func sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// med возвращает медиану списка чисел
func med(numbers []int) int {
	sort.Ints(numbers)
	mid := len(numbers) / 2
	return numbers[mid]
}
