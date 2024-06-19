package constant

import (
	"calc/src/operations"
)

type OperationsList map[string]func(a, b int) int

var Operations = OperationsList{
	"+": operations.Addition,
	"-": operations.Subtraction,
	"/": operations.Division,
	"*": operations.Multiplication,
}
