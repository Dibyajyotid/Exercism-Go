package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
	for i, s := range slice {
        if i == index {
            return s
        }
    }

    return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
    if index < 0 || index >= len(slice) {
    	slice = append(slice, value) //appending at the end of the slice
    } else {
       slice[index] = value //replacing the value at the index
    }

    return slice
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
    if len(values) > 0 {
        slice = append(make([]int, len(values)), slice...) // Grow the slice to fit new values
        copy(slice, values) // Copy new values into the start
        return slice
    }
	return slice //if no value is not given
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    if index < 0 || index >= len(slice) {
        return slice
    }
	return append(slice[:index], slice[index+1:]...)
}
