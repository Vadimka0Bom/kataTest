package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romToArab = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
	"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10}

var arabToRom = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}

func main() {
	fmt.Print("Введите выражение(через пробелы): ")
	scanner := bufio.NewScanner(os.Stdin)
	_ = scanner.Scan()
	input := scanner.Text()

	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор")
	}

	strNumOne := parts[0]
	operator := parts[1]
	strNumTwo := parts[2]

	numOne, errOne := strconv.Atoi(strNumOne)
	numTwo, errTwo := strconv.Atoi(strNumTwo)

	roman := false

	if errOne != nil && errTwo != nil {
		numOne = romToArab[strNumOne]
		numTwo = romToArab[strNumTwo]
		roman = true
	} else if errOne != nil || errTwo != nil {
		panic("Выдача паники, так как используются одновременно разные системы счисления.")
	}

	if (numOne < 1 || numOne > 10) || (numTwo < 1 || numTwo > 10) {
		panic("Выдача паники, числа должны находится в диапозоне от 1 до 10(включительно)")
	}

	var result int
	switch operator {
	case "+":
		result = numOne + numTwo
	case "-":
		result = numOne - numTwo
	case "*":
		result = numOne * numTwo
	case "/":
		if numOne == 0 || numTwo == 0 {
			panic("Выдача паники, на ноль делить нельзя") // Все равно не сможем ввести 0, так как диапозон ввода ограничен заданием от 1 до 10.
		}
		result = numOne / numTwo
	}

	if roman {
		if result < 1 {
			panic("Выдача паники, так как в римской системе нет отрицательных чисел.")
		}
		fmt.Print("Результат: ", arabToRom[result])
	} else {
		fmt.Print("Результат: ", result)
	}
}
