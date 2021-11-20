/**
  3. Longest Substring Without Repeating Characters

  Given a string s, find the length of the longest substring without repeating characters.



  Example 1:
  Input: s = "abcabcbb"
  Output: 3
  Explanation: The answer is "abc", with the length of 3.

  Example 2:
  Input: s = "bbbbb"
  Output: 1
  Explanation: The answer is "b", with the length of 1.

  Example 3:
  Input: s = "pwwkew"
  Output: 3
  Explanation: The answer is "wke", with the length of 3.
  Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

  Example 4:
  Input: s = ""
  Output: 0


  Constraints:

  0 <= s.length <= 5 * 104
  s consists of English letters, digits, symbols and spaces.
 */
 
 func lengthOfLongestSubstring(s string) int {
    start, maxLen := 0, 0
    record := make(map[string]int)
    
    for i := 0; i < len(s); i++ {
        item := s[i:i+1]
        
        endPos := i
        if _, ok := record[item]; ok {
            newStart := record[item] + 1
            for j := start; j < newStart; j++ {
                delete(record, s[j:j+1])
            }
            start = newStart
            endPos = i - 1
        }
        
        if maxLen < endPos-start+1 {
            maxLen = endPos - start + 1
        }
        record[item] = i
    }
    return maxLen
}
