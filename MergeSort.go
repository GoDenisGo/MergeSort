package MergeSort

// Sort - splits slice of length "n" elements, sorts them and returns the sorted slice (using Mergesort)
func Sort(Slice []int) []int {

	if len(Slice) <= 1 {
		return Slice
	} else {
		var LeftIndex = 0
		var RightIndex = 0
		var NewSlice []int
		var Mid = len(Slice) / 2

		// Splits the input slice into halves to recursively sort
		var Left = Slice[:Mid]
		var Right = Slice[Mid:]

		LeftSlice := Sort(Left)
		RightSlice := Sort(Right)

		for LeftIndex < len(LeftSlice) && RightIndex < len(RightSlice) {

			if LeftSlice[LeftIndex] <= RightSlice[RightIndex] {
				NewSlice = append(NewSlice, LeftSlice[LeftIndex])
				LeftIndex++
			} else {
				NewSlice = append(NewSlice, RightSlice[RightIndex])
				RightIndex++

			}
		}

		// Checks if any remaining elements need to be appended to the sorted slice
		if LeftIndex == len(LeftSlice) {
			NewSlice = append(NewSlice, RightSlice[RightIndex:]...)
		} else if RightIndex == len(RightSlice) {
			NewSlice = append(NewSlice, LeftSlice[LeftIndex:]...)
		}

		return NewSlice

	}

}
