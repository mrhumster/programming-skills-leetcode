/*
Package t1493

Учитывая двоичный массив nums, вам следует удалить из него один элемент.

Возвращает размер самого длинного непустого подмассива, содержащего только единицы, в
результирующем массиве. Возвращает 0, если такого подмассива нет.
*/
package t1493

func longestSubarray(nums []int) int {
	left, right := 0, 0
	zeroCount := 0
	maxLength := 0
	for right < len(nums) {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		currentLength := right - left
		if currentLength > maxLength {
			maxLength = currentLength
		}
		right++
	}
	return maxLength
}
