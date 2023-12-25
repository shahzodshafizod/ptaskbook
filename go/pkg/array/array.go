package array

////////////////////////////Int////////////////////////////

// Int struct type;
type Int []int

// NewInt is a constructor for Int struct type;
func NewInt(size int, value int) Int {
	var array = make([]int, size)
	if value != 0 {
		for index := range array {
			array[index] = value
		}
	}
	return array
}

// Reverse reverses the Int array;
func (array *Int) Reverse() {
	var length = len(*array)
	for index := 0; index < length/2; index++ {
		(*array)[index], (*array)[length-1-index] = (*array)[length-1-index], (*array)[index]
	}
}

// InsertBefore inserts new elements before the index position in the Int array;
func (array *Int) InsertBefore(index int, values ...int) {
	if index < 0 || index >= len(*array) {
		return
	}
	*array = append((*array)[:index], append(values, (*array)[index:]...)...)
}

// InsertAfter inserts new elements after the index position in the Int array;
func (array *Int) InsertAfter(index int, values ...int) {
	if index < 0 || index >= len(*array) {
		return
	}
	*array = append((*array)[:index+1], append(values, (*array)[index+1:]...)...)
}

// DeleteAt deletes the element of the Int array at the index position;
func (array *Int) DeleteAt(index int) (element int) {
	if index < 0 || index > len(*array) {
		return
	}
	element = (*array)[index]
	*array = append((*array)[:index], (*array)[index+1:]...)
	return
}

////////////////////////////Float64////////////////////////////

// Float64 struct type;
type Float64 []float64

// NewFloat64 is a constructor for Float64 struct type;
func NewFloat64(size int, value float64) Float64 {
	var array = make([]float64, size)
	if value != 0 {
		for index := range array {
			array[index] = value
		}
	}
	return array
}

// Reverse reverses the Float64 array;
func (array *Float64) Reverse() {
	var length = len(*array)
	for index := 0; index < length/2; index++ {
		(*array)[index], (*array)[length-1-index] = (*array)[length-1-index], (*array)[index]
	}
}

// InsertBefore inserts new elements before the index position in the Float64 array;
func (array *Float64) InsertBefore(index int, values ...float64) {
	if index < 0 || index >= len(*array) {
		return
	}
	*array = append((*array)[:index], append(values, (*array)[index:]...)...)
}

// InsertAfter inserts new elements after the index position in the Float64 array;
func (array *Float64) InsertAfter(index int, values ...float64) {
	if index < 0 || index >= len(*array) {
		return
	}
	*array = append((*array)[:index+1], append(values, (*array)[index+1:]...)...)
}

// DeleteAt deletes the element of the Float64 array at the index position;
func (array *Float64) DeleteAt(index int) float64 {
	if index < 0 || index > len(*array) {
		return 0
	}
	var element = (*array)[index]
	*array = append((*array)[:index], (*array)[index+1:]...)
	return element
}
