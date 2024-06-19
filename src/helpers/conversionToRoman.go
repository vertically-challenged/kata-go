package helpers

func IntToRoman(num int) string {
	values := []int{
		100,
		90,
		50,
		40,
		10,
		9,
		5,
		4,
		1,
	}

	symbols := []string{
		"C",
		"XC",
		"L",
		"XL",
		"X",
		"IX",
		"V",
		"IV",
		"I",
	}

	roman := ""

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			roman += symbols[i]
			num -= values[i]
		}
	}

	if roman == "" {
		panic("Выдача паники, так как в римской системе нет отрицательных чисел")
	}
	return roman
}
