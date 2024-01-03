package trie

/* 208. 实现 Trie (前缀树)
Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补完和拼写检查。
请你实现 Trie 类：
Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false
*/

/* 示例：
输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True
*/

// Trie trieNode
type Trie struct {
	isEnd    bool
	children [26]*Trie
}

// Constructor 初始化一个根节点
func Constructor() *Trie {
	return &Trie{}
}

// Insert 添加一个单词
func (t *Trie) Insert(word string) {
	// 首先将【根结点】作为第一个【当前结点】
	currentNode := t
	// 循环单词的字符，依次判断循环到的字符是否在【当前结点】的【子结点】中
	for _, v := range word {
		// 判断循环到的【字符】是否在【当前结点】的子结点中
		// 不存在的话，就重新创建一个子结点
		ch := v - 'a'
		if currentNode.children[ch] == nil { // 不存在创建一个子结点
			currentNode.children[ch] = &Trie{}
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
func (t *Trie) Search(word string) bool {
	// 首先将【根结点】作为第一个【当前结点】
	currentNode := t
	// 循环单词的字符，依次判断循环到的字符是否在【当前结点】的【子结点】中
	for _, v := range word {
		ch := v - 'a'
		// 如果循环到的字符，在【当前结点】的【子结点】中不存在，直接返回 false
		if currentNode.children[ch] == nil {
			return false
		}
		// 然后将循环到【当前结点】的【指定子结点】，复制给当前结点，方便进行下次循环判断
		currentNode = currentNode.children[ch]
	}
	return currentNode.isEnd
}

// StartsWith 搜索指定前缀字符串，是否存在于 trie 中
func (t *Trie) StartsWith(prefix string) bool {
	// 首先将【根结点】作为第一个【当前结点】
	currentNode := t
	// 循环单词的字符，依次判断循环到的字符是否在【当前结点】的【子结点】中
	for _, v := range prefix {
		ch := v - 'a'
		// 如果循环到的字符，在【当前结点】的【子结点】中不存在，直接返回 false
		if currentNode.children[ch] == nil {
			return false
		}
		// 然后将循环到【当前结点】的【指定子结点】，复制给当前结点，方便进行下次循环判断
		currentNode = currentNode.children[ch]
	}
	return true
}
