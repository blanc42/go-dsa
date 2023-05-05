package main

import (
	"dsa/ds/stack"
	"fmt"
)

func MinutesToSeconds(x int) int {
	return x / 60
}

func bitwise(x, y int) {
	fmt.Println("Bitwise AND", x&y)
	fmt.Println("Bitwise OR", x|y)
	fmt.Println("Bitwise XOR", x^y)
	fmt.Println("Right Shift by 2", x>>2)
	fmt.Println("Left shift by 1", x<<1)
}

func exits(x, y int) bool {
	if x >= 0 && y >= 0 {
		return true
	}
	return false
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	testval := image[sr][sc]
	image[sr][sc] = color
	points := [][]int{{sr - 1, sc}, {sr - 1, sc}, {sr, sc - 1}, {sr, sc + 1}}
	for _, p := range points {
		if exits(p[0], p[1]) {
			if testval == image[p[0]][p[1]] {
				image[p[0]][p[1]] = color
				floodFill(image, p[0], p[1], color)
			}
		}
	}
	return image
}

func tryStack() {
	s := new(stack.Stack)
	fmt.Println(s)
	s.Push(10)
	fmt.Println(s)
	s.Push(11)
	fmt.Println(s)
	s.Push(12)
	fmt.Println(s)
	s.Push(13)
	fmt.Println(s)
	for !s.IsEmpty() {
		fmt.Println(s.Pop())
	}
}

type edge struct {
	to   int
	cost int
}

type Node struct {
	vertex int
	cost   int
}

// 724
func pivotIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	} else if len(nums) == 0 {
		return -1
	}
	start := -1
	end := len(nums) - 1
	leftSum := 0
	rightSum := nums[end]
	for start < end {
		if leftSum == rightSum && (end-start) == 2 {
			return (end + start) / 2
		} else if leftSum > rightSum {
			end--
			rightSum += nums[end]
		} else {
			start++
			leftSum += nums[start]
		}
	}
	return -1
}

// problem 1
func twoSum(nums []int, target int) []int {
	var x, y int
	for i, v := range nums {
		for j, u := range nums[i+1:] {
			if v+u == target {
				x, y = i, j+i+1
			}
		}
	}
	return []int{x, y}
}
func maxArea(height []int) int {
	l := len(height)
	var max int
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	max := len(nums1) + len(nums2)
	i, j := 0, 0
	var arr []int
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			arr = append(arr, nums1[i])
			i++
		} else {
			arr = append(arr, nums2[j])
			j++
		}
	}
	if i <= len(nums1)-1 {
		arr = append(arr, nums1[i:]...)
	}
	if j <= len(nums2)-1 {
		arr = append(arr, nums2[j:]...)
	}
	if max%2 == 0 {
		return (float64(arr[max/2] + arr[max/2-1])) / 2
	} else {
		return float64(arr[max/2])
	}
}

func findMaxAverage(nums []int, k int) float64 {
	if len(nums) < 1 {
		return float64(nums[0])
	}
	start := 1
	sum := 0
	for _, v := range nums[0:k] {
		sum += v
	}
	prev_avg := float64(sum) / float64(k)
	for start+k-1 < len(nums) {
		curr_avg := float64(prev_avg*float64(k)+float64(nums[start+k-1])-float64(nums[start-1])) / float64(k)
		if curr_avg > prev_avg {
			prev_avg = curr_avg
		}
		start++
	}
	return prev_avg
}

func totalFruit(fruits []int) int {
	var max int
	i := 0
	for i < len(fruits) {
		t1 := fruits[i]
		t2 := -1
		j := i + 1
		var t2_index int
		for j < len(fruits) {
			if fruits[j] == t1 {
				j++
				continue
			} else if fruits[j] != t1 && t2 == -1 {
				t2 = fruits[j]
				t2_index = j
				j++
				continue
			} else if fruits[j] == t1 || fruits[j] == t2 {
				j++
				continue
			} else {
				break
			}
		}
		if j-i > max {
			max = j - i
		}
		i = t2_index
		if j == len(fruits) {
			return max
		}

	}
	return max
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	var res int64 = 0
	var minFound, maxFound bool = false, false
	var start, minStart, maxStart int = 0, 0, 0

	for i := 0; i < len(nums); i++ {
		num := nums[i]

		if num < minK || num > maxK {
			minFound = false
			maxFound = false
			start = i + 1
		}

		if num == minK {
			minFound = true
			minStart = i
		}

		if num == maxK {
			maxFound = true
			maxStart = i
		}

		if minFound && maxFound {
			res += int64(min(minStart, maxStart) - start + 1)
		}
	}
	return res
}

func main() {
	fmt.Println(countSubarrays([]int{2, 1, 2, 3, 5, 2, 7, 5, 1, 1}, 1, 5))
}
