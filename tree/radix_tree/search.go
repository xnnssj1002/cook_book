package radix_tree

// Search 查看一个单词在 radix 当中是否存在
// 查询的核型逻辑
// • 调用 radix tree 根节点 root 对应的 search 方法，开启查询流程
// • 倘若 word 内容刚好和节点 path 相等，则返回节点，并在上层通过节点的 end 标识，判断是否存在单词以当前节点结尾
// • 倘若 word 以节点 path 作为前缀，则查看后继节点是否还与 path 存在公共前缀，如果是的话，将节点指针指向后继节点，递归开启后续流程
// • 除此之外，都说明 word 不存于 radix tree 中，直接终止流程
func (r *Radix) Search(word string) bool {
	node := r.root.search(word)
	return node != nil && node.fullPath == word && node.end
}

func (rn *radixNode) search(word string) *radixNode {
walk:
	for {
		prefix := rn.path
		// word 长于 path
		if len(word) > len(prefix) {
			// 没匹配上，直接返回 nil
			if word[:len(prefix)] != prefix {
				return nil
			}
			// word 扣除公共前缀后的剩余部分
			word = word[len(prefix):]
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				// 后继节点还有公共前缀，继续匹配
				if c == rn.indices[i] {
					rn = rn.children[i]
					continue walk
				}
			}
			// word 还有剩余部分，但是 prefix 不存在后继节点和 word 剩余部分有公共前缀了
			// 必然不存在
			return nil
		}

		// 和当前节点精准匹配上了
		if word == prefix {
			return rn
		}

		// 走到这里意味着 len(word) <= len(prefix) && word != prefix
		return rn
	}
}
