// You have a long flowerbed in which some of the plots are planted, and some are not.
// However, flowers cannot be planted in adjacent plots.

// Given an integer array flowerbed containing 0's and 1's, where 0 means empty and 1
// means not empty, and an integer n, return true if n new flowers can be planted in the
// flowerbed without violating the no-adjacent-flowers rule and false otherwise.

// Example 1:

// Input: flowerbed = [1,0,0,0,1], n = 1
// Output: true
// Example 2:

// Input: flowerbed = [1,0,0,0,1], n = 2
// Output: false

package leetcodeArrayChallenges

func CanPlaceFlowers(flowerbed []int, n int) bool {
	spacesAvailableToPlant := 0

	for i := 0; i < len(flowerbed); i++ {
		prev := 0
		next := 0

		if i > 0 {
			prev = flowerbed[i-1]

		}
		if i < len(flowerbed)-1 {

			next = flowerbed[i+1]
		}
		if prev == 0 && flowerbed[i] == 0 && next == 0 {
			flowerbed[i] = 1
			spacesAvailableToPlant++
		}
	}

	return spacesAvailableToPlant >= n

}