/*
Package t11

You are given an integer array height of length n.
There are n vertical lines drawn such that the two
endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container,
such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.
*/
package t11

func MyMin(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func getVolume(a, b int, nums []int) int {
	return (b - a) * MyMin(nums[a], nums[b])
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxV := 0
	for left < right && right < len(height) {
		currentV := getVolume(left, right, height)
		if currentV > maxV {
			maxV = currentV
		}
		// Внимание. Вот вся соль
		// сканирование окошком, с самого большого
		// окна и к центу по логике:
		if height[left] < height[right] {
			left++
		} else if height[left] >= height[right] {
			right--
		}
	}
	return maxV
}
