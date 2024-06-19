package controller

import (
	"strconv"
	"strings"

	"calc/src/constant"
	"calc/src/helpers"
)

func calculation(a, b int, function func(a, b int) int) int {
	rangeExpression := a >= 0 && a <= 10 && b >= 0 && b <= 10
	if rangeExpression {
		return function(a, b)
	}
	panic("Выдача паники, так как формат математической операции не удовлетворяет заданию")
}

func Controller(expression string) interface{} {
	for operation, function := range constant.Operations {
		operationIndex := strings.Index(expression, operation)

		if operationIndex > 0 {
			expressionSplit := strings.Split(expression, operation)

			if len(expressionSplit) > 2 {
				panic("Выдача паники, так как формат математической операции не удовлетворяет заданию")
			}

			operand1, err1 := strconv.Atoi(expressionSplit[0])
			operand2, err2 := strconv.Atoi(expressionSplit[1])

			if err1 == nil && err2 == nil {
				return calculation(operand1, operand2, function)
			}
			if err1 != nil && err2 != nil {
				operand1 := helpers.Reducing(strings.ToLower(expressionSplit[0]))
				operand2 := helpers.Reducing(strings.ToLower(expressionSplit[1]))

				if operand1 > 0 && operand2 > 0 {
					return helpers.IntToRoman(calculation(operand1, operand2, function))
				}

			}
		}
	}
	panic("Выдача паники, так как используются одновременно разные системы счисления")
}
