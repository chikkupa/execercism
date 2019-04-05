package listops

type IntList []int
type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

func (list IntList)Foldl(fn binFunc, initial int) int {
	for _, item := range list {
		initial = fn(initial, item)
	}
	return initial
}

func (list IntList)Foldr(fn binFunc, initial int) int {
	for i := len(list) - 1; i >= 0; i-- {
		initial = fn(initial, list[i])
	}
	return initial
}

func (list IntList)Filter(fn predFunc) IntList {
	filterList := []int{}

	for _, item := range list {
		if fn(item) {
			filterList = append(filterList, item)
		}
	}

	return filterList
}

func (list IntList)Length() int {
	var length int

	for _ = range list {
		length++
	}

	return length
}

func (list IntList)Map(fn unaryFunc) IntList {
	for i, item := range list {
		list[i] = fn(item)
	}
	return list
}

func (list IntList)Reverse() IntList {
	for i := 0; i < (len(list) + 1) / 2; i++ {
		list[i], list[len(list) - (i + 1)] = list[len(list) - (i + 1)], list[i]
	}
	return list
}

func (list IntList)Append(newList IntList) IntList {
	for _, item := range newList {
		list = append(list, item)
	}

	return list
}

func (list IntList)Concat(newLists []IntList) IntList {
	for _, newList := range newLists {
		for _, item := range newList {
			list = append(list, item)
		}
	}

	return list
}