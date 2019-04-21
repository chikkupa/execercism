package listops

// IntList custom type for list of integers
type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Foldl foldl function
func (list IntList) Foldl(fn binFunc, initial int) int {
	for _, item := range list {
		initial = fn(initial, item)
	}
	return initial
}

// Foldr foldr function
func (list IntList) Foldr(fn binFunc, initial int) int {
	for i := len(list) - 1; i >= 0; i-- {
		initial = fn(list[i], initial)
	}
	return initial
}

// Filter filter list satisfies the predicate function
func (list IntList) Filter(fn predFunc) IntList {
	filterList := []int{}

	for _, item := range list {
		if fn(item) {
			filterList = append(filterList, item)
		}
	}

	return filterList
}

// Length fuction to find the length of a list
func (list IntList) Length() int {
	var length int

	for range list {
		length++
	}

	return length
}

// Map equivalent to map function in functional programming
func (list IntList) Map(fn unaryFunc) IntList {
	for i, item := range list {
		list[i] = fn(item)
	}
	return list
}

// Reverse reverse the items in the list
func (list IntList) Reverse() IntList {
	for i := 0; i < (len(list)+1)/2; i++ {
		list[i], list[len(list)-(i+1)] = list[len(list)-(i+1)], list[i]
	}
	return list
}

// Append append a new item to the list
func (list IntList) Append(newList IntList) IntList {
	for _, item := range newList {
		list = append(list, item)
	}

	return list
}

// Concat concatenate a list of list to another list
func (list IntList) Concat(newLists []IntList) IntList {
	for _, newList := range newLists {
		for _, item := range newList {
			list = append(list, item)
		}
	}

	return list
}
