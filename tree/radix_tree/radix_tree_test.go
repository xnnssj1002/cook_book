package radix_tree

import "testing"

func TestCommonPrefixLen(t *testing.T) {
	wordA, wordB := "hello", "hero"
	i := commonPrefixLen("hello", "hero")
	t.Logf("[%s] 与 [%s] 公共前缀长度为：%d", wordA, wordB, i)

}
