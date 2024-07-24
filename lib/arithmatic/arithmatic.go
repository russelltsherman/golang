package arithmatic

func Subtract(args ...int) int {
	if len(args) < 2 {
		return 0
	}

	result := args[0]

	for i := 1; i < len(args); i++ {
		result -= args[i]
	}

	return result
}

func Sum(args ...int) (result int) {
	for _, v := range args {
		result += v
	}
	return result
}

func Multiply(args ...int) int {
	if len(args) < 2 {
		return 0
	}
	result := 1

	for i := 0; i < len(args); i++ {
		result *= args[i]
	}

	return result
}
