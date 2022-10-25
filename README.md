# MergeSort

### Introduction
This is a demonstration of how the recursive Mergesort can be implemented using Golang.

The Mergesort algorithm is an example of the typical divide-and-conquer algorithms that are used to sort data.

The Mergesort algorithm halves an intial slice of unsorted elements until there are sub-slices of size 1, which are already sorted.
The algorithm then compares the elements in each sub-slice and merges them together in a sorted order.
This process of sorting and merging repeats until there is one sorted slice containing all the original elements.

The time complexity of O(NlogN) for the Mergesort algorithm is consistent across its Worst, Average and Best times. This makes the algorithm useful when consistency is key.

### Installation
To run this program, clone it from git and import it to a main project.

### Usage
Define a slice of integers (`[]int`) and pass it into the `Sort()` function.
The `Sort()` function returns the slice of elements in a sorted, ascending order.
You can view the sorted slice by returning the result of the `Sort()` function into a variable and printing the new slice.
An example of how to use the MergeSort package can be observed below: 

```golang
package main

import (
  "fmt"
  
  "example.com/MergeSort"
)

func main() {
  var myNumbers = []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
  var sorted = Mergesort.Sort(myNumbers)
  
  fmt.Println(sorted)
}
```
