package main

// Specify Filter function here

func main() {

}

func Filter(input []int, predicate func(int, int) bool) []int {
	var output []int
	for index, element := range input {
		if predicate(element, index) {
			output = append(output, element)
		}
	}
	return output
}
