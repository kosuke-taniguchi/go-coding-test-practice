package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func InsertionSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 1; i < lenNumbers; i++ {
		temp := numbers[i]
		j := i - 1
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}

	return numbers
}

func Partition(numbers []int, low int, high int) int {
	pivot := numbers[high]
	i := low - 1
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]

	return i + 1
}

func QuickSort(numbers []int, low int, high int) []int {
	if low < high {
		partitionIndex := Partition(numbers, low, high)
		QuickSort(numbers, low, partitionIndex-1)
		QuickSort(numbers, partitionIndex+1, high)
	}

	return numbers
}

func BinarySearch(numbers []int, value int) int {
	center := int(len(numbers) / 2)

	if numbers[center] == value {
		return center
	}

	if numbers[center] < value {
		right := numbers[center:]
		BinarySearch(right, value)
	}

	if numbers[center] > value {
		left := numbers[:center]
		BinarySearch(left, value)
	}

	return -1
}

func BinarySearchV2(numbers []int, value int) int {
	left := 0
	right := len(numbers) - 1

	for left <= right {
		mid := int((left + right) / 2)
		if numbers[mid] == value {
			return mid
		} else if numbers[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchV3(numbers []int, value int, left int, right int) int {
	if left > right {
		return -1
	}

	mid := int((left + right) / 2)
	if numbers[mid] == value {
		return mid
	} else if numbers[mid] < value {
		left = mid + 1
		return BinarySearchV3(numbers, value, left, right)
	} else {
		right = mid - 1
		return BinarySearchV3(numbers, value, left, right)
	}
}

func BubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		for j := 0; j < lenNumbers-1-i; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

func CocktailSort(numbers []int) []int {
	lenNumbers := len(numbers)
	start := 0
	end := lenNumbers - 1
	swapped := true
	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}
		end--
		swapped = false

		for i := end; i > start; i-- {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		start++
		swapped = true
	}

	return numbers
}

func SelectionSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		tempIndex := i
		for j := i; j < lenNumbers; j++ {
			if numbers[j] < numbers[tempIndex] {
				tempIndex = j
			}
		}
		numbers[i], numbers[tempIndex] = numbers[tempIndex], numbers[i]
	}

	return numbers
}

func CombSort(numbers []int) []int {
	lenNumbers := len(numbers)
	swapped := true
	gap := int(float32(lenNumbers) / float32(1.3))
	for swapped && gap > 0 {
		swapped = false
		gap = int(float32(gap) / float32(1.3))
		if gap < 1 {
			gap = 1
		}
		for i := 0; i < lenNumbers-gap; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}
	}

	return numbers
}

func GnomeSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers-1; i++ {
		j := i
		for j >= 0 && numbers[j] > numbers[j+1] {
			numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			j--
		}
	}

	return numbers
}

func GnomeSortV2(numbers []int) []int {
	lenNumbers := len(numbers)
	i := 0
	for i < lenNumbers-1 {
		if numbers[i] > numbers[i+1] {
			numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
			if i == 0 {
				i++
			} else {
				i--
			}
		} else {
			i++
		}
	}

	return numbers
}

func InsertionSortV2(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers-1; i++ {
		temp := numbers[i+1]
		j := i
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp
	}

	return numbers
}

func ShellSort(numbers []int) []int {
	lenNumbers := len(numbers)
	gap := int(lenNumbers / 2)
	for gap > 0 {
		if gap < 1 {
			gap = 1
		}

		for i := 0; i < lenNumbers-gap; i++ {
			temp := numbers[i+gap]
			j := i
			for j >= 0 && numbers[j] > temp {
				numbers[j+gap] = numbers[j]
				j--
			}
			numbers[j+gap] = temp
		}

		gap = int(gap / 2)
	}

	return numbers
}

func MaxNum(numbers []int) int {
	max := 0
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max
}

func CountingSort(numbers []int) []int {
	maxNum := MaxNum(numbers)
	countingList := make([]int, maxNum+1)
	result := make([]int, len(numbers))
	for _, num := range numbers {
		countingList[num]++
	}

	for i := 1; i < len(countingList); i++ {
		countingList[i] += countingList[i-1]
	}

	for _, num := range numbers {
		result[countingList[num]-1] = num
		countingList[num]--
	}

	return result
}

func PartitionV2(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}

	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func QuickSortV2(numbers []int, low int, high int) []int {
	if low < high {
		partitionIndex := PartitionV2(numbers, low, high)
		QuickSortV2(numbers, low, partitionIndex-1)
		QuickSortV2(numbers, partitionIndex+1, high)
	}

	return numbers
}

func MergeSort(numbers []int) []int {
	lenNumbers := len(numbers)
	if lenNumbers <= 1 {
		return numbers
	}

	center := int(lenNumbers / 2)
	left := numbers[:center]
	right := numbers[center:]

	MergeSort(left)
	MergeSort(right)

	i, j, k := 0, 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		numbers[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		numbers[k] = right[j]
		j++
		k++
	}

	return numbers
}

func MergeSortV2(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := int(len(numbers) / 2)
	left := numbers[:center]
	right := numbers[center:]

	MergeSortV2(left)
	MergeSortV2(right)

	return Merge(numbers, left, right)
}

func Merge(numbers []int, left []int, right []int) []int {
	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		numbers[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		numbers[k] = right[j]
		k++
		j++
	}

	return numbers
}

type Node struct {
	data     int
	nextNode *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(data int) {
	if l.head == nil {
		newNode := &Node{data, nil}
		l.head = newNode
		return
	}

	newNode := &Node{data, l.head}
	l.head = newNode
}

func (l *LinkedList) Append(data int) {
	newNode := &Node{data, nil}
	if l.head == nil {
		l.head = newNode
		newNode.nextNode = nil
		return
	}

	currentNode := l.head
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	currentNode.nextNode = newNode
}

func (l *LinkedList) Print() {
	currentNode := l.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func (l *LinkedList) Remove(data int) {
	currentNode := l.head
	if currentNode == nil {
		return
	}
	if currentNode != nil && currentNode.data == data {
		l.head = currentNode.nextNode
		currentNode = nil
		return
	}

	var prevNode *Node
	for currentNode != nil && currentNode.data != data {
		prevNode = currentNode
		currentNode = currentNode.nextNode
	}

	if currentNode == nil {
		return
	}

	prevNode.nextNode = currentNode.nextNode
	currentNode = nil
}

func (l *LinkedList) Reverse() {
	currentNode := l.head
	if currentNode == nil {
		return
	}

	var prevNode *Node
	for currentNode != nil {
		NextNode := currentNode.nextNode
		currentNode.nextNode = prevNode
		prevNode = currentNode
		currentNode = NextNode
	}

	l.head = prevNode
}

type NodeV2 struct {
	data     int
	prevNode *NodeV2
	nextNode *NodeV2
}

type DoublyLinkedList struct {
	head *NodeV2
}

func (d *DoublyLinkedList) Insert(data int) {
	newNode := &NodeV2{data, nil, nil}
	currentNode := d.head
	if currentNode == nil {
		d.head = newNode
		return
	}

	d.head.prevNode = newNode
	newNode.nextNode = d.head
	d.head = newNode
}

func (d *DoublyLinkedList) Print() {
	currentNode := d.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.nextNode
	}
}

func (d *DoublyLinkedList) Append(data int) {
	newNode := &NodeV2{data, nil, nil}
	currentNode := d.head
	if currentNode == nil {
		d.head = newNode
		return
	}
	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	newNode.prevNode = currentNode
	currentNode.nextNode = newNode
}

//This is not working
func (d *DoublyLinkedList) Remove(data int) {
	currentNode := d.head

	if currentNode != nil && currentNode.data == data {
		if currentNode.nextNode == nil {
			d.head = nil
			currentNode = nil
			return
		} else {
			NextNode := currentNode.nextNode
			d.head = NextNode
			NextNode.prevNode = nil
			currentNode = nil
			return
		}
	}

	for currentNode != nil && currentNode.data == data {
		currentNode = currentNode.nextNode
	}

	if currentNode == nil {
		return
	}

	if currentNode.nextNode == nil {
		currentNode.prevNode.nextNode = nil
		currentNode = nil
		return
	}

	PrevNode := currentNode.prevNode
	NextNode := currentNode.nextNode
	PrevNode.nextNode = NextNode
	NextNode.prevNode = PrevNode
	currentNode = nil
}

func (d *DoublyLinkedList) Reverse() {
	currentNode := d.head
	if currentNode == nil {
		return
	}

	var PrevNode *NodeV2
	for currentNode != nil {
		PrevNode = currentNode.prevNode
		currentNode.prevNode = currentNode.nextNode

		currentNode.nextNode = PrevNode
		currentNode = currentNode.prevNode
	}

	if PrevNode != nil {
		d.head = PrevNode.prevNode
	}
}

type BinarySearchNode struct {
	value int
	right *BinarySearchNode
	left  *BinarySearchNode
}

type BinarySearchTree struct {
	root *BinarySearchNode
}

func (bst *BinarySearchTree) Insert(value int) {
	if bst.root == nil {
		newNode := &BinarySearchNode{value: value, right: nil, left: nil}
		bst.root = newNode
	}

	insertNode(bst.root, value)
}

func insertNode(node *BinarySearchNode, value int) *BinarySearchNode {
	if node == nil {
		return &BinarySearchNode{value: value, right: nil, left: nil}
	}

	if node.value > value {
		node.left = insertNode(node.left, value)
	} else {
		node.right = insertNode(node.right, value)
	}

	return node
}

func (bst *BinarySearchTree) Inorder() {
	inorderNodes(bst.root)
}

func inorderNodes(node *BinarySearchNode) {
	if node != nil {
		inorderNodes(node.left)
		fmt.Println(node.value)
		inorderNodes(node.right)
	}
}

func (bst *BinarySearchTree) Search(value int) bool {
	if bst.root.value == value {
		return true
	}

	return searchNode(bst.root, value)
}

func searchNode(node *BinarySearchNode, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	}

	if node.value > value {
		return searchNode(node.left, value)
	} else {
		return searchNode(node.right, value)
	}
}

type BinarySearchTreeNode struct {
	value int
	left  *BinarySearchTreeNode
	right *BinarySearchTreeNode
}

type BinarySearchTreeV2 struct {
	root *BinarySearchTreeNode
}

func (bst *BinarySearchTreeV2) Insert(value int) {
	if bst.root == nil {
		bst.root = &BinarySearchTreeNode{value: value, right: nil, left: nil}
	}

	insertBSTNode(bst.root, value)
}

func insertBSTNode(node *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if node == nil {
		return &BinarySearchTreeNode{value: value, right: nil, left: nil}
	}

	if node.value > value {
		node.left = insertBSTNode(node.left, value)
	} else {
		node.right = insertBSTNode(node.right, value)
	}

	return node
}

func (bst *BinarySearchTreeV2) Inorder() {
	inorderBSTNodes(bst.root)
}

func inorderBSTNodes(node *BinarySearchTreeNode) {
	if node != nil {
		inorderBSTNodes(node.left)
		fmt.Println(node.value)
		inorderBSTNodes(node.right)
	}
}

func (bst *BinarySearchTreeV2) Search(value int) bool {
	return searchBSTNode(bst.root, value)
}

func searchBSTNode(node *BinarySearchTreeNode, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	}

	if value < node.value {
		return searchBSTNode(node.left, value)
	} else {
		return searchBSTNode(node.right, value)
	}
}

func (bst *BinarySearchTreeV2) Remove(value int) {
	removeBSTNode(bst.root, value)
}

func removeBSTNode(node *BinarySearchTreeNode, value int) *BinarySearchTreeNode {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = removeBSTNode(node.left, value)
	} else if value > node.value {
		node.right = removeBSTNode(node.right, value)
	} else {
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}

		minValue := searchMinValue(node.right)
		node.value = minValue.value
		node.right = removeBSTNode(node.right, minValue.value)
	}

	return node
}

func searchMinValue(node *BinarySearchTreeNode) *BinarySearchTreeNode {
	currentNode := node
	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode
}

type BinarySearchTreeNodeV3 struct {
	value int
	left  *BinarySearchTreeNodeV3
	right *BinarySearchTreeNodeV3
}

type BinarySearchTreeV3 struct {
	root *BinarySearchTreeNodeV3
}

func (bst *BinarySearchTreeV3) Insert(value int) {
	if bst.root == nil {
		bst.root = &BinarySearchTreeNodeV3{value: value, left: nil, right: nil}
	}
	insertNodeV3(bst.root, value)
}

func insertNodeV3(node *BinarySearchTreeNodeV3, value int) *BinarySearchTreeNodeV3 {
	if node == nil {
		return &BinarySearchTreeNodeV3{value: value, left: nil, right: nil}
	}

	if value < node.value {
		node.left = insertNodeV3(node.left, value)
	} else {
		node.right = insertNodeV3(node.right, value)
	}

	return node
}

func (bst *BinarySearchTreeV3) Inorder() {
	inorderNodesV3(bst.root)
}

func inorderNodesV3(node *BinarySearchTreeNodeV3) {
	if node != nil {
		inorderNodesV3(node.left)
		fmt.Println(node.value)
		inorderNodesV3(node.right)
	}
}

func (bst *BinarySearchTreeV3) Search(value int) bool {
	return searchNodeV3(bst.root, value)
}

func searchNodeV3(node *BinarySearchTreeNodeV3, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	}

	if value < node.value {
		return searchNodeV3(node.left, value)
	} else {
		return searchNodeV3(node.right, value)
	}
}

func (bst *BinarySearchTreeV3) Remove(value int) {
	removeNodeV3(bst.root, value)
}

func removeNodeV3(node *BinarySearchTreeNodeV3, value int) *BinarySearchTreeNodeV3 {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = removeNodeV3(node.left, value)
	} else if value > node.value {
		node.right = removeNodeV3(node.right, value)
	} else {
		if node.left == nil {
			node = node.right
		} else if node.right == nil {
			node = node.left
		} else {
			minValue := searchMinValueV3(node.right)
			node = minValue
			node.right = removeNodeV3(node.right, minValue.value)
		}
	}

	return node
}

func searchMinValueV3(node *BinarySearchTreeNodeV3) *BinarySearchTreeNodeV3 {
	for node.left != nil {
		node = node.left
	}

	return node
}

func MergeSortV3(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := int(len(numbers) / 2)
	left := numbers[:center]
	right := numbers[center:]

	MergeSortV3(left)
	MergeSortV3(right)

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		numbers[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		numbers[k] = right[j]
		k++
		j++
	}

	return numbers
}

func bubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		for j := 0; j < lenNumbers-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}

func cockTailSort(numbers []int) []int {
	lenNumbers := len(numbers)
	swapped := true
	start := 0
	end := lenNumbers - 1
	for swapped {
		swapped = false
		for i := start; i < end; i++ {
			if numbers[i] > numbers[i+1] {
				numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
				swapped = true
			}
		}

		if !swapped {
			break
		}

		start++
		swapped = false
		for i := end; i > start; i-- {
			if numbers[i-1] > numbers[i] {
				numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
				swapped = true
			}
		}

		end--
	}

	return numbers
}

func selectionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		tempIndex := i
		for j := i; j < len(numbers); j++ {
			if numbers[tempIndex] > numbers[j] {
				tempIndex = j
			}
		}
		numbers[tempIndex], numbers[i] = numbers[i], numbers[tempIndex]
	}
	return numbers
}

func gnomeSort(numbers []int) []int {
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers-1; i++ {
		j := i
		for j >= 0 && numbers[j] > numbers[j+1] {
			numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			j--
		}
	}
	return numbers
}

func insertionSort(numbers []int) []int {
	for i := 0; i < len(numbers)-1; i++ {
		temp := numbers[i+1]
		j := i
		for j >= 0 && numbers[j] > temp {
			numbers[j+1] = numbers[j]
			j--
		}
		numbers[j+1] = temp

	}
	return numbers
}

func shellSort(numbers []int) []int {
	lenNumbers := len(numbers)
	gap := int(lenNumbers / 2)
	for gap > 0 {
		if gap < 1 {
			gap = 1
		}

		for i := 0; i < lenNumbers-gap; i++ {
			temp := numbers[i+gap]
			j := i
			for j >= 0 && numbers[j] > temp {
				numbers[j+gap] = numbers[j]
				j -= gap
			}
			numbers[j+gap] = temp
		}
		gap /= 2
	}
	return numbers
}

func searchMax(numbers []int) int {
	maxNum := -1000000000000
	for _, i := range numbers {
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum
}

func countingSort(numbers []int) []int {
	lenNumbers := len(numbers)
	maxNum := searchMax(numbers)
	countingList := make([]int, maxNum+1)
	for _, num := range numbers {
		countingList[num]++
	}

	for i := 1; i < maxNum+1; i++ {
		countingList[i] += countingList[i-1]
	}

	result := make([]int, lenNumbers)
	for _, num := range numbers {
		indexNum := countingList[num] - 1
		result[indexNum] = num
		countingList[num]--
	}

	return result
}

func partition(numbers []int, low int, high int) int {
	pivot := numbers[high]
	i := low - 1
	for j := low; j < high; j++ {
		if numbers[j] < pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]
	return i + 1
}

func innerQuickSort(numbers []int, low int, high int) []int {
	if low < high {
		partitionIndex := partition(numbers, low, high)
		innerQuickSort(numbers, low, partitionIndex-1)
		innerQuickSort(numbers, partitionIndex+1, high)
	}

	return numbers
}

func quickSort(numbers []int) []int {
	return innerQuickSort(numbers, 0, len(numbers)-1)
}

func mergeSort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	center := int(len(numbers) / 2)
	left := mergeSort(numbers[:center])
	right := mergeSort(numbers[center:])

	i, j, k := 0, 0, 0
	for i < len(left) && j < len(right) && k < len(numbers) {
		if left[i] < right[j] {
			numbers[k] = left[i]
			i++
		} else {
			numbers[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		numbers[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		numbers[k] = right[j]
		j++
		k++
	}

	return numbers
}

func BinarySearchV4(numbers []int, target int) int {
	left := 0
	right := len(numbers) - 1
	for left <= right {
		mid := int((left + right) / 2)
		if numbers[mid] == target {
			return mid
		} else if numbers[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

func MergeSortV4(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := int(len(nums) / 2)
	left := nums[:mid]
	right := nums[mid:]

	left = MergeSortV4(left)
	right = MergeSortV4(right)

	result := make([]int, 0)
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}

	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}

func isPalindrome(x int) bool {
	xStr := strconv.Itoa(x)
	return xStr == reverse(xStr)
}

func reverse(str string) string {
	var result string
	for i := len(str); i > 0; i-- {
		result += string(str[i-1])
	}
	return result
}

func isPalindromeV2(value int) bool {
	digitNums := []int{}

	for value > 0 {
		digitNums = append(digitNums, int(value%10))
		value /= 10
	}

	for i := 0; i < int(len(digitNums)/2); i++ {
		if digitNums[i] != digitNums[len(digitNums)-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeV3(value int) bool {
	digitNums := []int{}

	for value > 0 {
		digitNums = append(digitNums, int(value%10))
		value /= 10
	}

	reverseDigitNums := []int{}
	for i := len(digitNums) - 1; i >= 0; i-- {
		reverseDigitNums = append(reverseDigitNums, digitNums[i])
	}

	for i := 0; i < len(digitNums); i++ {
		if digitNums[i] != reverseDigitNums[i] {
			return false
		}
	}

	return true
}

func longestCommonPrefix(strs []string) string {
	prefixCount := map[string]int{}

	for _, str := range strs {
		var tempStr string
		for _, s := range str {
			tempStr += string(s)
			prefixCount[tempStr]++
		}
	}

	var result string
	for k, v := range prefixCount {
		if v == len(strs) && len(k) > len(result) {
			result = k
		}
	}

	return result
}

func lengthOfLastWord(s string) int {
	splitedText := strings.Split(s, " ")
	lastWord := splitedText[len(splitedText)-1]

	i := 0
	for lastWord == "" {
		fmt.Println(lastWord)
		i++
		lastWord = splitedText[len(splitedText)-i-1]
	}

	return len(lastWord)
}

func plusOne(digits []int) []int {
	var result []int
	var num int
	digitNum := len(digits) - 1

	for _, digit := range digits {
		num += digit * int(math.Pow(float64(10), float64(digitNum)))
		digitNum--
	}
	num++

	dNum := int(math.Pow(float64(10), float64(len(digits)-1)))
	for dNum >= 1 {
		resultEl := int(num / dNum)
		result = append(result, resultEl)
		num -= resultEl * dNum
		dNum /= 10
	}

	return result
}

func plusOneV2(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		} else {
			digits[i] = 0
		}
	}

	result := make([]int, len(digits)+1)
	result[0] = 1
	return result
}

func climbStairs(n int) int {
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	currentNode := head
	for currentNode != nil {
		if currentNode.Val == currentNode.Next.Val {
			currentNode.Next = currentNode.Next.Next
			currentNode = currentNode.Next
		} else {
			currentNode = currentNode.Next
		}
	}
	return head
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := 0, 0, 0
	nums0 := nums1[:m]

	for i < m && j < n {
		if nums0[i] < nums2[j] {
			nums1[k] = nums0[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}

	for i < m {
		nums1[k] = nums0[i]
		i++
		k++
	}

	for j < n {
		nums1[k] = nums2[j]
		j++
		k++
	}

	fmt.Println(nums1)
}

func mergeV2(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m == 0 || nums1[m-1] < nums2[n-1] {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			nums1[m+n-1] = nums1[m-1]
			m--
		}
	}

	fmt.Println(nums1)
}

type BinarySearchTreeNodeV4 struct {
	value int
	left  *BinarySearchTreeNodeV4
	right *BinarySearchTreeNodeV4
}

func InsertBSTNodeV4(node *BinarySearchTreeNodeV4, value int) *BinarySearchTreeNodeV4 {
	if node == nil {
		return &BinarySearchTreeNodeV4{value: value, left: nil, right: nil}
	}

	if value < node.value {
		node.left = InsertBSTNodeV4(node.left, value)
	} else {
		node.right = InsertBSTNodeV4(node.right, value)
	}

	return node
}

func InorderBSTNodeV4(node *BinarySearchTreeNodeV4) {
	if node != nil {
		InorderBSTNodeV4(node.left)
		fmt.Println(node.value)
		InorderBSTNodeV4(node.right)
	}
}

func SearchBSTNodeV4(node *BinarySearchTreeNodeV4, value int) bool {
	if node == nil {
		return false
	}

	if node.value == value {
		return true
	} else if node.value > value {
		return SearchBSTNodeV4(node.left, value)
	} else {
		return SearchBSTNodeV4(node.right, value)
	}
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		inorderTraversal(root.Left)
		result = append(result, root.Val)
		inorderTraversal(root.Right)
	}

	return result
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return ish(root.Left, root.Right)
}

func ish(l, r *TreeNode) bool {
	if l == nil || r == nil {
		if l == nil && r == nil {
			return true
		} else {
			return false
		}
	}

	if l.Val != r.Val {
		return false
	}

	return ish(l.Left, r.Right) && ish(l.Right, r.Left)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)

	if leftDepth > rightDepth {
		rightDepth++
		return rightDepth
	} else {
		leftDepth++
		return leftDepth
	}
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j == i {
				continue
			}

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

func isPalindromeByString(value int) bool {
	valueStr := strconv.Itoa(value)
	lenValueStr := len(valueStr)

	for i := 0; i <= int(lenValueStr/2); i++ {
		if valueStr[i] != valueStr[len(valueStr)-i-1] {
			return false
		}
	}

	return true
}

func isPalindromeByInteger(value int) bool {
	valueList := []int{}
	for value > 0 {
		valueList = append(valueList, int(value%10))
		value /= 10
	}

	for i := 0; i <= len(valueList)-1; i++ {
		if valueList[i] != valueList[len(valueList)-1-i] {
			return false
		}
	}

	return true
}

func longestCommonPrefixV2(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}
	var prefix string
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[i-1]) && j < len(strs[i]) && j < len(prefix); j++ {
			if i == 1 && strs[i-1][:j] == strs[i][:j] {
				prefix = strs[i-1][:j]
			} else {
				if prefix[:j] == strs[i][:j] {
					prefix = prefix[:j]
				}
			}
		}
	}

	return prefix
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// incorrect
func sortedArrayToBST(nums []int) *TreeNode {
	mid := len(nums) / 2
	midNum := nums[mid]
	right := nums[mid+1:]
	left := nums[:mid]
	root := &TreeNode{
		Val: midNum,
	}

	var currentNode *TreeNode
	var newNode *TreeNode

	for i, num := range left {
		newNode = &TreeNode{Val: num}
		if i == 0 {
			currentNode = root
		}
		if num < currentNode.Val {
			currentNode.Left = newNode
		} else {
			currentNode.Right = newNode
		}

		currentNode = newNode
	}

	for i, num := range right {
		newNode = &TreeNode{Val: num}
		if i == 0 {
			currentNode = root
		}
		if num < currentNode.Val {
			currentNode.Left = newNode
		} else {
			currentNode.Right = newNode
		}

		currentNode = newNode
	}

	return root
}

func sortedArrayToBSTV2(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[len(nums)/2],
		Left:  sortedArrayToBSTV2(nums[:len(nums)/2]),
		Right: sortedArrayToBSTV2(nums[len(nums)/2+1:]),
	}
}

func generate(numRows int) [][]int {
	results := make([][]int, 0, numRows)
	for i := 0; i < numRows; i++ {
		if i == 0 {
			results = append(results, []int{1})
			continue
		}

		currentRow := make([]int, 0, i+1)
		prevRow := results[i-1]
		for j := 0; j < i+1; j++ {
			if j == 0 {
				currentRow = append(currentRow, prevRow[j])
			} else if j == i {
				currentRow = append(currentRow, prevRow[j-1])
			} else {
				currentRow = append(currentRow, prevRow[j]+prevRow[j-1])
			}
		}

		results = append(results, currentRow)
	}

	return results
}

func getRow(rowIndex int) []int {
	results := make([][]int, 0, rowIndex+1)
	for i := 0; i < rowIndex+1; i++ {
		elem := make([]int, 0, i+1)
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				elem = append(elem, 1)
			} else {
				elem = append(elem, results[i-1][j-1]+results[i-1][j])
			}
		}
		results = append(results, elem)
	}
	return results[rowIndex]
}

func maxProfit(prices []int) int {
	var profit int
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			currentProfit := prices[j] - prices[i]
			if profit < currentProfit {
				profit = currentProfit
			}
		}
	}

	if profit < 0 {
		return 0
	}
	return profit
}

func maxProfitV2(prices []int) int {
	var max int
	minSoFar := prices[0]
	for _, price := range prices {
		if minSoFar > price {
			minSoFar = price
		} else {
			currentProfit := price - minSoFar
			if max < currentProfit {
				max = currentProfit
			}
		}
	}

	return max
}

func singleNumber(nums []int) int {
	result := map[int]int{}
	for _, num := range nums {
		_, ok := result[num]
		if !ok {
			result[num] = 1
		} else {
			result[num]++
		}
	}

	for k, v := range result {
		if v == 1 {
			return k
		}
	}

	return 0
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	l1, l2 := head, head.Next
	for l1 != l2 {
		if l2 == nil || l2.Next == nil {
			return false
		}

		l1 = l1.Next
		l2 = l2.Next.Next
	}

	return true
}

func hasCycleV2(head *ListNode) bool {
	visitedMap := make(map[*ListNode]struct{})

	for head != nil {
		_, ok := visitedMap[head]
		if ok {
			return true
		}

		visitedMap[head] = struct{}{}
		head = head.Next
	}

	return false
}

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	result = append(result, root.Val)
	result = append(result, preorderTraversal(root.Left)...)
	result = append(result, preorderTraversal(root.Right)...)

	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var result []int
	result = append(result, postorderTraversal(root.Left)...)
	result = append(result, postorderTraversal(root.Right)...)
	result = append(result, root.Val)

	return result
}

func majorityElement(nums []int) int {
	numCount := make(map[int]int, 0)
	for _, num := range nums {
		_, ok := numCount[num]
		if !ok {
			numCount[num] = 1
			continue
		}

		numCount[num]++
	}

	for k, v := range numCount {
		if v > len(nums)/2 {
			return k
		}
	}

	return 0
}

func majorityElementV2(nums []int) int {
	numCount := make(map[int]int)
	for _, num := range nums {
		numCount[num]++
		if numCount[num] > len(nums)/2 {
			return num
		}
	}

	return -1
}

func isHappy(n int) bool {
	if n == 1 {
		return true
	}

	history := make(map[int]struct{}, 0)

	for {
		n, history = sumSqrDigits(n, history)
		if n == 1 {
			return true
		}
		if n == -1 {
			return false
		}
	}

	return false
}

func sumSqrDigits(n int, history map[int]struct{}) (int, map[int]struct{}) {
	if n%10 == 0 && n/10 == 0 {
		return 1, history
	}

	if _, ok := history[n]; ok {
		return -1, history
	} else {
		history[n] = struct{}{}
	}

	var result int
	for n > 0 {
		result += int(math.Pow(float64(n%10), 2))
		n /= 10
	}

	return result, history
}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sResult := strToInt(s)
	tResult := strToInt(t)

	for i := 0; i < len(s); i++ {
		if sResult[i] != tResult[i] {
			return false
		}
	}

	return true
}

func strToInt(strs string) []int {
	var idx int
	strsMap := make(map[string]int, 0)
	newStrs := strings.Split(strs, "")
	result := make([]int, 0)
	for _, str := range newStrs {
		if _, ok := strsMap[str]; !ok {
			strsMap[str] = idx
			idx++
		}

		result = append(result, strsMap[str])
	}

	return result
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	currentNode := head
	var prevNode *ListNode
	for currentNode.Next != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}

	currentNode.Next = prevNode

	return currentNode
}

func containsDuplicate(nums []int) bool {
	countMap := make(map[int]struct{}, 0)
	for _, num := range nums {
		if _, ok := countMap[num]; ok {
			return true
		} else {
			countMap[num] = struct{}{}
		}
	}

	return false
}

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < i+k+1; j++ {
			if j >= len(nums) {
				continue
			}
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsNearbyDuplicateV2(nums []int, k int) bool {
	contained := make(map[int]int, 0)
	for i, num := range nums {
		if n, ok := contained[num]; ok && i-n <= k {
			return true
		}
		contained[num] = i
	}

	return false
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 考え方はあってるけど下記だと、root.Leftが上書きされて、
	// root.Rightにも上書きされたroot.Left(もとはroot.Right)が入ってきちゃう
	root.Left = invertTree(root.Right)
	root.Right = invertTree(root.Left)

	return root
}

func invertTreeV2(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = invertTreeV2(root.Right), invertTreeV2(root.Left)
	return root
}

func invertTreeV3(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	// 一旦待避させる
	left := root.Left
	right := root.Right

	root.Left = invertTreeV3(right)
	root.Right = invertTreeV3(left)

	return root
}

func summaryRanges(nums []int) []string {
	var result []string
	if len(nums) == 0 {
		return result
	}

	var startTmp int
	var tmp int
	for i, num := range nums {
		if i == 0 {
			startTmp = num
			tmp = num
			continue
		}

		if tmp+1 == num {
			tmp = num
		} else {
			if startTmp == tmp {
				result = append(result, fmt.Sprintf("%d", startTmp))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", startTmp, tmp))
			}
			startTmp = num
			tmp = num
		}
	}

	if startTmp == tmp {
		result = append(result, fmt.Sprintf("%d", startTmp))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", startTmp, tmp))
	}

	return result
}

// 2進数に変換して1の個数で判定
func isPowerOfTwo(n int) bool {
	nBinary := strings.Split(strconv.FormatInt(int64(n), n), "")
	binList := make([]int, len(nBinary))
	for _, num := range nBinary {
		numInt, _ := strconv.Atoi(num)
		binList = append(binList, numInt)
	}

	var count int
	for _, bin := range binList {
		if bin == 1 {
			count++
		}
	}

	return count == 1
}

func isPowerOfTwoV2(n int) bool {
	if n == 0 {
		return false
	}

	var num float64
	var i int
	for num < float64(n) {
		num = math.Pow(2, float64(i))
		i++
	}

	return int(num) == n
}

// not correct
func isPalindromeListNode(head *ListNode) bool {
	if head == nil {
		return false
	}

	// count length linked list
	var lenNode int
	currentNode := head
	for currentNode != nil {
		lenNode++
		currentNode = currentNode.Next
	}

	currentNode = head

	lenHalf := lenNode / 2
	stackList := make([]int, 0, lenHalf)
	for i := 0; i < lenHalf; i++ {
		stackList = append([]int{currentNode.Val}, stackList...)
		currentNode = currentNode.Next
	}

	if lenNode%2 != 0 {
		currentNode = currentNode.Next
	}

	for _, val := range stackList {
		if val != currentNode.Val {
			return false
		}
		currentNode = currentNode.Next
	}

	return true
}

func isPalindromeListNodeV2(head *ListNode) bool {
	if head == nil {
		return false
	}

	currentNode := head
	valList := make([]int, 0)
	for currentNode != nil {
		valList = append(valList, currentNode.Val)
		currentNode = currentNode.Next
	}

	for i := 0; i < len(valList)/2; i++ {
		if valList[i] != valList[len(valList)-i-1] {
			return false
		}
	}

	return true
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	} else {
		return root
	}
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sStrs := strings.Split(s, "")
	for _, sStr := range sStrs {
		t = strings.Replace(t, sStr, "", 1)
	}

	return t == ""
}

// faster than v1
func isAnagramV2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[int32]int, 0)
	tMap := make(map[int32]int, 0)

	for _, sStr := range s {
		sMap[sStr]++
	}

	for _, tStr := range t {
		tMap[tStr]++
	}

	for k, _ := range sMap {
		if sMap[k] != tMap[k] {
			return false
		}
	}

	return true
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	binaryTreePathsHelper(root, fmt.Sprintf("%d", root.Val), &result)
	return result
}

func binaryTreePathsHelper(root *TreeNode, path string, result *[]string) {
	if root.Right == nil && root.Left == nil {
		*result = append(*result, path)
		return
	}

	if root.Left != nil {
		leftPath := path + fmt.Sprintf("->%d", root.Left.Val)
		binaryTreePathsHelper(root.Left, leftPath, result)
	}

	if root.Right != nil {
		rightPath := path + fmt.Sprintf("->%d", root.Right.Val)
		binaryTreePathsHelper(root.Right, rightPath, result)
	}
}

func addDigits(num int) int {
	result := addDigitsHelper(num)
	for result/10 != 0 {
		result = addDigitsHelper(result)
	}

	return result
}

func addDigitsHelper(tmp int) int {
	var result int
	for tmp >= 1 {
		result += tmp % 10
		tmp /= 10
	}

	return result
}

func addDigitsV2(num int) int {
	for num/10 > 0 {
		var sum int
		for num >= 1 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}

	return num
}

// not correct
func isUgly(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			if !isUglyHelper(i) {
				return false
			}
		}
	}

	return true
}

func isUglyHelper(n int) bool {
	nums := []int{2, 3, 5}
	for _, num := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func isUglyV2(n int) bool {
	primeFactors := []int{2, 3, 5}
	for _, pf := range primeFactors {
		for n%pf == 0 {
			n /= pf
		}
	}
	return n == 1
}

func isUglyV3(n int) bool {
	if n == 0 {
		return false
	}

	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else if n%3 == 0 {
			n /= 3
		} else if n%5 == 0 {
			n /= 5
		} else {
			return false
		}
	}

	return true
}

func missingNumber(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		if !isInMissingNumber(i, nums) {
			return i
		}
	}

	return -1
}

func isInMissingNumber(n int, nums []int) bool {
	for _, num := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func missingNumberV2(nums []int) int {
	var sum int
	for i := 0; i <= len(nums); i++ {
		sum += i
	}

	var numSum int
	for _, num := range nums {
		numSum += num
	}

	return sum - numSum
}

func missingNumberV3(nums []int) int {
	var sum int
	var sumNums int
	for i, num := range nums {
		sumNums += num
		sum += i + 1
	}

	return sum - sumNums
}

// incorrect
func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			for j := i; j < len(nums)-i-1; j++ {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func moveZeroesV2(nums []int) {
	var n int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[n] = nums[n], nums[i]
			n++
		}
	}

	fmt.Println(nums)
}

// 少し冗長
// reflect.DeepEqual()はずるいかな
func wordPattern(pattern string, s string) bool {
	sList := strings.Split(s, " ")
	sMap := make(map[string]int, 0)
	sIntList := make([]int, 0, len(sList))
	var n int
	for _, str := range sList {
		_, ok := sMap[str]
		if !ok {
			sMap[str] = n
			n++
		}
		sIntList = append(sIntList, sMap[str])
	}

	pList := strings.Split(pattern, "")
	pMap := make(map[string]int, 0)
	pIntList := make([]int, 0, len(pattern))
	var m int
	for _, str := range pList {
		_, ok := pMap[str]
		if !ok {
			pMap[str] = m
			m++
		}
		pIntList = append(pIntList, pMap[str])
	}

	if len(sIntList) != len(pIntList) {
		return false
	}

	for i := 0; i < len(sIntList); i++ {
		if sIntList[i] != pIntList[i] {
			return false
		}
	}

	return true
}

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	if n == 1 {
		return true
	}
	return false
}

// もっと良い方法ある
func countBits(n int) []int {
	results := make([]int, 0, n+1)
	for i := 0; i < n+1; i++ {
		binIStr := strconv.FormatInt(int64(i), 2)
		binIStrList := strings.Split(binIStr, "")
		var count int
		for _, b := range binIStrList {
			if b == "1" {
				count++
			}
		}
		results = append(results, count)
	}

	return results
}

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}

	if n != 1 {
		return false
	}

	return true
}

// easy solution
func reverseString(s []byte) {
	result := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	fmt.Println(result)
}

// correct answer
func reverseStringV2(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}

	fmt.Println(s)
}

func reverseVowels(s string) string {
	sSlice := strings.Split(s, "")
	var i int
	j := len(sSlice) - 1
	for i < j {
		if isInVowels(sSlice[i]) && isInVowels(sSlice[j]) {
			sSlice[i], sSlice[j] = sSlice[j], sSlice[i]
			i++
			j--
			continue
		}
		if !isInVowels(sSlice[i]) {
			i++
		}
		if !isInVowels(sSlice[j]) {
			j--
		}
	}

	return strings.Join(sSlice, "")
}

func isInVowels(s string) bool {
	vowels := []string{"a", "i", "u", "e", "o", "A", "I", "U", "E", "O"}
	for _, v := range vowels {
		if v == s {
			return true
		}
	}

	return false
}

// 少し汚い
func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]struct{}, 0)
	for _, n := range nums1 {
		_, ok := numMap[n]
		if !ok {
			numMap[n] = struct{}{}
		}
	}

	result := make([]int, 0)
	nums2Map := make(map[int]struct{}, 0)
	for _, num := range nums2 {
		_, ok := numMap[num]
		_, ok2 := nums2Map[num]
		if ok && !ok2 {
			result = append(result, num)
			nums2Map[num] = struct{}{}
		}
	}

	return result
}

// やってること同じだけどこっちのほうが綺麗
func intersectionV2(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]bool, 0)
	for _, n := range nums1 {
		numMap[n] = true
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if numMap[num] == true {
			result = append(result, num)
			numMap[num] = false
		}
	}

	return result
}

// leetCodeにpostした笑
func intersect(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, 0)
	for _, num := range nums1 {
		numMap[num]++
	}

	result := make([]int, 0)
	for _, num := range nums2 {
		if numMap[num] != 0 {
			result = append(result, num)
			numMap[num]--
		}
	}

	return result
}

// constraints: Do not use any built-in library function such as sqrt
func isPerfectSquare(num int) bool {
	if num == 0 {
		return false
	}

	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid > num {
			right = mid - 1
		} else if mid*mid < num {
			left = mid + 1
		} else {
			return true
		}
	}

	return false
}

// LeetCodeで用意された関数。コンパイルエラー防ぐために定義
func guess(n int) int {
	return 0
}

func guessNumber(n int) int {
	left, right := 0, n
	for left <= right {
		mid := (left + right) / 2
		if guess(mid) == -1 {
			right = mid - 1
		} else if guess(mid) == 1 {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

// incorrect
// 連続した値じゃなきゃいけないと思ってたら、そうじゃないらしい
func canConstruct(ransomNote string, magazine string) bool {
	lenRansomNote := len(ransomNote)
	for i := 0; i < len(magazine)-lenRansomNote+1; i++ {
		if ransomNote == magazine[i:i+lenRansomNote] {
			return true
		}
		if reverseStringForCanConstruct(ransomNote) == magazine[i:i+lenRansomNote] {
			return true
		}
	}
	return false
}

func reverseStringForCanConstruct(str string) string {
	var result string
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

// 正解だけど、少し処理が遅い
func canConstructV2(ransomNote string, magazine string) bool {
	magazineStrMap := make(map[string]int, 0)
	magazineSlice := strings.Split(magazine, "")
	ransomNoteSlice := strings.Split(ransomNote, "")
	for _, s := range magazineSlice {
		magazineStrMap[s]++
	}

	for _, s := range ransomNoteSlice {
		value, ok := magazineStrMap[s]
		if !ok {
			return false
		}
		if value == 0 {
			return false
		} else {
			magazineStrMap[s]--
		}
	}
	return true
}

// 後ろからみていく
// incorrect
// アイデアは良かったけど実装できず。。。
func firstUniqChar(s string) int {
	var resultIndex int
	strMap := make(map[uint8]struct{}, 0)
	for i := len(s) - 1; i >= 0; i-- {
		_, ok := strMap[s[i]]
		if !ok {
			resultIndex = i
			strMap[s[i]] = struct{}{}
		} else {
			if resultIndex == -1 {
				resultIndex = -1
			}
		}
	}

	return resultIndex
}

func firstUniqCharV2(s string) int {
	strMap := make(map[string]int, 0)
	for _, str := range s {
		strMap[string(str)]++
	}

	for i := 0; i < len(s); i++ {
		times, ok := strMap[string(s[i])]
		if ok && times == 1 {
			return i
		}
	}

	return -1
}

// Fail: s = "a" t = "aa"
// incorrect
func findTheDifference(s string, t string) byte {
	result := []rune{}
	for _, r := range t {
		if !isInForFindTheDifference(r, s) {
			result = append(result, r)
		}
	}

	fmt.Println(result)
	return byte(result[0])
}

func isInForFindTheDifference(target rune, list string) bool {
	for _, l := range list {
		if l == target {
			return true
		}
	}

	return false
}

func findTheDifferenceV2(s string, t string) byte {
	tmpMap := make(map[byte]int, 0)
	sList := []byte(s)
	tList := []byte(t)

	for _, str := range tList {
		tmpMap[str]++
	}

	for _, str := range sList {
		tmpMap[str]--
	}

	for k, v := range tmpMap {
		if v != 0 {
			return k
		}
	}
	return 0
}

func findTheDifferenceV3(s string, t string) byte {
	tmpMap := make(map[byte]int, 0)
	sList := []byte(s)
	tList := []byte(t)

	for i := 0; i < len(t); i++ {
		if i != len(t)-1 {
			tmpMap[sList[i]]++
			tmpMap[tList[i]]--
		} else {
			tmpMap[tList[i]]--
		}
	}

	for k, v := range tmpMap {
		if v != 0 {
			return k
		}
	}
	return 0
}

// ノーヒントでも結構できるようになってきた！
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	sArr := strings.Split(s, "")
	tArr := strings.Split(t, "")
	var result bool
	var i int
	for _, el := range sArr {
		result = false
		for i < len(t) {
			if tArr[i] == el {
				result = true
				i++
				break
			}
			i++
		}
	}

	return result
}

func isSubsequenceV2(s string, t string) bool {
	if s == "" {
		return true
	}

	var index int
	for i := 0; i < len(t); i++ {
		if t[i] == s[index] {
			index++
		}

		if index == len(s) {
			return true
		}
	}

	return false
}

func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	helperSumOfLeftLeaves(&result, root, false)
	return result
}

func helperSumOfLeftLeaves(result *int, root *TreeNode, isLeft bool) {
	if root == nil {
		return
	}

	if isLeft && root.Left == nil && root.Right == nil {
		*result += root.Val
	}

	helperSumOfLeftLeaves(result, root.Left, true)
	helperSumOfLeftLeaves(result, root.Right, false)
}

func isValid(s string) bool {
	strList := strings.Split(s, "")
	pStack := make([]string, 0)
	pMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, str := range strList {
		if str == "(" || str == "[" || str == "{" {
			pStack = append(pStack, str)
		} else {
			if len(pStack) == 0 || pStack[len(pStack)-1] != pMap[str] {
				return false
			} else {
				pStack = pStack[:len(pStack)-1]
			}
		}
	}

	return len(pStack) == 0
}

// LeetCode Level Medium
// おそらく初Medium
// 処理自体はあってるけどListNodeが大きくなったときに処理しきれない。非効率
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// calc sum of l1
	var sumL1 int
	currentL1 := l1
	digitL1 := 1
	for currentL1 != nil {
		sumL1 += currentL1.Val * digitL1
		digitL1 *= 10
		currentL1 = currentL1.Next
	}

	// calc sum of l2
	var sumL2 int
	currentL2 := l2
	digitL2 := 1
	for currentL2 != nil {
		sumL2 += currentL2.Val * digitL2
		digitL2 *= 10
		currentL2 = currentL2.Next
	}

	// make list node
	sum := sumL1 + sumL2
	if sum == 0 {
		return &ListNode{Val: 0}
	}

	var prevNode *ListNode
	var root *ListNode
	for sum > 0 {
		newNode := &ListNode{
			Val: sum % 10,
		}
		if prevNode == nil {
			root = newNode
		} else {
			prevNode.Next = newNode
		}
		sum /= 10
		prevNode = newNode
	}

	return root
}

// ListNodeの全てを参照するのではなく、
func addTwoNumbersV2(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	currentNode := root
	var tmp int
	for l1 != nil || l2 != nil {
		sum := tmp
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		currentNode.Next = &ListNode{Val: sum % 10}
		currentNode = currentNode.Next
		tmp = sum / 10
	}

	if tmp != 0 {
		currentNode.Next = &ListNode{Val: tmp % 10}
	}

	return root.Next
}

// incorrect "dvdf"のような場合に対応できてない
func lengthOfLongestSubstring(s string) int {
	var count int
	var result int
	tmpMap := make(map[byte]struct{}, 0)
	for i := 0; i < len(s); i++ {
		_, ok := tmpMap[s[i]]
		if !ok {
			count++
		} else {
			if count > result {
				result = count
			}
			tmpMap = make(map[byte]struct{}, 0)
			count = 1
		}

		if i == len(s)-1 && count > result {
			result = count
		}
		tmpMap[s[i]] = struct{}{}
	}

	return result
}

func longestPalindrome(s string) int {
	strMap := make(map[rune]int, 0)
	for _, str := range s {
		strMap[str]++
	}

	var result int
	var oddNum int
	for _, v := range strMap {
		if v%2 == 0 {
			result += v
		} else {
			result += v - 1
			oddNum = 1
		}
	}

	return result + oddNum
}

func fizzBuzz(n int) []string {
	result := make([]string, n, n)
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			result[i] = "FizzBuzz"
		} else if (i+1)%3 == 0 {
			result[i] = "Fizz"
		} else if (i+1)%5 == 0 {
			result[i] = "Buzz"
		} else {
			result[i] = strconv.Itoa(i + 1)
		}
	}

	return result
}

func thirdMax(nums []int) int {
	numsMap := make(map[int]struct{}, 0)
	resultList := make([]int, 0)
	for _, num := range nums {
		if _, ok := numsMap[num]; ok {
			continue
		}
		numsMap[num] = struct{}{}
		resultList = append(resultList, num)
	}

	sort.Ints(resultList)

	if len(resultList) < 3 {
		return resultList[len(resultList)-1]
	} else {
		return resultList[len(resultList)-3]
	}
}

func arrangeCoins(n int) int {
	count := n
	i := 1
	for count-i-1 > 0 {
		count = count - i - 1
		i++
	}
	return i
}

// これでも正解だが、効率が悪い. mapでのインデックスアクセスはスライスの数十倍遅いらしい
func findDisappearedNumbers(nums []int) []int {
	numsMap := make(map[int]struct{}, 0)
	for _, num := range nums {
		numsMap[num] = struct{}{}
	}

	result := make([]int, 0)
	for i := 1; i <= len(nums); i++ {
		_, ok := numsMap[i]
		if !ok {
			result = append(result, i)
		}
	}

	return result
}

// faster
func findDisappearedNumbersV2(nums []int) []int {
	numsList := make([]int, len(nums))
	for _, num := range nums {
		numsList[num-1] = num
	}

	result := make([]int, 0)
	for i := 0; i < len(numsList); i++ {
		if numsList[i] == 0 {
			result = append(result, i+1)
		}
	}

	return result
}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	var result int
	for _, gNum := range g {
		for i, sNum := range s {
			if gNum <= sNum {
				result++
				s = s[i+1:]
				break
			}

		}
	}

	return result
}

// これめっちゃ考えたけどわからなかった
func repeatedSubstringPattern(s string) bool {
	next := getNext(s)
	n := next[len(s)]
	m := len(s) - n
	if n >= m && len(s)%m == 0 {
		return true
	}
	return false
}

func getNext(s string) []int {
	res := make([]int, len(s)+1)

	i := 0
	j := -1
	res[0] = -1

	for i < len(s) {
		if j == -1 || s[i] == s[j] {
			i++
			j++
			res[i] = j
		} else {
			j = res[j]
		}
	}
	return res
}

func islandPerimeter(grid [][]int) int {
	var result int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				result += 4
				if i > 0 && grid[i-1][j] == 1 {
					result -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					result -= 2
				}
			}
		}
	}

	return result
}

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var count int
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if count > max {
				max = count
			}
			count = 0
		}
	}

	if count > max {
		max = count
	}

	return max
}

// めっちゃ効率悪い
func constructRectangle(area int) []int {
	result := []int{area, 1}
	for i := 1; i < area; i++ {
		if area%i != 0 {
			continue
		}

		devided := area / i
		if devided < i {
			continue
		}

		result[0] = devided
		result[1] = i
	}

	return result
}

// 最大が2乗付近だからSqrtを使う
func constructRectangleV2(area int) []int {
	maxNum := math.Sqrt(float64(area))
	result := make([]int, 2)
	for i := 1; i <= int(maxNum); i++ {
		if area%i == 0 {
			result[0] = area / i
			result[1] = i
		}
	}

	return result
}

func findPoisonedDuration(timeSeries []int, duration int) int {
	var count int
	for i := 0; i < len(timeSeries)-1; i++ {
		plusedDuration := timeSeries[i] + duration
		if plusedDuration > timeSeries[i+1] {
			count += (timeSeries[i+1] - timeSeries[i])
		} else {
			count += duration
		}
	}

	return count + duration
}

// 遅い
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := make([]int, 0, len(nums1))
	for i := 0; i < len(nums1); i++ {
		var iIdx int
		var exist bool
		var greaterIdx int
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				iIdx = j
				exist = true
			}
			if nums1[i] < nums2[j] && exist {
				greaterIdx = j
				break
			}
		}
		if iIdx < greaterIdx {
			result = append(result, nums2[greaterIdx])
		} else {
			result = append(result, -1)
		}
	}

	return result
}

// faster
func nextGreaterElementV2(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int, 0)
	result := make([]int, 0, len(nums2))
	for i := 0; i < len(nums2); i++ {
		numsMap[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		idx := numsMap[nums1[i]]
		var maxIdx int
		for j := idx; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				maxIdx = j
				break
			}
		}

		if maxIdx > idx {
			result = append(result, nums2[maxIdx])
		} else {
			result = append(result, -1)
		}
	}

	return result
}

// v2より読みやすく
func nextGreaterElementV3(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]int, 0)
	result := make([]int, 0, len(nums2))
	for i := 0; i < len(nums2); i++ {
		numsMap[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		idx := numsMap[nums1[i]]
		for j := idx; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				result = append(result, nums2[j])
				break
			}
			if j == len(nums2)-1 {
				result = append(result, -1)
			}
		}
	}

	return result
}

func findWords(words []string) []string {
	result := make([]string, 0)
	rowMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 2,
		"d": 1,
		"e": 0,
		"f": 1,
		"g": 1,
		"h": 1,
		"i": 0,
		"j": 1,
		"k": 1,
		"l": 1,
		"m": 2,
		"n": 2,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 1,
		"t": 0,
		"u": 0,
		"v": 2,
		"w": 0,
		"x": 2,
		"y": 0,
		"z": 2,
	}

	for _, word := range words {
		row := -1
		var isSameRow bool
		for i, w := range word {
			str := strings.ToLower(string(w))
			if i == 0 {
				row = rowMap[str]
				isSameRow = true
				continue
			}

			if rowMap[str] != row {
				isSameRow = false
				break
			}

			isSameRow = true
		}

		if isSameRow {
			result = append(result, word)
		}
	}

	return result
}

// 考え方は惜しかった。けど間違ってる
func findMode(root *TreeNode) []int {
	var result []int
	var count int
	helperFindMode(root, &result, &count)
	return result
}

func helperFindMode(root *TreeNode, result *[]int, count *int) {
	if root == nil {
		return
	}

	if root.Left != nil {
		if root.Left.Val == root.Val {
			*count++
		} else {
			*count = 0
		}
		if *count != 0 {
			*result = append(*result, root.Val)
		}
		helperFindMode(root.Left, result, count)
	}
	if root.Right != nil {
		if root.Right.Val == root.Val {
			*count++
		} else {
			*count = 0
		}
		if *count != 0 {
			*result = append(*result, root.Val)
		}
		helperFindMode(root.Right, result, count)
	}
}

// ヒント見た後
// これも惜しいけど不正解ぽい
// わからなかった...
func findModeV2(root *TreeNode) []int {
	var count int
	prev := 0
	max := 1
	result := make([]int, 0)
	helperFindModeV2(root, &result, &count, &max, &prev)
	return result
}

func helperFindModeV2(root *TreeNode, result *[]int, count *int, max *int, prev *int) {
	if root == nil {
		return
	}

	helperFindModeV2(root.Left, result, count, max, &root.Val)
	if root.Val == *prev {
		*count++
	} else {
		*count = 1
	}

	if *max < *count {
		*max = *count
		*result = []int{root.Val}
	} else if *max == *count {
		*result = append(*result, root.Val)
	}
	helperFindModeV2(root.Right, result, count, max, &root.Val)
}

// 1, 2, 3位以降はスコアを表にするものかと思ってた。やりなおしはV2で
func findRelativeRanks(scores []int) []string {
	result := make([]string, 0, len(scores))
	rankList := make([]int, len(scores))

	for i, score := range scores {
		if i == 0 {
			rankList[0] = score
			continue
		}

		if score > rankList[0] {
			rankList[1] = rankList[0]
			rankList[0] = score
		} else if score > rankList[1] {
			rankList[2] = rankList[1]
			rankList[1] = score
		} else if score > rankList[2] {
			rankList[2] = score
		}
	}

	for _, score := range scores {
		if score == rankList[0] {
			result = append(result, "Gold Medal")
		} else if score == rankList[1] {
			result = append(result, "Silver Medal")
		} else if score == rankList[2] {
			result = append(result, "Bronze Medal")
		} else {
			result = append(result, strconv.Itoa(score))
		}
	}

	return result
}

func findRelativeRanksV2(scores []int) []string {
	sortedScore := make([]int, len(scores))
	_ = copy(sortedScore, scores)
	sort.Ints(sortedScore)
	resultMap := make(map[int]string, 0)
	for i := len(sortedScore) - 1; i >= 0; i-- {
		if i == len(sortedScore)-1 {
			resultMap[sortedScore[i]] = "Gold Medal"
		} else if i == len(sortedScore)-2 {
			resultMap[sortedScore[i]] = "Silver Medal"
		} else if i == len(sortedScore)-3 {
			resultMap[sortedScore[i]] = "Bronze Medal"
		} else {
			resultMap[sortedScore[i]] = fmt.Sprintf("%d", len(sortedScore)-i)
		}
	}

	result := make([]string, 0, len(scores))
	for _, score := range scores {
		result = append(result, resultMap[score])
	}

	return result
}

func detectCapitalUse(words string) bool {
	if len(words) == 1 {
		return true
	}

	result := true
	var prevWord rune
	for i, word := range words {
		if i == 0 {
			prevWord = word
			continue
		}

		if (unicode.IsUpper(word) && unicode.IsLower(prevWord)) || (unicode.IsLower(word) && unicode.IsUpper(prevWord)) {
			if i == 1 && unicode.IsUpper(prevWord) {
				result = true
			} else {
				return false
			}
		}
		prevWord = word
	}

	return result
}

func getMinimumDifference(root *TreeNode) int {
	values := make([]int, 0)
	getValuesForGetMinimumDifference(root, &values)

	sort.Ints(values)

	result := math.MaxInt
	for i := 1; i < len(values); i++ {
		twoDifference := values[i] - values[i-1]
		if twoDifference < result {
			result = twoDifference
		}
	}

	return result
}

func getValuesForGetMinimumDifference(root *TreeNode, list *[]int) {
	if root == nil {
		return
	}
	getValuesForGetMinimumDifference(root.Right, list)
	*list = append(*list, root.Val)
	getValuesForGetMinimumDifference(root.Left, list)
}

// ヒントあり. 相変わらずBinary Treeは苦手
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return -1
	}

	var diameter float64
	helperForDiameterOfBT(root, &diameter)
	return int(diameter)
}

func helperForDiameterOfBT(root *TreeNode, diameter *float64) float64 {
	if root == nil {
		return 0
	}

	left := helperForDiameterOfBT(root.Left, diameter)
	right := helperForDiameterOfBT(root.Right, diameter)
	*diameter = math.Max(*diameter, left+right)
	return math.Max(left, right) + 1
}

func checkRecord(s string) bool {
	attendances := strings.Split(s, "")
	var numOfLate int
	var numOfAbsent int
	for _, attendance := range attendances {
		if attendance == "A" {
			numOfAbsent++
			numOfLate = 0
			if numOfAbsent > 1 {
				return false
			}
		} else if attendance == "L" {
			numOfLate++
			if numOfLate > 2 {
				return false
			}
		} else {
			numOfLate = 0
		}
	}

	return true
}

// Correct but slowly.
// Runtime: faster than abount 10%
func reverseWords(s string) string {
	splited := strings.Split(s, " ")
	var result string
	for _, word := range splited {
		var reversed string
		for i := len(word) - 1; i >= 0; i-- {
			reversed += string(word[i])
		}

		result += reversed + " "
	}

	result = result[:len(result)-1]

	return result
}

// a little bit faster
func reverseWordsV2(s string) string {
	splited := strings.Split(s, " ")
	var result string
	for _, word := range splited {
		wordRune := []rune(word)
		j := len(word) - 1
		for i := 0; i < len(word)/2; i++ {
			wordRune[i], wordRune[j] = wordRune[j], wordRune[i]
			j--
		}

		result += string(wordRune) + " "
	}

	result = result[:len(result)-1]

	return result
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	var result int
	for i, num := range nums {
		if i%2 == 0 {
			result += num
		}
	}

	return result
}

func arrayPairSumV2(nums []int) int {
	sort.Ints(nums)
	var result int
	for i := 0; i < len(nums); i = i + 2 {
		result += nums[i]
	}
	return result
}

func main() {
	fmt.Println(reverseWords("I am Kosuke"))
	fmt.Println(reverseWordsV2("I am Kosuke"))
	fmt.Println(arrayPairSum([]int{6, 2, 6, 5, 1, 2}))
	fmt.Println(arrayPairSumV2([]int{6, 2, 6, 5, 1, 2}))
}
