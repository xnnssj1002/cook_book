package sliding_window

import "testing"

var lengthOfLongestSubstringSources = []struct {
	name string
	s    string
	hope int
}{
	{
		name: "test0",
		s:    "abcabcbb",
		hope: 3,
	},
	{
		name: "test1",
		s:    "bbbbb",
		hope: 1,
	},
	{
		name: "test1",
		s:    "pwwkew",
		hope: 3,
	},
}

func TestLengthOfLongestSubstring(t *testing.T) {

	for _, source := range lengthOfLongestSubstringSources {
		r := lengthOfLongestSubstring(source.s)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
