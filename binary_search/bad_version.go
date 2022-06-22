package binary_search

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

// WANT is a package variable to manipulate the `isBadVersion` within a test - only doing so because of weird leetcode setup
// that doesn't utilize function injection for testing
var WANT int

func firstBadVersion(n int) int {
	var low, mid int
	var high = n
	var isBad bool

	for low <= high {
		mid = (low + high) / 2
		isBad = isBadVersion(mid)
		if !isBad {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !isBad {
		mid++
	}

	return mid

}

func isBadVersion(version int) bool {
	if version >= WANT {
		return true
	}

	return false
}
