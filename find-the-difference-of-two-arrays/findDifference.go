/* Package t2215
* Given two 0-indexed integer arrays nums1 and nums2, return a list answer of size 2 where:
* answer[0] is a list of all distinct integers in nums1 which are not present in nums2.
* answer[1] is a list of all distinct integers in nums2 which are not present in nums1.
* Note that the integers in the lists may be returned in any order.
 */
package t2215

func Unique[T comparable](slice []T) []T {
	uniqueMap := make(map[T]struct{})
	var uniqueSlice []T

	for _, v := range slice {
		if _, exists := uniqueMap[v]; !exists {
			uniqueMap[v] = struct{}{}
			uniqueSlice = append(uniqueSlice, v)
		}
	}

	return uniqueSlice
}

func Intersection[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	var result []T

	for _, item := range a {
		set[item] = true
	}

	for _, item := range b {
		if set[item] {
			result = append(result, item)
			set[item] = false
		}
	}
	return result
}

func Subtract[T comparable](a, b []T) []T {
	set := make(map[T]bool)
	var result []T
	for _, item := range b {
		set[item] = true
	}

	for _, item := range a {
		if !set[item] {
			result = append(result, item)
		}
	}
	return result
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	nums1 = Unique(nums1)
	nums2 = Unique(nums2)
	intercsection := Intersection(nums1, nums2)
	nums1 = Subtract(nums1, intercsection)
	nums2 = Subtract(nums2, intercsection)
	return [][]int{nums1, nums2}
}
