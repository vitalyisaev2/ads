package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
	assert.Equal(t, "bab", longestPalindrome("babad"))
	assert.Equal(t, "a", longestPalindrome("ac"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "bb", longestPalindrome("bb"))
	assert.Equal(t, "abccba", longestPalindrome("abccbacba"))
	assert.Equal(t, "abcdcba", longestPalindrome("abcdcbacba"))
	assert.Equal(t, "bb", longestPalindrome("abb"))
	assert.Equal(t, "aaabaaa", longestPalindrome("aaabaaaa"))
	assert.Equal(t, "adada", longestPalindrome("babadada"))
	assert.Equal(t, "a", longestPalindrome("abcda"))
	assert.NotPanics(t, func() {
		longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth")
	})
}
