package strain

// Ints Collection of integers
type Ints []int

// Strings Collection of strings
type Strings []string

// Lists Collection of lists
type Lists [][]int

// Keep function returns the list of values  matching the predicate
func (list Ints) Keep(pred func(int) bool) Ints {
	var keepList Ints

	for _, item := range list {
		if pred(item) {
			keepList = append(keepList, item)
		}
	}

	return keepList
}

// Discard function returns the list of values doesn't matching the predicate
func (list Ints) Discard(pred func(int) bool) Ints {
	var discardList Ints

	for _, item := range list {
		if !pred(item) {
			discardList = append(discardList, item)
		}
	}

	return discardList
}

// Keep function returns a list of list matching the predicate
func (lists Lists) Keep(pred func([]int) bool) Lists {
	var keepLists Lists

	for _, itemList := range lists {
		if pred(itemList) {
			keepLists = append(keepLists, itemList)
		}
	}

	return keepLists
}

// Keep function returns the list of values  matching the predicate
func (list Strings) Keep(pred func(string) bool) Strings {
	var keepList Strings

	for _, item := range list {
		if pred(item) {
			keepList = append(keepList, item)
		}
	}

	return keepList
}
