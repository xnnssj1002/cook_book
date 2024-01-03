package trie

// 前缀树实现-优化版本

// OptimizationTrie trieNode
type OptimizationTrie struct {
	isEnd    bool
	children [26]*OptimizationTrie
}

// NewOptimizationTrie 初始化一个根节点
func NewOptimizationTrie() *OptimizationTrie {
	return &OptimizationTrie{}
}

// Insert 添加一个单词
func (t *OptimizationTrie) Insert(word string) {
	// 首先将【根结点】作为第一个【当前结点】
	currentNode := t
	// 循环单词的字符，依次判断循环到的字符是否在【当前结点】的【子结点】中
	for _, ch := range word {
		// 判断循环到的【字符】是否在【当前结点】的子结点中
		// 不存在的话，就重新创建一个子结点
		ch -= 'a'
		if currentNode.children[ch] == nil { // 不存在创建一个子结点
			currentNode.children[ch] = &OptimizationTrie{}
		}
		// 然后将循环到【当前结点】的【指定子结点】，复制给当前结点，方便进行下次循环判断
		currentNode = currentNode.children[ch] // 将当前结点的子结点，重新复制为
	}

	// 当前字符串循环处理完成后，将【最后子结点】即最后【当前结点】的 isEnd 设置为 true，表示该结点为一个单词结尾
	currentNode.isEnd = true
}

// Search 搜索一个单词是否存在于 trie 树中
// 存在的要求为：
// 1、单词的所有字符都可以在 trie 树中找到
// 2、最后一个字符对应【子结点】的 isEnd 为 true
func (t *OptimizationTrie) Search(word string) bool {
	searchNode := t.searchPrefix(word)
	return searchNode != nil && searchNode.isEnd
}

// StartsWith 搜索指定前缀字符串，是否存在于 trie 中
func (t *OptimizationTrie) StartsWith(prefix string) bool {
	return t.searchPrefix(prefix) != nil
}

// searchPrefix 获取指定前缀字符串所在的结点
func (t *OptimizationTrie) searchPrefix(prefix string) *OptimizationTrie {
	// 首先将【根结点】作为第一个【当前结点】
	currentNode := t
	// 循环单词的字符，依次判断循环到的字符是否在【当前结点】的【子结点】中
	for _, ch := range prefix {
		ch -= 'a'
		// 如果循环到的字符，在【当前结点】的【子结点】中不存在，直接返回 nil。表示没有找到【前缀字符串】所在的子结点
		if currentNode.children[ch] == nil {
			return nil
		}
		// 然后将循环到【当前结点】的【指定子结点】，复制给当前结点，方便进行下次循环判断
		currentNode = currentNode.children[ch]
	}
	return currentNode
}
