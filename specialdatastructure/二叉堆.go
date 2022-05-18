package specialdatastructure

func heapSort2(arr []int) {
	buildHeap(arr)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		downAdjust(arr, 0, i)
	}
}

func downAdjust(arr []int, parent, length int) {
	leftChild := parent*2 + 1
	rightChild := parent*2 + 2
	for leftChild < length && rightChild < length {
		biggerIndex := leftChild
		if arr[rightChild] > arr[leftChild] {
			biggerIndex = rightChild
		}
		if arr[parent] >= arr[biggerIndex] {
			break
		}
		arr[parent], arr[biggerIndex] = arr[biggerIndex], arr[parent]
		parent = biggerIndex
		leftChild = parent*2 + 1
		rightChild = parent*2 + 1
	}
}

func buildHeap(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	lastNotLeafNodePos := (length - 2) / 2
	for i := lastNotLeafNodePos; i >= 0; i-- {
		downAdjust(arr, i, length)
	}
}
