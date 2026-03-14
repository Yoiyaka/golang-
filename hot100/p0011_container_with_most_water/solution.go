package p0011_container_with_most_water

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > max {
			max = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}
