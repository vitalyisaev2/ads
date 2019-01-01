package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, 0, lengthOfLongestSubstring(""))
	assert.Equal(t, 1, lengthOfLongestSubstring(" "))
	assert.Equal(t, 2, lengthOfLongestSubstring("au"))
	assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	assert.Equal(t, 2, lengthOfLongestSubstring("aab"))
}
