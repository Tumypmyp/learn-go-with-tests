package arrays

func Sum(numbers []int) int {
	add := func(res, x int) int {
		return res + x
	}
	return Reduce(numbers, add, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	count := func(res, x []int) []int {
		return append(res, Sum(x))
	}
	return Reduce(numbersToSum, count, []int{})
}

func SumAllTails(numbersToSum ...[]int) []int {
	sumTail := func(res, x []int) []int {
		if len(x) == 0 {
			return append(res, 0)
		} else {
			tail := x[1:]
			return append(res, Sum(tail))
		}
	}
	return Reduce(numbersToSum, sumTail, []int{})
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = accumulator(result, x)
	}
	return result
}

func Find[A any](collection []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range collection {
		if predicate(v) {
			return v, true
		}
	}
	return
}
