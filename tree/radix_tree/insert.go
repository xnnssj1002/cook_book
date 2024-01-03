package radix_tree

// Insert 插入一个单词
// 插入的核心逻辑：
// • 首先，检查一下 word 是否已存在，如果存在，则终止流程，无需重复操作
// • 以 radix tree 的根节点 root 作为起点，调用 radixNode.insert 方法，执行插入流程
// • 如果此前 radix tree 中从未插入过任何单词，则直接将首个单词插入到 root 当中
// • 开启 for 循环，根据节点的相对路径 path 与插入 word 之间的关系进，进行分类处理：
// • 求出 path 与 word 的公共前缀长度 i
// • 如果公共前缀 i 大于 0，说明 word 必然经过当前节点，需要对 passCnt 计数器加 1
// • 如果公共前缀 i 小于 path 长度，则要将当前节点拆分为公共前缀部分 + 后继剩余部分两个节点
// • 如果公共前缀长度小于 word 长度，则需要继续检查，word 和后继节点是否还有公共前缀，如果有的话，则递归对后继节点执行相同流程
// • 当保证 word 和后继节点不再有公共前缀，则直接将 word 包装成一个新的节点，插入到当前节点的子节点列表 children 当中
func (r *Radix) Insert(word string) {
	// 不重复插入
	if r.Search(word) {
		return
	}
	r.root.insert(word)
}

func (rn *radixNode) insert(word string) {
	fullWord := word

	// 如果当前节点为 root，此之前没有注册过子节点，则直接插入并返回
	if rn.path == "" && len(rn.children) == 0 {
		rn.insertWord(word, word)
		return
	}

walk:
	for {
		// 获取到 word 和当前节点 path 的公共前缀长度
		i := commonPrefixLen(word, rn.path)
		// 只要公共前缀大于 0，则一定经过当前节点，需要累加 passCnt
		if i > 0 {
			rn.passCnt++
		}

		// 公共前缀小于当前节点的相对路径，需要对节点进行分解
		if i < len(rn.path) {
			// 需要进行节点切割
			child := radixNode{
				// 进行相对路径切分
				path: rn.path[i:],
				// 继承完整路径
				fullPath: rn.fullPath,
				// 对当前节点的后继节点进行委托
				children: rn.children,
				indices:  rn.indices,
				end:      rn.end,
				// 传承给孩子节点时，需要把之前累加上的 passCnt 计数扣除
				passCnt: rn.passCnt - 1,
			}

			// 续接上孩子节点
			rn.indices = string(rn.path[i])
			rn.children = []*radixNode{&child}
			// 调整原节点的 full path
			rn.fullPath = rn.fullPath[:len(rn.fullPath)-(len(rn.path)-i)]
			// 调整原节点的 path
			rn.path = rn.path[:i]
			// 原节点是新拆分出来的，目前不可能有单词以该节点结尾
			rn.end = false
		}

		// 公共前缀小于插入 word 的长度
		if i < len(word) {
			// 对 word 扣除公共前缀部分
			word = word[i:]
			// 获取 word 剩余部分的首字母
			c := word[0]
			for i := 0; i < len(rn.indices); i++ {
				// 如果与后继节点还有公共前缀，则将 rn 指向子节点，然后递归执行流程
				if rn.indices[i] == c {
					rn = rn.children[i]
					continue walk
				}
			}

			// 到了这里，意味着 word 剩余部分与后继节点没有公共前缀了
			// 此时直接构造新的节点进行插入
			rn.indices += string(c)
			child := radixNode{}
			child.insertWord(word, fullWord)
			rn.children = append(rn.children, &child)
			return
		}

		// 倘若公共前缀恰好是 path，需要将 end 置为 true
		rn.end = true
		return
	}
}

// 求取两个单词的公共前缀
func commonPrefixLen(wordA, wordB string) int {
	var move int
	for move < len(wordA) && move < len(wordB) && wordA[move] == wordB[move] {
		move++
	}
	return move
}

// 传入相对路径和完整路径，补充一个新生成的节点信息
func (rn *radixNode) insertWord(path, fullPath string) {
	rn.path, rn.fullPath = path, fullPath
	rn.passCnt = 1
	rn.end = true
}
