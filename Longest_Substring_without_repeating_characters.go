
func lengthOfLongestSubstring(s string) int {
    var lastSeen [256]int

    longest := 0
    left := 0

    for right := 0; right < len(s); right++ {
        char := s[right]

        if lastSeen[char] > left {
            left = lastSeen[char]
        }

        lastSeen[char] = right + 1

        if right-left+1 > longest {
            longest = right - left + 1
        }
    }

    return longest
}
