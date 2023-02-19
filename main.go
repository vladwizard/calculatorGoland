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
	op := members[1]

	romanCount := 0
	num1, err := strconv.Atoi(members[0])
	if err != nil {
		romanCount += 1
	}
	num2, err := strconv.Atoi(members[2])
	if err != nil {
		romanCount += 1
	}
	if romanCount == 1 {
		log.Panicln(errors.New("Вывод ошибки, так как используются одновременно разные системы счисления."))
	} else if romanCount == 2 {
		num1 = toDecimal(members[0])
		num2 = toDecimal(members[2])
	}

	if !(1 <= num1 && num1 <= 10 && 1 <= num2 && num2 <= 10) {
		log.Panicln(errors.New("Вывод ошибки, так как входные числа не в диапазоне 1 - 10."))
	}

	var res int
	if op == "+" {
		res = num1 + num2
	} else if op == "-" {
		res = num1 - num2
	} else if op == "*" {
		res = num1 * num2
	} else if op == "/" {
		res = num1 / num2
	} else {
		log.Panicln(errors.New("Вывод ошибки, так как строка не является математической операцией."))
	}

	if romanCount == 2 {
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
	value int
	digit string
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

func toDecimal(num string) int {
	decimal := 0
	for i := 0; i < len(num); i++ {
		if num[i] == 'I' && i+1 != len(num) && num[i+1] == 'V' {
			decimal += 4
			i++
		} else if num[i] == 'I' && i+1 != len(num) && num[i+1] == 'X' {
			decimal += 9
			i++
		} else {
			for _, el := range dictionary {
				if el.digit == string(num[i]) {
					decimal += el.value
					break
				}
			}
		}
	}
	return decimal
}
func toRoman(num int) string {
	roman := ""
	for _, val := range dictionary {
		for num >= val.value {
			roman += val.digit
			num -= val.value
		}
	}
	return roman
}
