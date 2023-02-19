package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	members := strings.Fields(input)
	if len(members) > 3 {
		log.Panicln(errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."))
	} else if len(members) < 3 {
		log.Panicln(errors.New("Вывод ошибки, так как строка не является математической операцией."))
	}

	num1, num2, isRoman := handleNumbers(members[0], members[2])

	var res int
	switch op := members[1]; op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		log.Panicln(errors.New("Вывод ошибки, так как строка не является математической операцией."))
	}

	if isRoman {
		if res == 0 {
			log.Panicln(errors.New("Вывод ошибки, так как в римской системе нет нуля."))
		} else if res < 0 {
			log.Panicln(errors.New("Вывод ошибки, так как в римской системе нет отрицательных чисел."))
		}
		fmt.Println(toRoman(res))
	} else {
		fmt.Println(res)
	}
}

var dictionary = []struct {
	decimal int
	roman   string
}{
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func findInDictionary(romanDigit string) int {
	for _, el := range dictionary {
		if el.roman == romanDigit {
			return el.decimal
		}
	}
	return 0
}

func handleNumbers(text1 string, text2 string) (int, int, bool) {
	romanCount := 0
	num1, err := strconv.Atoi(text1)
	if err != nil {
		romanCount += 1
	}
	num2, err := strconv.Atoi(text2)
	if err != nil {
		romanCount += 1
	}
	if romanCount == 1 {
		log.Panicln(errors.New("Вывод ошибки, так как используются одновременно разные системы счисления."))
	} else if romanCount == 2 {
		num1 = toDecimal(text1)
		num2 = toDecimal(text2)
	}
	if !(1 <= num1 && num1 <= 10 && 1 <= num2 && num2 <= 10) {
		log.Panicln(errors.New("Вывод ошибки, так как входные числа не в диапазоне 1 - 10."))
	}
	return num1, num2, romanCount != 0
}
func toDecimal(num string) int {
	decimal := 0
	i := 0
	for ; i < len(num)-1; i++ {
		if num[i] == 'I' && num[i+1] == 'V' {
			decimal += 4
			i++
		} else if num[i] == 'I' && num[i+1] == 'X' {
			decimal += 9
			i++
		} else {
			decimal += findInDictionary(string(num[i]))
		}
	}
	decimal += findInDictionary(string(num[i]))
	return decimal
}
func toRoman(num int) string {
	roman := ""
	for _, val := range dictionary {
		for num >= val.decimal {
			roman += val.roman
			num -= val.decimal
		}
	}
	return roman
}
