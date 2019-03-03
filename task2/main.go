package main

// Specify MapTo function here

func MapTo(arr []int, converter func(int, int) string) []string {
	var output []string
	for index, element := range arr {
		output = append(output, converter(element, index))
	}
	return output
}

func Convert(arr []int) []string {
	var digits = map[int]string{
		1: "one",
		2: "two",
		3: "three",
		4: "four",
		5: "five",
		6: "six",
		7: "seven",
		8: "eight",
		9: "nine",
	}

	cb := func(elem, _ int) string {
		digit, ok := digits[elem]
		if ok {
			return digit
		}
		return "unknown"
	}

	return MapTo(arr, cb)
}

func main() {
}
