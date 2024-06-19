package helpers

import (
	"calc/src/constant"
)

func Reducing(operand string) int {
	return constant.Numbers[operand]
}
