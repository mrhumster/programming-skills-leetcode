/*
Package t605

You have a long flowerbed in which some of the plots are planted,
and some are not. However, flowers cannot be planted in adjacent plots.

Given an integer array flowerbed containing 0's and 1's, where 0 means
empty and 1 means not empty, and an integer n, return true if n new flowers
can be planted in the flowerbed without violating the no-adjacent-flowers
rule and false otherwise.
*/
package t605

func canPlaceFlowers(flowerbed []int, n int) bool {
	isBusy := func(pos int) bool {
		return flowerbed[pos] == 1
	}

	isEthical := func(pos int) bool {
		if pos == 0 {
			if len(flowerbed) == 1 {
				return true
			}
			return flowerbed[pos+1] == 0
		}
		if pos == len(flowerbed)-1 {
			return flowerbed[pos-1] == 0
		}
		return flowerbed[pos-1] == 0 && flowerbed[pos+1] == 0
	}

	for i := range flowerbed {
		busy := isBusy(i)
		if busy {
			continue
		}
		if isEthical(i) {
			flowerbed[i] = 1
			n -= 1
		}
	}
	return n <= 0
}

