package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the value of n: ")
	fmt.Scanf("%d", &n)

	arr := make([]int, n)

	acceptInput(arr, n)
	bubbleSort(arr, n)
	displayArray(arr, n)
}

func acceptInput(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Enter element %v :", i+1)
		fmt.Scanf("%v", &arr[i])
	}
}

func bubbleSort(arr []int, n int) {
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swapping elements
			}
		}
	}
}

func displayArray(arr []int, n int) {
	fmt.Println("Sorted array:")
	for i := 0; i < n; i++ {
		fmt.Println(arr[i])
	}
}
