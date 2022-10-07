package helpers

import (
	"food/models"
)

func partition(arr []models.Response_restaurants, low, high int) ([]models.Response_restaurants, int) {

	pivot := arr[high-1]

	i := low
	for j := low; j < high; j++ {
		if arr[j].Distance < pivot.Distance {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high-1] = arr[high-1], arr[i]
	return arr, i
}

func QuickSort(arr []models.Response_restaurants, low, high int) []models.Response_restaurants {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = QuickSort(arr, low, p-1)
		arr = QuickSort(arr, p+1, high)
	}
	return arr
}
