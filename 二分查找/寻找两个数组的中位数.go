func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	// 保证 m <= n，二分时会方便
	if m > n {
		return findMedianSortedArrays(nums2, nums1)
	}
	left, right := 0, m
	totalLeft := (m + n + 1) / 2
	for left <= right {
		i := (left + right) / 2
		j := totalLeft - i

		nums1LeftMax := ^int(^uint(0) >> 1) // min int
		if i > 0 {
			nums1LeftMax = nums1[i-1]
		}
		nums1RightMin := int(^uint(0) >> 1) // max int
		if i < m {
			nums1RightMin = nums1[i]
		}
		nums2LeftMax := ^int(^uint(0) >> 1)
		if j > 0 {
			nums2LeftMax = nums2[j-1]
		}
		nums2RightMin := int(^uint(0) >> 1)
		if j < n {
			nums2RightMin = nums2[j]
		}

		if nums1LeftMax <= nums2RightMin && nums2LeftMax <= nums1RightMin {
			// 已经划分完毕
			if (m+n)%2 == 1 {
				return float64(max(nums1LeftMax, nums2LeftMax))
			}
			return float64(max(nums1LeftMax, nums2LeftMax)+min(nums1RightMin, nums2RightMin)) / 2.0
		} else if nums1LeftMax > nums2RightMin {
			right = i - 1
		} else {
			left = i + 1
		}
	}
	return 0.0 // 不会走到
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}