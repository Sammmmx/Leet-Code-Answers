
func lengthOfLongestSubstring(s string) int {
    lastSeen := make(map[byte]int)
    longest := 0
    left := 0

    for right := 0; right < len(s); right++ {
        char := s[right]

        if index, exists := lastSeen[char]; exists && index >= left {
            left = index + 1
        }

        lastSeen[char] = right

        current := right - left + 1

        if current > longest {
            longest = current
        }
    }

    return longest
}
