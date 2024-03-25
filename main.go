package main

//тестовое задание (Стрельцов Георгий)
import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
	X, Y      int
	Operation Operate
}

type Operate func(int, int) int

var operators = map[string]Operate{
	"+": func(x, y int) int { return x + y },
	"-": func(x, y int) int { return x - y },
	"*": func(x, y int) int { return x * y },
	"/": func(x, y int) int { return x / y },
}

func (exp Expression) calculate() int {
	return exp.Operation(exp.X, exp.Y)
}

func main() {
	var x, operator, y, overkill string
	fmt.Fscanln(os.Stdin, &x, &operator, &y, &overkill)

	if x == "" || y == "" || operator == "" || overkill != "" {
		panic("Переданные данные не прошли валидацию")
	}

	if xInt, err := strconv.Atoi(x); err == nil {
		yInt, err := strconv.Atoi(y)
		_, ok := operators[operator]
		if err != nil || ok == false || (xInt < 1 || xInt > 10) || (yInt < 1 || yInt > 10) {
			panic("Переданные данные не прошли валидацию")
		}

		exp := Expression{xInt, yInt, operators[operator]}
		fmt.Println(exp.calculate())
	} else {
		exp := RomanToInt(x, operator, y)
		if exp.X == 0 || exp.Y == 0 || exp.Operation == nil {
			panic("Переданные данные не прошли валидацию")
		}
		if exp.calculate() < 1 {
			panic("Ответ меньше одного")
		} else {
			fmt.Println(IntToRoman(exp.calculate()))
		}
	}

}

func RomanToInt(x, operator, y string) Expression {

	var conversions = map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	var exp Expression

	exp.Operation = operators[operator]
	exp.X = conversions[x]
	exp.Y = conversions[y]

	return exp
}

func IntToRoman(number int) string {
	conversions := []struct {
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()

}
