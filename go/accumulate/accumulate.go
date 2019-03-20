package accumulate

type function func(string) string

// Accumulate function return a list of strings
// after applying given operation on each element
func Accumulate(list []string, converter function) []string {
	for index, item := range list {
		list[index] = converter(item)
	}
	return list
}
