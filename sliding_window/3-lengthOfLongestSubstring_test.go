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
		name: "test2",
		s:    "pwwkew",
		hope: 3,
	},
	{
		name: "test3",
		s:    " ",
		hope: 1,
	},
	{
		name: "test4",
		s:    "cdd",
		hope: 2,
	},
}

func TestLengthOfLongestSubstring1(t *testing.T) {

	for _, source := range lengthOfLongestSubstringSources {
		r := lengthOfLongestSubstring1(source.s)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestLengthOfLongestSubstring2(t *testing.T) {

	for _, source := range lengthOfLongestSubstringSources {
		r := lengthOfLongestSubstring2(source.s)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}

func TestLengthOfLongestSubstring3(t *testing.T) {

	for _, source := range lengthOfLongestSubstringSources {
		r := lengthOfLongestSubstring3(source.s)
		if r != source.hope {
			t.Errorf("【%s】exec err, hope %v, got %v\n", source.name, source.hope, r)

		} else {
			t.Logf("【%s】exec ok\n", source.name)
		}
	}

}
